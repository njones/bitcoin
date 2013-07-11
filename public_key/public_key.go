package public_key

import (
    "bytes"
    "errors"
    "base58"
    "checksum"
    "network"
)

type PublicKey struct {
    *network.Network
    Address          []byte
}

const AddressSize = 20 // bytes
const MainAddressPrefix = byte(0x00)
const TestAddressPrefix = byte(0x6F)

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
    switch public_key.Network.NetworkID {
    case network.MainID: address_prefix = MainAddressPrefix
    case network.TestID: address_prefix = TestAddressPrefix
    default:             err = errors.New("Unknown NetworkID")
    }
    return
}

func DecodeAddressPrefix(address_prefix byte) (network_ *network.Network, err error) {
    switch address_prefix {
    case MainAddressPrefix: network_ = network.MainNetwork
    case TestAddressPrefix: network_ = network.TestNetwork
    default:                err = errors.New("Unknown Network Address Prefix")
    }
    return
}

func (public_key *PublicKey) Encode() (encoded string, err error) {
    var address_prefix byte
    if address_prefix, err = public_key.AddressPrefix(); nil != err {
        return
    }
    data := append([]byte{address_prefix}, public_key.Address...)
    data = append(data, checksum.Checksum(data)...)
    encoded = base58.Encode(data)
    return
}

func Decode(encoded string) (public_key *PublicKey, err error) {
    var decoded []byte
    if decoded, err = base58.Decode(encoded); nil != err {
        return
    }
    var network *network.Network
    if network, err = DecodeAddressPrefix(decoded[0]); nil != err {
        return
    }
    actual_checksum := checksum.Checksum(decoded[0:len(decoded)-4])
    expected_checksum := decoded[len(decoded)-4:]
    if 0 != bytes.Compare(actual_checksum, expected_checksum) {
        err = errors.New("Checksum failure")
        return
    }
    address := decoded[1:len(decoded)-4]
    if public_key, err = NewFromAddress(address); nil != err {
        return
    }
    public_key.Network = network
    return
}
