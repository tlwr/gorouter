package registry_test

import (
	"code.cloudfoundry.org/gorouter/registry"
	"code.cloudfoundry.org/gorouter/route"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("route_table", func() {
	var route_table registry.RouteTable

	BeforeEach(func() {
		//route_table = new implementation
	})

	Describe("MatchingEndpoints", func() {
		Context("When there are no elements in the table", func() {
			It("should return an empty endpoint pool", func() {
				pool, err := route_table.MatchingEndpoints("some.url")
				Expect(err).NotTo(HaveOccurred())

				Expect(pool.IsEmpty())
			})
		})

		Context("when we have a route with an endpoint in the table", func() {

			BeforeEach(func() {
				endpoint := route.NewEndpoint(&route.EndpointOpts{
					Host: "192.168.1.1",
				})
				err := route_table.Insert("some.url", endpoint)
				Expect(err).NotTo(HaveOccurred())

			})

			It("returns the enpoint when queried with the exact url", func() {
				pool, err := route_table.MatchingEndpoints("some.url")

				/*
						Pool1 [10.3.3.3:800]
						Pool2 [10.3.3.4:800]
				*/
				Expect(err).NotTo(HaveOccurred())
				Expect(pool.   Each(func() {} {
					Expect()
				}))

			})
		})
	})
})
