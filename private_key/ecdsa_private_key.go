package private_key

import (
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"github.com/sour-is/koblitz/kelliptic"
)

const ExponentSize = 32 // bytes

var s256 = kelliptic.S256()

type ECDSAPrivateKey struct {
	PublicKey ECDSAPublicKey
	priv      []byte
}

func NewECDSAPrivateKey() (pk ECDSAPrivateKey, err error) {
	pk.priv, pk.PublicKey.X, pk.PublicKey.Y, err = elliptic.GenerateKey(s256, rand.Reader)
	return
}

func NewECDSAPrivateKeyFromExponent(exponent []byte) (pk *ECDSAPrivateKey, err error) {
	if len(exponent) != ExponentSize {
		return nil, fmt.Errorf("Exponent of wrong size, expected: %d, actual: %d", ExponentSize, len(exponent))
	}
	pk = &ECDSAPrivateKey{}
	pk.PublicKey.X, pk.PublicKey.Y = elliptic.Unmarshal(s256, exponent)
	pk.priv = make([]byte, ExponentSize, ExponentSize)
	copy(pk.priv, exponent)
	return
}
