package inversemap

import (
	"reflect"
	"testing"
)

var (
	m0    = map[int]int{1: 5}
	invM0 = map[int]int{5: 1}
)
var (
	m1    = map[int]int{1: 5, 2: 3}
	invM1 = map[int]int{5: 1, 3: 2}
)

func assertInvMap(t *testing.T, m, invM interface{}) {
	if !reflect.DeepEqual(InverseMap(m), invM) {
		t.Errorf("%v != InverseMap(%v)", m, invM)
	}
	if !reflect.DeepEqual(InverseMap(invM), m) {
		t.Errorf("InverseMap(%v) != %v", m, invM)
	}
}

func TestInverseMap(t *testing.T) {
	assertInvMap(t, m0, invM0)
	assertInvMap(t, m1, invM1)
}
