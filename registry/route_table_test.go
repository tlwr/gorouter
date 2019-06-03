package registry_test

import (
	"code.cloudfoundry.org/gorouter/registry"
	"code.cloudfoundry.org/gorouter/route"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/deckarep/golang-set"
)

var _ = Describe("route_table", func() {
	var routeTable registry.RouteTable

	BeforeEach(func() {
		//route_table = new implementation
	})

	Describe("MatchingEndpoints and Insert", func() {
		const someUrl = "some.url"
		Context("When there are no elements in the table", func() {
			It("should return an empty endpoint pool", func() {
				pool, err := routeTable.MatchingEndpoints(someUrl)
				Expect(err).NotTo(HaveOccurred())

				Expect(pool.IsEmpty())//.To(BeTrue())
			})

			It("should return an error when queried with an invalid url", func() {
				_, err := routeTable.MatchingEndpoints("so*e.url")
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError("Invalid uri"))
			})
		})

		Context("when we have a route with an endpoint in the table", func() {
			var endpoint *route.Endpoint

			BeforeEach(func() {
				endpoint = route.NewEndpoint(&route.EndpointOpts{
					Host: "192.168.1.1",
				})
				err := routeTable.Insert(someUrl, endpoint)
				Expect(err).NotTo(HaveOccurred())
			})

			It("returns the endpoint when queried with the exact url", func() {
				pool, err := routeTable.MatchingEndpoints(someUrl)

				Expect(err).NotTo(HaveOccurred())

				routeEndpoints()

				Expect(endpoints).To(HaveLen(1))
				Expect(endpoints[0]).To(Equal(endpoint))
			})

			It("returns the endpoint when queried with a wildcard url", func() {
				pool, err := routeTable.MatchingEndpoints("*.url")

				Expect(err).NotTo(HaveOccurred())

				var endpoints []*route.Endpoint
				pool.Each(func(endpoint *route.Endpoint) {
					endpoints = append(endpoints, endpoint)
				})

				Expect(endpoints).To(HaveLen(1))
				Expect(endpoints[0]).To(Equal(endpoint))
			})

			It("returns an empty pool when queried with a mismatched url", func() {
				pool, err := routeTable.MatchingEndpoints("other.url")

				Expect(err).NotTo(HaveOccurred())
				Expect(pool.IsEmpty()).To(BeTrue())
			})

			It("returns an empty pool when queried with a mismatched wildcard url", func() {
				pool, err := routeTable.MatchingEndpoints("*.some.url")

				Expect(err).NotTo(HaveOccurred())
				Expect(pool.IsEmpty()).To(BeTrue())
			})

			Context("when we have another endpoint for the same route", func() {
				var otherEndpoint *route.Endpoint

				BeforeEach(func() {
					otherEndpoint = route.NewEndpoint(&route.EndpointOpts{
						Host: "192.168.1.2",
					})
					err := routeTable.Insert(someUrl, otherEndpoint)
					Expect(err).NotTo(HaveOccurred())
				})

				It("returns both endpoints", func() {
					pool, err := routeTable.MatchingEndpoints(someUrl)
					Expect(err).NotTo(HaveOccurred())

				})
			})

			Context("when we have another endpoint for a different route", func() {
				Context("when the 2 routes are independent", func() {

				})

				Context("when a route matches the other's wildcard", func() {

				})
			})
		})

		Context("when there are lots of routes in the table", func() {

		})
	})

	Describe("Delete", func() {
		Context("when attempting to delete a route that exists", func() {
			It("deletes that route", func() {

			})

			It("doesn't delete other routes", func() {

			})
		})

		Context("when attempting to delete a route that does not exist", func() {

		})
	})
})

type EndpointSet struct {
	endpoints map[*route.Endpoint]bool
}

func NewEndpointSet(pool route.EndpointPool) *EndpointSet {
	var endpoints = make(map[*route.Endpoint]bool)
	pool.Each(func(endpoint *route.Endpoint) {
		endpoints[endpoint] = true
	})

	return &EndpointSet{
		endpoints: endpoints,
	}
}

func (e *EndpointSet) Contains(endpoint *route.Endpoint) bool {
	_, ok := e.endpoints[endpoint]
	return ok
}

func routeEndpoints(pool route.EndpointPool) map[*route.Endpoint]bool {



	return endpoints
}
