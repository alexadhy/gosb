package dto

import "encoding/json"

// Error is the json response type returned by SendBird API
type Error struct {
	Code    ErrorCode `json:"code,omitempty"`
	Message string    `json:"message,omitempty"`
}

func (e Error) Error() string {
	return e.Message
}

func NewErrFromJSON(b []byte) *Error {
	var e Error
	if err := json.Unmarshal(b, &e); err != nil {
		return &Error{
			Code:    ErrUnmarshalJSONResponse,
			Message: err.Error(),
		}
	}
	return &e
}

func NewErr(code ErrorCode, msg string) *Error {
	return &Error{Message: msg, Code: code}
}

// ErrorCode maps error code from
type ErrorCode int

const (
	ErrUnmarshalJSONResponse = ErrorCode(iota + 900000)
	ErrEmptyURLValue
	ErrParamShouldTypeString = ErrorCode(iota + 400100)
	ErrParamShouldTypeNumber
	ErrParamShouldTypeList
	ErrParamShouldTypeJSON
	ErrParamShouldTypeBool
	ErrMissingRequiredSymbol
	ErrNegativeNumberNotAllowed

	ErrUnauthorizedRequest = ErrorCode(400108)

	ErrParamValueLengthExceeded = ErrorCode(iota + 400110)
	ErrInvalidValue
	ErrIncompatibleValues
	ErrParamValueOutOfRange
	ErrInvalidURLOfResources

	ErrNotAllowedCharacter = ErrorCode(400151)

	ErrResourceNotFound = ErrorCode(iota + 400201)
	ErrResourceAlreadyExists
	ErrTooManyItemsInParameter

	ErrDeactivatedUserNotAccessible = ErrorCode(iota + 400300)
	ErrUserNotFound
	ErrInvalidAccessToken
	ErrInvalidSessionKeyValue
	ErrApplicationNotFound
	ErrUserIdLengthExceeded
	ErrPaidQuotaExceeded
	ErrDomainNotAllowed

	ErrInvalidApiToken = ErrorCode(iota + 400400)
	ErrMissingSomeParameters
	ErrInvalidJsonRequestBody
	ErrInvalidRequestURL

	ErrTooManyUserWebsocketConnections = ErrorCode(iota + 400500)
	ErrTooManyApplicationWebsocketConnections

	ErrBlockedUserSendNotAllowed = ErrorCode(iota + 40700)
	ErrBlockedUserInvitedNotAllowed
	ErrBlockedUserInviteNotAllowed

	ErrBannedUserEnterChannelNotAllowed = ErrorCode(iota + 400750)
	ErrBannedUserEnterCustomChannelNotAllowed

	ErrMisc400920      = ErrorCode(400920)
	ErrInvalidEndpoint = ErrorCode(400930)

	ErrApplicationNotAvailable = ErrorCode(403100) // 403

	ErrRateLimitExceeded                = ErrorCode(500910)        // 429
	ErrInternalPushTokenNotRegistered   = ErrorCode(iota + 500601) // 500
	ErrInternalPushTokenNotUnregistered                            // 500

	ErrInternal = ErrorCode(500901) // 500

	ErrNotAvailable = ErrorCode(500902)
)
