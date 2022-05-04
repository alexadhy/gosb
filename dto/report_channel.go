package dto

import "fmt"

type ReportChannel struct{}

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
