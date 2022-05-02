package events

import "github.com/alexadhy/gosb/dto"

// AlertUserMessageRateLimitExceeded is a webhook event is fired on user exceeding message rate_limit
type AlertUserMessageRateLimitExceeded struct {
	Category string            `json:"category"`
	User     *dto.UserResponse `json:"user"`
	AppID    string            `json:"app_id"`
}
