package messages

import "encoding/json"

type NameResponseMessage struct {
    Name string
}

const NameRequestMessageString string = "{\"" + TypeJsonKey +  "\": \"namerequest\"}"

func GetNameResponseMessage(m *map[string]any) (*NameResponseMessage, error) {
    name, err := assertStringFromJson(NameJsonKey, m)
    if err != nil {
        return nil, err
    }
    return &NameResponseMessage{name}, nil
}

func (message *NameResponseMessage) MarshalJSON() ([]byte, error) {
    var m map[string]any = map[string]any{}

    m[TypeJsonKey] = SendText.String()
    m[NameJsonKey] = message.Name

    res, _ := json.Marshal(m)

    return res, nil
}

