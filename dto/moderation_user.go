package dto

import "fmt"

// ModerationUserResource is the resource definition for user moderation
type ModerationUserResource struct {
	UserID     string    `json:"user_id"`
	Nickname   string    `json:"nickname"`
	ProfileURL string    `json:"profile_url"`
	IsOnline   bool      `json:"is_online"`
	LastSeenAt int       `json:"last_seen_at"`
	Metadata   *Metadata `json:"metadata"`
}

// ModerationMuteUserInOpenChannelRequest is the request payload to mute user in an open channel
type ModerationMuteUserInOpenChannelRequest struct {
	ChannelURL string `json:"channel_url"`
	// body
	UserID string `json:"user_id"`
	// optional
	Seconds     int    `json:"seconds,omitempty"`
	Description string `json:"description,omitempty"`
}

// ModerationMuteUserInOpenChannelResponse returns OpenChannelResponse
type ModerationMuteUserInOpenChannelResponse OpenChannelResponse

// URL returns URL to ModerationMuteUserInOpenChannelRequest endpoint
func (m ModerationMuteUserInOpenChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/mute", baseURL, openChannelEndpoint, m.ChannelURL)
}

// ModerationMuteUserInGroupChannelRequest  is the request payload to mute user in a group channel
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

type ModerationMuteUserInCustomChannelResponse EmptyResponse

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

type ModerationUnmuteUserInOpenChannelResponse EmptyResponse

func (m ModerationUnmuteUserInOpenChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/mute/%s", baseURL, openChannelEndpoint, m.ChannelURL, m.MutedUserID)
}

// ModerationUnmuteUserInGroupChannelRequest unmutes member in a group channel request params
type ModerationUnmuteUserInGroupChannelRequest struct {
	ChannelURL  string `json:"channel_url"`
	MutedUserID string `json:"muted_user_id"`
}

// URL returns URL string to unmute member in group channel
func (m ModerationUnmuteUserInGroupChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/mute/%s", baseURL, groupChannelEndpoint, m.ChannelURL, m.MutedUserID)
}

// ModerationUnmuteUserInGroupChannelResponse is the response type to unmute user in a group channel
type ModerationUnmuteUserInGroupChannelResponse EmptyResponse

// ModerationUnmuteUsersInCustomChannelRequest  is the request URL params
type ModerationUnmuteUsersInCustomChannelRequest struct {
	CustomType string   `qs:"-"`
	UserIds    []string `qs:"user_ids,comma"`
}

// URL returns URL string to unmute users in custom channels with certain type at once
func (m ModerationUnmuteUsersInCustomChannelRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/mute", baseURL, customChannelEndpoint, m.CustomType)
	s, _ := encodeQS[ModerationUnmuteUsersInCustomChannelRequest](base, m)
	return s
}

// ModerationUnmuteUsersInCustomChannelResponse is the response type to unmute users in a custom channel type at once
type ModerationUnmuteUsersInCustomChannelResponse EmptyResponse

// ModerationListUsersMutedChannelsRequest is the request URL params, and query value to list channels where user is muted
type ModerationListUsersMutedChannelsRequest struct {
	UserID string `json:"user_id"`
	// optional
	Token string `qs:"token,omitempty"`
	Limit string `qs:"limit,omitempty"`
}

// ModerationListUsersMutedChannelsResponse is the response type to list of open and group channels where user is muted
type ModerationListUsersMutedChannelsResponse struct {
	MutedChannels []Channel `json:"muted_channels,omitempty"`
	Next          string    `json:"next,omitempty"`
	Error         *Error    `json:"error,omitempty"`
}

// URL returns URL string to list of open & group channels where the user is muted.
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

// ModerationListMutedUsersInGroupChannel is the response payload to list muted members in  group channel
type ModerationListMutedUsersInGroupChannel EmptyResponse

// ModerationListMutedUsersInGroupChannelRequest  is the request URL params, and query values to list muted members in
// a group channel.
type ModerationListMutedUsersInGroupChannelRequest struct {
	ChannelURL string `json:"channel_url"`
	//optional
	ShowTotalMuteCount bool   `qs:"show_total_mute_count,omitempty"`
	Token              string `qs:"token,omitempty"`
	Limit              int    `qs:"limit,omitempty"`
}

