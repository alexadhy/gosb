package dto

import "fmt"

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
