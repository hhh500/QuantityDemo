package bnSdk

import "context"

type secType int

const (
	secTypeNone secType = iota
	secTypeAPIKey
	secTypeSigned // if the 'timestamp' parameter is required
)

const (
	apiKey           = "apiKey"
	timestampKey     = "timestamp"
	signatureKey     = "signature"
	receiveWindowKey = "receiveWindow"
)

var (
	ctx = context.Background()
)
