package dto

import "fmt"

type ReportUserRequest struct {
	OffendingUserID string `json:"offending_user_id"`
	// body contents
	ChannelType       string `json:"channel_type"`
	ChannelURL        string `json:"channel_url"`
	ReportCategory    string `json:"report_category"`
	ReportingUserID   string `json:"reporting_user_id,omitempty"`
	ReportDescription string `json:"report_description,omitempty"`
}

type ReportUserResponse struct {
	ReportResource
}

func (r ReportUserRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/users/%s", baseURL, reportEndpoint, r.OffendingUserID)
}

type ReportChannelRequest struct {
	ChannelType string `json:"channel_type"`
	ChannelURL  string `json:"channel_url"`
	// body contents
	ReportCategory    string `json:"report_category"`
	ReportingUserID   string `json:"reporting_user_id,omitempty"`
	ReportDescription string `json:"report_description,omitempty"`
}

type ReportChannelResponse struct {
	ReportResource
}

func (r ReportChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/%s", baseURL, reportEndpoint, r.ChannelType, r.ChannelURL)
}

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
