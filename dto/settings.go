package dto

// Settings is the response DTO to any request on custom channel / settings endpoint
type Settings struct {
	DomainFilter                   *DomainFilter                 `json:"domain_filter,omitempty"`
	ProfanityFilter                *ProfanityFilter              `json:"profanity_filter,omitempty"`
	MaxMessageLength               *int64                        `json:"max_message_length,omitempty"`
	DisplayPastMessage             *bool                         `json:"display_past_message,omitempty"`
	ImageModeration                *ImageModeration              `json:"image_moderation,omitempty"`
	AllowLinks                     *bool                         `json:"allow_links,omitempty"`
	UserMessagesPerChannelDuration *int64                        `json:"user_messages_per_channel_duration,omitempty"`
	UserMessagesPerChannel         *int64                        `json:"user_messages_per_channel,omitempty"`
	ProfanityTriggeredModeration   *ProfanityTriggeredModeration `json:"profanity_triggered_moderation,omitempty"`
	MessageRetentionHours          *int64                        `json:"message_retention_hours,omitempty"`
	Error                          Error                         `json:"error,omitempty"`
}

type DomainFilter struct {
	Domains           []string `json:"domains,omitempty"`
	Type              *int64   `json:"type,omitempty"`
	ShouldCheckGlobal *bool    `json:"should_check_global,omitempty"`
}

type ImageModeration struct {
	Type      *int64  `json:"type,omitempty"`
	SoftBlock *bool   `json:"soft_block,omitempty"`
	Limits    *Limits `json:"limits,omitempty"`
	CheckUrls *bool   `json:"check_urls,omitempty"`
}

type Limits struct {
	Adult    *int64 `json:"adult,omitempty"`
	Spoof    *int64 `json:"spoof,omitempty"`
	Medical  *int64 `json:"medical,omitempty"`
	Violence *int64 `json:"violence,omitempty"`
	Racy     *int64 `json:"racy,omitempty"`
}

type ProfanityFilter struct {
	Keywords          *string       `json:"keywords,omitempty"`
	RegexFilters      []RegexFilter `json:"regex_filters,omitempty"`
	Type              *int64        `json:"type,omitempty"`
	ShouldCheckGlobal *bool         `json:"should_check_global,omitempty"`
}

type RegexFilter struct {
	Regex *string `json:"regex,omitempty"`
}

type ProfanityTriggeredModeration struct {
	Count    *int64 `json:"count,omitempty"`
	Duration *int64 `json:"duration,omitempty"`
	Action   *int64 `json:"action,omitempty"`
}
