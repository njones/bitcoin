package network

import "testing"

func Test0(t *testing.T) {
	if Main != 0 {
		t.Errorf("Main network ID should be 0")
	}
	if Test != 1 {
		t.Errorf("Test network ID should be 1")
	}
	if Network(3) != 3 {
		t.Errorf("Created network ID should be 3")
	}
}
