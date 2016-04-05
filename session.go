package smitego

import (
	"errors"
	"fmt"
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

// GetGods returns all Gods and their various attributes
func (s *Session) GetGods(ctx context.Context, lang LanguageCode) ([]God, error) {
	var r []God
	if err := s.parent.doReqURL(ctx, fmt.Sprintf("%s/%d", s.urlSession("getgods"), lang), &r); err != nil {
		return nil, err
	}
	return r, nil
}

// GetDemoDetails returns information regarding a particular match.
// Rarely used in lieu of getmatchdetails().
func (s *Session) GetDemoDetails(ctx context.Context, matchID int) ([] OldMatchDetails, error) {
	var r []OldMatchDetails
	if err := s.parent.doReqURL(ctx, fmt.Sprintf("%s/%d", s.urlSession("getdemodetails"), matchID), &r); err != nil {
		return nil, err
	}
	return r, nil
}

// GetEsportsproleaguedetails returns the matchup information for each matchup for the current
// eSports Pro League season.  An important return value is “match_status” which represents a
// match being scheduled (1), in-progress (2), or complete (3)
func (s *Session) GetEsportsproleaguedetails(ctx context.Context) ([]Match, error) {
	var r []Match
	if err := s.parent.doReqURL(ctx, s.urlSession("getesportsproleaguedetails"), &r); err != nil {
		return nil, err
	}
	return r, nil
}

// GetFriends returns the Smite User names of each of the player’s friends.
func (s *Session) GetFriends(ctx context.Context, playerID int) ([]Match, error) {
	var r []Match
	if err := s.parent.doReqURL(ctx, s.urlSession("getfriends"), &r); err != nil {
		return nil, err
	}
	return r, nil
}
