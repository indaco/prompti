package huhfield

import (
	"errors"
	"testing"
)

func TestBase_GetKey(t *testing.T) {
	b := &Base{Key: "test-key"}
	if got := b.GetKey(); got != "test-key" {
		t.Errorf("GetKey() = %q, want %q", got, "test-key")
	}
}

func TestBase_Error(t *testing.T) {
	b := &Base{}
	if b.Error() != nil {
		t.Error("expected nil error initially")
	}
	e := errors.New("some error")
	b.Err = e
	if b.Error() != e {
		t.Error("expected error to match")
	}
}

func TestBase_Skip(t *testing.T) {
	b := &Base{}
	if b.Skip() {
		t.Error("expected Skip() = false")
	}
}

func TestBase_Zoom(t *testing.T) {
	b := &Base{}
	if b.Zoom() {
		t.Error("expected Zoom() = false")
	}
}
