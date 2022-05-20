package events

import "github.com/alexadhy/gosb/dto"

// ImageModerationBlock webhook event is invoked when a message
// with explicit images or inappropriate image URLs is blocked
type ImageModerationBlock struct {
	Category    string            `json:"category"`
	BlockedAt   int64             `json:"blocked_at"`
	Sender      *dto.UserResponse `json:"sender"`
	Channel     *dto.Channel      `json:"channel"`
	MessageText string            `json:"message_text"`
	AppID       string            `json:"app_id"`
}

func (i ImageModerationBlock) EventCategory() string {
	return i.Category
}
