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

type ModerationUserMuteInOpenChannelRequest struct {
	ChannelURL string `json:"channel_url"`
	// body
	UserID string `json:"user_id"`
	// optional
	Seconds     int    `json:"seconds,omitempty"`
	Description string `json:"description,omitempty"`
}

type ModerationUserMuteInOpenChannelResponse struct {
	OpenChannelResponse
}

func (m ModerationUserMuteInOpenChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/mute", baseURL, openChannelEndpoint, m.ChannelURL)
}

type ModerationUserMuteInGroupChannelRequest struct {
	ChannelURL string `json:"channel_url"`
	// body
	UserID string `json:"user_id"`
	// optional
	Seconds     int    `json:"seconds,omitempty"`
	Description string `json:"description,omitempty"`
}

func (m ModerationUserMuteInGroupChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/mute", baseURL, groupChannelEndpoint, m.ChannelURL)
}

type ModerationUserMuteInCustomChannelRequest struct {
	UserID string `json:"user_id"`
	// body
	ChannelCustomTypes []string `json:"channel_custom_types"`
}

type ModerationUserMuteInCustomChannelResponse struct{}

func (m ModerationUserMuteInCustomChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/muted_channel_custom_types", baseURL, userEndpoint, m.UserID)
}

type ModerationUsersMuteInCustomChannelRequest struct {
	CustomType string `json:"custom_type"`
	// body
	UserIds []string `json:"user_ids"`
	// optional
	Seconds        int    `json:"seconds,omitempty"`
	Description    string `json:"description,omitempty"`
	OnDemandUpsert bool   `json:"on_demand_upsert,omitempty"`
}

type ModerationUsersMuteInCustomChannelResponse struct{}

func (m ModerationUsersMuteInCustomChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/mute", baseURL, customChannelEndpoint, m.CustomType)
}

type ModerationUserUnmuteInOpenChannelRequest struct {
	ChannelURL  string `json:"channel_url"`
	MutedUserID string `json:"muted_user_id"`
}

type ModerationUserUnmuteInOpenChannelResponse struct{}

func (m ModerationUserUnmuteInOpenChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/mute/%s", baseURL, openChannelEndpoint, m.ChannelURL, m.MutedUserID)
}

type ModerationUserUnmuteInGroupChannelRequest struct {
	ChannelURL  string `json:"channel_url"`
	MutedUserID string `json:"muted_user_id"`
}

func (m ModerationUserUnmuteInGroupChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/mute/%s", baseURL, groupChannelEndpoint, m.ChannelURL, m.MutedUserID)
}

type ModerationUsersUnmuteInCustomChannelRequest struct {
	CustomType string `json:"custom_type"`
	// body
	UserIds []string `json:"user_ids"`
}

func (m ModerationUsersUnmuteInCustomChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/mute", baseURL, customChannelEndpoint, m.CustomType)
}

type ModerationUserListMutedChannelsRequest struct {
	UserID string `json:"user_id"`
	// optional
	Token string `qs:"token,omitempty"`
	Limit string `qs:"limit,omitempty"`
}

type ModerationUserListMutedChannelsResponse struct {
	MutedChannels []Channel `json:"muted_channels"`
	Next          string    `json:"next"`
}

func (m ModerationUserListMutedChannelsRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/mute", baseURL, userEndpoint, m.UserID)
	base, _ = encodeQS[ModerationUserListMutedChannelsRequest](base, m)
	return base
}

type ModerationUsersListMutedInOpenChannelRequest struct {
	ChannelURL string `json:"channel_url"`
	//optional
	ShowTotalMuteCount bool   `qs:"show_total_mute_count,omitempty"`
	Token              string `qs:"token,omitempty"`
	Limit              int    `qs:"limit,omitempty"`
}

type ModerationUsersListMutedInOpenChannelResponse struct {
	MutedList      []ModerationUserResource `json:"muted_list"`
	TotalMuteCount int                      `json:"total_mute_count"`
	Next           string                   `json:"next"`
}

func (m ModerationUsersListMutedInOpenChannelRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/mute", baseURL, openChannelEndpoint, m.ChannelURL)
	base, _ = encodeQS[ModerationUsersListMutedInOpenChannelRequest](base, m)
	return base
}

type ModerationUsersListMutedInGroupChannel struct{}

