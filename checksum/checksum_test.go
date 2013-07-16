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

func testCompare(t *testing.T, a, b []byte, expected bool) {
    if expected != Compare(a, b) {
        t.Errorf("Should have failed")
        return
    }
}

func testCompareSuccess(t *testing.T, a, b []byte) {
    testCompare(t, a, b, true)
}

func testCompareFail(t *testing.T, a, b []byte) {
    testCompare(t, a, b, false)
}

func TestCompare(t *testing.T) {
    data1, _ := hex.DecodeString("00bb73ad508e67b211ecaba7d3d592ab46983d50dc")
    data2, _ := hex.DecodeString("01bb73ad508e67b211ecaba7d3d592ab46983d50dc")
    data3, _ := hex.DecodeString("00bb73ad508e67b211ecaba7d3d592ab46983d50de")
    data4, _ := hex.DecodeString("00bb73ad508e67b211ecaba7d3d592ab46983d50")
    data5, _ := hex.DecodeString("")

    testCompareFail(t, data1, data2)
    testCompareFail(t, data1, data3)
    testCompareFail(t, data1, data4)
    testCompareFail(t, data1, data5)
    testCompareSuccess(t, data1, data1)
    testCompareSuccess(t, data5, data5)
}
