package events

import "github.com/alexadhy/gosb/dto"

// AnnouncementCreateChannels is invoked when channels are created for sending an announcement
// to target users who do not already have a channel with the announcement sender
type AnnouncementCreateChannels struct {
	Category          string       `json:"category"`
	AnnouncementID    string       `json:"announcement_id"`
	AnnouncementGroup string       `json:"announcement_group"`
	CreatedAt         int64        `json:"created_at"`
	Channel           *dto.Channel `json:"channel"`
	AppID             string       `json:"app_id"`
}

func (a AnnouncementCreateChannels) EventCategory() string {
	return a.Category
}

// AnnouncementSendMessages is invoked when an announcement message is sent to target channels and users
type AnnouncementSendMessages struct {
	Category          string            `json:"category"`
	AnnouncementID    string            `json:"announcement_id"`
	AnnouncementGroup string            `json:"announcement_group"`
	CreatedAt         int64             `json:"created_at"`
	Sender            *dto.UserResponse `json:"sender"`
	Channel           *dto.Channel      `json:"channel"`
	Message           *dto.Message      `json:"message"`
	AppID             string            `json:"app_id"`
}

func (a AnnouncementSendMessages) EventCategory() string {
	return a.Category
}
