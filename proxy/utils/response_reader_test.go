package utils_test

import (
	"bufio"
	"bytes"
	"io"
	"runtime"
	"time"

	"code.cloudfoundry.org/gorouter/proxy/utils"
	"code.cloudfoundry.org/gorouter/test_util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = FDescribe("ResponseReader", func() {
	Describe("ReadResponseWithTimeout", func() {
		var (
			timeout time.Duration
			reader  *bufio.Reader
		)

		BeforeEach(func() {
			timeout = 50 * time.Millisecond
			reader = bufio.NewReader(io.MultiReader(bytes.NewBufferString("HTTP/1.1 200\r\n\r\n"), nil))
		})

		It("reads the response before the timeout", func() {
			resp, err := utils.ReadResponseWithTimeout(reader, nil, timeout)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
		})

		It("returns an error when response is invalid", func() {
			badReader := bufio.NewReader(io.MultiReader(bytes.NewBufferString("Invalid HTTP\r\n\r\n"), nil))
			resp, err := utils.ReadResponseWithTimeout(badReader, nil, timeout)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("malformed HTTP"))
			Expect(resp).To(BeNil())
		})

		Context("when read response times out", func() {
			var (
				slowReader    *bufio.Reader
				sleepDuration time.Duration
			)

			BeforeEach(func() {
				sleepDuration = 100 * time.Millisecond
				slowReader = bufio.NewReader(&test_util.SlowReadCloser{SleepDuration: sleepDuration})
			})

			It("returns an error", func() {
				resp, err := utils.ReadResponseWithTimeout(slowReader, nil, timeout)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("timeout"))
				Expect(resp).To(BeNil())
			})

			It("doesn't leak goroutines", func() {
				beforeGoroutineCount := runtime.NumGoroutine()
				utils.ReadResponseWithTimeout(slowReader, nil, timeout)

				Eventually(func() int {
					return runtime.NumGoroutine()
				}).Should(BeNumerically("<=", beforeGoroutineCount))
			})
		})
	})
})
