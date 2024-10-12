package models

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type ChatResponseChunk struct {
	Role      string `json:"role"`
	Message   string `json:"message"`
	CreatedAt string `json:"created"`
	Id        string `json:"id"`
	Action    string `json:"action"`
	Model     string `json:"model"`
}

type ErrorResponseChunk struct {
	Action string `json:"action"`
	Status string `json:"status"`
	Type   string `json:"type"`
}
