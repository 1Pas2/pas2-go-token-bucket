package pas2

import (
	"testing"
	"time"
)

func TestRefill(t *testing.T) {
	now := time.Unix(0, 0)
	b, _ := New(2, 1, now)
	if !b.Take(now, 2) || b.Take(now, 1) || !b.Take(now.Add(time.Second), 1) {
		t.Fatal("refill")
	}
}
func TestBackwards(t *testing.T) {
	now := time.Unix(10, 0)
	b, _ := New(1, 1, now)
	b.Take(now, 1)
	if b.Take(now.Add(-time.Second), 1) {
		t.Fatal("backward clock created tokens")
	}
}
func TestInvalid(t *testing.T) {
	if _, e := New(0, 1, time.Now()); e == nil {
		t.Fatal("invalid capacity")
	}
	b, _ := New(1, 1, time.Now())
	if b.Take(time.Now(), 0) {
		t.Fatal("invalid cost")
	}
}
