package describe

import (
	"encoding/json"
	"errors"
	log "github.com/Sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	descTypeInvalid = iota // invalid description type

	// DescTypeStandard is the standard description
	DescTypeStandard
	// DescTypeSysctl is the sysctl description (read from file system directly)
	DescTypeSysctl

	descTypeMax = 2

	descFileExt = ".des" // description file extension
)

// DescType defines the type of description
type DescType int

// global sysctl descriptions
var sysctlDesc []Description = []Description{}

// readDescriptions reads in descriptions from a input, typically a file
// this is a thin wrapper whose purpose is for this to be testable
func readDescriptions(r io.Reader) ([]Description, error) {
	var output []Description
	if err := json.NewDecoder(r).Decode(&output); err != nil {
		return nil, errors.New("Desciption read error: " + err.Error())
	}

	return output, nil
}

// ReadDescriptionPath reads all files in path and read in descriptions
func ReadDescriptionPath(path string, defn *DescDefn) error {

	descriptions := make(map[DescType][]Description, descTypeMax)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return errors.New("Description path error: " + err.Error())
	}

	for _, file := range files {
		// only process known suffix
		if !strings.HasSuffix(file.Name(), descFileExt) {
			continue
		}

		// standard description files
		var dt DescType = DescTypeStandard

		// got the file, try and open it
		f, err := os.Open(path + file.Name())
		if err != nil {
			return errors.New("Description file open error: " + err.Error())
		}
		defer f.Close()

		log.WithFields(log.Fields{
			"file": file.Name(),
		}).Info("Reading description file")

		// read descriptions from file
		d, err := readDescriptions(f)
		if err != nil {
			return err
		}

		if _, found := descriptions[dt]; found {
			// the list goes on...
			descriptions[dt] = append(descriptions[dt], d...)
		} else {
			descriptions[dt] = d
		}

		log.WithFields(log.Fields{
			"file":    file.Name(),
			"entries": len(d),
		}).Info("Processed description file")
	}

	// finally assign the descriptions
	defn.DescriptionMap = descriptions

	return nil
}

func doSysctlFile(path string, info os.FileInfo, err error) error {

	// ignore directory
	if info.IsDir() {
		return nil
	}

	writeflag := info.Mode() & 0200 // consider only root privilege

	log.WithFields(log.Fields{
		"path":      path,
		"read-only": writeflag == 0,
	}).Debug("do sysctl")

	// convert path to sysctl name
	// eg. /proc/sys/vm/nr_hugepages -> vm.nr_hugepages
	sysctlName := strings.Replace(strings.TrimPrefix(path, "/proc/sys/"), "/", ".", -1)

	// convert path to api path
	// eg. /proc/sys/vm/nr_hugepages -> /sysctl/vm/nr_hugepages
	apiPath := "/sysctl" + strings.TrimPrefix(path, "/proc/sys")

	var d Description

	// create read description
	if writeflag == 0 {
		d = Description{
			Name: sysctlName,
			System: DescriptionSystem{
				Source: DescriptionSource{
					Type: "sysctl",
				},
				ReadFormat: DescriptionReadFormat{
					Type:      "asis",
					Path:      path,
					Multiline: false,
				},
			},
			Api: DescriptionApi{
				Path:         apiPath,
				Methods:      []string{"GET"},
				Descriptions: []DescriptionApiDesc{},
			},
		}
	} else {
		d = Description{
			Name: sysctlName,
			System: DescriptionSystem{
				Source: DescriptionSource{
					Type: "sysctl",
				},
				ReadFormat: DescriptionReadFormat{
					Type:      "asis",
					Path:      path,
					Multiline: false,
				},
				WriteFormat: DescriptionWriteFormat{
					Type:      "asis",
					Multiline: false,
				},
			},
			Api: DescriptionApi{
				Path:         apiPath,
				Methods:      []string{"GET", "PUT"},
				Descriptions: []DescriptionApiDesc{},
			},
		}
	}

	sysctlDesc = append(sysctlDesc, d)

	return nil
}

// ReadSysctlDescriptions read from /proc/sys filesystem
// and generate the descriptions automatically
func ReadSysctlDescriptions(defn *DescDefn) error {

	filepath.Walk("/proc/sys", doSysctlFile)

	defn.DescriptionMap[DescTypeSysctl] = sysctlDesc

	return nil
}
