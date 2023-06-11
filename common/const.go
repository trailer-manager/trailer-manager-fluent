package common

const (
	HostDefault = "0.0.0.0"
	PortDefault = 1883

	ModeLocal       = "local"
	ModeDevelopment = "dev"
	ModeStaging     = "stg"
	ModeProduction  = "prd"

	ContextTypeNone = 0
	ContextTypeGo   = 1
	ContextTypeEcho = 2

	ContextLogType = "LOG_TYPE"
	ContextKey = "CTX_KEY_"

	ContextLogTypeStart = "START"
	ContextLogTypeEnd = "END"
	ContextLogTypeNormal = "NORMAL"

	HeaderTransactionId = "transactionId"
)
