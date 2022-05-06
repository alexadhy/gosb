package dto

import "fmt"

// DataExportRequest is the request payload to export data endpoint
type DataExportRequest struct {
	DataType string
	Token    string `qs:"token"`
	Limit    int    `qs:"limit"`
}

// DataExportResource is the response payload to endpoints in data export
type DataExportResource struct {
	StartTs            int      `json:"start_ts"`
	EndTs              int      `json:"end_ts"`
	Status             string   `json:"status"`
	RequestID          string   `json:"request_id"`
	Format             string   `json:"format"`
	CSVDelimiter       string   `json:"csv_delimiter"`
	Timezone           string   `json:"timezone"`
	CreatedAt          int      `json:"created_at"`
	ChannelUrls        []string `json:"channel_urls"`
	ExcludeChannelUrls []string `json:"exclude_channel_urls,omitempty"`
	File               *ZipFile `json:"file,omitempty"`
	SenderIds          []string `json:"sender_ids,omitempty"`
	UserIDs            []string `json:"user_ids,omitempty"`
}

// ZipFile contains a set of preferences to use on making
type ZipFile struct {
	Url       string `json:"url"`
	ExpiresAt int    `json:"expires_at"`
}

// URL returns URL string to data export request endpoint
// GET https://api-{application_id}.sendbird.com/v3/export/{data_type}
func (d DataExportRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/export/%s", baseURL, d.DataType)
	s, _ := encodeQS[DataExportRequest](base, d)
	return s
}
