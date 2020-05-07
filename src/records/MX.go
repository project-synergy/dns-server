package records

import (
	_"fmt"
	"strings"
)

type MX struct {
	preference [2]byte
	exchange []byte

	Preference int
	Exchange string
}

func (mx MX) ToBytes() []byte {

	bytes := []byte{}
	written := 0

	bytes = append(bytes, byte(mx.Preference >> 8 & 0xF))
	bytes = append(bytes, byte(mx.Preference & 0xF))
	written += 2

	parts := strings.Split(mx.Exchange, ".")

	for _, part:= range parts {

		stringBytes := []byte(part)
		length := len(part)

		bytes = append(bytes, byte(length))
		written += 1

		for i:=0; i<length; i++ {
			bytes = append(bytes, stringBytes[i])
			written += 1
		}
	}

	bytes = append(bytes, 0x0)
	written++


	//fmt.Println(bytes)
	
	return bytes
}