type ModerationUsersListMutedInGroupChannelRequest struct {
	ChannelURL string `json:"channel_url"`
	//optional
	ShowTotalMuteCount bool   `qs:"show_total_mute_count,omitempty"`
	Token              string `qs:"token,omitempty"`
	Limit              int    `qs:"limit,omitempty"`
}

type ModerationUsersListMutedInGroupChannelResponse struct {
	MutedList      []ModerationUserResource `json:"muted_list"`
	TotalMuteCount int                      `json:"total_mute_count"`
	Next           string                   `json:"next"`
}

func (m ModerationUsersListMutedInGroupChannelRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/mute", baseURL, groupChannelEndpoint, m.ChannelURL)
	base, _ = encodeQS[ModerationUsersListMutedInGroupChannelRequest](base, m)
	return base
}

type ModerationUsersListMutedInCustomChannelRequest struct {
	CustomType string `json:"custom_type"`
	// optional
	Token string `json:"token,omitempty"`
	Limit int    `json:"limit,omitempty"`
}

type ModerationUsersListMutedInCustomChannelResponse struct {
	MutedList []ModerationUserResource `json:"muted_list"`
	Next      string                   `json:"next"`
}

func (m ModerationUsersListMutedInCustomChannelRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/mute", baseURL, customChannelEndpoint, m.CustomType)
	base, _ = encodeQS[ModerationUsersListMutedInCustomChannelRequest](base, m)
	return base
}

type ModerationUserListMutedInGroupChannel struct{}

type ModerationUserListMutedInGroupChannelRequest struct {
	ChannelURL string `json:"channel_url"`
	MutedUrlId string `json:"muted_user_id"`
}

type ModerationUserListMutedInGroupChannelResponse struct {
	IsMuted           bool   `json:"is_muted"`
	RemainingDuration int    `json:"remaining_duration"`
	StartAt           int    `json:"start_at"`
	EndAt             int    `json:"end_at"`
	Description       string `json:"description"`
}

func (m ModerationUserListMutedInGroupChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/mute/%s", baseURL, groupChannelEndpoint, m.ChannelURL, m.MutedUrlId)
}

type ModerationUserListMutedInOpenChannelRequest struct {
	ChannelURL string `json:"channel_url"`
	MutedUrlId string `json:"muted_user_id"`
}

type ModerationUserListMutedInOpenChannelResponse struct {
	IsMuted           bool   `json:"is_muted"`
	RemainingDuration int    `json:"remaining_duration"`
	StartAt           int    `json:"start_at"`
	EndAt             int    `json:"end_at"`
	Description       string `json:"description"`
}

func (m ModerationUserListMutedInOpenChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/mute/%s", baseURL, openChannelEndpoint, m.ChannelURL, m.MutedUrlId)
}

type ModerationUserBlockRequest struct {
	UserID string `json:"user_id"`
	//body
	TargetId string         `json:"target_id,omitempty"`
	UserIds  []string       `json:"user_ids,omitempty"`
	Users    []UserResponse `json:"users,omitempty"`
}

type ModerationUserBlockResponse struct {
	UserResponse
}

func (m ModerationUserBlockRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/block", baseURL, userEndpoint, m.UserID)
}

type ModerationUserUnblockRequest struct {
	UserID   string `json:"user_id"`
	TargetID string `json:"target_id"`
}

type ModerationUserUnblockResponse struct{}

func (m ModerationUserUnblockRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/block/%s", baseURL, userEndpoint, m.UserID, m.TargetID)
}

type ModerationUsersListBlockedRequest struct {
	UserID string `json:"user_id"`
	//optional
	Token            string `qs:"token,omitempty"`
	Limit            int    `qs:"limit,omitempty"`
	UserIds          string `qs:"user_ids,omitempty"`
	MetadataKey      string `qs:"metadatakey,omitempty"`
	MetadataValuesIn string `qs:"metadatavalues_in,omitempty"`
}

type ModerationUsersListBlockedResponse struct {
	Users []UserResponse `json:"users"`
	Next  string         `json:"next"`
}

func (m ModerationUsersListBlockedRequest) URL(baseURl string) string {
	base := fmt.Sprintf("%s/%s/%s/block", baseURl, userEndpoint, m.UserID)
	base, _ = encodeQS[ModerationUsersListBlockedRequest](base, m)
	return base
}

type ModerationUserBanFromOpenChannelRequest struct {
	ChannelURL string `json:"channel_url"`
	// body
	UserID string `json:"user_id"`
	// optional
	AgentID     string `json:"agent_id"`
	Seconds     string `json:"seconds"`
	Description string `json:"description"`
}

