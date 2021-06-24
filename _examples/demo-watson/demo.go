package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"time"
	watson "watsonconn"
)

func main() {

	var (
		ws          watson.Session
		region      string
		instances   string
		assistants  string
		versions    string
		credentials string
	)

	// flags declaration
	flag.StringVar(&region, "r", "us-south", "Specify region zone")
	flag.StringVar(&instances, "i", "", "Specify ibm cloud instance")
	flag.StringVar(&assistants, "a", "", "Specify watson assistant")
	flag.StringVar(&versions, "v", time.Now().Format("2006-01-02"), "Specify watson version")
	flag.StringVar(&credentials, "p", "", "Specify credentials")

	flag.Parse()

	ws, err := watson.New(assistants, instances, versions, region, credentials)
	if err != nil {
		log.Printf("Cannot create session with params -r '%s' (region) -i '%s' (instance) -a '%s' (assistant) -v '%s' (version) -p '%s' (credentials)\n",
			region, instances, assistants, versions, credentials)

		log.Panicf("Error response: %v", err)
	}

	// conversation
	input := bufio.NewScanner(os.Stdin)
	log.Println("Watson assistant is ready, write something.")
	for input.Scan() {
		response, err := ws.SendStatefulMessage(input.Text())
		if err != nil || len(response.Output.Generic) == 0 {
			log.Println("Something went wrong: ", err)
			return
		}
		for _, v := range response.Output.Generic {
			log.Println("Watson replied:\n", v.Text)
		}
	}
}
