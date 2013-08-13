package public_key

import (
    "bytes"
    "errors"
    "github.com/steakknife/bitcoin/base58"
    "github.com/steakknife/bitcoin/checksum"
    "github.com/steakknife/bitcoin/network"
)

type PublicKey struct {
    *network.Network
    Address          []byte
}

const (
    AddressSize = 20 // bytes
    MainAddressPrefix = byte(0x00)
    TestAddressPrefix = byte(0x6F)
)

func NewFromAddress(address []byte) (public_key *PublicKey, err error) {
    if AddressSize != len(address) {
        err = errors.New("Bitcoin address has incorrect number of bytes")
        return
    }
    public_key = &PublicKey {
        Network: network.MainNetwork,
        Address: address,
    }
    return
}

func (public_key *PublicKey) AddressPrefix() (address_prefix byte, err error) {
    switch public_key.Network {
        case network.Main: address_prefix = MainAddressPrefix
        case network.Test: address_prefix = TestAddressPrefix
        default:           err = errors.New("Unknown NetworkID")
    }
    return
}

func DecodeAddressPrefix(address_prefix byte) (network_ *network.Network, err error) {
    switch address_prefix {
        case MainAddressPrefix: network_ = network.Main
        case TestAddressPrefix: network_ = network.Test
        default:                err = errors.New("Unknown Network Address Prefix")
    }
    return
}

func (public_key *PublicKey) Encode() (encoded string, err error) {
    address_prefix, err := public_key.AddressPrefix()
    if err != nil {
        return
    }
    data := append([]byte{address_prefix}, public_key.Address...)
    data = append(data, checksum.Checksum(data)...)
    encoded = base58.Encode(data)
    return
}

func Decode(encoded string) (public_key *PublicKey, err error) {
    decoded, err := base58.Decode(encoded)
    if err != nil {
        return
    }
    network, err := DecodeAddressPrefix(decoded[0])
    if err != nil {
        return
    }
    actual_checksum := checksum.Checksum(decoded[:len(decoded)-4])
    expected_checksum := decoded[len(decoded)-4:]
    if bytes.Compare(actual_checksum, expected_checksum) != 0 {
        err = errors.New("Checksum failure")
        return
    }
    address := decoded[1:len(decoded)-4]
    public_key, err := NewFromAddress(address)
    if err != nil {
        return
    }
    public_key.Network = network
    return
}
