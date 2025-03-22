package main

import "QuantityDemo/utils/msgUtils/dingMsg"

func main() {
	secret := ""
	webhook := ""
	dingMsg.SendDingMessage(webhook, secret, "test")
}
