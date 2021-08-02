package store

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/diggs/go-backoff"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

// DefaultDatabaseConnect is the default redis connect scheme
// ([[db:]password@]host:port)
const DefaultDatabaseConnect = "localhost:6379"

var (
	rdb *redis.Client
	ctx context.Context
)

// parseCnxNfo format:  "[[db:]password@]host:port"
//  ex :
//    "0:password@localhost:6379"
//    "password@localhost:6379"
//    "localhost:6379"
func parseCnxNfo(s string) (*redis.Options, error) {
	if s == "" {
		return nil, errors.New("Bad connection parameter: empty string")
	}

	var (
		addr string
		pwd  string
		db   int
		err  error
	)
	v := strings.Split(s, "@")
	lv := len(v)

	if lv > 2 {
		return nil, fmt.Errorf("Bad connection parameter: got (%s)", s)
	}

	if lv == 1 {
		addr = v[0]
	} else {
		addr = v[1]

		dbPwd := strings.Split(v[0], ":")
		lv = len(dbPwd)

		if lv > 2 {
			return nil, fmt.Errorf("Bad connection parameter: got (%s)", s)
		}

		if lv == 1 {
			pwd = dbPwd[0]
		} else {
			pwd = dbPwd[1]

			db, err = strconv.Atoi(dbPwd[0])
			if err != nil {
				return nil, fmt.Errorf("Bad connection parameter: got (%s)", s)
			}
		}
	}

	return &redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	}, nil
}

// Setup redis context and client
//  store.Setup(ctx.Background(), "db:password@host:port")
//  store.Setup(ctx.Background(), "0:password@localhost:6379")
//  store.Setup(ctx.Background(), "password@localhost:6379")
//  store.Setup(ctx.Background(), "localhost:6379")
func Setup(cx context.Context, c string) error {
	opts, err := parseCnxNfo(c)
	if err != nil {
		return err
	}
	ctx = cx
	rdb = redis.NewClient(opts)

	return nil
}

// Ping the database
func BlockedPingBakkoff() (pong string, err error) {
	// Back off exponentially, starting at 3 milliseconds, capping at 320 seconds
	exp := backoff.NewExponentialFullJitter(3*time.Millisecond, 320*time.Second)
	for {
		pong, err = Ping()
		if err != nil {
			log.Errorf("%s - backing off %d second(s)", err.Error(), exp.NextDuration/time.Second)
			exp.Backoff()
		} else {
			exp.Reset()
			return pong, nil
		}
	}
}

func Ping() (string, error) {
	pong := rdb.Ping(ctx).Val()

	if pong == "" {
		return "", errors.New("unreachable redis server")
	}
	return pong, nil
}

// Hitable is the Hits interface
type Hitable interface {
	Add(k string, i int64)
	Top() (string, int64, error)
	Reset()
}

// Hits is a hits bag identified by a Key
type Hits struct {
	Key string
}

// NewHits returns a new hits bag.
func NewHits(k string) Hitable {
	return Hits{Key: k}
}

// Add an hit in hits bag
func (h Hits) Add(k string, i int64) {
	// TODO make a binarie string to compact the key
	rresp, err := rdb.ZIncrBy(ctx, h.Key, float64(i), k).Result()
	if err != nil {
		log.Errorf("error to add hit (%s)", err.Error())
	}
	log.Debugf("success stores hit to %s count(%d)", k, int64(rresp))
}

// Top returns the most popular hit in hits bag
func (h Hits) Top() (string, int64, error) {
	vals, err := rdb.ZRevRangeWithScores(ctx, h.Key, 0, 0).Result()
	if err != nil {
		return "", 0, err
	}
	if len(vals) < 1 {
		return "", 0, nil
	}

	return vals[0].Member.(string), int64(vals[0].Score), nil
}

// Reset deletes key of hits bag
func (h Hits) Reset() {
	del := rdb.Del(ctx, h.Key).Val()
	log.Debugf("reset hits bag to %s redis(%d)", h.Key, del)
}
