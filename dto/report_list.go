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
