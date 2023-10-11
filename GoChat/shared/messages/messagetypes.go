package messages

import (
	"encoding/json"
	"errors"
)

const (
    TypeJsonKey string = "messagetype"
    TextJsonKey string = "text"
    TimeJsonKey string = "time"
    SenderJsonKey string = "sender"
    NameJsonKey string = "name"
)

type MessageType int
const (
	Undefined MessageType = iota
	SendText
	ReceiveText
	NameRequest
	NameResponse
)

func (mt MessageType) String() string {
	switch mt {
	case SendText:
		return "SendText"
	case ReceiveText:
		return "ReceiveText"
	case NameRequest:
		return "NameRequest"
	case NameResponse:
		return "NameResponse"
	}
	return "Undefined"
}

func GetMessageTypeFromBytes(message []byte) MessageType {
	var m map[string]any
	err := json.Unmarshal(message, &m) //TODO Learn
	if err != nil {
		return Undefined
	}
	return GetMessageType(m)
}

func GetMessageType(message map[string]any) MessageType {
	s, ok := message[TypeJsonKey].(string)
	if !ok {
		return Undefined
	}
	return StringToMessageType(s)
}

func StringToMessageType(s string) MessageType {
	switch (s) {
	case "SendText":
		return SendText
	case "ReceiveText":
		return ReceiveText
	case "NameRequest":
		return NameRequest
	case "NameResponse":
		return NameResponse
	}
	return Undefined
}

func assertStringFromJson(key string, m *map[string]any) (string, error) {
    var ok bool
    var temp any
    var s string
    temp, ok = (*m)[key]
    if !ok {
        return "", errors.New("error getting string from map. Map does not contain the " + key + " key")
    }
    s, ok = temp.(string)
    if !ok {
        return "", errors.New("error getting string from map. " + key + " is not a string")
    }
    return s, nil
}
