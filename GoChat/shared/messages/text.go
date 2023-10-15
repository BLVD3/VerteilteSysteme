package messages

const (
    SenderJsonKey   string = "from"
    TimeJsonKey     string = "sent"
    TextJsonKey     string = "message"
)

type TextMessage struct {
    Sender  string  `json:"from"`
    Time    int64   `json:"sent"`
    Text    string  `json:"message"`
}
