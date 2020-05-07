package server

import (
	_ "fmt"
	"bytes"
)

type Question struct {
	QNAME string
	QTYPE string
	QCLASS string

	//custom variables
	labels []string
	labelCount int

	//custom public variables
	Domain string
}

type posRange struct {
	start int
	end int
}


func getLabelIndeces(buf *[]byte, terminator int) []int {
	/*********************************************
	*	Returns Array of position which splits
	*	each label.
	*	Example : 
	*	On : .www.domain.com.
	*	Returns : [ 0, 4, 11, 15]
	**********************************************/

	positions := []int{}


	for i:=0; i<= terminator; i++ {
		positions = append(positions, i)

		if terminator - i <= 0 {
			break
		}

		i += int((*buf)[i])

	}

	return positions
}



func getRange( positions []int) []posRange{
	
	/*
	*	Convert this Positions to ranges
	*	On: [ 0, 4, 11, 15]
	*	Returns: [ { 1, 4 }, { 5, 10 }, { 12, 14 }]
	*/



	ranges := []posRange{}

	posCount := len(positions)

	for i:=0; i < posCount; i++ {

		if posCount - i <= 1 {
			break
		}

		a := positions[i] + 1
		b := positions[i + 1] - 1

		ranges = append(ranges, posRange{ a, b })

	}

	return ranges

}

func getLabels(buf *[]byte, terminator int) ([]string, int) {

	positions := getLabelIndeces(buf, terminator)

	ranges := getRange(positions)
	rangeLength := len(ranges)

	labels := make([][]byte, rangeLength)
	count := 0
	

	for i:=0; i< rangeLength; i++ {
		
		myRange := ranges[i]

		for j:=myRange.start; j <= myRange.end; j++ {
			labels[count] = append(labels[count], (*buf)[j])

		}

		count++
	}

	labelString := []string{}

	labelLength := len(labels)
	for i:=0; i< labelLength; i++ {

		labelString = append(labelString, string(labels[i]))
	}

	return labelString, labelLength
}

func getDomain(labels []string, labelCount int) string {

	domain := labels[labelCount - 2] + "." + labels[labelCount-1]

	return domain

}

func getQNAME(labels []string, labelCount int) string {
	
	qname := ""

	for i:=0; i<labelCount; i++ {
		qname += "." + labels[i]
	}


	return qname + "."
}

func getQTYPE(buf *[]byte) string {
	
	return getQTYPEvalues(buf)
}

func getQCLASS(buf *[]byte) string {
	return string(*buf)
}

func getQuestion(buf *[]byte) {


	terminator := bytes.IndexByte(*buf, 0x00)

	buffer := (*buf)[:terminator]					

	qtype := (*buf)[terminator + 1:terminator+3]
	qclass := (*buf)[terminator + 3:terminator + 5]




	/*
	* Get each labels 
	*/
	
	labels, labelCount := getLabels(&buffer, terminator)

	req.Question.Domain = getDomain(labels, labelCount)

	req.Question.QNAME = getQNAME(labels, labelCount) 

	req.Question.QTYPE = getQTYPE(&qtype)

	req.Question.QCLASS = getQCLASS(&qclass)

}