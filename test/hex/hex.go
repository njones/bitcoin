package hex

import hx "encoding/hex"

func MustDecode(s string) (buf []byte) {
	buf, err := hx.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return
}

func String(b []byte) string {
	return hx.EncodeToString(b)
}
