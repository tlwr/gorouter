
import "code.cloudfoundry.org/gorouter/route"

type RouteTable interface {
	Find(uri route.Uri) *route.Pool
	Register(uri, endpoint)
	Insert(uri route.Uri, endpoint *route.Endpoint)
	Delete(uri route.Uri, endpoint *route.Endpoint)
}
