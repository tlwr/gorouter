package integration

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Backend Pruning", func() {
	var testState *testState

	BeforeEach(func() {
		testState = NewTestState()
	})

	AfterEach(func() {
		if testState != nil {
			testState.StopAndCleanup()
		}
	})

	Context("when apps fail for a short amount of time", func() {
		It("handles requests", func() {
			hostname := "some.basic.app"
			testState.StartGorouter()
			appInstanceOne := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
			appInstanceTwo := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

			appInstanceOne.Listener.Close()
			appInstanceTwo.Listener.Close()

			testState.registerTLS(appInstanceOne, hostname)
			testState.registerTLS(appInstanceTwo, hostname)

			Consistently(func() (int, error) {
				req := testState.newRequest(fmt.Sprintf("https://%s", hostname))
				resp, err := testState.client.Do(req)

				return resp.StatusCode, err
			}).Should(Equal(http.StatusBadGateway))
		})
	})
})
