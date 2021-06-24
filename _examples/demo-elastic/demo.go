package main

import (
	"fmt"
	"github.com/olivere/elastic/v7/config"
	"monitor"
	"strings"
)

func main() {
	fmt.Println(strings.Repeat("~", 37))
	fmt.Println("creating elastic session...")

	/*
		cfg, err := config.Unmarshal(url)
		if err != nil {
			fmt.Println(err)
			return
		}
	*/
	cfg := &config.Config{}
	session, err := monitor.NewEslasticSession(cfg)
	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}

	fmt.Println(strings.Repeat("~", 37))
	fmt.Println("Indexing document...")

	const (
		userID = "001"
		name   = "Bobby Axelrod"
		phone  = "001122334455"
	)

	view := monitor.NewView360(userID, name, phone)
	view.AddWhatsAppWebMessage("10", "Hello!", []string{"#Greetings"}, []*monitor.Entity{})

	err = session.Index(view.IndexName(), userID, view)
	if err != nil {
		fmt.Println("Error indexing:", err)
		return
	}

	fmt.Println(strings.Repeat("~", 37))
	fmt.Println("Updating document...")

	var view1 monitor.View360
	if j := session.Get(view1.IndexName(), userID); j != nil {
		if err := view1.Unmarshal(j); err == nil {
			view1.AddWhatsAppWebMessage("20", "bye", []string{"#Ending"}, []*monitor.Entity{})
			err = session.Index(view1.IndexName(), userID, view1)
			if err != nil {
				fmt.Println("Error indexing:", err)
				return
			}
		} else {
			fmt.Println("Error parsing:", err)
			return
		}
	} else {
		fmt.Println("Document not found:", userID)
		return
	}

	fmt.Printf("%s\n", strings.Repeat("~", 37))
	fmt.Println("Everything ran smoothly!")
}

/*
To delete run in Elasticsearch:
DELETE /360view/_doc/001
*/
