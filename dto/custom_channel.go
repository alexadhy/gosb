package dto

import (
	"fmt"
)

// CustomChannelListOperatorsRequest is the DTO struct for listing custom channel operators
type CustomChannelListOperatorsRequest struct {
	CustomType string  `qs:"-"`
	Token      *string `qs:"token,omitempty"`
	Limit      *int    `qs:"limit,omitempty"`
}

// URL creates custom channel list operator URL string
func (o CustomChannelListOperatorsRequest) URL(baseURL string) string {
	base := fmt.Sprintf(
		"%s/%s/%s/operators",
		baseURL,
		customChannelEndpoint,
		o.CustomType,
	)
	s, err := encodeQS[CustomChannelListOperatorsRequest](base, o)
	if err != nil {
		return ""
	}
	return s
}

// CustomChannelListOperatorsResponse is the return response for listing operator endpoint
type CustomChannelListOperatorsResponse struct {
	Operators     []UserResponse `json:"operators,omitempty"`
	NextPageToken *string        `json:"next,omitempty"`
	*Error
}

// CustomChannelRegisterOperatorRequest is the request payload for registering operator to an open channel
type CustomChannelRegisterOperatorRequest struct {
	ChannelType    string   `json:"-"`
	OperatorIDs    []string `json:"operator_ids"`
	OnDemandUpsert *bool    `json:"on_demand_upsert,omitempty"`
}

// URL creates custom channel register operator URL string
func (o CustomChannelRegisterOperatorRequest) URL(baseURL string) string {
	return fmt.Sprintf(
		"%s/%s/%s/operators",
		baseURL,
		customChannelEndpoint,
		o.ChannelType,
	)
}

// CustomChannelUnregisterOperatorRequest creates a request query string to unregister custom channel operator endpoint.
type CustomChannelUnregisterOperatorRequest struct {
	ChannelType string   `qs:"-"`
	OpIDs       []string `qs:"operator_ids,comma"`
}

// URL returns URL string for custom channel unregister operator endpoint
func (o CustomChannelUnregisterOperatorRequest) URL(baseURL string) string {
	base := fmt.Sprintf(
		"%s/%s/%s/operators",
		baseURL,
		customChannelEndpoint,
		o.ChannelType,
	)
	s, err := encodeQS[CustomChannelUnregisterOperatorRequest](base, o)
	if err != nil {
		return ""
	}
	return s
}

// CustomChannelListSettingsRequest is the query string to list all settings in custom channel
type CustomChannelListSettingsRequest struct {
	Token *string `qs:"token,omitempty"`
	Limit *int    `qs:"limit,omitempty"`
}

// URL returns URL string to custom channel list settings endpoint
// GET https://api-{application_id}.sendbird.com/v3/applications/settings_by_channel_custom_type
func (o CustomChannelListSettingsRequest) URL(baseURL string) string {
	base := fmt.Sprintf(
		"%s/%s",
		baseURL,
		customChannelEndpoint,
	)
	s, err := encodeQS[CustomChannelListSettingsRequest](base, o)
	if err != nil {
		return ""
	}
	return s
}

// CustomChannelListSettingsResponse is the response to list custom channel settings endpoint
type CustomChannelListSettingsResponse struct {
	SettingList   []Settings `json:"channel_custom_type_settings,omitempty"`
	NextPageToken *string    `json:"token,omitempty"`
	*Error
}

// CustomChannelGetSettingsRequest creates a request to get custom channel's settings
type CustomChannelGetSettingsRequest struct {
	CustomType string `qs:"-" json:"-"`
}

// URL returns URL string to viww custom channel settings
// GET https://api-{application_id}.sendbird.com/v3/applications/settings_by_channel_custom_type/{custom_type}
func (o CustomChannelGetSettingsRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s", baseURL, customChannelEndpoint, o.CustomType)
}

// CustomChannelCreateSettingRequest creates a custom channel setting request
type CustomChannelCreateSettingRequest struct {
	CustomType            string           `json:"custom_type"`
	DisplayPastMessage    *bool            `json:"display_past_message,omitempty"`
	MessageRetentionHours *int64           `json:"message_retention_hours,omitempty"`
	DomainFilter          *DomainFilter    `json:"domain_filter,omitempty"`
	ProfanityFilter       *ProfanityFilter `json:"profanity_filter,omitempty"`
}

// URL returns URL string to custom channel create setting endpoint
func (o CustomChannelCreateSettingRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s", baseURL, customChannelEndpoint)
}

// CustomChannelUpdateSettingRequest creates a custom channel setting request
type CustomChannelUpdateSettingRequest struct {
	CustomType            string           `json:"-"`
	DisplayPastMessage    *bool            `json:"display_past_message,omitempty"`
	MessageRetentionHours *int64           `json:"message_retention_hours,omitempty"`
	DomainFilter          *DomainFilter    `json:"domain_filter,omitempty"`
	ProfanityFilter       *ProfanityFilter `json:"profanity_filter,omitempty"`
}

// URL returns URL string for custom channel update setting endpoint
func (o CustomChannelUpdateSettingRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s", baseURL, customChannelEndpoint, o.CustomType)
}

// CustomChannelDeleteSettingRequest creates a custom channel setting delete request
type CustomChannelDeleteSettingRequest struct {
	CustomType string `json:"-"`
}

// URL returns URL string for custom channel delete endpoint
func (o CustomChannelDeleteSettingRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s", baseURL, customChannelEndpoint, o.CustomType)
}

// CustomChannelDefaultInvitationPreference is the response type to default invitation preference endpoint
type CustomChannelDefaultInvitationPreference struct {
	AutoAccept *bool `json:"auto_accept,omitempty"`
	*Error
}

// URL returns URL string for custom channel default invitation preference response
func (o CustomChannelDefaultInvitationPreference) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s", baseURL, defaultInvitationEndpoint)
}
