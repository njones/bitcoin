package base58

import (
    "bytes"
    "encoding/hex"
    "testing"
)

func testCase(t *testing.T, expected_encoded string, expected_decoded []byte) {
    encoded := Encode(expected_decoded)
    if expected_encoded != encoded {
        t.Errorf("Encode failed.  expected = %s, actual = %s", expected_encoded, encoded)
        return
    }
    decoded, err := Decode(expected_encoded)
    if nil != err {
        t.Errorf("Decode error %s", err)
        return
    }
    if 0 != bytes.Compare(expected_decoded, decoded) {
        t.Errorf("Decode failed.  expected = %s, actual = %s", expected_decoded, decoded)
        return
    }
}

func Test0(t *testing.T) {
    data, _ := hex.DecodeString("8018E14A7B6A307F426A94F8114701E7C8E774E7F9A47E2C2035DB29A206321725d91ea8a6")
    encoded := "5J1F7GHadZG3sCCKHCwg8Jvys9xUbFsjLnGec4H125Ny1V9nR6V"
    testCase(t, encoded, data)
}

func Test1(t *testing.T) {
    data, _ := hex.DecodeString("00000000")
    encoded := "1111"
    testCase(t, encoded, data)
}

func Test2(t *testing.T) {
    data, _ := hex.DecodeString("00000000FF")
    encoded := "11115Q"
    testCase(t, encoded, data)
}
