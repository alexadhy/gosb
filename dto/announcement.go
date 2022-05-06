package dto

import "fmt"

// Announcement is the resource type for SendBird's announcement object
type Announcement struct {
	UniqueID           *string  `json:"unique_id,omitempty"`
	AnnouncementGroup  *string  `json:"announcement_group,omitempty"`
	Message            *Message `json:"message,omitempty"`
	EnablePush         *bool    `json:"enable_push,omitempty"`
	TargetAt           *string  `json:"target_at,omitempty"`
	TargetUserCount    *int64   `json:"target_user_count,omitempty"`
	TargetChannelCount *int64   `json:"target_channel_count,omitempty"`
	Status             *string  `json:"status,omitempty"`
	ScheduledAt        *int64   `json:"scheduled_at,omitempty"`
	CeaseAt            *string  `json:"cease_at,omitempty"`
	ResumeAt           *string  `json:"resume_at,omitempty"`
	CompletedAt        *int64   `json:"completed_at,omitempty"`
	SentUserCount      *int64   `json:"sent_user_count,omitempty"`
	OpenCount          *int64   `json:"open_count,omitempty"`
	OpenRate           *int64   `json:"open_rate,omitempty"`
	Error              *Error   `json:"error,omitempty"`
}

// AnnouncementListResponse is the response type to list announcement endpoint
type AnnouncementListResponse struct {
	Announcements []Announcement `json:"announcements,omitempty"`
	NextPageToken *string        `json:"next,omitempty"`
	Error         *Error         `json:"error,omitempty"`
}

// AnnouncementListRequest is the request params & query value(s) to list announcements endpoint
type AnnouncementListRequest struct {
	Token             *string `qs:"token,omitempty"`
	Limit             *int    `qs:"limit,omitempty"`
	Order             *string `qs:"order,omitempty"`
	Status            *string `qs:"status,omitempty"`
	AnnouncementGroup *string `qs:"announcement_group,omitempty"`
}

// URL returns URL string to list announcements
func (a AnnouncementListRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s", baseURL, announcementEndpoint)
	s, _ := encodeQS[AnnouncementListRequest](base, a)
	return s
}

// AnnouncementGroupListRequest is the request URL params to list announcement group
type AnnouncementGroupListRequest struct {
	Token *string `qs:"token,omitempty"`
	Limit *int    `qs:"limit,omitempty"`
}

// URL returns URL string to list announcement groups
func (a AnnouncementGroupListRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/announcement_group", baseURL)
	s, _ := encodeQS[AnnouncementGroupListRequest](base, a)
	return s
}

// AnnouncementGroupListResponse is the response type to list announcement_group endpoint
type AnnouncementGroupListResponse struct {
	AnnouncementGroup []string `json:"announcement_group,omitempty"`
	NextPageToken     *string  `json:"next,omitempty"`
	Error             *Error   `json:"error,omitempty"`
}

// AnnouncementViewRequest is the request URL params needed to call view announcement endpoint
type AnnouncementViewRequest struct {
	UniqueID string `json:"-"`
}

// URL returns URL string to view single announcement
// GET https://api-{application_id}.sendbird.com/v3/announcements/{unique_id}
func (a AnnouncementViewRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s", baseURL, announcementEndpoint, a.UniqueID)
}

// AnnouncementViewResponse is the response type to view announcement request
type AnnouncementViewResponse Announcement

// AnnouncementCreateRequest is the request payload to create an announcement
type AnnouncementCreateRequest struct {
	UniqueID             *string                           `json:"-"`
	AnnouncementGroup    *string                           `json:"announcement_group,omitempty"`
	Message              *AnnouncementCreateMessageOption  `json:"message,omitempty"`
	EnablePush           *bool                             `json:"enable_push,omitempty"`
	TargetAt             *string                           `json:"target_at,omitempty"`
	TargetList           []string                          `json:"target_list,omitempty"`
	TargetChannelType    *string                           `json:"target_channel_type,omitempty"`
	CreateChannel        *bool                             `json:"create_channel,omitempty"`
	CreateChannelOptions *AnnouncementCreateChannelOptions `json:"create_channel_options,omitempty"`
	ScheduledAt          *int64                            `json:"scheduled_at,omitempty"`
	CeaseAt              *string                           `json:"cease_at,omitempty"`
	ResumeAt             *string                           `json:"resume_at,omitempty"`
}

