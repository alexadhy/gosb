package dto

import "fmt"

type ModerationFreezeOpenChannelRequest struct {
	ChannelURL string `json:"channel_url"`
	// body
	Freeze bool `json:"freeze"`
}

type ModerationFreezeOpenChannelResponse OpenChannelResponse

// URL returns the url to freeze or unfreeze an open channel
// PUT https://api-{application_id}.sendbird.com/v3/open_channels/{channel_url}/freeze
func (m ModerationFreezeOpenChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/freeze", baseURL, openChannelEndpoint, m.ChannelURL)
}

type ModerationFreezeGroupChannelRequest ModerationFreezeOpenChannelRequest

type ModerationFreezeGroupChannelResponse GroupChannelResponse

// URL returns the url to freeze or unfreeze a group channel
// PUT https://api-{application_id}.sendbird.com/v3/group_channels/{channel_url}/freeze
func (m ModerationFreezeGroupChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/freeze", baseURL, groupChannelEndpoint, m.ChannelURL)
}
