package events

import "github.com/alexadhy/gosb/dto"

// OperatorsCommon is the common object on every event(s) related to Operators
type OperatorsCommon struct {
	Category    string             `json:"category"`
	RequestedAt int64              `json:"requested_at"`
	Requester   *dto.UserResponse  `json:"requester"`
	ChannelType string             `json:"channel_type"`
	Channel     *dto.Channel       `json:"channel"`
	Operators   []dto.UserResponse `json:"operators"`
	AppID       string             `json:"app_id"`
}

// OperatorsRegisterByOperator  is invoked when one or more operators are registered by another operator's client app
type OperatorsRegisterByOperator struct {
	OperatorsCommon
}

// OperatorsUnregisterByOperator  is invoked when the registration of one or more operators
// is canceled by another operator's client app
type OperatorsUnregisterByOperator struct {
	OperatorsCommon
}
