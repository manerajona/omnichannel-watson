package watsonconn

import (
	"encoding/json"
	"time"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/watson-developer-cloud/go-sdk/v2/assistantv2"
)

// Session has the param to comunicate with Watson.
type Session struct {
	AssistantID string
	assistant   *assistantv2.AssistantV2
	SessionID   *string
	CreatedAt   time.Time
}

// NewSession creates a new Session for watson assistant.
// Returns the Session or an error.
func NewSession(cfg *Config) (Session, error) {
	return New(cfg.AssistantID, cfg.Instance, cfg.Version, cfg.Region, cfg.Credentials)
}

// New creates a new Session for watson assistant.
// Returns the Session or an error.
func New(assistantID string, instance string, version string, region string, credentials string) (session Session, err error) {

	// create assistant
	a, err := newAssistant(instance, version, region, credentials)
	if err != nil {
		return
	}

	// Call the assistant CreateSession method
	createSessionResult, _, err := a.CreateSession(&assistantv2.CreateSessionOptions{
		AssistantID: core.StringPtr(assistantID),
	})
	if err != nil {
		return
	}

	// create session
	session = Session{
		AssistantID: assistantID,
		assistant:   a,
		SessionID:   createSessionResult.SessionID,
		CreatedAt:   time.Now(),
	}
	return
}

func newAssistant(instance string, version string, region string, apikey string) (a *assistantv2.AssistantV2, err error) {

	// Instantiate the Watson AssistantV2 service
	authenticator := &core.IamAuthenticator{
		ApiKey: apikey,
	}

	a, err = assistantv2.NewAssistantV2(&assistantv2.AssistantV2Options{
		URL:           getServiceURL(region, instance),
		Version:       core.StringPtr(version),
		Authenticator: authenticator,
	})
	return
}

func getServiceURL(region string, instance string) string {
	return "https://api." + region + ".assistant.watson.cloud.ibm.com/instances/" + instance
}

// DeleteSession method delete sessionID for watson assistant.
func (w *Session) DeleteSession() (err error) {

	// Call the assistant DeleteSession method
	_, err = w.assistant.DeleteSession(&assistantv2.DeleteSessionOptions{
		AssistantID: core.StringPtr(w.AssistantID),
		SessionID:   w.SessionID,
	})
	return
}

// WatsonResponse parse the json response from Watson.
type WatsonResponse struct {
	Output struct {
		Intents []struct {
			Intent     string  `json:"intent"`
			Confidence float64 `json:"confidence"`
		} `json:"intents"`
		Entities []struct {
			Entity     string `json:"entity"`
			Location   []int  `json:"location"`
			Value      string `json:"value"`
			Confidence int    `json:"confidence"`
		} `json:"entities"`
		Generic []struct {
			ResponseType string `json:"response_type"`
			Text         string `json:"text"`
		} `json:"generic"`
	} `json:"output"`
}

// SendStatefulMessage method sends an string input to watson assistant.
// Return the WatsonResponse or an error.
func (w *Session) SendStatefulMessage(input string) (WatsonResponse, error) {

	result, _, err := w.assistant.Message(&assistantv2.MessageOptions{
		AssistantID: core.StringPtr(w.AssistantID),
		SessionID:   w.SessionID,
		Input: &assistantv2.MessageInput{
			MessageType: core.StringPtr("text"),
			Text:        core.StringPtr(input),
		},
		Context: &assistantv2.MessageContext{
			Global: &assistantv2.MessageContextGlobal{
				System: &assistantv2.MessageContextGlobalSystem{
					UserID: core.StringPtr("anonymous"), // TODO: add user's name
				},
			},
		},
	},
	)
	if err != nil {
		return WatsonResponse{}, err
	}
	body, _ := json.MarshalIndent(result, "", "  ")

	var response WatsonResponse
	json.Unmarshal(body, &response)

	return response, nil
}

// SendStatelessMessage method sends an string input to watson assistant.
// Return the WatsonResponse or an error.
func (w *Session) SendStatelessMessage(input string) (WatsonResponse, error) {

	result, _, err := w.assistant.MessageStateless(&assistantv2.MessageStatelessOptions{
		AssistantID: core.StringPtr(w.AssistantID),
		Input: &assistantv2.MessageInputStateless{
			MessageType: core.StringPtr("text"),
			Text:        core.StringPtr(input),
		},
	})
	if err != nil {
		return WatsonResponse{}, err
	}

	body, _ := json.MarshalIndent(result, "", "  ")

	var response WatsonResponse
	json.Unmarshal(body, &response)

	return response, nil
}
