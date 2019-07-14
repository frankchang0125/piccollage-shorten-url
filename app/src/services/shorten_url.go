package services

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"pic-collage.com/shorten_url/models"
)

var shortenURLDao models.ShortenURLDao
var start, end, counter uint64
var lock = sync.Mutex{}
var dispatcherURL string

func InitShortenURLService(db *sql.DB) (err error) {
	shortenURLDao = models.NewShortenURLSQLDao(db)
	dispatcherHost := viper.GetString("dispatcher.host")
	dispatcherPort := viper.GetInt("dispatcher.port")
	dispatcherURL = fmt.Sprintf("http://%s:%d",
		dispatcherHost, dispatcherPort)
	return fetchCounterRange()
}

func CreateShortenURL(url string) (shorten string, err error) {
	c := incrCounter()
	for c >= end {
		err = fetchCounterRange()
		if err != nil {
			return
		}
		c = incrCounter()
	}

	shorten = base62Encode(c)
	err = shortenURLDao.CreateShortenURL(url, shorten)
	if err != nil {
		if err == models.ErrDuplicateURL {
			// URL already shorten.
			shorten, err = shortenURLDao.GetShortenURL(url)
			if err != nil {
				return
			}
			return shorten, nil
		}
		return
	}

	// Save shorten url mapping to cache.
	err = redisClient.HSet(shortenURLKey, shorten, url).Err()
	if err != nil {
		log.WithError(err).Error("Fail to save shorten url to cache")
	}

	return
}

func GetURL(shorten string) (url string, err error) {
	// Try to get original url from cache.
	url, err = redisClient.HGet(shortenURLKey, shorten).Result()
	if err != nil {
		log.WithError(err).Warn("Fail to get shorten url from cache")
	} else {
		return url, nil
	}

	// Cache not found, try to get url from database.
	return shortenURLDao.GetURL(shorten)
}

// incrCounter increase the current counter
// and return the counter value before increasing.
func incrCounter() uint64 {
	for {
		v := atomic.LoadUint64(&counter)
		if atomic.CompareAndSwapUint64(&counter, v, v+1) {
			return v
		}
	}
}

type dispatchResp struct {
	Start uint64 `json:"start"`
	End   uint64 `json:"end"`
}

func fetchCounterRange() (err error) {
	lock.Lock()
	defer lock.Unlock()

	log.Info("Fetching new counter ranges")

	// Double check counter range,
	// if it's been updated by other thread already, do nothing.
	v := atomic.LoadUint64(&counter)
	if v < end {
		return
	}

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	request, err := http.NewRequest(http.MethodGet, dispatcherURL, nil)
	if err != nil {
		log.WithError(err).Error("Fail to create counter dispatcher service request")
		return
	}

	resp, err := client.Do(request)
	if err != nil {
		log.WithError(err).Error("Fail to request counter dispatcher service")
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.WithError(err).Error("Fail to read response body")
		return
	}

	var r = dispatchResp{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.WithError(err).Error("Fail to unmarshal response")
		return
	}

	atomic.StoreUint64(&start, max(r.Start, 62))
	atomic.StoreUint64(&end, max(r.End, 62))
	atomic.StoreUint64(&counter, max(r.Start, 62))

	log.WithFields(log.Fields{
		"start": r.Start,
		"end":   r.End,
	}).Debug("Current counter range")

	return
}

func base62Encode(c uint64) string {
	s := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var encode bytes.Buffer
	for c > 0 {
		encode.WriteByte(s[c%62])
		c /= 62
	}

	return reverse(encode.Bytes())
}

func reverse(s []byte) string {
	l := 0
	r := len(s) - 1

	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}

	return string(s)
}

func max(x, y uint64) uint64 {
	if x > y {
		return x
	}
	return y
}
