package network

import (
	"fmt"
	"github.com/steakknife/bitcoin/util/inversemap"
)

var (
	publicNetworksToPrefixes = map[Network]byte{Main: 0x00, Test: 0x6F}
	publicPrefixesToNetworks = inversemap.InverseMap(publicNetworksToPrefixes).(map[byte]Network)
)

func (net Network) PublicAddressPrefix() (addrPrefix byte, err error) {
	a, ok := publicNetworksToPrefixes[net]
	if !ok {
		return 0, fmt.Errorf("Unknown bitcoin public network ID (enum): %d", net)
	}
	return a, nil
}

func DecodePublicAddressPrefix(addrPrefix byte) (net Network, err error) {
	n, ok := publicPrefixesToNetworks[addrPrefix]
	if !ok {
		return 0, fmt.Errorf("Unknown bitcoin public address prefix %02X", addrPrefix)
	}
	return n, nil
}
