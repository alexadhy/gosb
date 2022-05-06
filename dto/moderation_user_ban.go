package dto

import "fmt"

type ModerationBanUserFromOpenChannelRequest struct {
	ChannelURL string `json:"channel_url"`
	// body
	UserID string `json:"user_id"`
	// optional
	AgentID     string `json:"agent_id"`
	Seconds     string `json:"seconds"`
	Description string `json:"description"`
}

type ModerationBanUserFromOpenChannelResponse struct {
	User        SimplifiedUser `json:"user"`
	StartAt     int            `json:"start_at"`
	EndAt       int            `json:"end_at"`
	Description string         `json:"description"`
}

// URL returns the url to ban a user from an open channel.
// POST https://api-{application_id}.sendbird.com/v3/open_channels/{channel_url}/ban
func (m ModerationBanUserFromOpenChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/ban", baseURL, openChannelEndpoint, m.ChannelURL)
}

type ModerationBanUserFromGroupChannelRequest ModerationBanUserFromOpenChannelRequest

type ModerationBanUserFromGroupChannelResponse ModerationBanUserFromOpenChannelResponse

// URL returns the url to ban a user from a group channel
// POST https://api-{application_id}.sendbird.com/v3/group_channels/{channel_url}/ban
func (m ModerationBanUserFromGroupChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/ban", baseURL, groupChannelEndpoint, m.ChannelURL)
}

type ModerationBanUserFromCustomChannelRequest struct {
	UserID string `json:"user_id"`
	// body
	ChannelCustomTypes []string `json:"channel_custom_types"`
}

type ModerationBanUserFromCustomChannelResponse struct{}

// URL returns the url to ban a user from channels with the specified custom channel types
// POST https://api-{application_id}.sendbird.com/v3/users/{user_id}/banned_channel_custom_types
func (m ModerationBanUserFromCustomChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/banned_channel_custom_types", baseURL, userEndpoint, m.UserID)
}

type ModerationBanUsersFromCustomChannelRequest struct {
	CustomType string `json:"custom_type"`
	// body
	BannedList     []BanUser `json:"banned_list"`
	OnDemandUpsert bool      `json:"on_demand_upsert"`
}

type BanUser struct {
	UserID      string `json:"user_id"`
	Seconds     int    `json:"seconds"`
	Description string `json:"description"`
}

type ModerationBanUsersFromCustomChannelResponse struct{}

// URL returns the url to ban users from channels with the specified custom channel type at once
// POST https://api-{application_id}.sendbird.com/v3/applications/settings_by_channel_custom_type/{custom_type}/ban
func (m ModerationBanUsersFromCustomChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/ban", baseURL, customChannelEndpoint, m.CustomType)
}

type ModerationUpdateBannedUserStatusInOpenChannelRequest struct {
	ChannelURL   string `json:"channel_url"`
	BannedUserID string `json:"banned_user_id"`
	// body
	Seconds int `json:"seconds"`
	// optional
	Description string `json:"description"`
}

type ModerationUpdateBannedUserStatusInOpenChannelResponse ModerationBanUserFromOpenChannelResponse

// URL returns the url to update the information about a banned user of an open channel
// PUT https://api-{application_id}.sendbird.com/v3/open_channels/{channel_url}/ban/{banned_user_id}
func (m ModerationUpdateBannedUserStatusInOpenChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/ban/%s", baseURL, openChannelEndpoint, m.ChannelURL, m.BannedUserID)
}

type ModerationUpdateBannedUserStatusInGroupChannelRequest ModerationUpdateBannedUserStatusInOpenChannelRequest

type ModerationUpdateBannedUserStatusInGroupChannelResponse ModerationBanUserFromOpenChannelResponse

// URL returns the url to update information about a banned user of a group channel
// PUT https://api-{application_id}.sendbird.com/v3/group_channels/{channel_url}/ban/{banned_user_id}
func (m ModerationUpdateBannedUserStatusInGroupChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/ban/%s", baseURL, groupChannelEndpoint, m.ChannelURL, m.BannedUserID)
}

type ModerationUnbanUserFromOpenChannelRequest struct {
	ChannelURL   string `json:"channel_url"`
	BannedUserID string `json:"banned_user_id"`
}

type ModerationUnbanUserFromOpenChannelResponse struct{}

// URL returns the url to unban a user who had been banned from an open channel
// DELETE https://api-{application_id}.sendbird.com/v3/open_channels/{channel_url}/ban/{banned_user_id}
func (m ModerationUnbanUserFromOpenChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/ban/%s", baseURL, openChannelEndpoint, m.ChannelURL, m.BannedUserID)
}

type ModerationUnbanUserFromGroupChannelRequest ModerationUnbanUserFromOpenChannelRequest

type ModerationUnbanUserFromGroupChannelResponse struct{}

// URL returns the url to unban a user who had been banned from an group channel
// DELETE https://api-{application_id}.sendbird.com/v3/group_channels/{channel_url}/ban/{banned_user_id}
func (m ModerationUnbanUserFromGroupChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/ban/%s", baseURL, groupChannelEndpoint, m.ChannelURL, m.BannedUserID)
}

type ModerationUnbanUserFromCustomChannelRequest struct {
	CustomType string `json:"custom_type"`
	// parameters
	UserIDs []string `json:"user_ids"`
}

// URL returns the url to unban users who had been banned from channels with the specified channel type at once
// DELETE https://api-{application_id}.sendbird.com/v3/applications/settings_by_channel_custom_type/{custom_type}/ban
func (m ModerationUnbanUserFromCustomChannelRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/ban", baseURL, customChannelEndpoint, m.CustomType)
	base, _ = encodeQS[ModerationUnbanUserFromCustomChannelRequest](base, m)
	return base
}
