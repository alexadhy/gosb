package dto

import "fmt"

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
