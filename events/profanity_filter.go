package events

import "github.com/alexadhy/gosb/dto"

// ProfanityFilterReplace webhook event is invoked when explicit words in a message are replaced with asterisks (*).
// The following shows a webhook payload of a profanity_filter:replace event
type ProfanityFilterReplace struct {
	Category     string             `json:"category"`
	Sender       *dto.UserResponse  `json:"sender"`
	Members      []dto.UserResponse `json:"members,omitempty"`
	CustomType   string             `json:"custom_type"`
	MessageType  string             `json:"type"`
	ReplacedText string             `json:"replaced_text"`
	Payload      *dto.Message       `json:"payload"`
	Channel      *dto.Channel       `json:"channel"`
	Sdk          string             `json:"sdk"`
	AppID        string             `json:"app_id"`
}

// ProfanityFilterBlock webhook event is invoked when a message with explicit words is blocked.
// The following shows a webhook payload for a profanity_filter:block event.
type ProfanityFilterBlock struct {
	Category  string            `json:"category"`
	BlockedAt int64             `json:"blocked_at"`
	Sender    *dto.UserResponse `json:"sender"`
	Message   string            `json:"message"`
	Channel   *dto.Channel      `json:"channel"`
	AppID     string            `json:"app_id"`
}

// ProfanityFilterModerate webhook event is invoked when a user is imposed
// with one of moderation penalties among mute, kick, and ban
type ProfanityFilterModerate struct {
	Category         string               `json:"category"`
	ModeratedAt      int64                `json:"moderated_at"`
	ModerationAction dto.ModerationAction `json:"moderation_action"`
	Channel          *dto.Channel         `json:"channel"`
	AppID            string               `json:"app_id"`
}
