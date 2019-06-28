package handlers_test

import (
	"bytes"
	"code.cloudfoundry.org/gorouter/common/uuid"
	"code.cloudfoundry.org/gorouter/handlers"
	logger_fakes "code.cloudfoundry.org/gorouter/logger/fakes"
	"code.cloudfoundry.org/gorouter/registry/fakes"
	"code.cloudfoundry.org/gorouter/route"
	"code.cloudfoundry.org/gorouter/test_util"
	"encoding/binary"
	"errors"
	"github.com/cloudfoundry/dropsonde/emitter/fake"
	"github.com/cloudfoundry/sonde-go/events"
	gouuid "github.com/nu7hatch/gouuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/urfave/negroni"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("HTTPStartStop Handler", func() {
	var (
		vcapHeader  string
		handler     *negroni.Negroni
		nextHandler http.HandlerFunc

		resp http.ResponseWriter
		req  *http.Request

		fakeEmitter *fake.FakeEventEmitter
		fakeLogger  *logger_fakes.FakeLogger
		fakeRegistry *fakes.FakeRegistry

		nextCalled bool
	)

	BeforeEach(func() {
		body := bytes.NewBufferString("What are you?")
		req = test_util.NewRequest("GET", "example.com", "/", body)
		resp = httptest.NewRecorder()

		var err error
		vcapHeader, err = uuid.GenerateUUID()
		Expect(err).ToNot(HaveOccurred())
		req.Header.Set(handlers.VcapRequestIdHeader, vcapHeader)

		fakeEmitter = fake.NewFakeEventEmitter("fake")
		fakeLogger = new(logger_fakes.FakeLogger)
		fakeRegistry = &fakes.FakeRegistry{}


		fakeRegistry.LookupStub = func(uri route.Uri) *route.EndpointPool {
			routePool := route.NewPool(&route.PoolOpts{})
			endpointWithTags := route.NewEndpoint(&route.EndpointOpts{
				Tags: map[string]string{"component": "route-emitter"},
			})
			routePool.Put(endpointWithTags)

			return routePool
		}

		nextHandler = http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			_, err := ioutil.ReadAll(req.Body)
			Expect(err).NotTo(HaveOccurred())

			rw.WriteHeader(http.StatusTeapot)
			rw.Write([]byte("I'm a little teapot, short and stout."))
			nextCalled = true
		})
		nextCalled = false
	})

	JustBeforeEach(func() {
		handler = negroni.New()
		handler.Use(handlers.NewRequestInfo())
		handler.Use(handlers.NewProxyWriter(fakeLogger))
		handler.Use(handlers.NewHTTPStartStop(fakeEmitter, fakeLogger, fakeRegistry))
		handler.UseHandlerFunc(nextHandler)
	})

	It("emits an HTTP StartStop event", func() {
		handler.ServeHTTP(resp, req)
		var startStopEvent *events.HttpStartStop
		var startStopEnvelope *events.Envelope
		findStartStopEnvelope := func() *events.Envelope {
			for _, ev := range fakeEmitter.GetEnvelopes() {
				if ev.GetEventType() == events.Envelope_HttpStartStop {
					startStopEnvelope = ev
					return startStopEnvelope
				}
			}
			return nil
		}

		Eventually(findStartStopEnvelope).ShouldNot(BeNil())
		startStopEvent = startStopEnvelope.GetHttpStartStop()
		reqID := startStopEvent.GetRequestId()
		var reqUUID gouuid.UUID
		binary.LittleEndian.PutUint64(reqUUID[:8], reqID.GetLow())
		binary.LittleEndian.PutUint64(reqUUID[8:], reqID.GetHigh())
		Expect(reqUUID.String()).To(Equal(vcapHeader))
		Expect(startStopEvent.GetMethod().String()).To(Equal("GET"))
		Expect(startStopEvent.GetStatusCode()).To(Equal(int32(http.StatusTeapot)))
		Expect(startStopEvent.GetContentLength()).To(Equal(int64(37)))
		Expect(startStopEnvelope.GetTags()).To(Equal(map[string]string{"component": "route-emitter"}))

		Expect(nextCalled).To(BeTrue(), "Expected the next handler to be called.")
	})

	Context("when the response writer is not a proxy response writer", func() {
		var badHandler *negroni.Negroni
		BeforeEach(func() {
			badHandler = negroni.New()
			badHandler.Use(handlers.NewHTTPStartStop(fakeEmitter, fakeLogger, fakeRegistry))
		})
		It("calls Fatal on the logger", func() {
			badHandler.ServeHTTP(resp, req)
			Expect(fakeLogger.FatalCallCount()).To(Equal(1))

			Expect(nextCalled).To(BeFalse())
		})
	})

	Context("when VcapRequestIdHeader is not provided", func() {
		BeforeEach(func() {
			req.Header.Set(handlers.VcapRequestIdHeader, "")
		})
		It("calls Fatal on the logger", func() {
			handler.ServeHTTP(resp, req)
			Expect(fakeLogger.FatalCallCount()).To(Equal(1))

			Expect(nextCalled).To(BeFalse())
		})
	})

	Context("when the emitter fails to emit", func() {
		BeforeEach(func() {
			fakeEmitter.ReturnError = errors.New("foo-error")
		})
		It("calls Info on the logger, but does not fail the request", func() {
			handler.ServeHTTP(resp, req)
			Expect(fakeLogger.InfoCallCount()).To(Equal(1))

			Expect(nextCalled).To(BeTrue())
		})
	})
})
