package messages

import "fmt"

type TextMessage struct {
    Sender  string  `json:"from"`
    Time    int64   `json:"sent"`
    Text    string  `json:"message"`
}

func (message *TextMessage) String() string {
    return fmt.Sprintf("%d %s: %s", message.Time, message.Sender, message.Text)
}
