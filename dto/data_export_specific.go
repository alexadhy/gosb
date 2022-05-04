package dto

import "fmt"

//http request: GET https://api-{application_id}.sendbird.com/v3/export/{data_type}/{request_id}

type dataExportSpecific struct {
}

type dataExportSpecificRequest struct {
	dataType  string `json:"data_type"`
	requestID string `json:"request_id"`
}

type dataExportSpecificResponse struct {
	dataExportResource
}

func (d dataExportSpecificRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/export/%s/%s", baseURL, d.dataType, d.requestID)
}