type ModerationUserBanFromOpenChannelResponse struct {
	User        SimplifiedUser `json:"user"`
	StartAt     int            `json:"start_at"`
	EndAt       int            `json:"end_at"`
	Description string         `json:"description"`
}

// URL returns the url to ban a user from an open channel.
// POST https://api-{application_id}.sendbird.com/v3/open_channels/{channel_url}/ban
func (m ModerationUserBanFromOpenChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/ban", baseURL, openChannelEndpoint, m.ChannelURL)
}

type ModerationUserBanFromGroupChannelRequest ModerationUserBanFromOpenChannelRequest

type ModerationUserBanFromGroupChannelResponse ModerationUserBanFromOpenChannelResponse

// URL returns the url to ban a user from a group channel
// POST https://api-{application_id}.sendbird.com/v3/group_channels/{channel_url}/ban
func (m ModerationUserBanFromGroupChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/ban", baseURL, groupChannelEndpoint, m.ChannelURL)
}

type ModerationUserBanFromCustomChannelRequest struct {
	UserID string `json:"user_id"`
	// body
	ChannelCustomTypes []string `json:"channel_custom_types"`
}

type ModerationUserBanFromCustomChannelResponse struct{}

// URL returns the url to ban a user from channels with the specified custom channel types
// POST https://api-{application_id}.sendbird.com/v3/users/{user_id}/banned_channel_custom_types
func (m ModerationUserBanFromCustomChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/banned_channel_custom_types", baseURL, userEndpoint, m.UserID)
}