// ModerationListMutedUsersInGroupChannelResponse is the response payload to list muted members in group channel endpoint.
type ModerationListMutedUsersInGroupChannelResponse struct {
	MutedList      []ModerationUserResource `json:"muted_list"`
	TotalMuteCount int                      `json:"total_mute_count"`
	Next           string                   `json:"next"`
	Error          *Error                   `json:"error,omitempty"`
}

// URL returns URL string to list muted members in group channel endpoint
func (m ModerationListMutedUsersInGroupChannelRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/mute", baseURL, groupChannelEndpoint, m.ChannelURL)
	base, _ = encodeQS[ModerationListMutedUsersInGroupChannelRequest](base, m)
	return base
}

// ModerationListMutedUsersInCustomChannelRequest is the request URL param & query values to list muted users in custom channel
type ModerationListMutedUsersInCustomChannelRequest struct {
	CustomType string `json:"custom_type"`
	// optional
	Token string `json:"token,omitempty"`
	Limit int    `json:"limit,omitempty"`
}

// ModerationListMutedUsersInCustomChannelResponse is the response type
type ModerationListMutedUsersInCustomChannelResponse struct {
	MutedList []ModerationUserResource `json:"muted_list"`
	Next      string                   `json:"next"`
	Error     *Error                   `json:"error,omitempty"`
}

func (m ModerationListMutedUsersInCustomChannelRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/mute", baseURL, customChannelEndpoint, m.CustomType)
	base, _ = encodeQS[ModerationListMutedUsersInCustomChannelRequest](base, m)
	return base
}

// ModerationViewMutedUserInGroupChannelRequest is the request payload to view muted participant in open channel
type ModerationViewMutedUserInGroupChannelRequest struct {
	ChannelURL string `json:"channel_url"`
	MutedUrlId string `json:"muted_user_id"`
}

// ModerationViewMutedUserInGroupChannelResponse is the response payload to view muted participant in open channel
type ModerationViewMutedUserInGroupChannelResponse struct {
	IsMuted           bool   `json:"is_muted"`
	RemainingDuration int    `json:"remaining_duration"`
	StartAt           int    `json:"start_at"`
	EndAt             int    `json:"end_at"`
	Description       string `json:"description"`
	Error             *Error `json:"error,omitempty"`
}

// URL returns URL string to list muted users in group channel
func (m ModerationViewMutedUserInGroupChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/mute/%s", baseURL, groupChannelEndpoint, m.ChannelURL, m.MutedUrlId)
}

// ModerationViewMutedUserInOpenChannelRequest is the request URL params to list muted users in open channel
type ModerationViewMutedUserInOpenChannelRequest struct {
	ChannelURL string `json:"channel_url"`
	MutedUrlId string `json:"muted_user_id"`
}

// ModerationViewMutedUserInOpenChannelResponse is the response payload to list muted users in open channel API
type ModerationViewMutedUserInOpenChannelResponse struct {
	IsMuted           bool   `json:"is_muted"`
	RemainingDuration int    `json:"remaining_duration"`
	StartAt           int    `json:"start_at"`
	EndAt             int    `json:"end_at"`
	Description       string `json:"description"`
	Error             *Error `json:"error,omitempty"`
}

// URL returns URL string to view muted user in open channel endpoint
func (m ModerationViewMutedUserInOpenChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/mute/%s", baseURL, openChannelEndpoint, m.ChannelURL, m.MutedUrlId)
}

// ModerationBlockUserRequest is the request payload to block user
type ModerationBlockUserRequest struct {
	UserID string `json:"user_id"`
	//body
	TargetId string         `json:"target_id,omitempty"`
	UserIds  []string       `json:"user_ids,omitempty"`
	Users    []UserResponse `json:"users,omitempty"`
}

// ModerationBlockUserResponse is the response type to block user
type ModerationBlockUserResponse UserResponse

// URL returns URL string to block user request
func (m ModerationBlockUserRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/block", baseURL, userEndpoint, m.UserID)
}

// ModerationUnblockUserRequest is the request URL params for unblocking user
type ModerationUnblockUserRequest struct {
	UserID   string `json:"user_id"`
	TargetID string `json:"target_id"`
}

// ModerationUnblockUserResponse is an alias for EmptyResponse
type ModerationUnblockUserResponse EmptyResponse

// URL returns URL string to unblock user Sendbird API
func (m ModerationUnblockUserRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/block/%s", baseURL, userEndpoint, m.UserID, m.TargetID)
}

