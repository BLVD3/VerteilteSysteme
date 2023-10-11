package messages

type SendTextMessage struct {
    Text string
    Time string
}

type ReceiveTextMessage struct {
    Text string
    Time string
    Sender string
}

func GetSendTextMessage(m *map[string]any) (*SendTextMessage, error) {
    var message SendTextMessage
    var text string
    var time string
    var err error

    text, err = assertStringFromJson(TextJsonKey, m)
    if err != nil {
        return nil, err
    }

    time, err = assertStringFromJson(TimeJsonKey, m)
    if err != nil {
        return nil, err
    }

    message.Text = text
    message.Time = time
    return &message, nil
}

func GetReceiveTextMessage(m *map[string]any) (*ReceiveTextMessage, error) {
    var message ReceiveTextMessage
    var text string
    var time string
    var sender string
    var err error

    text, err = assertStringFromJson(TextJsonKey, m)
    if err != nil {
        return nil, err
    }

    time, err = assertStringFromJson(TimeJsonKey, m)
    if err != nil {
        return nil, err
    }

    sender, err = assertStringFromJson(SenderJsonKey, m)
    if err != nil {
        return nil, err
    }

    message.Text = text
    message.Time = time
    message.Sender = sender
    return &message, nil
}
