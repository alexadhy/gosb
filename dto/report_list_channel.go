package dto

import "fmt"

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
