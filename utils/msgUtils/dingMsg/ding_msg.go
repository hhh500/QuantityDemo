package dingMsg

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func hmacSha256(stringToSign string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func getSignUrl(webhook, secret string) string {
	timestamp := time.Now().UnixNano() / 1e6
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, secret)
	sign := hmacSha256(stringToSign, secret)
	url := fmt.Sprintf("%s&timestamp=%d&sign=%s", webhook, timestamp, sign)
	return url
}

func SendDingMessage(webhook, secret, msg string) error {
	payload := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": msg,
		},
	}
	byteArr, _ := json.Marshal(payload)
	resp, err := http.Post(getSignUrl(webhook, secret), "application/json", bytes.NewBuffer(byteArr))
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}
