package dto

import "fmt"

// ModerationListBannedUsersRequest stores the parameters the request supports
type ModerationListBannedUsersRequest struct {
	UserID string `json:"user_id"`
	//optional
	Token string `qs:"token"`
	Limit string `qs:"limit"`
}

// ModerationListBannedUsersResponse stores the list of channels
type ModerationListBannedUsersResponse struct {
	BannedChannels []BannedChannel `json:"banned_channel"`
	Next           string          `json:"next"`
}

type BannedChannel struct {
	StartAt     int               `json:"start_at"`
	EndAt       int               `json:"end_at"`
	Description string            `json:"description"`
	Channel     SimplifiedChannel `json:"channel"`
	Next        string            `json:"next"`
}

// SimplifiedChannel is a smaller version of open/group channel.
type SimplifiedChannel struct {
	Name       string `json:"name"`
	CustomType string `json:"custom_type"`
	ChannelURL string `json:"channel_url"`
	CreatedAt  int    `json:"created_at"`
	CoverURL   string `json:"cover_url"`
	Data       string `json:"data"`
}

// URL returns the url to retrieve a list of open channels and group channels where a user is banned
// GET https://api-{application_id}.sendbird.com/v3/users/{user_id}/ban
func (m ModerationListBannedUsersRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/ban", baseURL, userEndpoint, m.UserID)
	base, _ = encodeQS[ModerationListBannedUsersRequest](base, m)
	return base
}

type ModerationListBannedUsersInOpenChannelRequest struct {
	ChannelURL string `json:"channel_url"`
	// optional
	Token             string `json:"token"`
	Limit             int    `json:"limit"`
	ShowTotalBanCount bool   `json:"show_total_ban_count"`
}

type ModerationListBannedUsersInOpenChannelResponse struct {
	BannedList    []BannedUser `json:"banned_list"`
	TotalBanCount int          `json:"total_ban_count"`
	Next          string       `json:"next"`
}

type BannedUser struct {
	User        SimplifiedUser `json:"user"`
	StartAt     int            `json:"start_at"`
	EndAt       int            `json:"end_at"`
	Description string         `json:"description"`
}

type SimplifiedUser struct {
	UserID     string   `json:"user_id"`
	Nickname   string   `json:"nickname"`
	ProfileURL string   `json:"profile_url"`
	IsOnline   bool     `json:"is_online"`
	LastSeenAt int      `json:"last_seen_at"`
	Metadata   Metadata `json:"metadata"`
}

// URL returns the url to retrieve the list of users who are banned from the specified open channel
// GET https://api-{application_id}.sendbird.com/v3/open_channels/{channel_url}/ban
func (m ModerationListBannedUsersInOpenChannelRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/ban", baseURL, openChannelEndpoint, m.ChannelURL)
	base, _ = encodeQS[ModerationListBannedUsersInOpenChannelRequest](base, m)
	return base
}

type ModerationListBannedUsersInGroupChannelRequest ModerationListBannedUsersInOpenChannelRequest

type ModerationListBannedUsersInGroupChannelResponse ModerationListBannedUsersInOpenChannelResponse

// URL returns the url to retrieve the list of users who are banned from a group channel
// GET https://api-{application_id}.sendbird.com/v3/group_channels/{channel_url}/ban
func (m ModerationListBannedUsersInGroupChannelRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/ban", baseURL, groupChannelEndpoint, m.ChannelURL)
	base, _ = encodeQS[ModerationListBannedUsersInGroupChannelRequest](base, m)
	return base
}

type ModerationListBannedUsersInCustomChannelRequest struct {
	CustomType string `json:"custom_type"`
	// optional
	Token string `json:"token"`
	Limit string `json:"limit"`
}

type ModerationListBannedUsersInCustomChannelResponse struct {
	BannedList []BannedUser `json:"banned_list"`
	Next       string       `json:"next"`
}

// URL returns the url to retrieve a list of users banned from channels with the specified custom channel type
// GET https://api-{application_id}.sendbird.com/v3/applications/settings_by_channel_custom_type/{custom_type}/ban
func (m ModerationListBannedUsersInCustomChannelRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/ban", baseURL, customChannelEndpoint, m.CustomType)
	base, _ = encodeQS[ModerationListBannedUsersInCustomChannelRequest](base, m)
	return base
}

type ModerationListBannedUserInOpenChannelRequest struct {
	ChannelURL   string `json:"channel_url"`
	BannedUserID string `json:"banned_user_id"`
}

type ModerationListBannedUserInOpenChannelResponse struct {
	User        SimplifiedChannel `json:"user"`
	StartAt     int               `json:"start_at"`
	EndAt       int               `json:"end_at"`
	Description string            `json:"description"`
}

// URL returns the url to retrieve information about a user banned from an open channel
// GET https://api-{application_id}.sendbird.com/v3/open_channels/{channel_url}/ban/{banned_user_id}
func (m ModerationListBannedUserInOpenChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/ban/%s", baseURL, openChannelEndpoint, m.ChannelURL, m.BannedUserID)
}

type ModerationListBannedUserInGroupChannelRequest ModerationListBannedUserInOpenChannelRequest

type ModerationListBannedUserInGroupChannelResponse ModerationListBannedUserInOpenChannelResponse

// URL returns the url to retrieve information about a user banned from a gorup channel
// GET https://api-{application_id}.sendbird.com/v3/group_channels/{channel_url}/ban/{banned_user_id}
func (m ModerationListBannedUserInGroupChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/ban/%s", baseURL, groupChannelEndpoint, m.ChannelURL, m.BannedUserID)
}
