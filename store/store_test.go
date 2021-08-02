package store

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/go-redis/redis/v8"
	redismock "github.com/go-redis/redismock/v8"
)

var (
	mock redismock.ClientMock
)

func init() {
	ctx = context.TODO()
	rdb, mock = redismock.NewClientMock()
}

type suiteT struct {
	nfo string
	opt *redis.Options
	err error
}

func Test_parseCnxNfo(t *testing.T) {
	var casesT [7]suiteT

	casesT[0].nfo = "1:password@host:port"
	casesT[0].opt = &redis.Options{Addr: "host:port", Password: "password", DB: 1}
	casesT[1].nfo = "password@host:port"
	casesT[1].opt = &redis.Options{Addr: "host:port", Password: "password", DB: 0}
	casesT[2].nfo = "host:port"
	casesT[2].opt = &redis.Options{Addr: "host:port", Password: "", DB: 0}
	casesT[3].nfo = ""
	casesT[3].err = errors.New("Bad connection parameter: empty string")
	casesT[4].nfo = "password@host:port@plop"
	casesT[4].err = errors.New("Bad connection parameter: got (password@host:port@plop)")
	casesT[5].nfo = "db:password@host:port"
	casesT[5].err = errors.New("Bad connection parameter: got (db:password@host:port)")
	casesT[6].nfo = "1:password:toto@host:port"
	casesT[6].err = errors.New("Bad connection parameter: got (1:password:toto@host:port)")

	for _, caseT := range casesT {
		opts, err := parseCnxNfo(caseT.nfo)
		if !reflect.DeepEqual(err, caseT.err) {
			t.Fatalf("errors do not match; got(%#v) want (%#v)", err, caseT.err)
		}
		if !reflect.DeepEqual(opts, caseT.opt) {
			t.Fatalf("options do not match; got(%#v) want (%#v)", opts, caseT.opt)
		}
	}
}

func Test_redis(t *testing.T) {
	mock.ExpectPing().SetVal("PONG")
	Ping()

	myHits := NewHits("myhits")

	mock.ExpectZIncrBy("myhits", float64(1), "foo").SetVal(float64(1))
	myHits.Add("foo", 1)
	mock.ExpectZIncrBy("myhits", float64(1), "bar").SetVal(float64(1))
	myHits.Add("bar", 1)
	mock.ExpectZIncrBy("myhits", float64(1), "bar").SetVal(float64(2))
	myHits.Add("bar", 1)

	mock.ExpectZRevRangeWithScores("myhits", 0, 0).SetVal([]redis.Z{
		{Score: float64(2), Member: "bar"},
	})
	top, count, err := myHits.Top()
	if err != nil {
		t.Fatalf("got an unexpected error; (%s)", err.Error())
	}
	if got, want := top, "bar"; got != want {
		t.Fatalf("top do not match; got(%s) want (%s)", got, want)
	}

	if got, want := count, int64(2); got != want {
		t.Fatalf("top do not match; got(%d) want (%d)", got, want)
	}

	mock.ExpectDel("myhits").SetVal(1)
	myHits.Reset()
}
