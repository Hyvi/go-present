package byte2string

import (
	"fmt"
)

func B2S() {
	var bys [20]byte

	bys[0] = 'h'
	bys[1] = 'e'
	bys[2] = 'i'
	// expain not equal to "hei", because bys[3] is not 0, is 0x00, so the string is "hei\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00"
	// eplain bys[3] is not 0, because bys is a array, and bys is a value type, so bys[3] is not assigned, so it is 0x00
	if string(bys[:]) == "hei" {
		fmt.Println("[20]byte('h','e','i') === \"hei\"")
	}
}
