package dto

import "fmt"

type Metrics struct {
	Messages         int `json:"messages"`
	MessagesPerUser  int `json:"messages_per_user"`
	CreatedChannels  int `json:"created_channels"`
	ActiveChannels   int `json:"active_channels"`
	MessageSenders   int `json:"message_senders"`
	MessageViewers   int `json:"message_viewers"`
	CreatedUsers     int `json:"created_users"`
	DeactivatedUsers int `json:"deactivated_users"`
	DeletedUsers     int `json:"deleted_users"`
}

type Value struct {
	Date    string `json:"date"`
	Channel string `json:"channel"`
	Value   int    `json:"value"`
}

type StatisticsAdvancedAnalayticsRequest struct {
	MetricType    string `qs:"metric_type"`
	TimeDimension string `qs:"time_dimension"`
	StartYear     int    `qs:"start_year"`
	EndYear       int    `qs:"end_year"`
	StartMonth    int    `qs:"start_month"`
	EndMonth      int    `qs:"end_month"`
	// optional
	Limit      int    `qs:"limit,omitempty"`
	Token      string `qs:"token,omitempty"`
	Segments   string `qs:"segments,omitempty"`
	StartDay   int    `qs:"start_day,omitempty"`
	EndDay     int    `qs:"end_day,omitempty"`
	ExportAsCV bool   `qs:"export_as_cv,omitempty"`
}

type StatisticsAdvancedAnalayticsResponse struct {
	MetricType string   `json:"metric_type"`
	Segments   []string `json:"segments"`
	Values     []Value  `json:"values"`
	Next       string   `json:"string"`
}

// URL returns the url to retrieve advanced analytics metrics based on the specified parameters
// GET https://api-{application_id}.sendbird.com/v3/statistics/metric
func (m StatisticsAdvancedAnalayticsRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/metric", baseURL, statisticsEndpoint)
	base, _ = encodeQS(base, m)
	return base
}

type StatisticsConcurrentConnectionsRequest struct{}

type StatisticsConcurrentConnectionsResponse struct {
	Ccu int `json:"ccu"`
}

// URL returns the url to retrueve the number of devices and opened browser tabs currently connected to Sendbird server.
// GET https://api-{application_id}.sendbird.com/v3/applications/ccu
func (m StatisticsConcurrentConnectionsRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/ccu", baseURL, applicationsEndpoint)
}

type StatisticsDailyActiveUsersRequest struct {
	Date string `qs:"date"`
}

type StatisticsDailyActiveUsersResponse struct {
	Dau int `json:"dau"`
}

// URL returns the url to retrieve the numebr of daily users of an application
// GET https://api-{application_id}.sendbird.com/v3/applications/dau
func (m StatisticsDailyActiveUsersRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/dau", baseURL, applicationsEndpoint)
	base, _ = encodeQS(base, m)
	return base
}

type StatisticsMonthlyActiveUsersRequest struct {
	Date string `qs:"date"`
}

type StatisticsMonthlyActiveUsersResponse struct {
	Mau int `json:"mau"`
}

// URL returns the url to retrieve the number of monthly users of an application
// GET https://api-{application_id}.sendbird.com/v3/applications/mau
func (m StatisticsMonthlyActiveUsersRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/mau", baseURL, applicationsEndpoint)
	base, _ = encodeQS(base, m)
	return base
}

type StatisticsPeakConnectionsRequest struct {
	TimeDimension string `qs:"time_dimension"`
	StartYear     int    `qs:"start_year"`
	StartMonth    int    `qs:"start_month"`
	EndYear       int    `qs:"end_year"`
	EndMonth      int    `qs:"end_month"`
	// Optional
	StartDay int `qs:"start_day"`
	EndDay   int `qs:"end_day"`
}

type StatisticsPeakConnectionsResponse struct {
	PeakConnections []PeakConnection `json:"peak_connections"`
}

type PeakConnection struct {
	Date            string `json:"date"`
	PeakConnections int    `json:"peak_connections"`
}

// URL returns the url to retrieve the number of devies that were concurrently connected
// to the Sendbird server during the requested time period.
// GET https://api-{application_id}.sendbird.com/v3/applications/peak_connections
func (m StatisticsPeakConnectionsRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/peak_connections", baseURL, applicationsEndpoint)
}
