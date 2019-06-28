package handlers

import (
	"code.cloudfoundry.org/gorouter/registry"
	"code.cloudfoundry.org/gorouter/route"
	"fmt"
	"github.com/cloudfoundry/dropsonde/emitter"
	"net/http"
	"time"

	"code.cloudfoundry.org/gorouter/logger"
	"code.cloudfoundry.org/gorouter/proxy/utils"
	"github.com/cloudfoundry/dropsonde"
	"github.com/cloudfoundry/dropsonde/factories"
	"github.com/cloudfoundry/sonde-go/events"
	"github.com/gogo/protobuf/proto"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/uber-go/zap"
	"github.com/urfave/negroni"
)

type httpStartStopHandler struct {
	emitter dropsonde.EventEmitter
	logger  logger.Logger
	registry registry.Registry
}

// NewHTTPStartStop creates a new handler that handles emitting frontent
// HTTP StartStop events
func NewHTTPStartStop(emitter dropsonde.EventEmitter, logger logger.Logger, registry registry.Registry) negroni.Handler {
	return &httpStartStopHandler{
		emitter: emitter,
		logger:  logger,
		registry: registry,
	}
}

// ServeHTTP handles emitting a StartStop event after the request has been completed
func (hh *httpStartStopHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	prw, ok := rw.(utils.ProxyResponseWriter)
	if !ok {
		hh.logger.Fatal("request-info-err", zap.String("error", "ProxyResponseWriter not found"))
		return
	}
	requestID, err := uuid.ParseHex(r.Header.Get(VcapRequestIdHeader))
	if err != nil {
		hh.logger.Fatal("start-stop-handler-err", zap.String("error", "X-Vcap-Request-Id not found"))
		return
	}

	startTime := time.Now()

	next(rw, r)

	startStopEvent := factories.NewHttpStartStop(r, prw.Status(), int64(prw.Size()), events.PeerType_Server, requestID)
	startStopEvent.StartTimestamp = proto.Int64(startTime.UnixNano())

	// Wrap event in an envelope
	envelope, err := emitter.Wrap(startStopEvent, hh.emitter.Origin())
	if err != nil {
		hh.logger.Info("failed-to-emit-startstop-event", zap.Error(err))
	}

	// Look up route in registry
	uri := route.Uri(hostWithoutPort(r.Host) + r.RequestURI)
	hh.logger.Debug(fmt.Sprintf("Looking up URI: %s", uri.String()))
	endpointPool := hh.registry.Lookup(uri)
	if endpointPool != nil {
		hh.logger.Debug("endpointPool not nil")
		hh.logger.Debug(fmt.Sprintf("endpointPool.isEmpty(): %t", endpointPool.IsEmpty()))
		endpointPool.Each(func(endpoint *route.Endpoint) {
			if endpoint.PrivateInstanceId == startStopEvent.GetInstanceId() {
				// Update envelope to include tags
				envelope.Tags = endpoint.Tags
			}
		})
	}

	err = hh.emitter.EmitEnvelope(envelope)
	if err != nil {
		hh.logger.Info("failed-to-emit-startstop-event", zap.Error(err))
	}
}
