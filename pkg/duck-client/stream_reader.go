package duck_client

import (
	"bufio"
	"encoding/json"
	"io"
	"strings"

	"github.com/bug-breeder/duckai/models"
)

// ProcessStream reads the streaming response and sends ResponseData to the provided channel.
func ProcessStream(body io.Reader, responseChan chan<- models.ResponseData) error {
	reader := bufio.NewReader(body)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "data: ") {
			dataStr := strings.TrimPrefix(line, "data: ")
			if dataStr == "[DONE]" {
				break
			}
			var responseData models.ResponseData
			err := json.Unmarshal([]byte(dataStr), &responseData)
			if err != nil {
				continue
			}
			responseChan <- responseData
		}
	}
	return nil
}
