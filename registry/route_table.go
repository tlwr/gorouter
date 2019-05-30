package registry
import "code.cloudfoundry.org/gorouter/route"

type RouteTable interface {
	MatchingEndpoints(uri route.Uri) (*route.EndpointPool, error)
	Insert(uri route.Uri, endpoint *route.Endpoint) error
	Delete(uri route.Uri, endpoint *route.Endpoint) error
}
