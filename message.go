package ws_chat

const (
	MsgEvent      MessageEvent = "msg"
	DownloadEvent MessageEvent = "download"
)

type MessageEvent string

const (
	Success MessageType = "success"
	Warning MessageType = "warning"
	Info    MessageType = "info"
	Error   MessageType = "error"
)

type MessageType string

type message struct {
	to    string
	Event MessageEvent `json:"event"`
	Type  MessageType  `json:"type"`
	Data  interface{}  `json:"data"`
}
