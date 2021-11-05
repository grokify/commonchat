package slack

import (
	"encoding/json"
	"net/url"
	"strings"
)

type Message struct {
	Attachments []Attachment `json:"attachments,omitempty"`
	IconEmoji   string       `json:"icon_emoji,omitempty"`
	IconURL     string       `json:"icon_url,omitempty"`
	Mrkdwn      bool         `json:"mrkdwn,omitempty"`
	Text        string       `json:"text,omitempty"`
	Username    string       `json:"username,omitempty"`
}

func ParseMessageJSON(bytes []byte) (Message, error) {
	msg := Message{}
	err := json.Unmarshal(bytes, &msg)
	return msg, err
}

func ParseMessageURLEncoded(data []byte) (Message, error) {
	qry, err := url.ParseQuery(string(data))
	if err != nil {
		return Message{}, err
	}
	return ParseMessageJSON([]byte(qry.Get("payload")))
}

func ParseMessageAny(data []byte) (Message, error) {
	if strings.Index(strings.TrimSpace(string(data)), "{") == 0 {
		return ParseMessageJSON(data)
	}
	return ParseMessageURLEncoded(data)
}

type Attachment struct {
	AuthorIcon   string   `json:"author_icon,omitempty"`
	AuthorLink   string   `json:"author_link,omitempty"`
	AuthorName   string   `json:"author_name,omitempty"`
	Color        string   `json:"color,omitempty"`
	Fallback     string   `json:"fallback,omitempty"`
	Fields       []Field  `json:"fields,omitempty"`
	MarkdownIn   []string `json:"mrkdwn_in,omitempty"`
	Pretext      string   `json:"pretext,omitempty"`
	Text         string   `json:"text,omitempty"`
	ThumbnailURL string   `json:"thumbnail_url,omitempty"`
	Title        string   `json:"title,omitempty"`
}

type Field struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short bool   `json:"short,omitempty"`
}
