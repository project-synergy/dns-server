package records

import(
	"strconv"
	"strings"
)

type A struct {
	record [4]byte

	IPAddress string
}

func (a A) ToBytes() []byte {

	bytes := make([]byte, 4)

	parts := strings.Split(a.IPAddress, ".")

	for i:=0; i<len(parts); i++ {
		part, _ := strconv.ParseUint(parts[i], 10, 8)

		bytes[i] = byte(part)
	}

	return bytes

}