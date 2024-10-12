package models

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestBody struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type ResponseData struct {
	Role      string `json:"role"`
	Message   string `json:"message"`
	CreatedAt string `json:"created"`
	Id        string `json:"id"`
	Action    string `json:"action"`
	Model     string `json:"model"`
}

type ErrorResponseData struct {
	Action string `json:"action"`
	Status string `json:"status"`
	Type   string `json:"type"`
}
