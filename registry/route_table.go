
import "code.cloudfoundry.org/gorouter/route"

type RouteTable interface {
	Find(uri route.Uri) *route.Pool                 // lookup
	Insert(uri route.Uri, endpoint *route.Endpoint) // register
	Delete(uri route.Uri, endpoint *route.Endpoint) // deregister
}
