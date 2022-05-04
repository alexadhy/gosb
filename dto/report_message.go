package dto

import "fmt"

// POST https://api-{application_id}.sendbird.com/v3/report/{channel_type}/{channel_url}/messages/{message_id}

type ReportMessage struct{}

type ReportMessageRequest struct {
	ChannelType string `json:"channel_type"`
	ChannelURL  string `json:"channel_url"`
	MessageID   string `json:"message_id"`
	// body contents
	ReportCategory    string `json:"report_category"`
	OffendingUserID   string `json:"offending_user_id"`
	ReportingUserID   string `json:"reporting_user_id,omitempty"`
	ReportDescription string `json:"report_description,omitempty"`
}

type ReportMessageResponse struct {
	ReportResource
}

func (r ReportMessageRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/%s/messages/%s", baseURL, reportEndpoint, r.ChannelType, r.ChannelURL, r.MessageID)
}
