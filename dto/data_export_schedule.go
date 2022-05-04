package dto

import "fmt"

// POST https://api-{application_id}.sendbird.com/v3/export/{data_type}

type dataExportSchedule struct {
}

type dataExportScheduleRequest struct {
	dataType                string   `json:"data_type"`
	startTs                 int      `json:"start_ts"`
	endTs                   int      `json:"end_ts"`
	format                  string   `json:"format,"`
	csvDelimiter            string   `json:"csv_delimiter,csv"`
	timezone                string   `json:"timezone"`
	senderIds               []string `json:"sender_ids"`
	excludeSenderIds        []string `json:"exclude_sender_ids"`
	channelUrls             []string `json:"channel_urls"`
	exclueChannelUrls       []string `json:"exclude_channel_urls"`
	userIds                 []string `json:"user_ids"`
	showReadReceipt         bool     `json:"show_read_receipt"`
	showChannelMetada       bool     `json:"show_channel_metadata`
	neighboringMessageLimit int      `json:"neighboring_message_limit"`
}

type dataExportScheduleResponse struct {
	dataExportResource
}

func (d dataExportScheduleRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/export/%s")
}
