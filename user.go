package gosb

import (
	"context"

	"github.com/alexadhy/gosb/dto"
)

// UserList takes *dto.UserListRequest as an input and will return *dto.UserListResponse, and error if any
// it is used to list all users on your application
func (c *Client) UserList(ctx context.Context, input *dto.UserListRequest) (*dto.UserListResponse, error) {
	result, err := getRequest[dto.UserListResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, err
	}
	return result, nil
}

// UserView retrieves information of a user.
func (c *Client) UserView(ctx context.Context, input *dto.UserGetRequest) (*dto.UserResponse, error) {
	result, err := getRequest[dto.UserResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, err
	}
	return result, nil
}

// UserCreate creates a new user in the application. You can choose to authenticate the user with just their user ID,
// or with either an access token or a session token.
func (c *Client) UserCreate(ctx context.Context, input *dto.UserCreateRequest) (*dto.UserResponse, error) {
	result, err := postRequest[dto.UserResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, err
	}
	return result, nil
}

// UserUpdate updates information of a user.
func (c *Client) UserUpdate(ctx context.Context, input *dto.UserUpdateRequest) (*dto.UserResponse, error) {
	result, err := putRequest[dto.UserResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, err
	}
	return result, nil
}

// DeleteUser updates information of a user.
func (c *Client) DeleteUser(ctx context.Context, input *dto.UserDeleteRequest) (*dto.EmptyResponse, error) {
	result, err := deleteRequest[dto.EmptyResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, err
	}
	return result, nil
}

// UserIssueSessionToken creates a session token for user
func (c *Client) UserIssueSessionToken(ctx context.Context, input *dto.UserCreateSessionTokenRequest) (*dto.UserCreateSessionTokenResponse, error) {
	result, err := postRequest[dto.UserCreateSessionTokenResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, err
	}
	return result, nil
}

// UserRevokeAllSessionTokens revokes all session token in the application
func (c *Client) UserRevokeAllSessionTokens(ctx context.Context, input *dto.UserRevokeAllSessionTokenRequest) (*dto.EmptyResponse, error) {
	result, err := deleteRequest[dto.EmptyResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, err
	}
	return result, nil
}

// UserListUserGroupChannels retrieves a list of the user's group channels.
func (c *Client) UserListUserGroupChannels(ctx context.Context, input *dto.UserListGroupChannelRequest) (*dto.UserListGroupChannelResponse, error) {
	result, err := getRequest[dto.UserListGroupChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, err
	}
	return result, nil
}

// UserLeaveGroupChannels makes the user leave all group channels.
func (c *Client) UserLeaveGroupChannels(ctx context.Context, input *dto.UserLeaveAllGroupChannelRequest) (*dto.EmptyResponse, error) {
	result, err := putRequest[dto.EmptyResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, err
	}
	return result, nil
}

// UserGetUnreadMessagesCount retrieves the total number of a user’s unread messages in group channels.
func (c *Client) UserGetUnreadMessagesCount(ctx context.Context, input *dto.MessageViewUnreadCountRequest) (*dto.MessageViewUnreadCountResponse, error) {
	result, err := getRequest[dto.MessageViewUnreadCountResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, err
	}
	return result, nil
}

// UserGetUnreadItemsCount retrieves a set of total number of a user's unread messages, unread mentioned messages,
// or received invitations in either Supergroup channels or non-Supergroup channels.
// TODO @alexadhy
// func (c *Client) UserGetUnreadItemsCount
