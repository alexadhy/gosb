package events

import "github.com/alexadhy/gosb/dto"

// ReportCommon common report object
type ReportCommon struct {
	Category       string            `json:"category"`
	CreatedAt      int64             `json:"created_at"`
	ReportType     string            `json:"report_type"`
	ReportCategory string            `json:"report_category"`
	ReportingUser  *dto.UserResponse `json:"reporting_user"`
	Channel        *dto.Channel      `json:"channel"`
	AppID          string            `json:"app_id"`
}

// MessageReport is invoked when a message is reported by a user.
type MessageReport struct {
	ReportCommon
	OffendingUser   *dto.UserResponse `json:"offending_user"`
	ReportedMessage *dto.Message      `json:"reported_message"`
}

func (m MessageReport) EventCategory() string {
	return m.Category
}

// UserReport webhook event is invoked when a user is reported by another user.
// The following shows a webhook payload of a user:report event
type UserReport struct {
	ReportCommon
	OffendingUser *dto.UserResponse `json:"offending_user"`
}

func (u UserReport) EventCategory() string {
	return u.Category
}

// OpenChannelReport webhook event is invoked when an open channel is reported by a user.
// The following shows a webhook payload of a open_channel:report event
type OpenChannelReport struct {
	ReportCommon
}

func (o OpenChannelReport) EventCategory() string {
	return o.Category
}

// GroupChannelReport webhook event is invoked when a group channel is reported by a user.
// The following shows a webhook payload of a group_channel:report event
type GroupChannelReport struct {
	ReportCommon
}

func (g GroupChannelReport) EventCategory() string {
	return g.Category
}
