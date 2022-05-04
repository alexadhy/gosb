package dto

import "fmt"

type DataExportScheduleRequest struct {
	DataType                string   `json:"data_type"`
	StartTs                 int      `json:"start_ts"`
	EndTs                   int      `json:"end_ts"`
	Format                  string   `json:"format,"`
	CSVDelimiter            string   `json:"csv_delimiter,csv"`
	Timezone                string   `json:"timezone"`
	SenderIDs               []string `json:"sender_ids"`
	ExcludeSenderIDs        []string `json:"exclude_sender_ids"`
	ChannelUrls             []string `json:"channel_urls"`
	ExclueChannelUrls       []string `json:"exclude_channel_urls"`
	UserIds                 []string `json:"user_ids"`
	ShowReadReceipt         bool     `json:"show_read_receipt"`
	ShowChannelMetada       bool     `json:"show_channel_metadata`
	NeighboringMessageLimit int      `json:"neighboring_message_limit"`
}

type DataExportScheduleResponse struct {
	DataExportResource
}

// URL returns URL string to export data endpoint
// POST https://api-{application_id}.sendbird.com/v3/export/{data_type}
func (d DataExportScheduleRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/export/%s", baseURL, d.DataType)
}
