package dto

import "fmt"

// ChannelGetInvitationPrefRequest is the request param to get invitation preference of a channel
type ChannelGetInvitationPrefRequest struct {
	UserID string `json:"-"`
}

// URL returns URL string to get invitation preferences endpoint
// GET https://api-{application_id}.sendbird.com/v3/users/{user_id}/channel_invitation_preference
func (o ChannelGetInvitationPrefRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/channel_invitation_preferences", baseURL, userEndpoint, o.UserID)
}

// ChannelGetInvitationPrefResponse is the response payload to get invitation preference endpoint
type ChannelGetInvitationPrefResponse struct {
	AutoAccept *bool  `json:"auto_accept,omitempty"`
	Error      *Error `json:"error,omitempty"`
}

// ChannelUpdateInvitationPrefRequest is the request payload to update invitation preference endpoint
type ChannelUpdateInvitationPrefRequest struct {
	UserID     string `json:"-"`
	AutoAccept bool   `json:"auto_accept"`
}

// URL returns URL string to get invitation preferences endpoint
// PUT https://api-{application_id}.sendbird.com/v3/users/{user_id}/channel_invitation_preference
func (o ChannelUpdateInvitationPrefRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/channel_invitation_preferences", baseURL, userEndpoint, o.UserID)
}

type ChannelUpdateInvitationPrefResponse ChannelGetInvitationPrefResponse
