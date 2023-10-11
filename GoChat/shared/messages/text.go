package messages

import "errors"

type SendTextMessage struct {
    Text string
    Time string
}
const (
    TextJsonKey string = "text"
    TimeJsonKey string = "time"
    SenderJsonKey string = "sender"
)

type ReceiveTextMessage struct {
    Text string
    Time string
    Sender string
}

func assertStringFromJson(key string, m map[string]any) (string, error) {}

func GetSendTextMessage(m map[string]any) (*SendTextMessage, error) {
    if (GetMessageType(m) != SendText) {
        return nil, errors.New("Error parsing map to SendTextMessage. Map is not a SendText-message")
    }
    var message SendTextMessage
    var ok bool
    var temp any
    var text string
    var time string
    temp, ok = m[TextJsonKey]
    if !ok {
        return nil, errors.New("Error parsing map to SendTextMessage. Map does not contain " + TextJsonKey)
    }
    text, ok = temp.(string)
    if !ok {
        return nil, errors.New("Error parsing map to SendTextMessage. " + TextJsonKey + " is not a String")
    }
    temp, ok = m[TimeJsonKey]
    if !ok {
        return nil, errors.New("Error parsing map to SendTextMessage. Map does not contain " + TimeJsonKey)
    }
    time, ok = temp.(string)
    if !ok {
        return nil, errors.New("Error parsing map to SendTextMessage. " + TimeJsonKey + " is not a String")
    }
    message.Text = text
    message.Time = time
    return &message, nil
}

func GetReceiveTextMessage(m map[string]any) (ReceiveTextMessage, error) {
    return nil, nil
}
