package main

import (
	"server"
	_ "fmt"
	"strings"
	"strconv"
)

var APP *server.App
var req *server.Request
var res *server.Response

type A struct {
	record [4]byte
}

func sendMyDomain() bool {

	answer := server.Answer{}
	
	if req.Question.QRecord == "A" {

		if req.Question.Domain == "manushamil.com" {
			

			bytes := []byte{}

			record := "34.93.61.81"

			parts := strings.Split(record, ".")

			for i:=0; i<len(parts); i++ {
				part, _ := strconv.ParseUint(parts[i], 10, 8)

				bytes = append(bytes, byte(part))
			}

			answer.SetName()
			answer.SetTYPE("A")
			answer.SetClass()
			answer.SetTTL(400)
			answer.SetRDATA(bytes)

			res.AddAnswer(answer)
			res.AddAnswer(answer)
		}
	}

	return true
}


func main() {

	APP = server.Server()
	req = &APP.Request
	res = &APP.Response

	APP.Use(sendMyDomain)

	APP.Listen(53)
}
