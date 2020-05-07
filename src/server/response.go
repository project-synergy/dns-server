package server

import (
	_ "fmt"
	_ "strconv"
)

type Response struct {
	Header Header
	Question Question
	Answers []Answer

	bytesWritten int
}

func (res *Response) AddAnswer(answer Answer) {

	res.Answers = append(res.Answers, answer)	

	res.Header.setANCount(res.Header.ANCOUNT + 1)
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










/* 
*
*	Below functions does not need much modification.
*
*/

func (res *Response) setHeader(buf *[]byte) int {

	res.bytesWritten = 0

	header := res.Header.toBytes()

	for i:=0; i<12; i++ {
		res.bytesWritten += 1
		(*buf)[i] = header[i]
	}

	return res.bytesWritten
}

func (res *Response) setQuestion(buf *[]byte) int {
	/*
		Question must not be modified
	*/

	len, bytes := req.Question.toBytes()

	for i:=0; i<len; i++ {
		(*buf)[res.bytesWritten] = bytes[i]
		res.bytesWritten++
	}

	return res.bytesWritten

	//fmt.Println((*buf)[:res.bytesWritten])
	

	/*fmt.Println((*buf)[req.Question.byteRange.start:req.Question.byteRange.end])

	for i:= req.Question.byteRange.start; i< req.Question.byteRange.end; i++ {

		fmt.Println((*buf)[i])

	}*/

	//Leave it as it is.
}

func (res *Response) setAnswers(buf *[]byte) int {

	for i:=0; i< len(res.Answers); i++ {

		len, bytes := res.Answers[i].toBytes()

		for j:=0; j< len; j++ {
			(*buf)[res.bytesWritten] = bytes[j]
			res.bytesWritten++
		}

	}

	return res.bytesWritten
	
}