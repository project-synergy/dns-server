package server

import (
	"fmt"
	"net"
	"log"
)

func listenTo(PORT int) {
	
	pc, err := net.ListenPacket("udp", fmt.Sprintf(":%d",PORT))
	if err != nil {
		log.Fatal(err)
	}

	defer pc.Close()

	fmt.Println("App running on PORT :", PORT)

	for {
		buf := make([]byte, 512)
		_, addr, err := pc.ReadFrom(buf)
		if err != nil {
			continue
		}
	
		/*
		*	Get Request and Build Response Here
		*/
		getRequest(&buf)
		buildResponse(&buf)

		/*
		*	Process the request to Response - Apply Middlewares etc..
		*/
		APP.ProcessRequest()

		/*
		*	Send back the response and reset Request and Response
		*/

		newBuf := make([]byte, 512)

		res.setHeader(&newBuf) //default Header
		res.setQuestion(&newBuf) // No modification Done
		bytesWritten := res.setAnswers(&newBuf)

		pc.WriteTo(newBuf[:bytesWritten], addr)	//send response

		APP.Clear() //Clear request and response

	}



}
