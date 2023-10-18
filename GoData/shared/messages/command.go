package messages

type Command struct {
    Action  string  `json:"action"`
    Key     string  `json:"key"`
    Value   string  `json:"value"`
}

