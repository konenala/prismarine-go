package chat

import (
	"encoding/json"
	"strings"
)

// Component represents a chat component (matches Minecraft's JSON chat format)
type Component struct {
	Text          string      `json:"text,omitempty"`
	Bold          bool        `json:"bold,omitempty"`
	Italic        bool        `json:"italic,omitempty"`
	Underlined    bool        `json:"underlined,omitempty"`
	Strikethrough bool        `json:"strikethrough,omitempty"`
	Obfuscated    bool        `json:"obfuscated,omitempty"`
	Color         string      `json:"color,omitempty"`
	ClickEvent    *ClickEvent `json:"clickEvent,omitempty"`
	HoverEvent    *HoverEvent `json:"hoverEvent,omitempty"`
	Extra         []Component `json:"extra,omitempty"`
}

// ClickEvent represents a click event on a chat component
type ClickEvent struct {
	Action string `json:"action"` // "open_url", "run_command", "suggest_command", etc.
	Value  string `json:"value"`
}

// HoverEvent represents a hover event on a chat component
type HoverEvent struct {
	Action string      `json:"action"` // "show_text", "show_item", "show_entity"
	Value  interface{} `json:"value"`  // Can be string or complex object
}

// Message represents a chat message
type Message struct {
	Component Component // Root component
}

// NewMessage creates a new message from a plain text string
func NewMessage(text string) *Message {
	return &Message{
		Component: Component{Text: text},
	}
}

// ToPlainText converts the message to plain text (removes formatting)
func (m *Message) ToPlainText() string {
	return componentToPlainText(m.Component)
}

func componentToPlainText(c Component) string {
	var sb strings.Builder
	sb.WriteString(c.Text)
	for _, extra := range c.Extra {
		sb.WriteString(componentToPlainText(extra))
	}
	return sb.String()
}

// ToJSON converts the message to JSON format
func (m *Message) ToJSON() (string, error) {
	data, err := json.Marshal(m.Component)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ParseJSON parses a JSON chat message
func ParseJSON(jsonStr string) (*Message, error) {
	var component Component
	err := json.Unmarshal([]byte(jsonStr), &component)
	if err != nil {
		return nil, err
	}
	return &Message{Component: component}, nil
}
