package whatsappconn

import (
	"github.com/Rhymen/go-whatsapp"
	"github.com/olivere/elastic/v7/config"
	"log"
	"monitor"
	"strings"
	watson "watsonconn"
)

var exists = struct{}{}

type messages map[string]struct{}

// History type
type History struct {
	es    *monitor.ElasticSession
	cache map[string]messages
}

// NewHistory creates a map of messages
func NewHistory() *History {

	session, err := monitor.NewEslasticSession(&config.Config{})
	if err != nil {
		log.Fatalln("Error creating elasticsearch session:", err)
	}

	cache := make(map[string]messages)
	return &History{cache: cache, es: session}
}

// AddToCache new whatsapp.TextMessage to cached history
func (h History) AddToCache(message whatsapp.TextMessage) {
	origin := message.Info.RemoteJid

	if _, ok := h.cache[origin]; !ok {
		h.cache[origin] = make(messages)
	}
	h.cache[origin][message.Info.Id] = exists
}

// IsCached returns TRUE if whatsapp.TextMessage is fund in cached history
func (h History) IsCached(message whatsapp.TextMessage) (present bool) {
	if messages, ok := h.cache[message.Info.RemoteJid]; ok {
		_, present = messages[message.Info.Id]
	}
	return
}

// Add new whatsapp.TextMessage to Metrics
func (h History) AddMetric(message whatsapp.TextMessage, response watson.WatsonResponse) {

	jid := message.Info.RemoteJid // {phone}@s.whatsapp.net
	phone := strings.Split(jid, "@")[0]
	messageID := message.Info.Id
	text := message.Text

	var intents []string
	for _, v := range response.Output.Intents {
		intents = append(intents, v.Intent)
	}

	var entities []*monitor.Entity
	for _, v := range response.Output.Entities {
		entities = append(entities, monitor.NewEntity(v.Entity, v.Value))
	}

	// document "360View"
	var view monitor.View360
	if j := h.es.Get(view.IndexName(), jid); j != nil {
		// document exists
		if err := view.Unmarshal(j); err != nil {
			log.Println("Record lost, error getting document:", err)
			return
		}
	} else {
		// new document
		vp := monitor.NewView360(jid, "anonymous", phone) // TODO: add user's name
		view = *vp
	}
	view.AddWhatsAppWebMessage(messageID, text, intents, entities)

	// Index document
	err := h.es.Index(view.IndexName(), jid, view)
	if err != nil {
		log.Println("Record lost, error indexing 360View:", err)
	}
}
