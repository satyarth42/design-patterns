package appendix

import "fmt"

type Request struct {
	reqType string
}

type IHandler interface {
	setNext(handler IHandler)
	handleRequest(req Request)
}

type BaseHandler struct {
	successor IHandler
}
func (b *BaseHandler) setNext(handler IHandler) {
	b.successor = handler
}

type HandleAll struct {
	BaseHandler
}
func (h *HandleAll) handleRequest(req Request) {
	fmt.Println("no handler found for:",req.reqType)
}

type HandleFanMail struct {
	BaseHandler
}
func (h *HandleFanMail) handleRequest(req Request) {
	switch req.reqType {
	case "fan":
		fmt.Println("fan mail")
	default:
		if h.BaseHandler.successor==nil {
			fmt.Println("no successor")
		} else {
			h.BaseHandler.successor.handleRequest(req)
		}
	}
}

type HandleSpam struct {
	BaseHandler
}
func (h *HandleSpam) handleRequest(req Request) {
	switch req.reqType {
	case "spam":
		fmt.Println("fan mail")
	default:
		if h.BaseHandler.successor==nil {
			fmt.Println("no successor")
		} else {
			h.BaseHandler.successor.handleRequest(req)
		}
	}
}

func GetCOR() IHandler {
	fanMailHandler := &HandleFanMail{}
	spamHandler := &HandleSpam{}
	handleAll := &HandleAll{}

	spamHandler.setNext(handleAll)
	fanMailHandler.setNext(spamHandler)

	return fanMailHandler
}