package main

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/mattheath/base62"
	"hash"
	"time"
)

const (
	URLIDKEY           = "next.url.id"
	ShortLinkKey       = "shortLink:%s:url"
	URLHashKey         = "urlHash:%s:url"
	ShortLinkDetailKey = "shortLink:%s:detail"
)

type RedisCli struct {
	Cli *redis.Client
}

type URLDetail struct {
	URL                 string        `json:"url"`
	CreateAt            string        `json:"create_at"`
	ExpirationInMinutes time.Duration `json:"expiration_in_minutes"`
}

// init redis client
func NewRedisCli(addr string, passwd string, db int) *RedisCli {
	var (
		c   *redis.Client
		err error
	)
	c = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       db,
	})
	if _, err = c.Ping().Result(); err != nil {
		panic(err)
	}
	return &RedisCli{Cli: c}

}

// shorten method
func (r *RedisCli) Shorten(url string, exp int64) (string, error) {
	//	convert url to sha1 hash
	var (
		h      string
		d      string
		id     int64
		eid    string
		err    error
		detail []byte
	)
	h = toSha1(url)
	if d, err = r.Cli.Get(fmt.Sprintf(URLHashKey, h)).Result(); err == redis.Nil {
		//	no existed and nothing to do
	} else if err != nil {
		return "", err
	} else {
		if d == "{}" {
			// expiration and nothing to do

		} else {
			return d, nil
		}
	}

	// increase global counter
	if err = r.Cli.Incr(URLIDKEY).Err(); err != nil {
		return "", err
	}

	if id, err = r.Cli.Get(URLIDKEY).Int64(); err != nil {
		return "", err
	}
	//	encode global counter to base62 string
	eid = base62.EncodeInt64(id)
	if err = r.Cli.Set(fmt.Sprintf(ShortLinkKey, eid), url, time.Minute*time.Duration(exp)).Err(); err != nil {
		return "", err
	}

	if err = r.Cli.Set(fmt.Sprintf(URLHashKey, h), eid, time.Minute*time.Duration(exp)).Err(); err != nil {
		return "", err
	}

	if detail, err = json.Marshal(
		&URLDetail{
			URL:                 url,
			CreateAt:            time.Now().String(),
			ExpirationInMinutes: time.Duration(exp),
		}); err != nil {
		return "", err
	}
	if err = r.Cli.Set(fmt.Sprintf(ShortLinkDetailKey, eid), detail, time.Minute*time.Duration(exp)).Err(); err != nil {
		return "", err
	}

	return eid, nil

}

// hash method
func toSha1(s string) string {
	var (
		sha hash.Hash
	)
	sha = sha1.New()
	return string(sha.Sum([]byte(s)))
}

// short link info method
func (r *RedisCli) ShortLinkInfo(eid string) (interface{}, error) {
	var (
		d   string
		i   interface{}
		err error
	)
	if d, err = r.Cli.Get(fmt.Sprintf(ShortLinkDetailKey, eid)).Result(); err == redis.Nil {
		return "", StatusError{
			Code: 404,
			Err:  errors.New("unknown short url"),
		}
	} else if err != nil {
		return "", err
	} else {
		if err = json.Unmarshal([]byte(d), &i); err != nil {
			return "", err
		}
		return i, nil
	}
}

// un shorten method -> get source url
func (r *RedisCli) UnShorten(eid string) (string, error) {
	var (
		url string
		err error
	)
	if url, err = r.Cli.Get(fmt.Sprintf(ShortLinkKey, eid)).Result(); err == redis.Nil {
		return "", StatusError{
			Code: 404,
			Err:  errors.New("unknown short url"),
		}
	} else if err != nil {
		return "", err
	} else {
		return url, nil
	}

}
