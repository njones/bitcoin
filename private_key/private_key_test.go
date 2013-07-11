package private_key

import (
    "fmt"
    "bytes"
    "encoding/hex"
    "testing"
)


func testCase(t *testing.T, expected_encoded_pk string, expected_decoded_exponent []byte) {

    var err error
    var pk *PrivateKey
    
    if pk, err = NewFromExponent(expected_decoded_exponent); nil != err {
        t.Errorf("NewFromAddress failed, err=", err)
        return
    }
    var encoded_pk string
    if encoded_pk, err = pk.Encode(); nil != err {
        t.Errorf("PrivateKey Encode() failed, err=", err)
        return
    }

    if expected_encoded_pk != encoded_pk {
        t.Errorf("Encode failed: expected = %s, actual = %s", expected_encoded_pk, encoded_pk)
        return
    }

    pk, err = Decode(expected_encoded_pk)
    if nil != err {
        t.Errorf("Decode error %s", err)
        return
    }

    if 0 != bytes.Compare(pk.PrivateKey.D.Bytes(), expected_decoded_exponent) {
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
        private_key, err := Generate()
        if nil != err {
            panic(fmt.Sprint("Could not generate bitcoin private key, err =", err))
        }
        address, err2 := private_key.PublicAddressString()
        if nil != err2 {
            panic(fmt.Sprint("Could not generate public address from private key, err = ", err))

        }
        wif, _ := private_key.Encode()
        fmt.Println("WIF pk  = ", wif)
        fmt.Println("address = ", address)
    }
}

func BenchmarkSetKey(b *testing.B) {
    exponent, _ := hex.DecodeString("594cfd670ac0453236816deec9ee2a4924c58e95fba992472846c1000e0adf80")
    for i := 0; i < b.N; i++ {
        private_key, err := NewFromExponent(exponent)
        if nil != err || nil == private_key {
            panic(fmt.Sprint("Could not generate bitcoin private key, err =", err))
        }
    }
}
