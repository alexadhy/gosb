package dto

import "fmt"

type DataExportSpecificRequest struct {
	DataType  string `json:"data_type"`
	RequestID string `json:"request_id"`
}

// DataExportSpecificResponse is an alias to DataExportResource
type DataExportSpecificResponse DataExportResource

// URL returns URL string
//GET https://api-{application_id}.sendbird.com/v3/export/{data_type}/{request_id}
func (d DataExportSpecificRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/export/%s/%s", baseURL, d.DataType, d.RequestID)
}
