package private

import "math/big"

type ECDSAPublicKey struct {
	X, Y *big.Int
}

func (pubKey ECDSAPublicKey) XBytes() []byte {
	return pubKey.X.Bytes()
}

func (pubKey ECDSAPublicKey) YBytes() []byte {
	return pubKey.Y.Bytes()
}
