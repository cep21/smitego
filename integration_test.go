// +build integration

package smitego

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/net/context"
	"os"
	"testing"
	"time"
)

// Create a file named info.json and put it at the root of your project.  That file should have your
// devId and authKey.  The file should be inside .gitignore and not checked into git.  Then,
// run `go test -v --tags=integration .` to start integration tests using your auth key.
type devInfo struct {
	DevID   int64  `json:"devId"`
	AuthKey string `json:"authKey"`
}

func mustLoad(filename string) devInfo {
	f, err := os.Open(filename)
	if err != nil {
		panic("Unable to open " + filename + ": " + err.Error())
	}
	defer f.Close()
	var v devInfo
	if err := json.NewDecoder(f).Decode(&v); err != nil {
		panic("Unable to decode file: " + err.Error())
	}
	return v
}

var client Client
var session *Session

const debugMatchId = 237403351
const debugPlayerName = "cep21"
const debugPlayerId = 8718637
const debugTeamId = 497261
const debugGodId = 1737

func init() {
	// Normally each run would make its own client and session, but I'm using a single instance
	// for all integration tests because HiRez throttles client sessions.
	d := mustLoad("info.json")
	client = Client{
		DevID:   d.DevID,
		AuthKey: d.AuthKey,
	}
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	session, err = client.CreateSession(ctx)
	mustNotErr(err)
}

func TestPing(t *testing.T) {
	Convey("Ping should work", t, func() {
		So(client.Ping(context.Background()), ShouldBeNil)
	})
}

func TestDataUsed(t *testing.T) {
	Convey("DataUsed should work", t, func() {
		usage, err := session.GetDataUsed(context.Background())
		So(err, ShouldBeNil)
		t.Log(usage.String())
	})
}

func TestGetGods(t *testing.T) {
	Convey("GetGods should work", t, func() {
		gods, err := session.GetGods(context.Background(), English)
		So(err, ShouldBeNil)
		So(len(gods), ShouldBeGreaterThan, 10)
		for _, g := range gods {
			t.Log(g.String())
		}
	})
}

func TestGetEsportsproleaguedetails(t *testing.T) {
	Convey("GetEsportsproleaguedetails should work", t, func() {
		matches, err := session.GetEsportsproleaguedetails(context.Background())
		So(err, ShouldBeNil)
		So(len(matches), ShouldBeGreaterThan, 1)
		for _, m := range matches {
			t.Log(m)
			t.Log(m.String())
		}
	})
}

func TestGetDemoDetails(t *testing.T) {
	Convey("GetDemoDetails should work", t, func() {
		dets, err := session.GetDemoDetails(context.Background(), debugMatchId)
		So(err, ShouldBeNil)
		So(len(dets), ShouldEqual, 1)
	})
}

func TestGetFriends(t *testing.T) {
	Convey("GetFriends should work", t, func() {
		friends, err := session.GetFriends(context.Background(), debugPlayerName)
		So(err, ShouldBeNil)
		So(len(friends), ShouldBeGreaterThan, 1)
	})
}

func TestGetGodRecommendedItems(t *testing.T) {
	Convey("getgodrecommendeditems should work", t, func() {
		items, err := session.GetGodRecommendedItems(context.Background(), debugGodId, English)
		So(err, ShouldBeNil)
		So(len(items), ShouldBeGreaterThan, 12)
	})
}

func TestGetItems(t *testing.T) {
	Convey("GetItems should work", t, func() {
		items, err := session.GetItems(context.Background(), English)
		So(err, ShouldBeNil)
		So(len(items), ShouldBeGreaterThan, 12)
		for _, i := range items {
			t.Log(i.String())
		}
	})
}

func TestGetMatchDetails(t *testing.T) {
	Convey("GetMatchDetails should work", t, func() {
		dets, err := session.GetMatchDetails(context.Background(), debugMatchId)
		So(err, ShouldBeNil)
		So(len(dets), ShouldBeGreaterThan, 9)
		t.Log(dets)
	})
}

func TestGetMatchPlayerDetails(t *testing.T) {
	Convey("GetMatchPlayerDetails should work", t, func() {
		dets, err := session.GetMatchPlayerDetails(context.Background(), debugMatchId)
		So(err, ShouldBeNil)
		So(len(dets), ShouldEqual, 1)
	})
}

func TestGetMatchidsByQueue(t *testing.T) {
	Convey("GetMatchidsByQueue should work", t, func() {
		client.VerboseLog = t.Log
		dets, err := session.GetMatchidsByQueue(context.Background(), Joust3v3, 2016, 4, 5, 1)
		So(err, ShouldBeNil)
		So(len(dets), ShouldEqual, 1268)
	})
}

func TestGetDaysMatchIds(t *testing.T) {
	Convey("GetMatchidsByQueue should work", t, func() {
		client.VerboseLog = t.Log
		dets, err := session.GetMatchidsByQueue(context.Background(), Joust3v3, 2016, 4, 5, 1)
		So(err, ShouldBeNil)
		So(len(dets), ShouldEqual, 1268)
	})
}

