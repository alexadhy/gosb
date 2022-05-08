package dto

import "fmt"

type ModerationDomainFilterApplicationWideRequest struct {
	CustomType   string       `json:"custom_type"`
	DomainFilter DomainFilter `json:"domain_filter"`
}

type ModerationFilterApplicationWideResponse struct {
	DomainFilter                 DomainFilter                 `json:"domain_filter"`
	ProfanityFilter              ProfanityFilter              `json:"profanity_filter"`
	ProfanityTriggeredModeration ProfanityTriggeredModeration `json:"profanity_triggered_moderation"`
	ImageModeration              ImageModeration              `json:"image_moderation"`
}

// URL returns the url to filter out certain domains contained in text and file messages
// as well as users' profile images according to your policies and criteria
// PUT https://api-{application_id}.sendbird.com/v3/applications/settings_global
func (m ModerationDomainFilterApplicationWideRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s", baseURL, globalSettingsEndpoint)
}

type ModerationDomainFilterCustomChannelRequest ModerationDomainFilterApplicationWideRequest

type ModerationDomainFilterCustomChannelResponse struct {
	Channel
}

// To apply the filter to channels with a custom channel type.
// PUT https://api-{application_id}.sendbird.com/v3/applications/settings_by_channel_custom_type/{custom_type}
func (m ModerationDomainFilterCustomChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s", baseURL, customChannelEndpoint, m.CustomType)
}

type ModerationProfanityFilterApplicationWideRequest struct {
	CustomType      string          `json:"custom_type"`
	ProfanityFilter ProfanityFilter `json:"profanity_filter"`
}

type ModerationProfanityFilterApplicationWideResponse ModerationFilterApplicationWideResponse

// URL returns the url to detect and filter out profanity word or regular expressions
// contained in text and file messages according to your policies and criteria
// PUT https://api-{application_id}.sendbird.com/v3/applications/settings_global
func (m ModerationProfanityFilterApplicationWideRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s", baseURL, globalSettingsEndpoint)
}

type ModerationProfanityFilterCustomChannelRequest ModerationProfanityFilterApplicationWideRequest

type ModerationProfanityFilterCustomChannelResponse struct {
	Channel
}

// To apply the filter to channels with a custom channel type.
// PUT https://api-{application_id}.sendbird.com/v3/applications/settings_by_channel_custom_type/{custom_type}
func (m ModerationProfanityFilterCustomChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s", baseURL, customChannelEndpoint, m.CustomType)
}

type ModerationProfanityTriggeredApplicationWideRequest struct {
	CustomType                   string                       `json:"custom_type"`
	ProfanityTriggeredModeration ProfanityTriggeredModeration `json:"profanity_triggered_moderation"`
}

type ModerationProfanityTriggeredApplicationWideResponse ModerationFilterApplicationWideResponse

func (m ModerationProfanityTriggeredApplicationWideRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s", baseURL, globalSettingsEndpoint)
}

type ModerationProfanityTriggeredCustomChannelRequest ModerationProfanityTriggeredApplicationWideRequest

type ModerationProfanityTriggeredCustomChannelResponse struct {
	Channel
}

func (m ModerationProfanityTriggeredCustomChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s", baseURL, customChannelEndpoint, m.CustomType)
}

type ModerationImageApplicationWideRequest struct {
	CustomType                   string                       `json:"custom_type"`
	ProfanityTriggeredModeration ProfanityTriggeredModeration `json:"profanity_triggered_moderation"`
}

type ModerationImageApplicationWideResponse ModerationFilterApplicationWideResponse

// To apply the settings application-wide.
// PUT https://api-{application_id}.sendbird.com/v3/applications/settings_global
func (m ModerationImageApplicationWideRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s", baseURL, globalSettingsEndpoint)
}

type ModerationImageCustomChannelRequest ModerationImageApplicationWideRequest

type ModerationImageCustomChannelResponse struct {
	Channel
}

// To apply the filter to channels with a custom channel type.
// PUT https://api-{application_id}.sendbird.com/v3/applications/settings_by_channel_custom_type/{custom_type}
func (m ModerationImageCustomChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s", baseURL, customChannelEndpoint, m.CustomType)
}
