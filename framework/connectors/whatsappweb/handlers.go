package whatsappconn

import (
	"errors"
	whatsapp "github.com/Rhymen/go-whatsapp"
	"log"
	"time"
	watson "watsonconn"
)

/*
WatsonHandler:

This type handles recieved messages through the watson assistant api.
*/
type WatsonHandler struct {
	Conn      *whatsapp.Conn
	WatsonCfg *watson.Config
	Hist      *History
}

// HandleError method implemented to be a valid WhatsApp handler.
func (h *WatsonHandler) HandleError(err error) {
	if e, ok := err.(*whatsapp.ErrConnectionFailed); ok {
		log.Printf("Connection failed, underlying error: %v\n", e.Err)
		<-time.After(30 * time.Second) // Retry after 30s
		log.Println("Reconnecting...")
		err := h.Conn.Restore()
		if err != nil {
			log.Fatalf("Restore failed: %v", err)
		}
	} else {
		log.Printf("An error occoured: %v\n", err)
	}
}

var sessions = make(map[string]*watson.Session) // TODO: refactor this

// HandleTextMessage method is where responses are handled by watson.
func (h *WatsonHandler) HandleTextMessage(message whatsapp.TextMessage) {

	if message.Info.Status == 0 {
		// Async
		go func(mge whatsapp.TextMessage) {
			// Only if is a new message
			if !h.Hist.IsCached(message) {

				mgeInID := mge.Info.Id
				origin := mge.Info.RemoteJid
				input := mge.Text

				watsonSession, err := h.getOrCreateWatsonSession(origin)
				if err != nil {
					log.Fatalf("ERROR|Watson|%v\n", err)
				}

				response, err := watsonSession.SendStatefulMessage(input)
				if err != nil {
					log.Printf("ERROR|Watson|%v\n", err)
					// delete session and try again...
					delete(sessions, origin)
					h.HandleTextMessage(mge)
					return
				}

				log.Printf("RECEIVED|id=%s|origin=%s|text:\n%s\n\n", mgeInID, origin, input)

				for _, output := range response.Output.Generic {
					mgeOutID, err := h.sendToDest(origin, output.Text)
					if err == nil {
						log.Printf("SENT|id=%s|dest=%s|text:\n%s\n\n", mgeOutID, origin, output.Text)
					} else {
						log.Printf("ERROR|Whatsapp|%v\n", err)
					}
				}
				// Async
				go h.Hist.AddMetric(message, response)
			}
		}(message)
	}
}

func (h *WatsonHandler) sendToDest(dest string, input string) (string, error) {

	if dest != "" && input != "" {
		text := whatsapp.TextMessage{
			Info: whatsapp.MessageInfo{
				RemoteJid: dest,
			},
			Text: input,
		}
		return h.Conn.Send(text)
	}
	return "", errors.New("destination and input should be not empty")
}

func (h *WatsonHandler) getOrCreateWatsonSession(origin string) (*watson.Session, error) {

	// return session if exists for origin
	if session, ok := sessions[origin]; ok {
		return session, nil
	}

	// create watson session and add to sessions
	session, err := watson.NewSession(h.WatsonCfg)
	sessions[origin] = &session
	return &session, err
}

/*END WatsonHandler *******************************************************/

/*
HistoryHandler:

This type is used in order to acquiring the user's chat history.
*/
type HistoryHandler struct {
	Conn *whatsapp.Conn
	Hist *History
}

// ShouldCallSynchronously returns allways true.
func (h *HistoryHandler) ShouldCallSynchronously() bool {
	return true
}

// HandleTextMessage handles and accumulates history's text messages.
func (h *HistoryHandler) HandleTextMessage(message whatsapp.TextMessage) {
	h.Hist.AddToCache(message)
}

// HandleError do nothing.
func (h *HistoryHandler) HandleError(err error) {
	log.Printf("Error occured while retrieving chat history: %s", err)
}

/*END HistoryHandler *******************************************************/