// ModerationListBlockedUsersRequest is the request URL param, query value, and / or payload to list blocked users
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
	Users []UserResponse `json:"users,omitempty"`
	Next  string         `json:"next,omitempty"`
	Error *Error         `json:"error,omitempty"`
}

func (m ModerationListBlockedUsersRequest) URL(baseURl string) string {
	base := fmt.Sprintf("%s/%s/%s/block", baseURl, userEndpoint, m.UserID)
	base, _ = encodeQS[ModerationListBlockedUsersRequest](base, m)
	return base
}

// ModerationBanUserFromOpenChannelRequest is the request URL param, and json body for banning user from a certain open channel
type ModerationBanUserFromOpenChannelRequest struct {
	ChannelURL string `json:"-"`
	// body
	UserID string `json:"user_id"`
	// optional
	AgentID     string `json:"agent_id"`
	Seconds     string `json:"seconds"`
	Description string `json:"description"`
}

// ModerationBanUserFromOpenChannelResponse is the response type for banning user from open channel
type ModerationBanUserFromOpenChannelResponse struct {
	User        UserResponse `json:"user"`
	StartAt     int          `json:"start_at"`
	EndAt       int          `json:"end_at"`
	Description string       `json:"description"`
	Error       *Error       `json:"error,omitempty"`
}

// URL returns the url to ban a user from an open channel.
// POST https://api-{application_id}.sendbird.com/v3/open_channels/{channel_url}/ban
func (m ModerationBanUserFromOpenChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/ban", baseURL, openChannelEndpoint, m.ChannelURL)
}

// ModerationBanUserFromGroupChannelRequest  is the request URL params and json body for banning user from a certain
// group channel
type ModerationBanUserFromGroupChannelRequest ModerationBanUserFromOpenChannelRequest

// ModerationBanUserFromGroupChannelResponse is the response payload for banning user from a certain group channel
type ModerationBanUserFromGroupChannelResponse ModerationBanUserFromOpenChannelResponse

// URL returns the url to ban a user from a group channel
// POST https://api-{application_id}.sendbird.com/v3/group_channels/{channel_url}/ban
func (m ModerationBanUserFromGroupChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/ban", baseURL, groupChannelEndpoint, m.ChannelURL)
}

// ModerationBanUserFromCustomChannelRequest is the request type to banning user from custom channels
type ModerationBanUserFromCustomChannelRequest struct {
	UserID string `json:"user_id"`
	// body
	ChannelCustomTypes []string `json:"channel_custom_types"`
}

// ModerationBanUserFromCustomChannelResponse is the response payload for banning user from custom channels
type ModerationBanUserFromCustomChannelResponse EmptyResponse

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

// ModerationUnbanUserFromOpenChannelRequest is the request URL params to unban user from an open channel
type ModerationUnbanUserFromOpenChannelRequest struct {
	ChannelURL   string `json:"channel_url"`
	BannedUserID string `json:"banned_user_id"`
}

// ModerationUnbanUserFromOpenChannelResponse is response type to unban user from an open channel
type ModerationUnbanUserFromOpenChannelResponse EmptyResponse

// URL returns the url to unban a user who had been banned from an open channel
// DELETE https://api-{application_id}.sendbird.com/v3/open_channels/{channel_url}/ban/{banned_user_id}
func (m ModerationUnbanUserFromOpenChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/ban/%s", baseURL, openChannelEndpoint, m.ChannelURL, m.BannedUserID)
}

// ModerationUnbanUserFromGroupChannelRequest  is the request URL params to unban user from group channel
type ModerationUnbanUserFromGroupChannelRequest ModerationUnbanUserFromOpenChannelRequest

// ModerationUnbanUserFromGroupChannelResponse is the response to unban user from group channel
type ModerationUnbanUserFromGroupChannelResponse EmptyResponse

// URL returns the url to unban a user who had been banned from an group channel
// DELETE https://api-{application_id}.sendbird.com/v3/group_channels/{channel_url}/ban/{banned_user_id}
func (m ModerationUnbanUserFromGroupChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/ban/%s", baseURL, groupChannelEndpoint, m.ChannelURL, m.BannedUserID)
}

// ModerationUnbanUserFromCustomChannelRequest is the request type to unban users from custom channel
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

// ModerationUnbanUserFromCustomChannelResponse is the response type to unban users from custom channel
type ModerationUnbanUserFromCustomChannelResponse EmptyResponse

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
	Error          *Error          `json:"error,omitempty"`
}

