package network

import "testing"

func TestPrivateAddressPrefix(t *testing.T) {
	if b, err := Main.PrivateAddressPrefix(); b != 0x80 && err == nil {
		t.Errorf("Main.PrivateAddrPrefix != 0x80 network:%v err: %s", b, err)
	}
	if b, err := Test.PrivateAddressPrefix(); b != 0xEF && err == nil {
		t.Errorf("Test.PrivateAddrPrefix != 0xEF network:%v err: %s", b, err)
	}

	if b, err := Network(123).PrivateAddressPrefix(); err == nil {
		t.Errorf("Invalid.PrivateAddrPrefix != err network:%v err: %s", b, err)
	}
}
