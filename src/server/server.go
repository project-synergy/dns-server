package server

import (
	_ "fmt"
	_ "net"
)

var APP App
var req *Request
var res *Response


type MiddleWare func() bool

type posRange struct {
	start int
	end int
}

type App struct {
	PORT int
	MiddleWares []MiddleWare
	Request Request
	Response Response
}

func (app *App) Listen(PORT int) {
	listenTo(PORT)
}

func (app *App) Use(MW MiddleWare) {
	app.MiddleWares = append( app.MiddleWares, MW)
}

func (app *App) Clear() {
	app.Request = Request{}
	app.Response = Response{}
}

func (app *App) ProcessRequest() {

	/*
	*	Run all Middlewares
	*/
	PROCEED := true

	for i := 0; i < len(app.MiddleWares); i++ {
		if PROCEED {
			PROCEED = app.MiddleWares[i]()	
		}
	}
	
	/*
	*	Send Response
	*/
	if PROCEED {
		//send only if last middleware returned true
		sendResponse()
	}


}

func sendResponse() {
	//Send App.Response back
}

func Server() *App {
	APP = App{}
	req = &APP.Request
	res = &APP.Response
	return &APP
}


