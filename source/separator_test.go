package source

import (
	"github.com/howardplus/lirest/source"
	"testing"
)

func TestColonSeparator(t *testing.T) {
	sep1 := source.NewSeparatorConverter(":")

	s1 := "foo:bar"
	k, v, err := sep1.ConvertLine(s1)
	if err != nil {
		t.Error("test", s1, "expect", nil, "got", err.Error())
	}
	if k != "foo" {
		t.Error("test", s1, "expect", "foo", "got", k)
	}
	if v != "bar" {
		t.Error("test", s1, "expect", "bar", "got", v)
	}

	s2 := "foo: bar"
	k, v, err = sep1.ConvertLine(s2)
	if err != nil {
		t.Error("test", s2, "expect", nil, "got", err.Error())
	}
	if k != "foo" {
		t.Error("test", s2, "expect", "foo", "got", k)
	}
	if v != "bar" {
		t.Error("test", s2, "expect", "bar", "got", v)
	}

	s3 := "foo    :   bar"
	k, v, err = sep1.ConvertLine(s3)
	if err != nil {
		t.Error("test", s3, "expect", nil, "got", err.Error())
	}
	if k != "foo" {
		t.Error("test", s3, "expect", "foo", "got", k)
	}
	if v != "bar" {
		t.Error("test", s3, "expect", "bar", "got", v)
	}

	s4 := "    foo    :   bar   "
	k, v, err = sep1.ConvertLine(s4)
	if err != nil {
		t.Error("test", s4, "expect", nil, "got", err.Error())
	}
	if k != "foo" {
		t.Error("test", s4, "expect", "foo", "got", k)
	}
	if v != "bar" {
		t.Error("test", s4, "expect", "bar", "got", v)
	}

	s5 := "foobar"
	_, _, err = sep1.ConvertLine(s5)
	if err == nil {
		t.Error("test", s5, "expect", "error", "got", err)
	}
}
