package dto

import "fmt"

type ModerationBlockUserRequest struct {
	UserID string `json:"user_id"`
	//body
	TargetId string         `json:"target_id,omitempty"`
	UserIds  []string       `json:"user_ids,omitempty"`
	Users    []UserResponse `json:"users,omitempty"`
}

type ModerationBlockUserResponse struct {
	UserResponse
}

func (m ModerationBlockUserRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/block", baseURL, userEndpoint, m.UserID)
}

type ModerationUnblockUserRequest struct {
	UserID   string `json:"user_id"`
	TargetID string `json:"target_id"`
}

type ModerationUnblockUserResponse struct{}

func (m ModerationUnblockUserRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/block/%s", baseURL, userEndpoint, m.UserID, m.TargetID)
}
