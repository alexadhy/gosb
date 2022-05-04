package dto

import "fmt"

type ListReportModerationMessageSpecific struct{}

type ListReportModerationMessageSpecificRequest struct {
	ChannelType string `json:"channel_type"`
	ChannelURL  string `json:"channel_url"`
	MessageID   string `json:"message_id"`
}

type ListReportModerationMessageSpecificResponse struct {
	Message // the documentation doesn't properly state the response
}

func (r ListReportModerationMessageSpecificRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/%s/profanity_messages/%s", baseURL, r.ChannelType, r.ChannelURL, r.MessageID)
}
