package base58

import (
    "math/big"
    "fmt"
)

/*
n   b58 n   b58 n   b58 n   b58
0   1   1   2   2   3   3   4
4   5   5   6   6   7   7   8
8   9   9   A   10  B   11  C
12  D   13  E   14  F   15  G
16  H   17  J   18  K   19  L
20  M   21  N   22  P   23  Q
24  R   25  S   26  T   27  U
28  V   29  W   30  X   31  Y
32  Z   33  a   34  b   35  c
36  d   37  e   38  f   39  g
40  h   41  i   42  j   43  k
44  m   45  n   46  o   47  p
48  q   49  r   50  s   51  t
52  u   53  v   54  w   55  x
56  y   57  z 
*/

const base58EncodeString = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

var (
    base58DecodeArray = []byte {
        // 48 .. 63
          0xff,  0,  1,  2,  3,  4,  5,  6,  7,  8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
        // 64 .. 79
          0xff,  9, 10, 11, 12, 13, 14, 15, 16, 0xff, 17, 18, 19, 20, 21, 0xff,
        // 80 .. 95
            22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 0xff, 0xff, 0xff, 0xff, 0xff,
        // 96 .. 111
          0xff, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 0xff, 44, 45, 46,
        // 112 ... 127
            47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 0xff, 0xff, 0xff, 0xff, 0xff, 
    }


    _base58 = func() *big.Int {
        b58 := new(big.Int)
        b58.SetBytes([]byte{58})
        return b58
    }()
)

func Encode(data []byte) (result string) {
    zero := big.NewInt(0)
    x := new(big.Int)
    x.SetBytes(data)
    remainder := big.NewInt(0)
    result = ""

    for x.Cmp(zero) > 0 {
        x.DivMod(x, _base58, remainder)
        encoded := string(base58EncodeString[remainder.Int64()])
        result = fmt.Sprint(encoded, result)
    }

    for i := 0; i < len(data) && 0 == data[i]; i++ {
        result = fmt.Sprint("1", result)
    }

    return
}

func Decode(encoded string) (result []byte, err error) {
    if len(encoded) == 0 {
        err = fmt.Errorf("Cannot decode empty string")
        return
    }
    pad_bytes := 0
    var i int
    for i = 0; i < len(encoded) && "1" == string(encoded[i]); i++ {
        pad_bytes++
    }
    sum := big.NewInt(0)
    var decoded byte

    for i, _ := range encoded {
        encoded_ascii := byte(encoded[i])
        if encoded_ascii >= 48 && encoded_ascii <= 127 {
            decoded = base58DecodeArray[encoded_ascii - 48]
        } else {
            decoded = byte(0xff)
        }
        if decoded == 0xff {
            err = fmt.Errorf("Bad character encountered")
            return
        }
        sum.Add(sum.Mul(sum, _base58), big.NewInt(int64(decoded)))
    }
    b := sum.Bytes()
    result = make([]byte, pad_bytes + len(b))
    for i := 0; i < pad_bytes; i++ {
        result = append(result, 0)
    }
    result = append(result, b...)
    return
}
