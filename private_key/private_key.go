package private_key

import (
//    "fmt"
    "bytes"
    "math/big"
    "errors"
    "code.google.com/p/go.crypto/ripemd160"
    "crypto/rand"
    "crypto/sha256"
    "github.com/steakknife/Golang-Koblitz-elliptic-curve-DSA-library/bitelliptic"
    "github.com/steakknife/Golang-Koblitz-elliptic-curve-DSA-library/bitecdsa"
    "base58"
    "network"
    "public_key"
    "checksum"
)

type PrivateKey struct {
    *network.Network
    *bitecdsa.PrivateKey 
}

const ExponentSize = 32 // bytes
const MainAddressPrefix = byte(0x80)
const TestAddressPrefix = byte(0xEF)

func Generate() (private_key *PrivateKey, err error) {
    var ecdsa_private_key *bitecdsa.PrivateKey
    if ecdsa_private_key, err = bitecdsa.GenerateKey(bitelliptic.S256(), rand.Reader); nil != err {
        return
    }
    private_key_ := &PrivateKey {
        Network: network.MainNetwork,
        PrivateKey: ecdsa_private_key,
    }
    private_key_encoded, err2 := private_key_.Encode()
    if nil != err2 {
        err = err2
        return
    }
    private_key_decoded, err3 := Decode(private_key_encoded)
    if nil != err3 || nil == private_key_decoded {
        err = err3
        return
    }
    private_key = private_key_
    return
}

func NewFromExponent(exponent []byte) (private_key *PrivateKey, err error) {
    if ExponentSize != len(exponent) {
        err = errors.New("Exponent of wrong size")
        return
    }
    
    private_key = &PrivateKey {
        Network: network.MainNetwork,
        PrivateKey: bitecdsa.SetKey(bitelliptic.S256(), new(big.Int).SetBytes(exponent)),
    }
    return
}

func (private_key *PrivateKey) PublicKey() (*public_key.PublicKey) {
    return &public_key.PublicKey {
        Network: private_key.Network,
        Address: private_key.PublicAddress(),
    }
}

func PadToSize(array []byte, size int) ([]byte) {
    if len(array) < size {
        result := append(bytes.Repeat([]byte{0}, size-len(array)), array...)
        return result
    }
    return array
}

func (private_key *PrivateKey) PublicAddressString() (public_addres_string string, err error) {
    return private_key.PublicKey().Encode()
}

func (private_key *PrivateKey) PublicAddress() (public_address []byte) {
    public_key := private_key.PrivateKey.PublicKey
    x := PadToSize(public_key.X.Bytes(), 32)
    y := PadToSize(public_key.Y.Bytes(), 32)

    first_round := sha256.New()
    first_round.Write([]byte{0x04})
    first_round.Write(x)
    first_round.Write(y)

    second_round := ripemd160.New()
    second_round.Write(first_round.Sum(nil))

    public_address = second_round.Sum(nil)
    return
}

func (private_key *PrivateKey) AddressPrefix() (address_prefix byte, err error) {
    switch private_key.Network.NetworkID {
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

func (private_key *PrivateKey) Encode() (result string, err error) {
    var address_prefix byte
    if address_prefix, err = private_key.AddressPrefix(); nil != err {
        return
    }
    exponent := PadToSize(private_key.PrivateKey.D.Bytes(), 32)
    data := append([]byte{address_prefix}, exponent...)
    data = append(data, checksum.Checksum(data)...)

    result = base58.Encode(data)
    return
}

func Decode(encoded string) (private_key *PrivateKey, err error) {
    var decoded []byte
    if decoded, err = base58.Decode(encoded); nil != err {
        return
    }
    var network_ *network.Network
    if network_, err = DecodeAddressPrefix(decoded[0]); nil != err {
        return
    }
    actual_checksum := checksum.Checksum(decoded[0:len(decoded)-4])
    expected_checksum := decoded[len(decoded)-4:]
    if 0 != bytes.Compare(actual_checksum, expected_checksum) {
        err = errors.New("Checksum failure")
        return
    }
    exponent := decoded[1:len(decoded)-4]
    if private_key, err = NewFromExponent(exponent); nil != err {
        return
    }
    private_key.Network = network_
    return
}
