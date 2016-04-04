package smitego

import (
	"golang.org/x/net/context"
	"errors"
)

type Session struct{
	parent *Client
	SessionID string
}

func (s *Session) urlSession(endpoint string) string {
	u := s.parent.URL(endpoint, s.SessionID)
	return u
}

func (s *Session) TestSession(ctx context.Context) (string, error) {
	var m string
	if err := s.parent.doReqUrl(ctx, s.urlSession("testsession"), &m); err != nil {
		return "", err
	}
	return m, nil
}

var ErrUnexpectedSize = errors.New("Unexpected data usage size")

func (s *Session) GetDataUsed(ctx context.Context) (*DataUsed, error) {
	r := make([]DataUsed, 1)
	if err := s.parent.doReqUrl(ctx, s.urlSession("getdataused"), &r); err != nil {
		return nil, err
	}
	if len(r) != 1 {
		return nil, ErrUnexpectedSize
	}
	return &r[0], nil
}
