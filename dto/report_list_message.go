package dto

import "fmt"

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
