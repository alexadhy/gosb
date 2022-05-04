package dto

import "fmt"

type ModeratedMessage struct {
}
type ListReportModerationMessage struct{}

type ListReportModerationMessageRequest struct {
	ChannelType string `json:"channel_type"`
	ChannelURL  string `json:"channel_url"`
	MessageTs   int    `qs:"message_ts"`
	PrevLimit   int    `qs:"prev_limit"`
	NextLimit   int    `qs:"next_limit"`
	UserID      string `qs:"user_id"`
}

type ListReportModerationMessageResponse struct {
	Messages []Message `json:"messages"`
}

func (r ListReportModerationMessageRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/%s/profanity_messages", baseURL, reportEndpoint, r.ChannelType, r.ChannelURL)
	base, _ = encodeQS[ListReportModerationMessageRequest](base, r)
	return base
}
