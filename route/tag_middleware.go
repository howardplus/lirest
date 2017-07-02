package route

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strings"
	"sync"
)

// URL that starts with underscore are special tags
// tags supported:
// _info: describs the REST info for this url
// _man: manual page (human-readable format) for this url

const (
	TagInfo = "_info" // describes the REST API usage

	TagHeaderName = "X-LIREST-TAG" // the HTTP tag used for the tag

	tagMax = 1
)

var TagSupported map[string]http.HandlerFunc
var once sync.Once

func init() {
	once.Do(func() {
		TagSupported = make(map[string]http.HandlerFunc, tagMax)
		TagSupported[TagInfo] = InfoHandler
	})
}

// return the tag
// or error if a tag is present but not found
// or continue processing if tag is not found
func checkTag(tag string) http.HandlerFunc {
	for t, f := range TagSupported {
		if t == tag {
			return f
		}
	}

	// returning an empty tag indicates that the tag is not processed
	return nil
}

// TagMiddleware handles the special tags
func TagMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	// the tag always appear on the last part
	s := strings.Split(r.URL.Path, "/")
	if last := s[len(s)-1]; len(last) > 1 && last[0] == '_' {
		if f := checkTag(strings.ToLower(last)); f != nil {
			// add the tag
			f(w, r)
		}
	}

	// go to next
	next(w, r)
}

// InfoHandler
// the handler when info tag is found
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"path": r.URL.Path,
	}).Debug("Add info tag")

	r.Header.Set(TagHeaderName, TagInfo)

	idx := strings.LastIndex(r.URL.Path, "/")
	r.URL.Path = r.URL.Path[0:idx]
}
