package private_key

import (
    "fmt"
    "bytes"
    "encoding/hex"
    "testing"
)


func testCase(t *testing.T, expected_encoded_pk string, expected_decoded_exponent []byte) {
    pk, err := NewFromExponent(expected_decoded_exponent)
    if err != nil {
        t.Errorf("NewFromAddress failed, err=", err)
        return
    }
    encoded_pk, err := pk.Encode()
    if err != nil {
        t.Errorf("PrivateKey Encode() failed, err=", err)
        return
    }

    if expected_encoded_pk != encoded_pk {
        t.Errorf("Encode failed: expected = %s, actual = %s", expected_encoded_pk, encoded_pk)
        return
    }

    pk, err = Decode(expected_encoded_pk)
    if err != nil {
        t.Errorf("Decode error %s", err)
        return
    }

    if bytes.Compare(pk.PrivateKey.D.Bytes(), expected_decoded_exponent) != 0 {
        t.Errorf("Decode failed.  expected = %s, actual = %s", expected_decoded_exponent, pk.PrivateKey.D.Bytes())
        return
    }
}

func Test0(t *testing.T) {
    exponent, _ := hex.DecodeString("594cfd670ac0453236816deec9ee2a4924c58e95fba992472846c1000e0adf80")
    pk := "5JVcebpmNXH4eekGXpYLL526G8x7fdm9yt6TtxsCUkyXWpkM58j"
    testCase(t, pk, exponent)
}

func BenchmarkGenerate(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _, err := Generate()
        if err != nil {
            b.Errorf(fmt.Sprint("Could not generate bitcoin private key, err =", err))
        }
    }
}

func BenchmarkSetKey(b *testing.B) {
    exponent, _ := hex.DecodeString("594cfd670ac0453236816deec9ee2a4924c58e95fba992472846c1000e0adf80")
    for i := 0; i < b.N; i++ {
        private_key, err := NewFromExponent(exponent)
        if err != nil || private_key == nil {
            b.Errorf(fmt.Sprint("Could not generate bitcoin private key, err =", err))
        }
    }
}
