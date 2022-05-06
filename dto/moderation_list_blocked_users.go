package dto

import "fmt"

type ModerationListBlockedUsersRequest struct {
	UserID string `json:"user_id"`
	//optional
	Token            string `qs:"token,omitempty"`
	Limit            int    `qs:"limit,omitempty"`
	UserIds          string `qs:"user_ids,omitempty"`
	MetadataKey      string `qs:"metadatakey,omitempty"`
	MetadataValuesIn string `qs:"metadatavalues_in,omitempty"`
}

type ModerationListBlockedUsersResponse struct {
	Users []UserResponse `json:"users"`
	Next  string         `json:"next"`
}

func (m ModerationListBlockedUsersRequest) URL(baseURl string) string {
	base := fmt.Sprintf("%s/%s/%s/block", baseURl, userEndpoint, m.UserID)
	base, _ = encodeQS[ModerationListBlockedUsersRequest](base, m)
	return base
}
