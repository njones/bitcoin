package checksum

import (
    "bytes"
    "encoding/hex"
    "testing"
)

func testCase(t *testing.T, data []byte, expected_checksum []byte) {
    checksum := Checksum(data)
    if 0 != bytes.Compare(checksum, expected_checksum) {
        t.Errorf("Checksum testcase failed: expected = %s, actual = %s", expected_checksum, checksum)
        return
    }
}

func TestChecksum0(t *testing.T) {
    data, _ := hex.DecodeString("00bb73ad508e67b211ecaba7d3d592ab46983d50dc")
    checksum, _ := hex.DecodeString("33164a23")
    testCase(t, data, checksum)
}
