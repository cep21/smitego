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
func (s *Session) GetDemoDetails(ctx context.Context, matchID int) ([]OldMatchDetails, error) {
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
func (s *Session) GetFriends(ctx context.Context, player string) ([]Player, error) {
	var r []Player
	if err := s.parent.doReqURL(ctx, fmt.Sprintf("%s/%s", s.urlSession("getfriends"), player), &r); err != nil {
		return nil, err
	}
	return r, nil
}

// GetGodRecommendedItems Returns the Recommended Items for a particular God. .
func (s *Session) GetGodRecommendedItems(ctx context.Context, godID int, lang LanguageCode) ([]RecommendedItem, error) {
	var r []RecommendedItem
	if err := s.parent.doReqURL(ctx, fmt.Sprintf("%s/%d/%d", s.urlSession("getgodrecommendeditems"), godID, lang), &r); err != nil {
		return nil, err
	}
	return r, nil
}

// GetItems returns all Items and their various attributes.
func (s *Session) GetItems(ctx context.Context, lang LanguageCode) ([]Item, error) {
	var r []Item
	if err := s.parent.doReqURL(ctx, fmt.Sprintf("%s/%d", s.urlSession("getitems"), lang), &r); err != nil {
		return nil, err
	}
	return r, nil
}

// GetMatchDetails returns the statistics for a particular completed match.
func (s *Session) GetMatchDetails(ctx context.Context, matchID int) ([]MatchPlayerInfo, error) {
	var r []MatchPlayerInfo
	if err := s.parent.doReqURL(ctx, fmt.Sprintf("%s/%d", s.urlSession("getmatchdetails"), matchID), &r); err != nil {
		return nil, err
	}
	return r, nil
}

// GetMatchPlayerDetails returns player information for a live match.
func (s *Session) GetMatchPlayerDetails(ctx context.Context, matchID int) ([]MatchPlayerDetails, error) {
	var r []MatchPlayerDetails
	if err := s.parent.doReqURL(ctx, fmt.Sprintf("%s/%d", s.urlSession("getmatchplayerdetails"), matchID), &r); err != nil {
		return nil, err
	}
	return r, nil
}

// GetMatchidsByQueue lists all Match IDs for a particular Match Queue; useful for API developers
// interested in constructing data by Queue.  To limit the data returned, an {hour} parameter was
// added (valid values: 0 - 23).  An {hour} parameter of -1 represents the entire day, but be
// warned that this may be more data than we can return for certain queues.  Also, a returned
// “active_flag” means that there is no match information/stats for the corresponding match.
// Usually due to a match being in-progress, though there could be other reasons..
func (s *Session) GetMatchidsByQueue(ctx context.Context, queue Queue, year int, month int, day int, hour int) ([]MatchQueueID, error) {
	var r []MatchQueueID
	dateFmt := fmt.Sprintf("%04d%02d%02d", year, month, day)
	if err := s.parent.doReqURL(ctx, fmt.Sprintf("%s/%d/%s/%d", s.urlSession("getmatchidsbyqueue"), queue, dateFmt, hour), &r); err != nil {
		return nil, err
	}
	return r, nil
}

// GetLeagueLeaderboard returns the top players for a particular league
// (as indicated by the queue/tier/season parameters).
func (s *Session) GetLeagueLeaderboard(ctx context.Context, queue Queue, tier Tier, season int) ([]LeaderboardPlayer, error) {
	var r []LeaderboardPlayer
	if err := s.parent.doReqURL(ctx, fmt.Sprintf("%s/%d/%d/%d", s.urlSession("getleagueleaderboard"), queue, tier, season), &r); err != nil {
		return nil, err
	}
	return r, nil
}

// GetLeagueSeasons provides a list of seasons (including the single active season) for a match queue.
func (s *Session) GetLeagueSeasons(ctx context.Context, queue Queue) ([]LeagueSeason, error) {
	var r []LeagueSeason
	if err := s.parent.doReqURL(ctx, fmt.Sprintf("%s/%d", s.urlSession("getleagueseasons"), queue), &r); err != nil {
		return nil, err
	}
	return r, nil
}

// GetMatchHistory gets recent matches and high level match statistics for a particular player
func (s *Session) GetMatchHistory(ctx context.Context, player string) ([]PlayerMatchHistory, error) {
	var r []PlayerMatchHistory
	if err := s.parent.doReqURL(ctx, fmt.Sprintf("%s/%s", s.urlSession("getmatchhistory"), player), &r); err != nil {
		return nil, err
	}
	return r, nil
}
