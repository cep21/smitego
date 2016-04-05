// +build !go1.5

package smitego

import (
	"golang.org/x/net/context"
	"net/http"
)

func withCancel(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error) {
	type canceler interface {
		CancelRequest(*http.Request)
	}
	canCancel, ok := client.Transport.(canceler)
	if !ok {
		return client.Do(req)
	}

	type doResponse struct {
		resp *http.Response
		err error
	}

	c := make(chan doResponse, 1)
	go func() {
		resp, err := client.Do(req)
		c <- doResponse {
			resp: resp,
			err: err,
		}
	}()
	select {
	case <-ctx.Done():
		canCancel.CancelRequest(req)
		<-c // Wait for f to return.
		return nil, ctx.Err()
	case r := <-c:
		return r.resp, r.err
	}
}
