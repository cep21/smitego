package smitego

import (
	"fmt"
	"golang.org/x/net/context"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"strings"
	"io/ioutil"
)

type errorRT struct {
}

func (t *errorRT) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, t
}

func (t *errorRT) Error() string {
	return "An error in transport"
}

type staticResp struct {
	Resp *http.Response
}

func (t *staticResp) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.Resp, nil
}

// Example of creating a session and making a function call.
func ExampleSession() {
	// First make a client to describe how you want to connect.  Each client returns a session
	// and primary function calls are done on the session.  Concurrent sessions are limited
	// by HiRez
	client := Client{
		DevID:   123,
		AuthKey: "AuthKey123",
	}

	// A context is how you can time out function calls
	ctx := context.Background()

	// Some functions don't require a session first and can be called on the client directly
	_ = client.Ping(ctx)

	// Most functions require a session

	session, _ := client.CreateSession(ctx)
	gods, _ := session.GetGods(ctx, English)
	fmt.Printf("Got %d gods\n", len(gods))
}

func TestClient(t *testing.T) {
	Convey("With a client", t, func() {
		c := Client {
		}
		ctx := context.Background()
		Convey("that errors http connections", func() {
			rt := &errorRT{}
			c.HTTPClient.Transport = rt
			Convey("Connections should error", func() {
				err := c.Ping(ctx)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, rt.Error())
			})
		})
		Convey("that 404s", func() {
			sr := &staticResp{
				Resp: &http.Response{
					StatusCode: http.StatusNotFound,
					Body: ioutil.NopCloser(strings.NewReader("")),
				},
			}
			c.HTTPClient.Transport = sr
			Convey("requests should error", func() {
				err := c.Ping(ctx)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, fmt.Sprintf("%d", http.StatusNotFound))
			})
		})

	})
}
