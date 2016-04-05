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
