package network

import (
    "testing"
)

func Test0(t *testing.T) {
    if 0 != MainNetwork.NetworkID {
        t.Errorf("Main network ID should be 0")
    }
    if 1 != TestNetwork.NetworkID {
        t.Errorf("Test network ID should be 1")
    }
    if 3 != NewNetwork(byte(3)).NetworkID {
        t.Errorf("Created network ID should be 3")
    }
}
