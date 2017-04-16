package route

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strings"
)

// URL that starts with underscore are special tags
// tags supported:
// _info: describs the REST info for this url
// _man: manual page (human-readable format) for this url

const (
	TagInfo = "_info"
	TagMan  = "_man"

	TagHeaderName = "X-LIREST-TAG"

	TagMax = 2
)

var TagSupported map[string]http.HandlerFunc

func init() {
	TagSupported = make(map[string]http.HandlerFunc, TagMax)
	TagSupported[TagInfo] = InfoHandler
	TagSupported[TagMan] = ManHandler
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

// Tag middleware handles the special tags
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

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"path": r.URL.Path,
	}).Debug("Add info tag")

	r.Header.Set(TagHeaderName, TagInfo)

	idx := strings.LastIndex(r.URL.Path, "/")
	r.URL.Path = r.URL.Path[0:idx]
}

func ManHandler(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"path": r.URL.Path,
	}).Debug("Add manual tag")

	r.Header.Set(TagHeaderName, TagMan)

	idx := strings.LastIndex(r.URL.Path, "/")
	r.URL.Path = r.URL.Path[0:idx]
}
