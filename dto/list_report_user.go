package dto

import "fmt"

type ListReportUserRequest struct {
	offendingUserId string
	token           string `qs:"token,omitempty"`
	limit           int    `qs:"limit,omitempty"`
}

type ListReportUserResponse struct {
	reportLogs []reportResource `json:"report_logs"`
	next       string           `json:"next"`
}

// GET https://api-{application_id}.sendbird.com/v3/report/users/{offending_user_id}

func (r ListReportUserRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/users/%s", baseURL, reportEndpoint, r.offendingUserId)
	s, err := encodeQS[ListReportUserRequest](base, r)
	if err != nil {
		return ""
	}
	return s
}
