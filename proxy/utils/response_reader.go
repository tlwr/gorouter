package utils

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"time"
)

type HTTPResponseAndError struct {
	response *http.Response
	error    error
}

type TimeoutError struct{}

func (t TimeoutError) Error() string {
	return fmt.Sprintf("timeout waiting for http response from backend")
}

func ReadResponseWithTimeout(r *bufio.Reader, req *http.Request, timeout time.Duration) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	read := make(chan HTTPResponseAndError)
	defer close(read)

	go readResponseUnlessContextCanceled(ctx, read, r, req)

	select {
	case s := <-read:
		return s.response, s.error
	case <-ctx.Done():
		return nil, TimeoutError{}
	}
}

func readResponseUnlessContextCanceled(ctx context.Context, c chan<- HTTPResponseAndError, r *bufio.Reader, req *http.Request) {
	resp, err := http.ReadResponse(r, req)

	select {
	case <-ctx.Done():
	default:
		c <- HTTPResponseAndError{resp, err}
	}
}
