package dto

import "fmt"

type ReportUser struct{}

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
