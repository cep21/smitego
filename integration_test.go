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
const debugPlayerId = "cep21"
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
		friends, err := session.GetFriends(context.Background(), debugPlayerId)
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
		p, err := session.GetMatchHistory(context.Background(), debugPlayerId)
		So(err, ShouldBeNil)
		So(len(p), ShouldBeGreaterThan, 1)
	})
}
