package duck_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bug-breeder/duckai/models"
)

// Client represents the DuckDuckGo Chat API client.
type Client struct {
	HTTPClient *http.Client
	Headers    map[string]string
	URL        string
}

// NewClient initializes and returns a new Client instance.
func NewClient() *Client {
	return &Client{
		HTTPClient: &http.Client{},
		URL:        "https://duckduckgo.com/duckchat/v1/chat",
		Headers: map[string]string{
			"Accept":          "text/event-stream",
			"Accept-Language": "en-US,en;q=0.9",
			"Content-Type":    "application/json",
			"Cookie":          "dcm=5",
			"Origin":          "https://duckduckgo.com",
			"Referer":         "https://duckduckgo.com/",
			"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64)",
			"x-vqd-4":         "4-107403175889066477207178046153825229413",
		},
	}
}

// SendMessage sends a message to the DuckDuckGo Chat API and returns a channel to read the response stream.
func (c *Client) SendMessage(requestBody models.RequestBody) (<-chan models.ResponseData, error) {
	requestData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %v", err)
	}
	log.Println("Request Data:", string(requestData))

	req, err := http.NewRequest("POST", c.URL, bytes.NewReader(requestData))
	log.Printf("Request: %v\n", req)
	log.Println("Request Body after serialize", req.Body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	for key, value := range c.Headers {
		req.Header.Set(key, value)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("received non-OK response: %s", resp.Status)
	}

	responseChan := make(chan models.ResponseData)
	go func() {
		defer resp.Body.Close()
		defer close(responseChan)
		err := ProcessStream(resp.Body, responseChan)
		if err != nil {
			fmt.Println("Error processing stream:", err)
			os.Exit(1)
		}
	}()

	return responseChan, nil
}
