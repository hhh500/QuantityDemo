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

// hmacSha256 用于对签名字符串进行 HMAC-SHA256 加密并 Base64 编码
func hmacSha256(signStr string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(signStr))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// getSignUrl 生成钉钉机器人的签名URL,包含时间戳和签名
func getSignUrl(webhook, secret string) string {
	// 获取当前毫秒时间戳
	timestamp := time.Now().UnixMilli()
	// 构建签名字符串(时间戳 + "\n" + 密钥)
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, secret)
	// 使用 HMAC-SHA256 生成签名
	sign := hmacSha256(stringToSign, secret)
	// 构造带签名的 webhook URL
	url := fmt.Sprintf("%s&timestamp=%d&sign=%s", webhook, timestamp, sign)
	return url
}

// SendDingMessage 向钉钉机器人发送文本消息
func SendDingMessage(webhook, secret, msg string) error {
	// 构建发送内容
	payload := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": msg,
		},
	}
	// 序列化为 JSON
	byteArr, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("JSON 序列化失败: %v", err)
	}
	// 构建签名后的 URL
	url := getSignUrl(webhook, secret)
	// 发送 HTTP POST 请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(byteArr))
	if err != nil {
		return fmt.Errorf("请求发送失败: %v", err)
	}
	defer resp.Body.Close()
	return nil
}
