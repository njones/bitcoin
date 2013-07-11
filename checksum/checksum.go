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
