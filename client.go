package smitego

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"hash"
	"io"
	"net/http"
	"time"
)

// DefaultBaseURL is where smite expects API calls.  Why the frick is this HTTP and not HTTPS.  (?????)
const DefaultBaseURL = "http://api.smitegame.com/smiteapi.svc"

// Client can create smite session objects and interact with the smite API
type Client struct {
	BaseURL         string
	DevID           int64
	AuthKey         string
	CurTime         func() time.Time
	HTTPClient      http.Client
	HashConstructor func() hash.Hash
	ErrCallback     func(err error)
}

// ErrResponseNotExpectedJSON is returned by API calls when the response isn't expected JSON
var ErrResponseNotExpectedJSON = errors.New("response not expected JSON")

// Ping is a quick way of validating access to the Hi-Rez API
func (c *Client) Ping(ctx context.Context) error {
	var m string
	if err := c.doReqURL(ctx, c.urlBase("ping"), &m); err != nil {
		return err
	}
	return nil
}

func (c *Client) doReqURL(ctx context.Context, u string, jsonInto interface{}) error {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return err
	}
	req.Cancel = ctx.Done()
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	var b bytes.Buffer
	if _, err := io.Copy(&b, resp.Body); err != nil {
		return err
	}
	if err := json.NewDecoder(&b).Decode(jsonInto); err != nil {
		return ErrResponseNotExpectedJSON
	}
	if err := resp.Body.Close(); err != nil {
		return err
	}
	return nil
}

// CreateSession is A required step to Authenticate the developerId/signature for further API use.
func (c *Client) CreateSession(ctx context.Context) (*Session, error) {
	var v createSessionResp
	if err := c.doReqURL(ctx, c.url("createsession", ""), &v); err != nil {
		return nil, err
	}
	return &Session{
		parent:    c,
		SessionID: v.SessionID,
	}, nil
}

func (c *Client) urlBase(endpoint string) string {
	return fmt.Sprintf("%s/%sjson", c.BaseURL, endpoint)
}

func (c *Client) url(endpoint string, session string) string {
	timeFmt := c.CurTime().UTC().Format("20060102150405")
	hasher := c.HashConstructor()
	sig := fmt.Sprintf("%d%s%s%s", c.DevID, endpoint, c.AuthKey, timeFmt)
	_, err := hasher.Write([]byte(sig))
	mustNotErr(err)
	signatureBytes := hasher.Sum(nil)
	signature := hex.EncodeToString(signatureBytes)
	if session != "" {
		session = session + "/"
	}
	ret := fmt.Sprintf("%s/%d/%s/%s%s", c.urlBase(endpoint), c.DevID, signature, session, timeFmt)
	return ret
}

func mustNotErr(err error) {
	if err != nil {
		panic("Unexpected error: " + err.Error())
	}
}
