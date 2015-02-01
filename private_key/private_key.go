package private_key

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"golang.org/x/crypto/ripemd160"
	"math/big"
	//"github.com/steakknife/Golang-Koblitz-elliptic-curve-DSA-library/bitelliptic"
	//    "github.com/steakknife/Golang-Koblitz-elliptic-curve-DSA-library/bitecdsa"
	"github.com/sour-is/koblitz/kelliptic"
	"github.com/steakknife/bitcoin/base58"
	"github.com/steakknife/bitcoin/checksum"
	"github.com/steakknife/bitcoin/network"
	"github.com/steakknife/bitcoin/public_key"
)

type PrivateKey struct {
	*network.Network
	*bitecdsa.PrivateKey
}

const (
	ExponentSize      = 32 // bytes
	MainAddressPrefix = byte(0x80)
	TestAddressPrefix = byte(0xEF)
)

func Generate() (private_key *PrivateKey, err error) {
	ecdsa_private_key, err := bitecdsa.GenerateKey(bitelliptic.S256(), rand.Reader)
	if err != nil {
		return
	}
	private_key_ := &PrivateKey{
		Network:    network.Main,
		PrivateKey: ecdsa_private_key,
	}
	private_key_encoded, err := private_key_.Encode()
	if err != nil {
		return
	}
	private_key_decoded, err := Decode(private_key_encoded)
	if err != nil || private_key_decoded == nil {
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

	private_key = &PrivateKey{
		Network:    network.Main,
		PrivateKey: bitecdsa.SetKey(bitelliptic.S256(), new(big.Int).SetBytes(exponent)),
	}
	return
}

func (private_key *PrivateKey) PublicKey() *public_key.PublicKey {
	return &public_key.PublicKey{
		Network: private_key.Network,
		Address: private_key.PublicAddress(),
	}
}

func PadToSize(array []byte, size int) []byte {
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
	switch private_key.Network {
	case network.Main:
		address_prefix = MainAddressPrefix
	case network.Test:
		address_prefix = TestAddressPrefix
	default:
		err = errors.New("Unknown NetworkID")
	}
	return
}

func DecodeAddressPrefix(address_prefix byte) (network_ *network.Network, err error) {
	switch address_prefix {
	case MainAddressPrefix:
		network_ = network.Main
	case TestAddressPrefix:
		network_ = network.Test
	default:
		err = errors.New("Unknown Network Address Prefix")
	}
	return
}

func (private_key *PrivateKey) Encode() (result string, err error) {
	address_prefix, err := private_key.AddressPrefix()
	if err != nil {
		return
	}
	exponent := PadToSize(private_key.PrivateKey.D.Bytes(), 32)
	data := append([]byte{address_prefix}, exponent...)
	data = append(data, checksum.Checksum(data)...)

	result = base58.Encode(data)
	return
}

func Decode(encoded string) (private_key *PrivateKey, err error) {
	decoded, err := base58.Decode(encoded)
	if err != nil {
		return
	}
	network_, err := DecodeAddressPrefix(decoded[0])
	if err != nil {
		return
	}
	actual_checksum := checksum.Checksum(decoded[:len(decoded)-4])
	expected_checksum := decoded[len(decoded)-4:]
	if bytes.Compare(actual_checksum, expected_checksum) != 0 {
		err = errors.New("Checksum failure")
		return
	}
	exponent := decoded[1 : len(decoded)-4]
	private_key, err = NewFromExponent(exponent)
	if err != nil {
		return
	}
	private_key.Network = network_
	return
}
