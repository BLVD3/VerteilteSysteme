package messages

import (
	"encoding/json"
)

const TypeJsonKey string = "messagetype"

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