type BannedChannel struct {
	StartAt     int     `json:"start_at"`
	EndAt       int     `json:"end_at"`
	Description string  `json:"description"`
	Channel     Channel `json:"channel"`
	Next        string  `json:"next"`
}

// URL returns the url to retrieve a list of open channels and group channels where a user is banned
// GET https://api-{application_id}.sendbird.com/v3/users/{user_id}/ban
func (m ModerationListBannedUsersRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/ban", baseURL, userEndpoint, m.UserID)
	base, _ = encodeQS[ModerationListBannedUsersRequest](base, m)
	return base
}

// ModerationListBannedUsersInOpenChannelRequest is the request query values & URL params to list all
// banned users in an open channel
type ModerationListBannedUsersInOpenChannelRequest struct {
	ChannelURL string `json:"channel_url"`
	// optional
	Token             string `json:"token"`
	Limit             int    `json:"limit"`
	ShowTotalBanCount bool   `json:"show_total_ban_count"`
}

// ModerationListBannedUsersInOpenChannelResponse is the response payload to list all banned participants
// in an open channel
type ModerationListBannedUsersInOpenChannelResponse struct {
	BannedList    []BannedUser `json:"banned_list"`
	TotalBanCount int          `json:"total_ban_count"`
	Next          string       `json:"next"`
	Error         *Error       `json:"error,omitempty"`
}

// BannedUser is the banned user response type
type BannedUser struct {
	User        UserResponse `json:"user"`
	StartAt     int          `json:"start_at"`
	EndAt       int          `json:"end_at"`
	Description string       `json:"description"`
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

// ModerationListBannedUsersInCustomChannelRequest is the request type to list banned users in a custom channel
// with a custom type
type ModerationListBannedUsersInCustomChannelRequest struct {
	CustomType string `json:"custom_type"`
	// optional
	Token string `qs:"token"`
	Limit string `qs:"limit"`
}

// ModerationListBannedUsersInCustomChannelResponse  list banned users in a custom channel
type ModerationListBannedUsersInCustomChannelResponse struct {
	BannedList []BannedUser `json:"banned_list"`
	Next       string       `json:"next"`
	Error      *Error       `json:"error,omitempty"`
}

// URL returns the url to retrieve a list of users banned from channels with the specified custom channel type
// GET https://api-{application_id}.sendbird.com/v3/applications/settings_by_channel_custom_type/{custom_type}/ban
func (m ModerationListBannedUsersInCustomChannelRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/ban", baseURL, customChannelEndpoint, m.CustomType)
	base, _ = encodeQS[ModerationListBannedUsersInCustomChannelRequest](base, m)
	return base
}

// ModerationViewBannedUserInOpenChannelRequest is the request URL params for listing users
// who are banned from a specific open channel
type ModerationViewBannedUserInOpenChannelRequest struct {
	ChannelURL   string `json:"channel_url"`
	BannedUserID string `json:"banned_user_id"`
}

// ModerationViewBannedUserInOpenChannelResponse is the response type to list users who are banned
// from a specific open channel
type ModerationViewBannedUserInOpenChannelResponse struct {
	User        Channel `json:"user"`
	StartAt     int     `json:"start_at"`
	EndAt       int     `json:"end_at"`
	Description string  `json:"description"`
	Error       *Error  `json:"error,omitempty"`
}

// URL returns the url to retrieve information about a user banned from an open channel
// GET https://api-{application_id}.sendbird.com/v3/open_channels/{channel_url}/ban/{banned_user_id}
func (m ModerationViewBannedUserInOpenChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/ban/%s", baseURL, openChannelEndpoint, m.ChannelURL, m.BannedUserID)
}

// ModerationViewBannedUserInGroupChannelRequest is the request type to list banned users in a group channel
type ModerationViewBannedUserInGroupChannelRequest ModerationViewBannedUserInOpenChannelRequest

// ModerationViewBannedUserInGroupChannelResponse  is the response type to list banned users in a group channel
type ModerationViewBannedUserInGroupChannelResponse ModerationViewBannedUserInOpenChannelResponse

// URL returns the url to retrieve information about a user banned from a gorup channel
// GET https://api-{application_id}.sendbird.com/v3/group_channels/{channel_url}/ban/{banned_user_id}
func (m ModerationViewBannedUserInGroupChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/ban/%s", baseURL, groupChannelEndpoint, m.ChannelURL, m.BannedUserID)
}
