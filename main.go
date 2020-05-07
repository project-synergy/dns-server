package main

import (
	"server"
	"fmt"
)

var APP *server.App
var req *server.Request
var res *server.Response

func sendMyDomain() bool {
	fmt.Println(req.Question.Domain)

	if req.Question.Domain == "manushamil.com" {
		fmt.Println("My Domain")
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
