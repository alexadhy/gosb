package dto

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Channel is the common Channel object
type Channel struct {
	*OpenChannelResponse
	*GroupChannelResponse
}

// UnmarshalJSON unmarshal Channel type to its member type (OpenChannelResponse / GroupChannelResponse)
func (c Channel) UnmarshalJSON(data []byte) error {
	msi := map[string]any{}
	if err := json.Unmarshal(data, &msi); err != nil {
		return err
	}
	if _, ok := msi["is_public"]; ok {
		g := GroupChannelResponse{}
		return json.Unmarshal(data, &g)
	}
	oc := OpenChannelResponse{}
	return json.Unmarshal(data, &oc)
}

// MarshalJSON marshal Message type to JSON byte
func (c Channel) MarshalJSON() ([]byte, error) {
	if c.OpenChannelResponse != nil {
		return json.Marshal(c.OpenChannelResponse)
	}
	if c.GroupChannelResponse != nil {
		return json.Marshal(c.GroupChannelResponse)
	}
	return nil, errors.New("unknown message type")
}

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

// ChannelUpdateInvitationPrefResponse is an alias for ChannelGetInvitationPrefResponse
type ChannelUpdateInvitationPrefResponse ChannelGetInvitationPrefResponse
