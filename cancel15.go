// +build go1.5

package smitego

import (
	"golang.org/x/net/context"
	"net/http"
)

func withCancel(ctx context.Context, client *http.Client, req *http.Request) (resp *http.Response, err error) {
	req.Cancel = ctx.Done()
	return client.Do(req)
}
