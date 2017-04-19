package describe

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

const (
	DescTypeInvalid = iota
	DescTypeStandard
	DescTypeCustom

	descTypeMax = 2

	descFileExt = ".des" // description file extension
)

var descStandardFiles []string = []string{
	"procfs",
	"sysctl",
}

type DescType int

type DescError struct {
	str string
}

func (d *DescError) Error() string {
	return d.str
}

// read in descriptions from a input, typically a file
// this is a thin wrapper whose purpose is for this to be testable
func ReadDescriptions(r io.Reader) ([]Description, error) {
	var output []Description = nil
	if err := json.NewDecoder(r).Decode(&output); err != nil {
		return nil, &DescError{str: "Desciption read error: " + err.Error()}
	}

	return output, nil
}

// select all files in path and read in descriptions
func ReadDescriptionPath(path string) (*DescDefn, error) {

	descriptions := make(map[DescType][]Description, descTypeMax)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, &DescError{str: "Description path error: " + err.Error()}
	}

	for _, file := range files {
		// only process known suffix
		if !strings.HasSuffix(file.Name(), descFileExt) {
			continue
		}

		// standard description files
		var dt DescType = DescTypeInvalid
		for _, fn := range descStandardFiles {
			if file.Name() == fn+descFileExt {
				dt = DescTypeStandard
			}
		}

		// custom description files
		if dt == DescTypeInvalid {
			dt = DescTypeCustom
		}

		// read descriptions from file
		f, err := os.Open(path + file.Name())
		if err != nil {
			return nil, &DescError{str: "Description file open error: " + err.Error()}
		}
		defer f.Close()

		d, err := ReadDescriptions(f)
		if err != nil {
			return nil, err
		}

		descriptions[dt] = d

		log.WithFields(log.Fields{
			"file": file.Name(),
		}).Info("Processed description file")
	}

	return &DescDefn{DescriptionMap: descriptions}, nil
}
