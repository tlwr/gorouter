package integration

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = XDescribe("route table being updated by NATS messages", func() {
	var testState *testState

	BeforeEach(func() {
		testState = NewTestState()
	})

	AfterEach(func() {
		if testState != nil {
			testState.StopAndCleanup()
		}
	})

	Context("When NATS emits messages with complete_uri_list set to true", func() {
		var testApp *httptest.Server

		BeforeEach(func() {
			testState.StartGorouter()
			testApp = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
			}))

			testState.registerWithCompleteUriList(testApp, []string{
				"one.some.domain",
				"two.some.domain",
			})

		})

		AfterEach(func() {
			testApp.Close()
		})

		It("Can access the app via both URIs", func() {

			req := testState.newRequest("http://one.some.domain")
			resp, err := testState.client.Do(req)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			resp.Body.Close()

			req = testState.newRequest("http://two.some.domain")
			resp, err = testState.client.Do(req)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			resp.Body.Close()
		})

		Context("when getting new routes", func() {
			testState.registerWithCompleteUriList(testApp, []string{
				"one.some.domain",
			})

			It("Can access the app via the new URI", func() {
				req := testState.newRequest("http://one.some.domain")
				resp, err := testState.client.Do(req)
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(200))
				resp.Body.Close()
			})

			It("cannot access the app via old URI", func() {
				req := testState.newRequest("http://two.some.domain")
				resp, err := testState.client.Do(req)
				Expect(err).NotTo(HaveOccurred())
				Expect(resp.StatusCode).To(Equal(404))
				resp.Body.Close()
			})

		})
	})
})
