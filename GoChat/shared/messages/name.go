package messages

type NameResponseMessage struct {
    Name string
}

func GetNameResponseMessage(m *map[string]any) (*NameResponseMessage, error) {
    name, err := assertStringFromJson(NameJsonKey, m)
    if err != nil {
        return nil, err
    }
    return &NameResponseMessage{name}, nil
}
