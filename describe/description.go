package describe

import (
	_ "github.com/Sirupsen/logrus"
)

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
//   "read-only": true,
//   "input": {
//     "source": {
//       "type": "filesystem",
//       "path": "/proc/cpuinfo",
//     },
//     "format": {
//       "type": "separated",
//       "delimiter": ":",
//       "multiline": true,
//       "multisection": true,
//     }
//   },
//   "output": {
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
//   }
// }
//

type Description struct {
	Name     string            `json:"name"`
	Readonly bool              `json:"read-only"`
	Input    DescriptionInput  `json:"input"`
	Output   DescriptionOutput `json:"output"`
}

type DescriptionSource struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

type DescriptionFormat struct {
	Type         string `json:"type"`
	Delimiter    string `json:"delimiter"`
	Multiline    bool   `json:"multiline"`
	Multisection bool   `json:"multisection"`
}

type DescriptionInput struct {
	Source DescriptionSource `json:"source"`
	Format DescriptionFormat `json:"format"`
}

type DescriptionOutput struct {
	Path         string                  `json:"path"`
	Methods      []string                `json:"methods"`
	Descriptions []DescriptionOutputDesc `json:"descriptions"`
}

type DescriptionOutputDesc struct {
	Method string `json:"method"`
	Short  string `json:"short"`
	Long   string `json:"long"`
}
