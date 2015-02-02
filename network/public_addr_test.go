package network

import "testing"

func TestPublicAddressPrefix(t *testing.T) {
	if b, err := Main.PublicAddressPrefix(); b != 0x00 && err == nil {
		t.Errorf("Main.PublicAddrPrefix != 0x00 network:%v err: %v", b, err)
	}
	if b, err := Test.PublicAddressPrefix(); b != 0x6F && err == nil {
		t.Errorf("Test.PublicAddrPrefix != 0x6F network:%v err: %v", b, err)
	}
	if b, err := Network(123).PublicAddressPrefix(); err == nil {
		t.Errorf("Invalid.PublicAddrPrefix != err network:%v err: %v", b, err)
	}
}
