package dto

import "fmt"

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
