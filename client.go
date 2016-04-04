package smitego

import (
	"net/http"
	"fmt"
	"time"
	"golang.org/x/net/context"
	"hash"
	"encoding/hex"
	"encoding/json"
	"errors"
	"bytes"
	"io"
)

// Why the frick is this http and not HTTPS.  (?????)
const DefaultBaseURL = "http://api.smitegame.com/smiteapi.svc"

type Client struct {
	BaseURL string
	DevID int64
	AuthKey string
	CurTime func() time.Time
	HTTPClient http.Client
	HashConstructor func() hash.Hash
	ErrCallback func(err error)
}

var ErrResponseNotExpectedJSON = errors.New("response not expected JSON")

func (c *Client) Ping(ctx context.Context) error {
	var m string
	if err := c.doReqUrl(ctx, c.urlBase("ping"), &m); err != nil {
		return err
	}
	return nil
}

func (c *Client) doReqUrl(ctx context.Context, u string, jsonInto interface{}) error {
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

func (c *Client) CreateSession(ctx context.Context) (*Session, error) {
	var v createSessionResp
	if err := c.doReqUrl(ctx, c.URL("createsession", ""), &v); err != nil {
		return nil, err
	}
	return &Session {
		parent: c,
		SessionID: v.SessionID,
	}, nil
}

func (c *Client) urlBase(endpoint string) string {
	return fmt.Sprintf("%s/%sjson", c.BaseURL, endpoint)
}

func (c *Client) URL(endpoint string, session string) string {
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
	ret :=  fmt.Sprintf("%s/%d/%s/%s%s", c.urlBase(endpoint), c.DevID, signature, session, timeFmt)
	return ret
}

func mustNotErr(err error) {
	if err != nil {
		panic("Unexpected error: " + err.Error())
	}
}
