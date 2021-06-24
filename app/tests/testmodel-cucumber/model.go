package main

import (
	"log"
	watson "watsonconn"
)

var (
	ws  watson.Session
	cfg watson.Config
)

const (
	path = "../../resources/watson.yaml"
)

func init() {
	// read watson.yaml from resources
	err := cfg.ParseFromYAML(path)
	if err != nil {
		log.Panicf("Error parsing %s : %v\n", path, err)
	}

	ws, err = watson.NewSession(&cfg)
	if err != nil {
		log.Printf("Cannot create session with params -r '%s' (region) -i '%s' (instance) -a '%s' (assistant) -v '%s' (version) -p '%s' (credentials)\n",
			cfg.Region, cfg.Instance, cfg.AssistantID, cfg.Version, cfg.Credentials)

		log.Panicf("Error response: %v", err)
	}
}

func askToWatson(input string) (r watson.WatsonResponse) {
	r, err := ws.SendStatefulMessage(input)
	if err != nil {
		log.Println("Something went wrong with Watson")
	}
	return
}

func main() { /* usual main func */ }
