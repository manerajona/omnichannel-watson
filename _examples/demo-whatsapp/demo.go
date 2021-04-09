package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	waconn "whatsappconn"

	"github.com/Rhymen/go-whatsapp"
)

type receiverHandler struct {
	Conn *whatsapp.Conn
}

func (*receiverHandler) HandleError(err error) {
	log.Printf("error occoured: %v\n", err)
}

func (*receiverHandler) HandleTextMessage(message whatsapp.TextMessage) {
	log.Printf("Received|id=%s|origin=%s|status=%d|text:\n%s\n\n",
		message.Info.Id, message.Info.RemoteJid, message.Info.Status, message.Text)
}

func main() {

	wac := waconn.New()

	//Add handler
	wac.AddHandler(&receiverHandler{wac})

	//login or restore your WhatsApp connection
	waconn.LoginOrRestore(wac)

	// Wait interruption (Ctrl + c)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	//Disconnect safe
	waconn.Disconect(wac)
}
