package gosb

import (
	"context"

	"github.com/alexadhy/gosb/dto"
)

// GroupChannelListOperators retrieves a list of operators of an group channel.
func (c *Client) GroupChannelListOperators(
	ctx context.Context,
	input *dto.GroupChannelListOperatorsRequest,
) (*dto.GroupChannelListOperatorsResponse, error) {
	u, err := makeRequestURL(input, c)
	if err != nil {
		return nil, err
	}

	result, err := getRequest[dto.GroupChannelListOperatorsResponse](ctx, c.hc, u)
	if err != nil {
		return nil, err
	}

	return result, result.Error
}

// GroupChannelRegisterOperators Registers one or more operators to a group channel
func (c *Client) GroupChannelRegisterOperators(
	ctx context.Context,
	input *dto.GroupChannelRegisterOperatorRequest,
) (*dto.GroupChannelResponse, error) {
	u, err := makeRequestURL(input, c)
	if err != nil {
		return nil, err
	}

	result, err := postRequest[dto.GroupChannelResponse](ctx, c.hc, u, input)
	if err != nil {
		return nil, err
	}

	return result, result.Error
}

// GroupChannelUnregisterOperators cancels the registration of operators in a group channel but leaves them as members
func (c *Client) GroupChannelUnregisterOperators(
	ctx context.Context,
	input *dto.GroupChannelUnregisterOperatorRequest,
) (*dto.EmptyResponse, error) {
	u, err := makeRequestURL(input, c)
	if err != nil {
		return nil, err
	}

	result, err := deleteRequest[dto.EmptyResponse](ctx, c.hc, u)
	if err != nil {
		return nil, err
	}

	return result, result.Error
}

// GroupChannelList Retrieves a list of group channels or Supergroup channels in the application.
// If you want to get a list of a specific user's group channels, use the User section's
func (c *Client) GroupChannelList(ctx context.Context, input *dto.GroupChannelListRequest) (*dto.GroupChannelListResponse, error) {
	u, err := makeRequestURL(input, c)
	if err != nil {
		return nil, err
	}

	result, err := getRequest[dto.GroupChannelListResponse](ctx, c.hc, u)
	if err != nil {
		return nil, err
	}

	return result, result.Error
}

// GroupChannelGet retrieves information on a specific group channel or a Supergroup channel.
func (c *Client) GroupChannelGet(ctx context.Context, input *dto.GroupChannelGetRequest) (*dto.GroupChannelResponse, error) {
	u, err := makeRequestURL(input, c)
	if err != nil {
		return nil, err
	}

	result, err := getRequest[dto.GroupChannelResponse](ctx, c.hc, u)
	if err != nil {
		return nil, err
	}

	return result, result.Error
}

// GroupChannelCreate creates a group channel or a Supergroup channel
func (c *Client) GroupChannelCreate(ctx context.Context, input *dto.GroupChannelCreateRequest) (*dto.GroupChannelCreateResponse, error) {
	u, err := makeRequestURL(input, c)
	if err != nil {
		return nil, err
	}

	result, err := postRequest[dto.GroupChannelCreateResponse](ctx, c.hc, u, input)
	if err != nil {
		return nil, err
	}

	return result, result.Error
}

// GroupChannelUpdate updates information on a specific group channel or a Supergroup channel.
func (c *Client) GroupChannelUpdate(ctx context.Context, input *dto.GroupChannelUpdateRequest) (*dto.GroupChannelUpdateResponse, error) {
	u, err := makeRequestURL(input, c)
	if err != nil {
		return nil, err
	}

	result, err := putRequest[dto.GroupChannelUpdateResponse](ctx, c.hc, u, input)
	if err != nil {
		return nil, err
	}

	return result, result.Error
}

// GroupChannelDelete deletes specific group channel
func (c *Client) GroupChannelDelete(ctx context.Context, input *dto.GroupChannelDeleteRequest) (*dto.GroupChannelDeleteResponse, error) {
	u, err := makeRequestURL(input, c)
	if err != nil {
		return nil, err
	}

	result, err := deleteRequest[dto.GroupChannelDeleteResponse](ctx, c.hc, u)
	if err != nil {
		return nil, err
	}

	return result, result.Error
}

// GroupChannelJoin allows a user to join a public group channel. This is only available
// for public group channels where the is_public property is true.
func (c *Client) GroupChannelJoin(ctx context.Context, input *dto.GroupChannelJoinRequest) (*dto.GroupChannelJoinResponse, error) {
	u, err := makeRequestURL(input, c)
	if err != nil {
		return nil, err
	}

	result, err := putRequest[dto.GroupChannelJoinResponse](ctx, c.hc, u, input)
	if err != nil {
		return nil, err
	}

	return result, result.Error
}

// GroupChannelLeave makes one or more users leave a specific group channel.
func (c *Client) GroupChannelLeave(ctx context.Context, input *dto.GroupChannelLeaveRequest) (*dto.GroupChannelLeaveResponse, error) {
	u, err := makeRequestURL(input, c)
	if err != nil {
		return nil, err
	}

	result, err := putRequest[dto.GroupChannelLeaveResponse](ctx, c.hc, u, input)
	if err != nil {
		return nil, err
	}

	return result, result.Error
}

