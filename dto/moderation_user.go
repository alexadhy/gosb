package dto

import "fmt"

type ModerationUserResource struct {
	UserID     string    `json:"user_id"`
	Nickname   string    `json:"nickname"`
	ProfileURL string    `json:"profile_url"`
	IsOnline   bool      `json:"is_online"`
	LastSeenAt int       `json:"last_seen_at"`
	Metadata   *Metadata `json:"metadata"`
}

type ModerationMuteUserInOpenChannelRequest struct {
	ChannelURL string `json:"channel_url"`
	// body
	UserID string `json:"user_id"`
	// optional
	Seconds     int    `json:"seconds,omitempty"`
	Description string `json:"description,omitempty"`
}

type ModerationMuteUserInOpenChannelResponse struct {
	OpenChannelResponse
}

func (m ModerationMuteUserInOpenChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/mute", baseURL, openChannelEndpoint, m.ChannelURL)
}

type ModerationMuteUserInGroupChannelRequest struct {
	ChannelURL string `json:"channel_url"`
	// body
	UserID string `json:"user_id"`
	// optional
	Seconds     int    `json:"seconds,omitempty"`
	Description string `json:"description,omitempty"`
}

func (m ModerationMuteUserInGroupChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/mute", baseURL, groupChannelEndpoint, m.ChannelURL)
}

type ModerationMuteUserInCustomChannelRequest struct {
	UserID string `json:"user_id"`
	// body
	ChannelCustomTypes []string `json:"channel_custom_types"`
}

type ModerationMuteUserInCustomChannelResponse struct{}

func (m ModerationMuteUserInCustomChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/muted_channel_custom_types", baseURL, userEndpoint, m.UserID)
}

type ModerationMuteUsersInCustomChannelRequest struct {
	CustomType string `json:"custom_type"`
	// body
	UserIds []string `json:"user_ids"`
	// optional
	Seconds        int    `json:"seconds,omitempty"`
	Description    string `json:"description,omitempty"`
	OnDemandUpsert bool   `json:"on_demand_upsert,omitempty"`
}

type ModerationMuteUsersInCustomChannelResponse struct{}

func (m ModerationMuteUsersInCustomChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/mute", baseURL, customChannelEndpoint, m.CustomType)
}

type ModerationUnmuteUserInOpenChannelRequest struct {
	ChannelURL  string `json:"channel_url"`
	MutedUserID string `json:"muted_user_id"`
}

type ModerationUnmuteUserInOpenChannelResponse struct{}

func (m ModerationUnmuteUserInOpenChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/mute/%s", baseURL, openChannelEndpoint, m.ChannelURL, m.MutedUserID)
}

type ModerationUnmuteUserInGroupChannelRequest struct {
	ChannelURL  string `json:"channel_url"`
	MutedUserID string `json:"muted_user_id"`
}

func (m ModerationUnmuteUserInGroupChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/mute/%s", baseURL, groupChannelEndpoint, m.ChannelURL, m.MutedUserID)
}

type ModerationUnmuteUsersInCustomChannelRequest struct {
	CustomType string `json:"custom_type"`
	// body
	UserIds []string `json:"user_ids"`
}

func (m ModerationUnmuteUsersInCustomChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/mute", baseURL, customChannelEndpoint, m.CustomType)
}

type ModerationListUsersMutedChannelsRequest struct {
	UserID string `json:"user_id"`
	// optional
	Token string `qs:"token,omitempty"`
	Limit string `qs:"limit,omitempty"`
}

type ModerationListUsersMutedChannelsResponse struct {
	MutedChannels []Channel `json:"muted_channels"`
	Next          string    `json:"next"`
}

func (m ModerationListUsersMutedChannelsRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/mute", baseURL, userEndpoint, m.UserID)
	base, _ = encodeQS[ModerationListUsersMutedChannelsRequest](base, m)
	return base
}

type ModerationListMutedUsersInOpenChannelRequest struct {
	ChannelURL string `json:"channel_url"`
	//optional
	ShowTotalMuteCount bool   `qs:"show_total_mute_count,omitempty"`
	Token              string `qs:"token,omitempty"`
	Limit              int    `qs:"limit,omitempty"`
}

type ModerationListMutedUsersInOpenChannelResponse struct {
	MutedList      []ModerationUserResource `json:"muted_list"`
	TotalMuteCount int                      `json:"total_mute_count"`
	Next           string                   `json:"next"`
}

