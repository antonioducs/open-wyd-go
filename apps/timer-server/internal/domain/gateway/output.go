package gateway

type GameOutput interface {
	SendMessage(sessionID uint32, message string)
	SendCharList(sessionID uint32)
}
