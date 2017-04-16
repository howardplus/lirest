package util

import (
	"github.com/howardplus/lirest/util"
	"testing"
)

func TestQueue(t *testing.T) {
	q1 := util.NewQueue()

	q1.Enqueue(1)
	q1.Enqueue(2)
	q1.Enqueue(3)

	if q1.Size() != 3 {
		t.Error("test size", "expect", 3, "got", q1.Size())
	}

	e1 := q1.Dequeue()
	if e1 != 1 {
		t.Error("test dequeue", "expect", 1, "got", e1)
	}
	if q1.Size() != 2 {
		t.Error("test size", "expect", 2, "got", q1.Size())
	}
}
