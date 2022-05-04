package dto

import "fmt"

type report struct {
}

type reportResource struct {
	reportingUser     UserResponse `json:"reporting_user"`
	reportType        string       `json:"report_type"`
	reportCategory    string       `json:"report_category"`
	reportedMessage   TextMessage  `json:"reported_message"` // guessing this will be a text message
	offendingUser     UserResponse `json:"offending_user"`
	channel           interface{}  `json:"channel"` // can be open or group channel, how do i handle it? generics?
	reportDescription string       `json:"report_description"`
	createdAt         int          `json:"created_at"`
}

// GET https://api-{application_id}.sendbird.com/v3/report

type ListReportRequest struct {
	token   string `qs:"token"`
	limit   int    `qs:"limit"`
	startTs int    `qs:"start_ts"`
	endTs   int    `qs:"end_ts"`
}

type ListReportResponse struct {
	reportLogs []reportResource `json:"report_logs"`
	next       string           `json:"next"`
}

func (r ListReportRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s", baseURL, reportEndpoint)
}