// GroupChannelHide hides or archives a specific group channel from a user’s group channel list.
func (c *Client) GroupChannelHide(ctx context.Context, input *dto.GroupChannelHideRequest) (*dto.GroupChannelHideResponse, error) {
	u, err := makeRequestURL(input, c)
	if err != nil {
		return nil, err
	}

	result, err := putRequest[dto.GroupChannelHideResponse](ctx, c.hc, u, input)
	if err != nil {
		return nil, err
	}

	return result, result.Error
}

// GroupChannelUnhide Unhides or unarchives a specific group channel and gets the channel appeared back in a user’s group channel list.
func (c *Client) GroupChannelUnhide(ctx context.Context, input *dto.GroupChannelUnhideRequest) (*dto.EmptyResponse, error) {
	u, err := makeRequestURL(input, c)
	if err != nil {
		return nil, err
	}

	result, err := deleteRequest[dto.EmptyResponse](ctx, c.hc, u)
	if err != nil {
		return nil, err
	}

	return result, result.Error
}

// GroupChannelResetChatHistory Resets a user's chat history in a specific group channel.
func (c *Client) GroupChannelResetChatHistory(ctx context.Context, input *dto.GroupChannelResetHistoryRequest) (*dto.GroupChannelResetHistoryResponse, error) {
	u, err := makeRequestURL(input, c)
	if err != nil {
		return nil, err
	}

	result, err := putRequest[dto.GroupChannelResetHistoryResponse](ctx, c.hc, u, input)
	if err != nil {
		return nil, err
	}

	return result, result.Error
}

// GroupChannelViewInvitationPref retrieves channel invitation preference for a user's group channels.
func (c *Client) GroupChannelViewInvitationPref(ctx context.Context, input *dto.ChannelGetInvitationPrefRequest) (*dto.ChannelGetInvitationPrefResponse, error) {
	u, err := makeRequestURL(input, c)
	if err != nil {
		return nil, err
	}

	result, err := getRequest[dto.ChannelGetInvitationPrefResponse](ctx, c.hc, u)
	if err != nil {
		return nil, err
	}

	return result, result.Error
}

// GroupChannelUpdateInvitationPref Updates channel invitation preference for a user's group channels.
func (c *Client) GroupChannelUpdateInvitationPref(
	ctx context.Context,
	input *dto.ChannelUpdateInvitationPrefRequest,
) (*dto.ChannelUpdateInvitationPrefResponse, error) {
	u, err := makeRequestURL(input, c)
	if err != nil {
		return nil, err
	}

	result, err := putRequest[dto.ChannelUpdateInvitationPrefResponse](ctx, c.hc, u, input)
	if err != nil {
		return nil, err
	}

	return result, result.Error
}

// GroupChannelListMembers retrieves the members of a specific group channel.
func (c *Client) GroupChannelListMembers(
	ctx context.Context,
	input *dto.GroupChannelListMembersRequest,
) (*dto.GroupChannelListMembersResponse, error) {
	u, err := makeRequestURL(input, c)
	if err != nil {
		return nil, err
	}

	result, err := putRequest[dto.GroupChannelListMembersResponse](ctx, c.hc, u, input)
	if err != nil {
		return nil, err
	}

	return result, result.Error
}

// GroupChannelCheckMembership checks whether the user is a member of a specific group channel.
func (c *Client) GroupChannelCheckMembership(
	ctx context.Context,
	input *dto.GroupChannelCheckMemberRequest,
) (*dto.GroupChannelCheckMemberResponse, error) {
	u, err := makeRequestURL(input, c)
	if err != nil {
		return nil, err
	}

	result, err := putRequest[dto.GroupChannelCheckMemberResponse](ctx, c.hc, u, input)
	if err != nil {
		return nil, err
	}

	return result, result.Error
}

// GroupChannelInviteMember invites one or more users as members to a group channel.
func (c *Client) GroupChannelInviteMember(
	ctx context.Context,
	input *dto.GroupChannelInviteMembersRequest,
) (*dto.GroupChannelInviteMembersResponse, error) {
	u, err := makeRequestURL(input, c)
	if err != nil {
		return nil, err
	}

	result, err := postRequest[dto.GroupChannelInviteMembersResponse](ctx, c.hc, u, input)
	if err != nil {
		return nil, err
	}

	return result, result.Error
}

// GroupChannelAcceptInvite Accepts an invitation from a group channel for a user to join.
func (c *Client) GroupChannelAcceptInvite(
	ctx context.Context,
	input *dto.GroupChannelAcceptInviteRequest,
) (*dto.GroupChannelAcceptInviteResponse, error) {
	u, err := makeRequestURL(input, c)
	if err != nil {
		return nil, err
	}

	result, err := putRequest[dto.GroupChannelAcceptInviteResponse](ctx, c.hc, u, input)
	if err != nil {
		return nil, err
	}

	return result, result.Error
}

// GroupChannelDeclineInvite Declines an invitation from a group channel for a user to join.
func (c *Client) GroupChannelDeclineInvite(
	ctx context.Context,
	input *dto.GroupChannelDeclineInviteRequest,
) (*dto.GroupChannelDeclineInviteResponse, error) {
	u, err := makeRequestURL(input, c)
	if err != nil {
		return nil, err
	}

	result, err := putRequest[dto.GroupChannelDeclineInviteResponse](ctx, c.hc, u, input)
	if err != nil {
		return nil, err
	}

	return result, result.Error
}
