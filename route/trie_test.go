package route

import (
	"github.com/howardplus/lirest/route"
	"testing"
)

func TestTrie(t *testing.T) {
	t1 := route.NewTrie()
	if t1.Depth() != 0 {
		t.Error("test", "", "expect", 0, "got", t1.Depth())
	}
	if err := t1.AddPath("/a", 1); err != nil {
		t.Error("expect", nil, "got", err.Error())
	}
	if t1.Depth() != 1 {
		t.Error("test", "/a", "expect", 1, "got", t1.Depth())
	}
	if t1.Count() != 2 {
		t.Error("test", "/a", "expect", 2, "got", t1.Count())
	}
	if err := t1.AddPath("/a/b", 1); err != nil {
		t.Error("expect", nil, "got", err.Error())
	}
	if t1.Depth() != 2 {
		t.Error("test", "/a/b", "expect", 2, "got", t1.Depth())
	}
	if t1.Count() != 3 {
		t.Error("test", "/a/b", "expect", 3, "got", t1.Count())
	}
	if err := t1.AddPath("/a/b", 1); err == nil {
		t.Error("expect error", "got nil")
	}
	t1.AddPath("/c/d/e/f/g", 1)
	if t1.Count() != 8 {
		t.Error("test", "/a/b", "expect", 8, "got", t1.Count())
	}

	t2 := route.NewTrie()
	t2.AddPath("/a/b", 1)
	if d := t2.Depth(); d != 2 {
		t.Error("test", "/a/b", "expect", 2, "got", d)
	}
	if c := t2.Count(); c != 3 {
		t.Error("test", "/a/b", "expect", 3, "got", c)
	}
}
