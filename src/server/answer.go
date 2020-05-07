package server

import (
	_ "fmt"
)

type Answer struct {
	Name [2]byte
	Type [2]byte
	Class [2]byte
	TTL [4]byte
	RDLength [2]byte
	RData []byte

	//custom variables
	rdLength int
}

func (ans *Answer) SetName() {
	ans.Name = [2]byte{ 0xc0, 0x0c } // default( [ 0xc0, 0x0c ])
}

func (ans *Answer) SetTYPE(recordSTR string) {

	var record uint8

	switch recordSTR {
		case "A": record = 1
		case "NS": record = 2
		case "MD": record = 3
		case "MF": record = 4
		case "CNAME": record = 5
		case "SOA": record = 6
		case "MB": record = 7
		case "MG": record = 8
		case "MR": record = 9
		case "NULL": record = 10
		case "WKS": record = 11
		case "PTR": record = 12
		case "HINFO": record = 13
		case "MINFO": record = 14
		case "MX": record = 15
		case "TXT": record = 16
		default: record = 0
	}

	ans.Type[0] = 0x0 // 1st byte doesn't get used ever, since the highest number 16 can be fit inside the 2nd byte alone.
	ans.Type[1] = record
	
}

func (ans *Answer) SetClass() {
	ans.Class = [2]byte{ 0x00, 0x01 } // default( IN )
}

func (ans *Answer) SetTTL(ttl uint64) {

	ans.TTL[0] = byte(ttl >> 24 & 0xFF)
	ans.TTL[1] = byte(ttl >> 16 & 0xFF)
	ans.TTL[2] = byte(ttl >> 8 & 0xFF)
	ans.TTL[3] = byte(ttl & 0xFF)

}

func (ans *Answer) SetRDATA(data []byte) {
	ans.rdLength = len(data)
	length := uint16(ans.rdLength)

	//fmt.Println(length)

	ans.RDLength[0] = byte((length >> 8) & 0xFF)
	ans.RDLength[1] = byte(length & 0xFF)

	//fmt.Println(ans.RDLength)


	ans.RData = data
}

func (ans *Answer) toBytes() (int, []byte) {
	length := 12 + ans.rdLength

	bytes := make([]byte, length)

	/*
	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
    |                                               |
    /                                               /
    /                      NAME                     /
    |                                               |
	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
	*/

	bytes[0] = ans.Name[0]
	bytes[1] = ans.Name[1]

	/*
	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
    |                      TYPE                     |
    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
	*/
	bytes[2] = ans.Type[0]
	bytes[3] = ans.Type[1]

	/*
	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
    |                     CLASS                     |
    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
	*/
	bytes[4] = ans.Class[0]
	bytes[5] = ans.Class[1]

	/*
	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
    |                      TTL                      |
    |                                               |
	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
	*/
	bytes[6] = ans.TTL[0]
	bytes[7] = ans.TTL[1]
	bytes[8] = ans.TTL[2]
	bytes[9] = ans.TTL[3]

	/*
	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
    |                   RDLENGTH                    |
    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--|
	*/
	bytes[10] = ans.RDLength[0]
	bytes[11] = ans.RDLength[1]

	/*
	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--|
    /                     RDATA                     /
    /                                               /
	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
	*/

	for i:=0; i<ans.rdLength; i++ {
		bytes[12 + i] = ans.RData[i]
	}

	return length, bytes

}