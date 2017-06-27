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
//     "format": {
//       "type": "separated",
//       "delimiter": ":",
//       "multiline": true,
//       "multisection": true
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
//       "data-type": "uint",
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
//     |              |- DescriptionFormat (read)
//     |              |               |- Type (string)
//     |              |               |- Delimiter (string)
//     |              |               |- Header (bool)
//     |              |               |- Title ([]string)
//     |              |               |- Regex (string)
//     |              |               |- Multiline (bool)
//     |              |               |- Multisection (bool)
//     |              |
//     |              |- DescriptionFormat (write)
//     |                              |- Type (string)
//     |                              |- Delimiter (string)
//     |                              |- Header (bool)
//     |                              |- Title ([]string)
//     |                              |- Regex (string)
//     |                              |- Multiline (bool)
//     |                              |- Multisection (bool)
//     |
//     |- DescriptionApi
//                    |- Path (string)
//                    |- Methods ([]string)
//                    |- []DescriptionApiDesc
//                                  |- Method (string)
//                                  |- Short (string)
//                                  |- Long (string)
//

// Description
type Description struct {
	Name   string            `json:"name"`
	System DescriptionSystem `json:"system"`
	Api    DescriptionApi    `json:"api"`
	Vars   []DescriptionVar  `json:"vars"`
}

// DescriptionSource
type DescriptionSource struct {
	Type    string `json:"type"`
	Path    string `json:"path"`
	Refresh string `json:"refresh"`
	Command string `json:"command"`
}

// DescriptionFormat
type DescriptionFormat struct {
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

// DescriptionSystem
type DescriptionSystem struct {
	Source      DescriptionSource `json:"source"`
	ReadFormat  DescriptionFormat `json:"rd-format"`
	WriteFormat DescriptionFormat `json:"wr-format"`
}

// DescriptionApi
type DescriptionApi struct {
	Path         string               `json:"path"`
	Methods      []string             `json:"methods"`
	Descriptions []DescriptionApiDesc `json:"descriptions"`
}

// DescriptionVar
type DescriptionVar struct {
	Name     string `json:"name"`
	DataType string `json:"data-type"`
}

// DescriptionApiDesc
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
