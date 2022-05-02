package dto

import "fmt"

// EventCommon is the common event object to embed to all event related request / response
type EventCommon struct {
	Enabled              bool     `json:"enabled" default:"false"`
	CallbackURL          *string  `json:"url,omitempty"`
	EnabledEvents        []string `json:"enabled_events,omitempty"`
	IncludeMembers       *bool    `json:"include_members,omitempty"`
	AllWebhookCategories []string `json:"all_webhook_categories,omitempty"`
}

// EventSubscribeRequest is the request payload to send to event subscription preferences endpoint
type EventSubscribeRequest struct {
	EventCommon
}

// URL is the URL string to event subscription preferences endpoint
// PUT https://api-{application_id}.sendbird.com/v3/applications/settings/webhook
func (e EventSubscribeRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/webhook", baseURL, settingsEndpoint)
}

// EventSubscribeResponse is the response type to event subscription pref endpoint
type EventSubscribeResponse struct {
	Webhook EventCommon `json:"webhook"`
}

// EventSubscriptionListRequest is the request parameter / URL values to list event subscription endpoint
type EventSubscriptionListRequest struct {
	DisplayAll bool `qs:"display_all_webhook_categories"`
}

// URL returns URL string for list event subscription endpoint
func (e EventSubscriptionListRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/webhook", baseURL, settingsEndpoint)
	s, err := encodeQS[EventSubscriptionListRequest](base, e)
	if err != nil {
		return ""
	}
	return s
}

// EventSubscriptionListResponse is an alias to EventSubscribeResponse
type EventSubscriptionListResponse EventSubscribeResponse
