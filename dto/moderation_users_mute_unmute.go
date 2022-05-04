package dto

import "fmt"

type ModerationMuteUserInOpenChannel struct{}

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

type ModerationMuteUserInGroupChannel struct{}

type ModerationMuteUserInGroupChannelRequest struct {
	ChannelURL string `json:"channel_url"`
	// body
	UserID string `json:"user_id"`
	// optional
	Seconds     int    `json:"seconds,omitempty"`
	Description string `json:"description,omitempty"`
}

type ModerationMuteUserInGroupChannelResponse struct{}

func (m ModerationMuteUserInGroupChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/mute", baseURL, groupChannelEndpoint, m.ChannelURL)
}

type ModerationMuteUserInCustomChannel struct{}

type ModerationMuteUserInCustomChannelRequest struct {
	UserID string `json:"user_id"`
	// body
	ChannelCustomTypes []string `json:"channel_custom_types"`
}

type ModerationMuteUserInCustomChannelResponse struct{}

func (m ModerationMuteUserInCustomChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/muted_channel_custom_types", baseURL, userEndpoint, m.UserID)
}

type ModerationMuteUsersInCustomChannel struct{}

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

type ModerationUnmuteUserInOpenChannel struct{}

type ModerationUnmuteUserInOpenChannelRequest struct {
	ChannelURL  string `json:"channel_url"`
	MutedUserID string `json:"muted_user_id"`
}

type ModerationUnmuteUserInOpenChannelResponse struct{}

func (m ModerationUnmuteUserInOpenChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/mute/%s", baseURL, openChannelEndpoint, m.ChannelURL, m.MutedUserID)
}

type ModerationUnmuteUserInGroupChannel struct{}

type ModerationUnmuteUserInGroupChannelRequest struct {
	ChannelURL  string `json:"channel_url"`
	MutedUserID string `json:"muted_user_id"`
}

type ModerationUnmuteUserInGroupChannelResponse struct{}

func (m ModerationUnmuteUserInGroupChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/mute/%s", baseURL, groupChannelEndpoint, m.ChannelURL, m.MutedUserID)
}

type ModerationUnmuteUsersInCustomChannel struct{}

type ModerationUnmuteUsersInCustomChannelRequest struct {
	CustomType string `json:"custom_type"`
	// body
	UserIds []string `json:"user_ids"`
}

func (m ModerationUnmuteUsersInCustomChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/mute", baseURL, customChannelEndpoint, m.CustomType)
}
