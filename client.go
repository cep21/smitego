package smitego

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"runtime"
	"time"
)

// DefaultBaseURL is where smite expects API calls.  Why the frick is this HTTP and not HTTPS.  (?????)
const DefaultBaseURL = "http://api.smitegame.com/smiteapi.svc"

// DefaultXboxURL is used for smite's XBox API
const DefaultXboxURL = "http://api.xbox.smitegame.com/smiteapi.svc"

// Version is the current library's version: sent with User-Agent
const Version = "0.1"

// Client can create smite session objects and interact with the smite API
type Client struct {
	DevID      int64
	AuthKey    string
	HTTPClient http.Client
	BaseURL    string
	CurTime    func() time.Time
	VerboseLog Log
}

// Log is a function that Client can take to optionally verbose log what it does internally
type Log func(...interface{})

// ErrNotExpectedJSON is returned by API calls when the response isn't expected JSON
type ErrNotExpectedJSON struct {
	OriginalBody string
	Err          error
}

// ErrBadStatusCode is returned when the API returns a non 200 error code
type ErrBadStatusCode struct {
	OriginalBody string
	Code         int
}

func (e *ErrBadStatusCode) Error() string {
	return fmt.Sprintf("Invalid status code: %d", e.Code)
}

func (c *Client) verboseLog(v ...interface{}) {
	if c.VerboseLog != nil {
		c.VerboseLog(v...)
	}
}

func (e *ErrNotExpectedJSON) Error() string {
	return fmt.Sprintf("Unexpected JSON: %s from %s", e.Err.Error(), e.OriginalBody)
}

func (c *Client) clientTime() time.Time {
	if c.CurTime == nil {
		return time.Now()
	}
	return c.CurTime()
}

// Ping is a quick way of validating access to the Hi-Rez API
func (c *Client) Ping(ctx context.Context) error {
	var m string
	if err := c.doReqURL(ctx, c.urlBase("ping"), &m); err != nil {
		return err
	}
	return nil
}

func (c *Client) doReqURL(ctx context.Context, u string, jsonInto interface{}) error {
	c.verboseLog("fetching", u)
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", fmt.Sprintf("github.com/cep21/smitego/%s (gover %s)", Version, runtime.Version()))
	resp, err := withCancel(ctx, &c.HTTPClient, req)
	if err != nil {
		return err
	}
	var b bytes.Buffer
	if _, err := io.Copy(&b, resp.Body); err != nil {
		return err
	}
	debug := b.String()
	if resp.StatusCode != http.StatusOK {
		c.verboseLog("Invalid status code", resp.StatusCode)
		return &ErrBadStatusCode{
			OriginalBody: debug,
			Code:         resp.StatusCode,
		}
	}
	c.verboseLog("Fetch result", debug)
	if err := json.NewDecoder(&b).Decode(jsonInto); err != nil {
		return &ErrNotExpectedJSON{
			OriginalBody: debug,
			Err:          err,
		}
	}
	if err := resp.Body.Close(); err != nil {
		return err
	}
	return nil
}

// CreateSession is a required step to Authenticate the developerId/signature for further API use.
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
	base := c.BaseURL
	if c.BaseURL == "" {
		base = DefaultBaseURL
	}
	return fmt.Sprintf("%s/%sjson", base, endpoint)
}

func (c *Client) url(endpoint string, session string) string {
	timeFmt := c.clientTime().UTC().Format("20060102150405")
	hasher := md5.New()
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