func TestGetTodaysMatchIds(t *testing.T) {
	Convey("Test fetching today's matches should work", t, func() {
		now := time.Now().UTC()
		y, m, d := now.Date()
		for h := 0; h < 24; h++ {
			dets, err := session.GetMatchidsByQueue(context.Background(), Joust3v3, y, int(m), d, h)
			So(err, ShouldBeNil)
			t.Logf("Hour = %d, len(items) = %d", h, len(dets))
			if len(dets) > 0 {
				t.Logf("A single match: %s", dets[0])
			}
		}
	})
}

func TestGetLeagueLeaderboard(t *testing.T) {
	Convey("GetLeagueLeaderboard should work", t, func() {
		client.VerboseLog = t.Log
		p, err := session.GetLeagueLeaderboard(context.Background(), JoustRanked3v3, SilverI, 5)
		So(err, ShouldBeNil)
		So(len(p), ShouldEqual, 500)
	})
}

func TestGetLeagueSeasons(t *testing.T) {
	Convey("GetLeagueSeasons should work", t, func() {
		client.VerboseLog = t.Log
		p, err := session.GetLeagueSeasons(context.Background(), ConquestRanked2)
		So(err, ShouldBeNil)
		So(len(p), ShouldBeGreaterThan, 2)
	})
}

func TestGetMatchHistory(t *testing.T) {
	Convey("GetMatchHistory should work", t, func() {
		client.VerboseLog = t.Log
		p, err := session.GetMatchHistory(context.Background(), debugPlayerName)
		So(err, ShouldBeNil)
		So(len(p), ShouldBeGreaterThan, 1)
	})
}

func TestGetMotd(t *testing.T) {
	Convey("GetMotd should work", t, func() {
		client.VerboseLog = t.Log
		p, err := session.GetMotd(context.Background())
		So(err, ShouldBeNil)
		So(len(p), ShouldBeGreaterThan, 1)
		t.Log(p[0])
	})
}

func TestGetPlayer(t *testing.T) {
	Convey("GetPlayer should work", t, func() {
		client.VerboseLog = t.Log
		p, err := session.GetPlayer(context.Background(), debugPlayerName)
		So(err, ShouldBeNil)
		So(len(p), ShouldEqual, 1)
		t.Log(p[0])
	})
}

func TestGetPlayerStatus(t *testing.T) {
	Convey("GetPlayerStatus should work", t, func() {
		client.VerboseLog = t.Log
		p, err := session.GetPlayerStatus(context.Background(), debugPlayerName)
		So(err, ShouldBeNil)
		So(len(p), ShouldEqual, 1)
		t.Log(p[0])
	})
}

func TestGetQueueStats(t *testing.T) {
	Convey("GetQueueStats should work", t, func() {
		client.VerboseLog = t.Log
		p, err := session.GetQueueStats(context.Background(), debugPlayerName, JoustRanked3v3)
		So(err, ShouldBeNil)
		So(len(p), ShouldBeGreaterThanOrEqualTo, 0)
	})
}

func TestGetTeamDetails(t *testing.T) {
	Convey("GetTeamDetails should work", t, func() {
		client.VerboseLog = t.Log
		p, err := session.GetTeamDetails(context.Background(), debugTeamId)
		So(err, ShouldBeNil)
		So(len(p), ShouldEqual, 1)
	})
}

func TestGetTeamPlayers(t *testing.T) {
	Convey("GetTeamPlayers should work", t, func() {
		client.VerboseLog = t.Log
		p, err := session.GetTeamPlayers(context.Background(), debugTeamId)
		So(err, ShouldBeNil)
		t.Log(p)
		So(len(p), ShouldBeGreaterThan, 1)
	})
}

func TestGetTopMatches(t *testing.T) {
	Convey("GetTopMatches should work", t, func() {
		client.VerboseLog = t.Log
		p, err := session.GetTopMatches(context.Background())
		So(err, ShouldBeNil)
		t.Log(p)
		So(len(p), ShouldBeGreaterThan, 1)
	})
}

func TestSearchTeams(t *testing.T) {
	Convey("SearchTeams should work", t, func() {
		client.VerboseLog = t.Log
		p, err := session.SearchTeams(context.Background(), "the")
		So(err, ShouldBeNil)
		t.Log(p)
		So(len(p), ShouldBeGreaterThan, 1)
	})
}

func TestGetPlayerAchievements(t *testing.T) {
	Convey("GetPlayerAchievements should work", t, func() {
		client.VerboseLog = t.Log
		p, err := session.GetPlayerAchievements(context.Background(), debugPlayerId)
		So(err, ShouldBeNil)
		t.Log(p)
		So(p.AssistedKills, ShouldBeGreaterThan, 1)
	})
}
