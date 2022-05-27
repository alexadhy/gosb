package events

import "github.com/alexadhy/gosb/dto"

// Change is the diff between vanilla and edited message
type Change struct {
	Key string `json:"key"`
	Old string `json:"old"`
	New string `json:"new"`
}

type OpenChannelCommon struct {
	Category string                   `json:"category"`
	Channel  *dto.OpenChannelResponse `json:"channel"`
	AppID    string                   `json:"app_id"`
}

// OpenChannelCreate event is fired on creation of open channel
type OpenChannelCreate struct {
	OpenChannelCommon
	CreatedAt int64              `json:"created_at"`
	Operators []dto.UserResponse `json:"operators"`
}

func (o OpenChannelCreate) EventCategory() string {
	return o.OpenChannelCommon.Category
}

// OpenChannelRemove event is fired on deletion of open_channel
type OpenChannelRemove struct {
	OpenChannelCommon
	RemovedAt int64 `json:"removed_at"`
}

func (o OpenChannelRemove) EventCategory() string {
	return o.OpenChannelCommon.Category
}

// OpenChannelEnter event is fired on user entering an open channel
type OpenChannelEnter struct {
	OpenChannelCommon
	EnteredAt int64             `json:"entered_at"`
	User      *dto.UserResponse `json:"user"`
}

func (o OpenChannelEnter) EventCategory() string {
	return o.OpenChannelCommon.Category
}

// OpenChannelExit event is fired on user exiting an open channel
type OpenChannelExit struct {
	OpenChannelCommon
	EnteredAt int64             `json:"entered_at"`
	User      *dto.UserResponse `json:"user"`
}

func (o OpenChannelExit) EventCategory() string {
	return o.OpenChannelCommon.Category
}

// OpenChannelMessageSend event is fired after user sent a message on open_channel
type OpenChannelMessageSend struct {
	OpenChannelCommon
	Sender         *dto.UserResponse  `json:"sender,omitempty"`
	Silent         bool               `json:"silent"`
	SenderIPAddr   string             `json:"sender_ip_addr"`
	CustomType     string             `json:"custom_type"`
	MentionType    string             `json:"mention_type"`
	MentionedUsers []dto.UserResponse `json:"mentioned_users"`
	MessageType    string             `json:"type"`
	Payload        *dto.Message       `json:"payload"`
	Sdk            string             `json:"sdk"`
}

func (o OpenChannelMessageSend) EventCategory() string {
	return o.OpenChannelCommon.Category
}

// OpenChannelMessageUpdate event is fired after a message is updated on open_channel
type OpenChannelMessageUpdate struct {
	OpenChannelCommon
	Sender         *dto.UserResponse  `json:"sender,omitempty"`
	Changes        []Change           `json:"changes"`
	CustomType     string             `json:"custom_type"`
	MentionType    string             `json:"mention_type"`
	MentionedUsers []dto.UserResponse `json:"mentioned_users"`
	MessageType    string             `json:"type"`
	Payload        *dto.Message       `json:"payload"`
	Sdk            string             `json:"sdk"`
}

func (o OpenChannelMessageUpdate) EventCategory() string {
	return o.OpenChannelCommon.Category
}

// OpenChannelMessageDelete event is fired after a message is deleted on open_channel
type OpenChannelMessageDelete struct {
	OpenChannelCommon
	Sender      *dto.UserResponse `json:"sender,omitempty"`
	CustomType  string            `json:"custom_type"`
	MessageType string            `json:"type"`
	Payload     *dto.Message      `json:"payload"`
	Sdk         string            `json:"sdk"`
}

func (o OpenChannelMessageDelete) EventCategory() string {
	return o.OpenChannelCommon.Category
}
