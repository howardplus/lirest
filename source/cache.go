package source

import (
	"crypto/sha256"
	"encoding/base64"
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/util"
	"time"
)

type cacheMsg struct {
	hash   string
	data   interface{}
	expire time.Duration
}

type cacheData struct {
	data interface{}
	time time.Time
	err  error
}

type cacheInternalData struct {
	data     interface{}
	create   time.Time
	lastused time.Time
	expire   time.Duration
}

// channels
var cacheReqChan chan string
var cacheDataChan chan *cacheData
var cacheSendChan chan *cacheMsg

const (
	sendChanQ    = 100
	cacheInitCap = 1000
)

// CacheManager is a background thread that
// expires cache entries periodically
func CacheManager() {
	// channels
	cacheReqChan = make(chan string, 1)
	cacheDataChan = make(chan *cacheData, 1)
	cacheSendChan = make(chan *cacheMsg, sendChanQ)

	// internal data for all the cache
	cacheMap := make(map[string]*cacheInternalData, 0)

	for {
		select {
		case req := <-cacheReqChan:
			// requesting cache data
			if data, found := cacheMap[req]; !found || data == nil {
				cacheDataChan <- &cacheData{
					err: util.NewError("Cache not found"),
				}
			} else {
				data.lastused = time.Now()
				cacheDataChan <- &cacheData{
					data: data.data,
					time: data.create,
				}
			}
		case msg := <-cacheSendChan:
			// creating cache
			cacheMap[msg.hash] = &cacheInternalData{
				data:     msg.data,
				expire:   msg.expire,
				create:   time.Now(),
				lastused: time.Now(),
			}
		case <-time.After(time.Second):
			// timer
			now := time.Now()
			for k, v := range cacheMap {
				if v.lastused.Add(v.expire).Unix() <= now.Unix() {
					log.WithFields(log.Fields{
						"key":    k,
						"expire": v.expire,
						"now":    now,
					}).Debug("cache entry expired")
					delete(cacheMap, k)
				}
			}
		}
	}
}

// CacheHash retusn the hash value of the cache
func CacheHash(path string) string {
	sum := sha256.Sum256([]byte(path))
	return base64.StdEncoding.EncodeToString(sum[:])
}

// Cache requests cache from the cache manager
func Cache(hash string) (interface{}, time.Time, error) {
	// blocking call to get the hash result
	cacheReqChan <- hash
	data := <-cacheDataChan
	if data.err != nil {
		return nil, time.Unix(0, 0), data.err
	}

	return data.data, data.time, nil
}

// SendCache sends a cache result to cache manager
func SendCache(hash string, data interface{}, expire time.Duration) error {

	log.WithFields(log.Fields{
		"hash":   hash,
		"expire": expire,
	}).Debug("Send cache")

	cacheSendChan <- &cacheMsg{
		hash:   hash,
		data:   data,
		expire: expire,
	}

	return nil
}
