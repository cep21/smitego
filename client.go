package smitego

import (
	"net/http"
	"fmt"
	"time"
	"golang.org/x/net/context"
	"bytes"
	"io"
)

type Client struct {
	BaseURL string
	DevID int64
	AuthKey string
	CurTime func() time.Time
	HTTPClient http.Client
}

func (c *Client) CreateSession(ctx context.Context) (*Session, error) {
	u := c.URL("createsession")
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	req.Cancel = ctx.Done()
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	b := bytes.Buffer{}
	io.Copy(b, resp.Body)
	resp.Body.Close()
	fmt.Println(b.String())
	return nil, nil
}

func (c *Client) URL(endpoint string) string {
	curTime := c.CurTime()
	year, month, day := curTime.Date()
	h := curTime.Hour()
	m := curTime.Minute()
	s := curTime.Second()
	timeFmt := fmt.Sprintf("%d%d%d%d%d%d", year, month, day, h, m, s)
	return fmt.Sprintf("%s/%sJson/%d/%s/%d", c.BaseURL, endpoint, c.DevID, c.AuthKey, timeFmt)
}
