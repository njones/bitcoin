package checksum

import "crypto/sha256"

func Checksum(data []byte) (checksum []byte) {
    first_round := sha256.New()
    first_round.Write(data)
    second_round := sha256.New()
    second_round.Write(first_round.Sum(nil))
    checksum = second_round.Sum(nil)[:4]
    return
}

func Compare(checksum1, checksum2 []byte) bool {
    if len(checksum1) != len(checksum2) {
        return false
    }
    result := byte(0)

    for i := 0; i < len(checksum1); i++ {
        result |= checksum1[i] ^ checksum2[i]
    }

    return (result == 0)
}
