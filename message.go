package ws_chat

const (
	MsgEvent MessageEvent = "msg"
)

type MessageEvent string

type message struct {
	to    string
	Event MessageEvent `json:"event"`
	Data  string       `json:"data"`
}
