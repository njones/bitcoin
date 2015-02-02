package network

import "fmt"

type Network byte

const (
	Main Network = iota
	Test
)

func (net Network) String() string {
	switch net {
	case Main:
		return "Main"
	case Test:
		return "Test"
	default:
		return fmt.Sprintf("Unknown btc network %d (0x%02X)", net, net)
	}
}
