package dto

import "fmt"

// http request: GET https://api-{application_id}.sendbird.com/v3/export/{data_type}

type dataExport struct {
}

type dataExportRequest struct {
	dataType string
	token    string `qs:"token"`
	limit    int    `qs:"limit"`
}

type messageResponse struct {
	exportedData []dataExportResource `json:"exported_data"`
	next         string               `json:"next"`
}

type dataExportResource struct {
	startTs            int      `json:"start_ts"`
	endTs              int      `json:"end_ts"`
	status             string   `json:"status"`
	requestID          string   `json:"request_id"`
	format             string   `json:"format"`
	csvDelimiter       string   `json:"csv_delimiter"`
	timezone           string   `json:"timezone"`
	createdAt          int      `json:"created_at"`
	channelUrls        []string `json:"channel_urls"`
	excludeChannelUrls []string `json:"exclude_channel_urls,omitempty"`
	file               zipFile  `json:"file,omitempty"`
	senderIds          []string `json:"sender_ids,omitempty"`
	userIds            []string `json:"user_ids,omitempty"`
}

type zipFile struct {
	url       string `json:"url"`
	expiresAt int    `json:"expires_at"`
}

func (d dataExportRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/export/%s", baseURL, d.dataType)
}