type ModerationUsersBanFromCustomChannelRequest struct {
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

type ModerationUsersBanFromCustomChannelResponse struct{}

// URL returns the url to ban users from channels with the specified custom channel type at once
// POST https://api-{application_id}.sendbird.com/v3/applications/settings_by_channel_custom_type/{custom_type}/ban
func (m ModerationUsersBanFromCustomChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/ban", baseURL, customChannelEndpoint, m.CustomType)
}

type ModerationUserUpdateBannedStatusInOpenChannelRequest struct {
	ChannelURL   string `json:"channel_url"`
	BannedUserID string `json:"banned_user_id"`
	// body
	Seconds int `json:"seconds"`
	// optional
	Description string `json:"description"`
}

type ModerationUserUpdateBannedStatusInOpenChannelResponse ModerationUserBanFromOpenChannelResponse

// URL returns the url to update the information about a banned user of an open channel
// PUT https://api-{application_id}.sendbird.com/v3/open_channels/{channel_url}/ban/{banned_user_id}
func (m ModerationUserUpdateBannedStatusInOpenChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/ban/%s", baseURL, openChannelEndpoint, m.ChannelURL, m.BannedUserID)
}

type ModerationUserUpdateBannedStatusInGroupChannelRequest ModerationUserUpdateBannedStatusInOpenChannelRequest

type ModerationUserUpdateBannedStatusInGroupChannelResponse ModerationUserBanFromOpenChannelResponse

// URL returns the url to update information about a banned user of a group channel
// PUT https://api-{application_id}.sendbird.com/v3/group_channels/{channel_url}/ban/{banned_user_id}
func (m ModerationUserUpdateBannedStatusInGroupChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/ban/%s", baseURL, groupChannelEndpoint, m.ChannelURL, m.BannedUserID)
}

type ModerationUserUnbanFromOpenChannelRequest struct {
	ChannelURL   string `json:"channel_url"`
	BannedUserID string `json:"banned_user_id"`
}

type ModerationUserUnbanFromOpenChannelResponse struct{}

// URL returns the url to unban a user who had been banned from an open channel
// DELETE https://api-{application_id}.sendbird.com/v3/open_channels/{channel_url}/ban/{banned_user_id}
func (m ModerationUserUnbanFromOpenChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/ban/%s", baseURL, openChannelEndpoint, m.ChannelURL, m.BannedUserID)
}

type ModerationUserUnbanFromGroupChannelRequest ModerationUserUnbanFromOpenChannelRequest

type ModerationUserUnbanFromGroupChannelResponse struct{}

// URL returns the url to unban a user who had been banned from an group channel
// DELETE https://api-{application_id}.sendbird.com/v3/group_channels/{channel_url}/ban/{banned_user_id}
func (m ModerationUserUnbanFromGroupChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/ban/%s", baseURL, groupChannelEndpoint, m.ChannelURL, m.BannedUserID)
}

type ModerationUserUnbanFromCustomChannelRequest struct {
	CustomType string `json:"custom_type"`
	// parameters
	UserIDs []string `json:"user_ids"`
}

// URL returns the url to unban users who had been banned from channels with the specified channel type at once
// DELETE https://api-{application_id}.sendbird.com/v3/applications/settings_by_channel_custom_type/{custom_type}/ban
func (m ModerationUserUnbanFromCustomChannelRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/ban", baseURL, customChannelEndpoint, m.CustomType)
	base, _ = encodeQS[ModerationUserUnbanFromCustomChannelRequest](base, m)
	return base
}

// ModerationUsersListBannedRequest stores the parameters the request supports
type ModerationUsersListBannedRequest struct {
	UserID string `json:"user_id"`
	//optional
	Token string `qs:"token"`
	Limit string `qs:"limit"`
}

// ModerationUsersListBannedResponse stores the list of channels
type ModerationUsersListBannedResponse struct {
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
func (m ModerationUsersListBannedRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/ban", baseURL, userEndpoint, m.UserID)
	base, _ = encodeQS[ModerationUsersListBannedRequest](base, m)
	return base
}

type ModerationUsersListBannedInOpenChannelRequest struct {
	ChannelURL string `json:"channel_url"`
	// optional
	Token             string `json:"token"`
	Limit             int    `json:"limit"`
	ShowTotalBanCount bool   `json:"show_total_ban_count"`
}

type ModerationUsersListBannedInOpenChannelResponse struct {
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
func (m ModerationUsersListBannedInOpenChannelRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/ban", baseURL, openChannelEndpoint, m.ChannelURL)
	base, _ = encodeQS[ModerationUsersListBannedInOpenChannelRequest](base, m)
	return base
}

type ModerationUsersListBannedInGroupChannelRequest ModerationUsersListBannedInOpenChannelRequest

type ModerationUsersListBannedInGroupChannelResponse ModerationUsersListBannedInOpenChannelResponse

// URL returns the url to retrieve the list of users who are banned from a group channel
// GET https://api-{application_id}.sendbird.com/v3/group_channels/{channel_url}/ban
func (m ModerationUsersListBannedInGroupChannelRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/ban", baseURL, groupChannelEndpoint, m.ChannelURL)
	base, _ = encodeQS[ModerationUsersListBannedInGroupChannelRequest](base, m)
	return base
}

type ModerationUsersListBannedInCustomChannelRequest struct {
	CustomType string `json:"custom_type"`
	// optional
	Token string `json:"token"`
	Limit string `json:"limit"`
}

type ModerationUsersListBannedInCustomChannelResponse struct {
	BannedList []BannedUser `json:"banned_list"`
	Next       string       `json:"next"`
}

// URL returns the url to retrieve a list of users banned from channels with the specified custom channel type
// GET https://api-{application_id}.sendbird.com/v3/applications/settings_by_channel_custom_type/{custom_type}/ban
func (m ModerationUsersListBannedInCustomChannelRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/ban", baseURL, customChannelEndpoint, m.CustomType)
	base, _ = encodeQS[ModerationUsersListBannedInCustomChannelRequest](base, m)
	return base
}

type ModerationUserListBannedInOpenChannelRequest struct {
	ChannelURL   string `json:"channel_url"`
	BannedUserID string `json:"banned_user_id"`
}

type ModerationUserListBannedInOpenChannelResponse struct {
	User        SimplifiedChannel `json:"user"`
	StartAt     int               `json:"start_at"`
	EndAt       int               `json:"end_at"`
	Description string            `json:"description"`
}

// URL returns the url to retrieve information about a user banned from an open channel
// GET https://api-{application_id}.sendbird.com/v3/open_channels/{channel_url}/ban/{banned_user_id}
func (m ModerationUserListBannedInOpenChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/ban/%s", baseURL, openChannelEndpoint, m.ChannelURL, m.BannedUserID)
}

type ModerationUserListBannedInGroupChannelRequest ModerationUserListBannedInOpenChannelRequest

type ModerationUserListBannedInGroupChannelResponse ModerationUserListBannedInOpenChannelResponse

// URL returns the url to retrieve information about a user banned from a gorup channel
// GET https://api-{application_id}.sendbird.com/v3/group_channels/{channel_url}/ban/{banned_user_id}
func (m ModerationUserListBannedInGroupChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/ban/%s", baseURL, groupChannelEndpoint, m.ChannelURL, m.BannedUserID)
}
