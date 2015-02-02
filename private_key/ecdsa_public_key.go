package private_key

import "math/big"

type ECDSAPublicKey struct {
	X, Y *big.Int
}
