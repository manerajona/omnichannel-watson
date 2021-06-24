package monitor

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	indexName = "360view"
)

// View360
type View360 struct {
	ID          string    `json:"idUser"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phoneNumber"`
	CreatedAt   time.Time `json:"createdAt"`
	Messages    []Message `json:"messages"`
}

// NewView360
func NewView360(id string, name string, phone string) *View360 {
	return &View360{
		ID:          id,
		Name:        name,
		PhoneNumber: phone,
		CreatedAt:   time.Now(),
		Messages:    []Message{},
	}
}

// IndexName of View360
func (v *View360) IndexName() string {
	return indexName
}

// Unmarshal json.RawMessage to View360
func (v *View360) Unmarshal(payload json.RawMessage) error {
	err := json.Unmarshal(payload, v)
	if err != nil {
		return fmt.Errorf("Cannot unmarshal to 360View due %v", err)
	}
	return nil
}

// Marshal View360 to json string
func (v *View360) Marshal() (string, error) {
	jsonList, err := json.Marshal(v)
	return string(jsonList), err
}

// AddWhatsAppWebMessage to View360
func (v *View360) AddWhatsAppWebMessage(id string, text string, intents []string, entities []*Entity) {
	v.Messages = append(v.Messages, *NewWhatsAppWebMessage(id, text, intents, entities))
}

// Message
type Message struct {
	ID        string       `json:"idMessage"`
	Text      string       `json:"text"`
	Intents   []string     `json:"intents"`
	Entities  []*Entity    `json:"entities"`
	Timestamp time.Time    `json:"timestamp"`
	Channel   *channelType `json:"channel"`
}

// NewWhatsAppWebMessage
func NewWhatsAppWebMessage(id string, text string, intents []string, entities []*Entity) *Message {
	return &Message{
		ID:        id,
		Text:      text,
		Intents:   intents,
		Entities:  entities,
		Timestamp: time.Now(),
		Channel:   whatsAppWeb,
	}
}

// Entity
type Entity struct {
	Entity string `json:"entity"`
	Value  string `json:"value"`
}

// NewEntity
func NewEntity(entity string, value string) *Entity {
	return &Entity{Entity: entity, Value: value}
}

var (
	whatsAppWeb = newChannelType("whatsappweb", "Whatsapp web")
)

type channelType struct {
	ChannelID string `json:"idChannel"`
	Name      string `json:"name"`
}

func newChannelType(channelID string, name string) *channelType {
	return &channelType{ChannelID: channelID, Name: name}
}
