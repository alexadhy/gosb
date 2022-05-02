package dto

import (
	"encoding/json"
	"errors"
)

type ThreadInfo struct {
	ReplyCount    *int64           `json:"reply_count,omitempty"`
	MostReplies   []map[string]any `json:"most_replies,omitempty"`
	LastRepliedAt *int64           `json:"last_replied_at,omitempty"`
	UpdatedAt     *int64           `json:"updated_at,omitempty"`
}

// Message may contain Text, File, or Admin Message
type Message struct {
	*TextMessage
	*FileMessage
	*AdminMessage
}

// UnmarshalJSON unmarshal Message type to its member type (FileMessage / TextMessage / AdminMessage)
func (m Message) UnmarshalJSON(data []byte) error {
	msi := map[string]any{}
	if err := json.Unmarshal(data, &msi); err != nil {
		return err
	}
	msgType, ok := messageType[msi["type"].(string)]
	if !ok {
		return errors.New("error unmarshaling json to Message format: unknown message type")
	}
	if err := json.Unmarshal(data, &msgType); err != nil {
		return err
	}
	switch msgType.(type) {
	case *FileMessage:
		m.FileMessage = msgType.(*FileMessage)
	case *TextMessage:
		m.TextMessage = msgType.(*TextMessage)
	case *AdminMessage:
		m.AdminMessage = msgType.(*AdminMessage)
	}
	return nil
}

// MarshalJSON marshal Message type to JSON byte
func (m Message) MarshalJSON() ([]byte, error) {
	if m.FileMessage != nil {
		return json.Marshal(m.FileMessage)
	}
	if m.AdminMessage != nil {
		return json.Marshal(m.AdminMessage)
	}
	if m.TextMessage != nil {
		return json.Marshal(m.TextMessage)
	}
	return nil, errors.New("unknown message type")
}

var (
	messageType = map[string]any{
		"FILE": &FileMessage{},
		"MESG": &TextMessage{},
		"ADMM": &AdminMessage{},
	}
)

// TextMessage is message of type text on sendbird
type TextMessage struct {
	CommonMessage
	Message              string      `json:"message"`
	Translations         any         `json:"translations,omitempty"`
	Data                 *string     `json:"data,omitempty"`
	OGTag                *OpenGraph  `json:"og_tag,omitempty"`
	IsAppleCriticalAlert *bool       `json:"is_apple_critical_alert,omitempty"`
	ThreadInfo           *ThreadInfo `json:"thread_info,omitempty"`
	ParentMessageID      *int64      `json:"parent_message_id,omitempty"`
	ParentMessageInfo    *ThreadInfo `json:"parent_message_info,omitempty"`
	IsReplyToChannel     *bool       `json:"is_reply_to_channel,omitempty"`
}

// File contains File request and response payload
type File struct {
	URL            string `json:"url"`
	Name           string `json:"name"`
	MIMEType       string `json:"type"`
	Size           int    `json:"size"`
	AdditionalData string `json:"data"`
}

// CommonMessage is common fields included in each object of type Message
type CommonMessage struct {
	MessageID       int64            `json:"message_id"`
	MessageType     string           `json:"type"`
	CustomType      string           `json:"custom_type"`
	ChannelURL      string           `json:"channel_url"`
	User            UserResponse     `json:"user"`
	MentionType     string           `json:"mention_type"`
	MentionedUsers  []UserResponse   `json:"mentioned_users"`
	SortedMetaArray []map[string]any `json:"sorted_metaarray,omitempty"`
	IsRemoved       *bool            `json:"is_removed"`
	CreatedAt       int64            `json:"created_at"`
	UpdatedAt       int64            `json:"updated_at,omitempty"`
}

// FileMessage is message of type file on sendbird
type FileMessage struct {
	CommonMessage
	File             *File       `json:"file,omitempty"`
	Thumbnails       []any       `json:"thumbnails,omitempty"`
	RequireAuth      *bool       `json:"require_auth,omitempty"`
	ThreadInfo       *ThreadInfo `json:"thread_info,omitempty"`
	ParentMessageID  *int64      `json:"parent_message_id,omitempty"`
	IsReplyToChannel *bool       `json:"is_reply_to_channel,omitempty"`
}

// AdminMessage is message of type admin on sendbird
type AdminMessage struct {
	CommonMessage
	Message string     `json:"message"`
	Data    *string    `json:"data,omitempty"`
	OGTag   *OpenGraph `json:"og_tag,omitempty"`
}
