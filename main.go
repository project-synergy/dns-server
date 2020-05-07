package main

import (
	"server"
	"records"
	_ "fmt"
)

var APP *server.App
var req *server.Request
var res *server.Response

func sendMyDomain() bool {

	if req.Question.QType == "A" {

		if req.Question.Domain == "manushamil.com" {
			var record records.IRecord

			answer := server.Answer{}
			record = records.A {
				IPAddress : "34.93.61.81",
			}

			answer.SetName()
			answer.SetTYPE("A")
			answer.SetClass()
			answer.SetTTL(400)
			answer.SetRDATA(record.ToBytes())

			res.AddAnswer(answer)
			res.AddAnswer(answer)
		}
	} 


	if req.Question.QType == "MX" {

		if req.Question.Domain == "manushamil.com" {
			var record records.IRecord

			answer := server.Answer{}

			record = records.MX {
				Preference: 10,
				Exchange: "mail.google.com",
			}

			answer.SetName()
			answer.SetTYPE("MX")
			answer.SetClass()
			answer.SetTTL(400)
			answer.SetRDATA(record.ToBytes())

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
