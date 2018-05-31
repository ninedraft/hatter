package hatter

type Request struct {
	Meta        Meta        `json:"meta"`
	RequestData RequestData `json:"request"`
	Session     Session     `json:"session"`
	Version     string      `json:"version"`
}

type Meta struct {
	Locale   string `json:"locale"`
	Timezone string `json:"timezone"`
	ClientID string `json:"client_id"`
}

type RequestData struct {
	Command           string  `json:"command"`
	OriginalUtterance string  `json:"original_utterance"`
	Type              Type    `json:"type"`
	Markup            Markup  `json:"markup"`
	Payload           Payload `json:"payload, omitempty"`
}

type Session struct {
	New       bool   `json:"new"`
	MessageID int    `json:"message_id"`
	SessionID string `json:"session_id"`
	SkillID   string `json:"skill_id"`
	UserID    string `json:"user_id"`
}

type Markup struct {
	DangerousContext bool `json:"dangerous_context, omitempty"`
}

type Type string

const (
	SimpleUtterance Type = "SimpleUtterance"
	ButtonPressed   Type = "ButtonPressed"
)
