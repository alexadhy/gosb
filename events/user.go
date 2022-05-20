package events

import "github.com/alexadhy/gosb/dto"

type UserCommon struct {
	Category string `json:"category"`
	AppID    string `json:"app_id"`
}

// UserBlock webhook event is invoked when a user blocks another user.
type UserBlock struct {
	UserCommon
	Blocker *dto.UserResponse `json:"blocker"`
	Blockee *dto.UserResponse `json:"blockee"`
}

func (u UserBlock) EventCategory() string {
	return u.Category
}

// UserUnblock webhook event is invoked when a user unblocks another user.
type UserUnblock struct {
	UserCommon
	Unblocker *dto.UserResponse `json:"unblocker"`
	Unblockee *dto.UserResponse `json:"unblockee"`
}

func (u UserUnblock) EventCategory() string {
	return u.Category
}
