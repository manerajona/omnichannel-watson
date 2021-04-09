package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	watson "watsonconn"
	waconn "whatsappconn"
)

const path = "./resources/watson.yaml"

var (
	cfg     watson.Config
	history = waconn.NewHistory()
)

func init() {
	// read watson.yaml from resources
	if err := cfg.ParseFromYAML(path); err != nil {
		log.Panicf("Error parsing %s : %v\n", path, err)
	}
}

func main() {

	// Create watson connection
	wac := waconn.New()

	// Add history handler
	handler := &waconn.HistoryHandler{
		Conn: wac,
		Hist: history,
	}

	wac.AddHandler(handler)
	waconn.LoginOrRestore(wac)
	waconn.Disconect(wac)
	wac.RemoveHandler(handler)

	// Add watson handler
	wac.AddHandler(&waconn.WatsonHandler{
		Conn:      wac,
		WatsonCfg: &cfg,
		Hist:      history,
	})

	waconn.LoginOrRestore(wac)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	waconn.Disconect(wac)
}
