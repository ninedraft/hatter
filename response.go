package hatter

import "time"

type Response struct {
	ResponseData ResponseData `json:"response"`
	Session      Session      `json:"session"`
	Version      string       `json:"version"`
}

type ResponseData struct {
	Text       string   `json:"text"`
	Tts        string   `json:"tts"`
	Buttons    []Button `json:"buttons"`
	EndSession bool     `json:"end_session"`
}

type Button struct {
	Title   string  `json:"title"`
	Payload Payload `json:"payload"`
	URL     string  `json:"url"`
	Hide    bool    `json:"hide"`
}

type Payload struct {
	Label     string                 `json:"label, omitempty"`
	ID        string                 `json:"id, omitempty"`
	Timestamp *time.Time             `json:"timestamp, omitempty"`
	Meta      map[string]interface{} `json:"meta, omitempty"`
}
