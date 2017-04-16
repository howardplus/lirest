package describe

import (
	"bytes"
	"github.com/howardplus/lirest/describe"
	"testing"
)

func TestDescribe(t *testing.T) {
	input1 := `[
{
  "short": "cpu",
  "full": "Display the CPU information of the system",
  "input": {
    "source": {
      "type": "filesystem",
      "path": "/proc/cpuinfo"
    },
    "format": {
      "type": "separated",
      "delimiter": ":",
      "multiline": true,
      "multisection": true
    }
  },
  "output": {
    "parent": "/sys",
    "path": "/cpu",
    "methods": ["GET"]
  }
}
]`
	b := bytes.NewBufferString(input1)
	d1, err := describe.ReadDescriptions(b)
	if err != nil {
		t.Error("test", "cpu", "expect", nil, "got", err.Error())
	}
	if len(d1) != 1 {
		t.Error("test", "cpu", "expect", 1, "got", len(d1))
	}
	if d1[0].Short != "cpu" {
		t.Error("test", "cpu", "expect", "cpu", "got", d1[0].Short)
	}
}

func TestPath(t *testing.T) {
	path := "./"

	def1, err := describe.ReadDescriptionPath(path, nil)
	if err != nil {
		t.Error("test", "cpu", "expect", nil, "got", err.Error())
	}
	d1 := def1.DescriptionMap[describe.DescTypeStandard]
	if d1 == nil {
		t.Error("test", "cpu", "expect", "non-nil", "got", d1)
	}
	if len(d1) != 1 {
		t.Error("test", "cpu", "expect", 1, "got", len(d1))
	}
	if d1[0].Short != "cpu" {
		t.Error("test", "cpu", "expect", "cpu", "got", d1[0].Short)
	}
}
