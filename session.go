package smitego

import (
	"errors"
	"golang.org/x/net/context"
)

// Session is an authenticated smite token that can take future actions
type Session struct {
	parent    *Client
	SessionID string
}

func (s *Session) urlSession(endpoint string) string {
	u := s.parent.url(endpoint, s.SessionID)
	return u
}

// TestSession is a means of validating that a session is established
func (s *Session) TestSession(ctx context.Context) (string, error) {
	var m string
	if err := s.parent.doReqURL(ctx, s.urlSession("testsession"), &m); err != nil {
		return "", err
	}
	return m, nil
}

// ErrUnexpectedSize is returned by GetDataUsed when the response doesn't look correct
var ErrUnexpectedSize = errors.New("unexpected data usage size")

// GetDataUsed returns API Developer daily usage limits and the current status against those limits
func (s *Session) GetDataUsed(ctx context.Context) (*DataUsed, error) {
	r := make([]DataUsed, 1)
	if err := s.parent.doReqURL(ctx, s.urlSession("getdataused"), &r); err != nil {
		return nil, err
	}
	if len(r) != 1 {
		return nil, ErrUnexpectedSize
	}
	return &r[0], nil
}