// AnnouncementCreateChannelOptions is the request payload to optionally create channel on making announcement
type AnnouncementCreateChannelOptions struct {
	Name       *string `json:"name,omitempty"`
	CoverURL   *string `json:"cover_url,omitempty"`
	CustomType *string `json:"custom_type,omitempty"`
	Data       *string `json:"data,omitempty"`
	Distinct   *bool   `json:"distinct,omitempty"`
}

// AnnouncementCreateMessageOption is the request payload for creating announcement message
// MessageType acceptable value is MESG or ADMM
type AnnouncementCreateMessageOption struct {
	MessageType string `json:"message_type"`
	UserID      string `json:"user_id"`
	Content     string `json:"content"`
	Data        string `json:"data"`
}

// URL returns URL string to create single announcement
// POST https://api-{application_id}.sendbird.com/v3/announcements
func (a AnnouncementCreateRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s", baseURL, announcementEndpoint)
}

// AnnouncementCreateResponse will return Announcement on success or Error on any error
type AnnouncementCreateResponse Announcement

// AnnouncementUpdateRequest is the request params and payload to send to update announcement endpoint
type AnnouncementUpdateRequest struct {
	UniqueID             string                            `json:"-"`
	Action               *AnnouncementAction               `json:"action,omitempty"`
	AnnouncementGroup    *string                           `json:"announcement_group,omitempty"`
	TargetCustomType     string                            `json:"target_custom_type,omitempty"`
	Message              *AnnouncementCreateMessageOption  `json:"message,omitempty"`
	EnablePush           *bool                             `json:"enable_push,omitempty"`
	CreateChannel        *bool                             `json:"create_channel,omitempty"`
	CreateChannelOptions *AnnouncementCreateChannelOptions `json:"create_channel_options,omitempty"`
	ScheduledAt          *int64                            `json:"scheduled_at,omitempty"`
	CeaseAt              *string                           `json:"cease_at,omitempty"`
	ResumeAt             *string                           `json:"resume_at,omitempty"`
}

// URL returns URL string to update announcement endpoint
// PUT https://api-{application_id}.sendbird.com/v3/announcements/{unique_id}
func (a AnnouncementUpdateRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s", baseURL, announcementEndpoint, a.UniqueID)
}

// AnnouncementStatisticPeriod can be between daily, weekly, and monthly
type AnnouncementStatisticPeriod string

const (
	STATISTIC_DAILY   = AnnouncementStatisticPeriod("daily")
	STATISTIC_WEEKLY  = AnnouncementStatisticPeriod("weekly")
	STATISTIC_MONTHLY = AnnouncementStatisticPeriod("monthly")
)

// AnnouncementViewStatisticsRequest is the URL query values and params for listing statistics to announcement group
type AnnouncementViewStatisticsRequest struct {
	Period            AnnouncementStatisticPeriod `qs:"-"`
	StartDate         *string                     `qs:"start_date,omitempty"`
	EndDate           *string                     `qs:"end_date,omitempty"`
	StartWeek         *string                     `qs:"start_week,omitempty"`
	EndWeek           *string                     `qs:"end_week,omitempty"`
	StartMonth        *string                     `qs:"start_month,omitempty"`
	EndMonth          *string                     `qs:"end_month,omitempty"`
	AnnouncementGroup string                      `qs:"announcement_group,omitempty"`
}

// URL returns URL string for viewing announcement statistics
func (a AnnouncementViewStatisticsRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/announcement_stats/", baseURL)
	switch a.Period {
	case STATISTIC_DAILY:
		base += "daily"
	case STATISTIC_WEEKLY:
		base += "weekly"
	case STATISTIC_MONTHLY:
		base += "monthly"
	}
	s, _ := encodeQS[AnnouncementViewStatisticsRequest](base, a)
	return s
}

// AnnouncementViewStatisticsResponse is the response type to view announcement statistics request
type AnnouncementViewStatisticsResponse struct {
	Announcements []Announcement `json:"announcements,omitempty"`
	Days          *int           `json:"days,omitempty"`
	Weeks         *int           `json:"weeks,omitempty"`
	Month         *int           `json:"months,omitempty"`
	Error         *Error         `json:"error,omitempty"`
}
