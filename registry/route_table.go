package registry
import "code.cloudfoundry.org/gorouter/route"

type RouteTable interface {
	MatchingEndpoints(route route.Uri) (*route.EndpointPool, error)
	Insert(route route.Uri, endpoint *route.Endpoint) error
	Delete(route route.Uri, endpoint *route.Endpoint) error
}
