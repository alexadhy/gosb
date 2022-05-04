package dto

import "fmt"

type Report struct {
}

type ReportResource struct {
	ReportingUser     UserResponse `json:"reporting_user"`
	ReportType        string       `json:"report_type"`
	ReportCategory    string       `json:"report_category"`
	ReportedMessage   TextMessage  `json:"reported_message"` // guessing this will be a text message
	OffendingUser     UserResponse `json:"offending_user"`
	Channel           interface{}  `json:"channel"` // can be open or group channel, how do i handle it? generics?
	ReportDescription string       `json:"report_description"`
	CreatedAt         int          `json:"created_at"`
}

type ListReportRequest struct {
	Token   string `qs:"token"`
	Limit   int    `qs:"limit"`
	StartTs int    `qs:"start_ts"`
	EndTs   int    `qs:"end_ts"`
}

type ListReportResponse struct {
	ReportLogs []ReportResource `json:"report_logs"`
	Next       string           `json:"next"`
}

func (r ListReportRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s", baseURL, reportEndpoint)
	s, _ := encodeQS[ListReportRequest](base, r)
	return s
}

type ListReportUserRequest struct {
	OffendingUserId string
	Token           string `qs:"token"`
	Limit           int    `qs:"limit"`
}

type ListReportUserResponse struct {
	ReportLogs []ReportResource `json:"report_logs"`
	Next       string           `json:"next"`
}

func (r ListReportUserRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/users/%s", baseURL, reportEndpoint, r.OffendingUserId)
	s, _ := encodeQS[ListReportUserRequest](base, r)
	return s
}

type listReportChannel struct{}

type ListReportChannelRequest struct {
	ChannelType string `json:"channel_type"`
	ChannelURL  string `json:"channel_url"`
	Token       string `qs:"token"`
	Limit       int    `qs:"limit"`
}

type listReportChannelResponse struct {
	ReportLogs []ReportResource `json:"report_logs"`
	Next       string           `json:"next"`
}

func (r ListReportChannelRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/%s", baseURL, reportEndpoint, r.ChannelType, r.ChannelURL)
	base, _ = encodeQS[ListReportChannelRequest](base, r)
	return base
}

type ListReportMessage struct{}

type ListReportMessageRequest struct {
	ChannelType string `json:"channel_type"`
	ChannelURL  string `json:"channel_url"`
	MessageID   string `json:"message_id:`
	Token       string `qs:"token"`
	Limit       int    `qs:"limit"`
}

type ListReportMessageResponse struct {
	ReportLogs []ReportResource `json:"report_logs"`
	Next       string           `json:"next"`
}

func (r ListReportMessageRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/%s/messages/%s", baseURL, reportEndpoint, r.ChannelType, r.ChannelURL, r.MessageID)
	base, _ = encodeQS[ListReportMessageRequest](base, r)
	return base
}

type ModeratedMessage struct{}

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
