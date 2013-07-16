package public_key

import (
    "bytes"
    "encoding/hex"
    "testing"
)

 // 20 + 4
 // 0123456789012345678901234567890123456789
 //           1         2         3
// "bb73ad508e67b211ecaba7d3d592ab46983d50dc 33164a23"
// "bb73ad508e67b211ecaba7d3d592ab46983d50dc"


func testCase(t *testing.T, expected_encoded_address string, expected_decoded_address []byte) {
    
    var err error
    var address *PublicKey
    if address, err = NewFromAddress(expected_decoded_address); nil != err {
        t.Errorf("NewFromAddress failed, err=", err)
        return
    }
    var encoded_address string
    if encoded_address, err = address.Encode(); nil != err {
        t.Errorf("PublicKey Encode() failed, err=", err)
        return
    }

    if expected_encoded_address != encoded_address {
        t.Errorf("Encode failed: expected = %s, actual = %s", expected_encoded_address, encoded_address)
        return
    }

    decoded_address, err := Decode(expected_encoded_address)
    if nil != err {
        t.Errorf("Decode error %s", err)
        return
    }

    if 0 != bytes.Compare(decoded_address.Address, expected_decoded_address) {
        t.Errorf("Decode failed.  expected = %s, actual = %s", expected_decoded_address, decoded_address.Address)
        return
    }
}

func Test0(t *testing.T) {
    address, _ := hex.DecodeString("bb73ad508e67b211ecaba7d3d592ab46983d50dc")
    encoded_address := "1J69wFcKYkABMsTyLyAxJmUnpRe5kiGWHg"
    testCase(t, encoded_address, address)
}
