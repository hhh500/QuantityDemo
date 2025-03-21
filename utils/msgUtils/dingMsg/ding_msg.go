package dingMsg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendDingMessage(webhook string, message string) error {
	payload := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": message,
		},
	}
	body, _ := json.Marshal(payload)
	resp, err := http.Post(webhook, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println("钉钉返回状态:", resp.Status)
	return nil
}
