package source

import (
	log "github.com/Sirupsen/logrus"
	"github.com/howardplus/lirest/util"
	"regexp"
	"strings"
)

// RegexValidator validates data based on regex
type RegexValidator struct {
	multiline bool
	regexExp  string
	regex     *regexp.Regexp
}

func NewRegexValidator(regex string, multiline bool) *RegexValidator {
	return &RegexValidator{
		regex:     regexp.MustCompile(regex), // panic if failed
		multiline: multiline,
	}
}

// Validate based on regex
func (v *RegexValidator) Validate(data string) error {

	log.WithFields(log.Fields{
		"regex":     v.regexExp,
		"multiline": v.multiline,
	}).Debug("Validate by regex")

	var lines []string
	if v.multiline {
		lines = strings.Split(data, "\n")
	} else {
		lines = make([]string, 1, 1)
		lines[0] = data
	}

	for _, l := range lines {
		if !v.regex.MatchString(l) {
			return util.NewError("Validation error")
		}
	}

	return nil
}
