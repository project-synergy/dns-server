package server

import (
	"strconv"
	"encoding/hex"
	_ "fmt"
	_ "bytes"
)

type Request struct {
	Header Header
	Question Question

	/*
	*	Custom Variables
	*/
	FullDomain string
	Domain string
	Extension string
}

func getTransactionID(buf []byte) uint16 {

	id, _ := strconv.ParseUint(hex.EncodeToString(buf), 16, 16)

	return uint16(id)
}

func getFlags(buf []byte) Flags{

	/*
	*	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
	*	|QR|   Opcode  |AA|TC|RD|RA|   Z    |   RCODE   |
	*	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
	*/
	QR := ((buf[0] >> 7) & 0x1)
	//fmt.Println("QR : ", req.Flags.QR)

	OPCODE := [4]byte{}

	j := 0
	for i := 7; i >= 4; i-- {
		OPCODE[j] = (buf[0] >> i) & 0x1
		j++
	}

	//fmt.Println("OPCODE : ", req.Flags.OPCODE)

	AA := buf[0] >> 2 & 0x1

	//fmt.Println("AA : ", req.Flags.AA)

	TC := buf[0] >> 1 & 0x1
	//fmt.Println("TC : ", req.Flags.TC)

	RD := (buf[0] & 0x1)
	//fmt.Println("RD : ", req.Flags.RD)

	RA := buf[1] >> 7 & 0xF

	//fmt.Println("RA : ", req.Flags.RA)

	Z := [3]byte{}
	j = 0
	for i := 7; i >= 5; i-- {
		Z[j] = (buf[0] >> i) & 0x1
		j++
	}

	//fmt.Println("Z : ", req.Flags.Z)

	RCODE := [4]byte{}

	j = 0
	for i := 3; i >= 0; i-- {
		RCODE[j] = (buf[0] >> i) & 0x1
		j++
	}

	//fmt.Println("RCODE : ", req.Flags.RCODE)

	return Flags{
		QR: QR,
		OPCODE: OPCODE,
		AA : AA,
		TC: TC,
		RD: RD,
		RA: RA,
		Z: Z,
		RCODE: RCODE,
	}
	
}

func getQuestionCount(buf []byte) uint16 {
	/*
	*	    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
	*		|                    QDCOUNT                    |
	*		+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
	*/

	question_count, _ := strconv.ParseUint(hex.EncodeToString(buf), 16, 16)

	return uint16(question_count)

	//fmt.Println("QDCOUNT : ", req.QDCOUNT)
}

func getAnswerCount(buf []byte) uint16 {
	/*
	*	    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
	*		|                    ANCOUNT                    |
	*		+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
	*/

	answer_count, _ := strconv.ParseUint(hex.EncodeToString(buf), 16, 16)

	return uint16(answer_count)

	//fmt.Println("ANCOUNT : ", req.ANCOUNT)
}

func getAuthorityRecordsServerCount(buf []byte) uint16 {
	/*
	*	    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
	*		|                    NSCOUNT                    |
	*		+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
	*/

	ns_count, _ := strconv.ParseUint(hex.EncodeToString(buf), 16, 16)

	return uint16(ns_count)

	//fmt.Println("NSCOUNT : ", req.NSCOUNT)
}

func getAdditionalRecordsCount(buf []byte) uint16 {
	/*
	*	    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
	*		|                    ARCOUNT                    |
	*		+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
	*/

	ar_count, _ := strconv.ParseUint(hex.EncodeToString(buf), 16, 16)

	return uint16(ar_count)

	//fmt.Println("ARCOUNT : ", req.ARCOUNT)
}

func getHeader(buf *[]byte) {

	req.Header.byteRange = posRange{ 0, 12 } //default

	req.Header.TransactionID = getTransactionID((*buf)[0:2])

	req.Header.Flags = getFlags((*buf)[2:4])

	req.Header.QDCOUNT = getQuestionCount((*buf)[4:6])

	req.Header.ANCOUNT = getAnswerCount((*buf)[6:8])

	req.Header.NSCOUNT = getAuthorityRecordsServerCount((*buf)[8:10])

	req.Header.ARCOUNT = getAdditionalRecordsCount((*buf)[10:12])

}


func getRequest(buf *[]byte) {

	header := (*buf)[:12]
	getHeader(&header)

	question := (*buf)[12:]
	getQuestion(&question)

}