// +build integration

package smitego

import (
	"crypto/md5"
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/net/context"
	"os"
	"testing"
	"time"
)

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

func TestIntegration(t *testing.T) {
	Convey("Loading a client", t, func() {
		d := mustLoad("info.json")
		c := Client{
			BaseURL:         DefaultBaseURL,
			DevID:           d.DevID,
			AuthKey:         d.AuthKey,
			CurTime:         time.Now,
			HashConstructor: md5.New,
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()
		So(c.Ping(ctx), ShouldBeNil)
		session, err := c.CreateSession(ctx)
		So(err, ShouldBeNil)
		So(session, ShouldNotBeNil)
		msg, err := session.TestSession(ctx)
		So(err, ShouldBeNil)
		t.Log(msg)
		usage, err := session.GetDataUsed(ctx)
		So(err, ShouldBeNil)
		t.Log(usage.String())
	})
}
