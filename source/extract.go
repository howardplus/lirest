package source

import (
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/describe"
	"github.com/howardplus/lirest/util"
	"os"
	"strconv"
	"strings"
	"time"
)

// Extractor returns a generic data based
// on the converter.
// An object that implements the Extractor interface needs
// to know where to get the data, which then feeds to the
// converter.
type Extractor interface {
	Extract() (*ExtractOutput, error)
}

// ExtractOutput is the output of the extracted data
// with json tags
type ExtractOutput struct {
	Name string      `json:"name"`
	Time time.Time   `json:"time"`
	Data interface{} `json:"data"`
}

// NewExtractor create a new extractor based on the description
func NewExtractor(s describe.DescriptionSource, rd describe.DescriptionReadFormat, c Converter, vars map[string]string) (Extractor, error) {
	var extractor Extractor

	refresh := time.Duration(0)
	switch s.Refresh {
	case "never":
		// never say never, 10 day is long enough
		refresh = 240 * time.Hour
	default:
		// something s/m/h
		v, err := strconv.Atoi(s.Refresh[:len(s.Refresh)-1])
		if err == nil {
			if strings.HasSuffix(s.Refresh, "s") {
				refresh = time.Duration(v) * time.Second
			} else if strings.HasSuffix(s.Refresh, "m") {
				refresh = time.Duration(v) * time.Minute
			} else if strings.HasSuffix(s.Refresh, "h") {
				refresh = time.Duration(v) * time.Hour
			}
		}
	case "":
		// Did not specify, which implies always refresh
	}

	switch s.Type {
	case "procfs", "sysfs", "sysctl":
		extractor = NewGenericExtractor(rd.Path, refresh, c, vars)
	case "command":
		extractor = NewCommandExtractor(rd.Command, c, vars)
	}

	// found an extractor, use it
	if extractor != nil {
		return extractor, nil
	}

	// return error on default
	return nil, util.NewError("Internal error: unknown input type")
}

// GenericExtractor extract data from reading from a file
// use this until it's not enough
type GenericExtractor struct {
	path    string
	conv    Converter
	refresh time.Duration
	vars    map[string]string
}

// NewGenericExtractor creates a GenericExtractor
func NewGenericExtractor(path string, refresh time.Duration, conv Converter, vars map[string]string) *GenericExtractor {
	return &GenericExtractor{path: path, refresh: refresh, conv: conv, vars: vars}
}

func (e *GenericExtractor) Extract() (*ExtractOutput, error) {
	log.WithFields(log.Fields{
		"path": e.path,
		"vars": e.vars,
	}).Debug("Extract from file system")

	// create path from variables
	path, err := util.FillVars(e.path, e.vars)
	if err != nil {
		return nil, util.NewError("Failed to generate path")
	}

	// ask data from cache
	var hash string
	if e.refresh != time.Duration(0) {
		hash = CacheHash("command" + path)
		if data, time, err := Cache(hash); err == nil {

			log.WithFields(log.Fields{
				"hash": hash,
				"path": e.path,
			}).Debug("Serve from cache")

			return &ExtractOutput{
				Name: e.conv.Name(),
				Time: time,
				Data: data,
			}, nil
		}
	}

	// open file from path
	f, err := os.Open(path)
	if err != nil {
		return nil, util.NewError("Failed to open system path")
	}
	defer f.Close()

	// TODO: verify the rw format on this path

	// give it to the converter
	data, err := e.conv.ConvertStream(f)
	if err != nil {
		return nil, err
	}

	// send to cache
	if e.refresh != time.Duration(0) {
		if err := SendCache(hash, data, e.refresh); err != nil {
			// cache error, non-fatal
			log.WithFields(log.Fields{
				"path": e.path,
			}).Debug("Failed to send cache")
		}
	}

	log.WithFields(log.Fields{
		"path": e.path,
	}).Debug("Convert successful")

	return &ExtractOutput{
		Name: e.conv.Name(),
		Time: time.Now(),
		Data: data,
	}, nil
}
