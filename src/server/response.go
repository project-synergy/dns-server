package server

import (
	_ "fmt"
	_ "strconv"
)

type Response struct {
	Header Header
}


func buildResponse(buf *[]byte) {

	// Set TransactionID
	res.Header.setTransactionID()

	// Set Flags
	res.Header.Flags.setQR(1)
	res.Header.Flags.setOPCODE(req.Header.Flags.OPCODE)
	res.Header.Flags.setAA(1)
	res.Header.Flags.setTC(0)
	res.Header.Flags.setRD(0)
	res.Header.Flags.setRA(0)
	res.Header.Flags.setZ(req.Header.Flags.Z)
	res.Header.Flags.setRD(0)
	res.Header.Flags.setRCODE([4]byte{0, 0, 0, 0,})

	// Set QDCount
	res.Header.setQDCount(1)
	res.Header.setANCount(0)
	res.Header.setNSCount(0)
	res.Header.setARCount(0)

}

func (res *Response) setHeader(buf *[]byte) {

	header := res.Header.toBytes()

	for i:=0; i<12; i++ {
		(*buf)[i] = header[i]
	}
}

func (res *Response) setQuestion(buf *byte) {
	/*
		Question must not be modified
	*/

	//Leave it as it is.

}