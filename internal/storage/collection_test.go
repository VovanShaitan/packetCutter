package storage

import (
	"testing"
)

func TestResultCollection(t *testing.T) {
	rc := NewResultCollection()

	rc.Add("ABC", "target1")
	rc.Add("ABC", "target2")
	rc.Add("DEF", "target3")

	targets, exists := rc.Get("ABC")
	if !exists || len(targets) != 2 {
		t.Errorf("Expected 2 targets for ABC, got %v", targets)
	}

	targets, exists = rc.Get("DEF")
	if !exists || len(targets) != 1 {
		t.Errorf("Expected 1 target for DEF, got %v", targets)
	}

	if rc.Len() != 2 {
		t.Errorf("Expected 2 keys, got %d", rc.Len())
	}

	if rc.Count() != 3 {
		t.Errorf("Expected 3 total targets, got %d", rc.Count())
	}
}
