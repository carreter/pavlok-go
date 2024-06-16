// Package pavlok provides a wrapper for the Pavlok API.
package pavlok

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	DefaultBaseURL = "https://api.pavlok.com/api/v5"
)

type StimulusType string

const (
	Zap  StimulusType = "zap"
	Beep StimulusType = "beep"
	Vibe StimulusType = "vibe"
)

type Stimulus struct {
	Type   StimulusType `json:"stimulusType"`
	Value  int          `json:"stimulusValue"` // 1-100, inclusive
	Reason string       `json:"reason,omitempty"`
}

// A Client is a client to the Pavlok API.
type Client struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
}

// NewClient creates a new client to the Pavlok API.
func NewClient(apiKey string, options ...OptionsFunc) *Client {
	c := &Client{
		httpClient: &http.Client{},
		baseURL:    DefaultBaseURL,
		apiKey:     strings.TrimPrefix(apiKey, "Bearer "),
	}

	for _, optionsFunc := range options {
		c = optionsFunc(c)
	}

	return c
}

func (c *Client) addHeaders(req *http.Request) {
	req.Header.Add("Authorization", "Bearer "+c.apiKey)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
}

// SendStimulus sends a stimulus.
func (c *Client) SendStimulus(stimulus Stimulus) error {
	bodyBytes, err := json.Marshal(stimulus)
	if err != nil {
		return err
	}
	body := fmt.Sprintf("{\"stimulus\":%s}", string(bodyBytes))

	req, err := http.NewRequest("POST", c.baseURL+"/stimulus/send", strings.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to create stimulus request: %w", err)
	}
	c.addHeaders(req)

	response, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send stimulus request: %w", err)
	} else if response.StatusCode != http.StatusOK {
		responseBody, _ := io.ReadAll(response.Body)
		_ = response.Body.Close()
		return fmt.Errorf("pavlok API returned error code %v: %v", response.Status, string(responseBody))
	}

	return nil
}

// Zap sends a zap stimulus.
func (c *Client) Zap(value int) error {
	return c.SendStimulus(Stimulus{
		Type:  Zap,
		Value: value,
	})
}

// Vibe sends a vibe stimulus.
func (c *Client) Vibe(value int) error {
	return c.SendStimulus(Stimulus{
		Type:  Vibe,
		Value: value,
	})
}

// Beep sends a beep stimulus.
func (c *Client) Beep(value int) error {
	return c.SendStimulus(Stimulus{
		Type:  Beep,
		Value: value,
	})
}