func (m ModerationListMutedUsersInOpenChannelRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/mute", baseURL, openChannelEndpoint, m.ChannelURL)
	base, _ = encodeQS[ModerationListMutedUsersInOpenChannelRequest](base, m)
	return base
}

type ModerationListMutedUsersInGroupChannel struct{}

type ModerationListMutedUsersInGroupChannelRequest struct {
	ChannelURL string `json:"channel_url"`
	//optional
	ShowTotalMuteCount bool   `qs:"show_total_mute_count,omitempty"`
	Token              string `qs:"token,omitempty"`
	Limit              int    `qs:"limit,omitempty"`
}

type ModerationListMutedUsersInGroupChannelResponse struct {
	MutedList      []ModerationUserResource `json:"muted_list"`
	TotalMuteCount int                      `json:"total_mute_count"`
	Next           string                   `json:"next"`
}

func (m ModerationListMutedUsersInGroupChannelRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/mute", baseURL, groupChannelEndpoint, m.ChannelURL)
	base, _ = encodeQS[ModerationListMutedUsersInGroupChannelRequest](base, m)
	return base
}

type ModerationListMutedUsersInCustomChannelRequest struct {
	CustomType string `json:"custom_type"`
	// optional
	Token string `json:"token,omitempty"`
	Limit int    `json:"limit,omitempty"`
}

type ModerationListMutedUsersInCustomChannelResponse struct {
	MutedList []ModerationUserResource `json:"muted_list"`
	Next      string                   `json:"next"`
}

func (m ModerationListMutedUsersInCustomChannelRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/mute", baseURL, customChannelEndpoint, m.CustomType)
	base, _ = encodeQS[ModerationListMutedUsersInCustomChannelRequest](base, m)
	return base
}

type ModerationListMutedUserInGroupChannel struct{}

type ModerationListMutedUserInGroupChannelRequest struct {
	ChannelURL string `json:"channel_url"`
	MutedUrlId string `json:"muted_user_id"`
}

type ModerationListMutedUserInGroupChannelResponse struct {
	IsMuted           bool   `json:"is_muted"`
	RemainingDuration int    `json:"remaining_duration"`
	StartAt           int    `json:"start_at"`
	EndAt             int    `json:"end_at"`
	Description       string `json:"description"`
}

func (m ModerationListMutedUserInGroupChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/mute/%s", baseURL, groupChannelEndpoint, m.ChannelURL, m.MutedUrlId)
}

type ModerationListMutedUserInOpenChannelRequest struct {
	ChannelURL string `json:"channel_url"`
	MutedUrlId string `json:"muted_user_id"`
}

type ModerationListMutedUserInOpenChannelResponse struct {
	IsMuted           bool   `json:"is_muted"`
	RemainingDuration int    `json:"remaining_duration"`
	StartAt           int    `json:"start_at"`
	EndAt             int    `json:"end_at"`
	Description       string `json:"description"`
}

func (m ModerationListMutedUserInOpenChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/mute/%s", baseURL, openChannelEndpoint, m.ChannelURL, m.MutedUrlId)
}

type ModerationBlockUserRequest struct {
	UserID string `json:"user_id"`
	//body
	TargetId string         `json:"target_id,omitempty"`
	UserIds  []string       `json:"user_ids,omitempty"`
	Users    []UserResponse `json:"users,omitempty"`
}

type ModerationBlockUserResponse struct {
	UserResponse
}

func (m ModerationBlockUserRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/block", baseURL, userEndpoint, m.UserID)
}

type ModerationUnblockUserRequest struct {
	UserID   string `json:"user_id"`
	TargetID string `json:"target_id"`
}

type ModerationUnblockUserResponse struct{}

func (m ModerationUnblockUserRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/block/%s", baseURL, userEndpoint, m.UserID, m.TargetID)
}

type ModerationListBlockedUsersRequest struct {
	UserID string `json:"user_id"`
	//optional
	Token            string `qs:"token,omitempty"`
	Limit            int    `qs:"limit,omitempty"`
	UserIds          string `qs:"user_ids,omitempty"`
	MetadataKey      string `qs:"metadatakey,omitempty"`
	MetadataValuesIn string `qs:"metadatavalues_in,omitempty"`
}

type ModerationListBlockedUsersResponse struct {
	Users []UserResponse `json:"users"`
	Next  string         `json:"next"`
}

func (m ModerationListBlockedUsersRequest) URL(baseURl string) string {
	base := fmt.Sprintf("%s/%s/%s/block", baseURl, userEndpoint, m.UserID)
	base, _ = encodeQS[ModerationListBlockedUsersRequest](base, m)
	return base
}

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
