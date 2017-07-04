package describe

import "strconv"

// describe takes in a set of description files
// and generate the paths/info/man from those
// description
//
// A description is in JSON format.
//
// Asample description looks like this:
//
// {
//   "short": "cpu"
//   "full": "Display the CPU information of the system"
//   "system": {
//     "source": {
//       "type": "filesystem",
//       "path": "/proc/cpuinfo",
//	     "refresh": "never|always|<time(s,m,h)>"
//     },
//     "rdFormat": {
//       "type": "separated",
//       "delimiter": ":",
//       "multiline": true,
//       "multisection": true
//     },
//     "wrFormat": {
//       "type": "regex",
//       "multiline": true
//     }
//   },
//   "api": {
//     "path": "/proc/cpuinfo",
//     "methods": ["GET"],
//	   "descriptions":
//     [
//       {
//         "method": "GET",
//         "short": "Get CPU info",
//         "long": "Get CPU information from the system"
//       }
//     ]
//   },
//   "vars": {
//     "pid": {
//       "dataType": "uint",
//     }
//   }
// }
//
// A description data structure tree:
//
// Description
//     |- Name (string)
//     |- DescriptionSystem
//     |              |- DescriptionSource
//     |              |              |- Type (string)
//     |              |              |- Path (string)
//     |              |              |- Refresh (string)
//     |              |
//     |              |- DescriptionReadFormat (read)
//     |              |               |- Type (string)
//     |              |               |- Delimiter (string)
//     |              |               |- Header (bool)
//     |              |               |- Title ([]string)
//     |              |               |- Regex (string)
//     |              |               |- Multiline (bool)
//     |              |               |- Multisection (bool)
//     |              |
//     |              |- DescriptionWriteFormat (write)
//     |                              |- Type (string)
//     |                              |- Regex (string)
//     |                              |- Multiline (bool)
//     |
//     |- DescriptionApi
//                    |- Path (string)
//                    |- Methods ([]string)
//                    |- []DescriptionApiDesc
//                                  |- Method (string)
//                                  |- Short (string)
//                                  |- Long (string)
//

// Description is the top level structure
type Description struct {
	Name   string            `json:"name"`
	System DescriptionSystem `json:"system"`
	Api    DescriptionApi    `json:"api"`
}

// DescriptionSource describes the source to be read
type DescriptionSource struct {
	Type    string `json:"type"`
	Path    string `json:"path"`
	Refresh string `json:"refresh"`
	Command string `json:"command"`
}

// DescriptionReadFormat describes how the source
// should be read
type DescriptionReadFormat struct {
	Type                string   `json:"type"`
	Delimiter           string   `json:"delimiter"`
	Header              bool     `json:"header"`
	Title               []string `json:"title"`
	Regex               string   `json:"regex"`
	Multiline           bool     `json:"multiline"`
	Multisection        bool     `json:"multisection"`
	HasTitle            bool     `json:"hasTitle"`
	HasHeading          bool     `json:"hasHeading"`
	TitleIncludeHeading bool     `json:"titleIncludeHeading"`
}

// DescriptionWriteFormat describes how the data
// can be written
type DescriptionWriteFormat struct {
	Type      string `json:"type"`
	Delimiter string `json:"delimiter"`
	Multiline bool   `json:"multiline"`
	Regex     string `json:"regex"`
	Min       int64  `json:"min"`
	Max       int64  `json:"max"`
}

// DescriptionSystem is the top level system variable descriptions
type DescriptionSystem struct {
	Source      DescriptionSource      `json:"source"`
	ReadFormat  DescriptionReadFormat  `json:"rdFormat"`
	WriteFormat DescriptionWriteFormat `json:"wrFormat"`
}

// DescriptionApi describes the REST api
type DescriptionApi struct {
	Path         string               `json:"path"`
	Methods      []string             `json:"methods"`
	Descriptions []DescriptionApiDesc `json:"descriptions"`
	Vars         []DescriptionVar     `json:"vars"`
}

// DescriptionVar describes the REST api variables
type DescriptionVar struct {
	Name     string `json:"name"`
	DataType string `json:"dataType"`
}

// DescriptionApiDesc is used for _info documentation
type DescriptionApiDesc struct {
	Method string `json:"method"`
	Short  string `json:"short"`
	Long   string `json:"long"`
}

// DescriptionVarValidate valiates the data type string
// of a variable, which is represented as string
func DescriptionVarValidate(v string, dataType string) bool {
	switch dataType {
	case "uint":
		if _, err := strconv.ParseUint(v, 10, 32); err != nil {
			return false
		}
	}
	// default is ok, including a string type
	return true
}
