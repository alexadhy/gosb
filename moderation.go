package gosb

import (
	"context"

	"github.com/alexadhy/gosb/dto"
)

// ModerationListBlockedUsers will retrieve a list of other users that a user has blocked.
func (c *Client) ModerationListBlockedUsers(
	ctx context.Context,
	input *dto.ModerationListBlockedUsersRequest,
) (*dto.ModerationListBlockedUsersResponse, error) {
	res, err := getRequest[dto.ModerationListBlockedUsersResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationBlockUser allows a user to block another user
func (c *Client) ModerationBlockUser(
	ctx context.Context,
	input *dto.ModerationBlockUserRequest,
) (*dto.ModerationBlockUserResponse, error) {
	res, err := postRequest[dto.ModerationBlockUserResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationUnblockUser unblocks a user
func (c *Client) ModerationUnblockUser(
	ctx context.Context,
	input *dto.ModerationUnblockUserRequest,
) (*dto.ModerationUnblockUserResponse, error) {
	res, err := deleteRequest[dto.ModerationUnblockUserResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationListUserMutedChannels retrieves a list of open and group channels where the user is muted.
func (c *Client) ModerationListUserMutedChannels(
	ctx context.Context,
	input *dto.ModerationListUsersMutedChannelsRequest,
) (*dto.ModerationListUsersMutedChannelsResponse, error) {
	res, err := getRequest[dto.ModerationListUsersMutedChannelsResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationListMutedUserInOpenChannel retrieves a list of open and group channels where the user is muted.
func (c *Client) ModerationListMutedUserInOpenChannel(
	ctx context.Context,
	input *dto.ModerationListUsersMutedChannelsRequest,
) (*dto.ModerationListUsersMutedChannelsResponse, error) {
	res, err := getRequest[dto.ModerationListUsersMutedChannelsResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationListMutedMemberInGroupChannel Retrieves a list of muted users in a group channel.
func (c *Client) ModerationListMutedMemberInGroupChannel(
	ctx context.Context,
	input *dto.ModerationListMutedUsersInGroupChannelRequest,
) (*dto.ModerationListMutedUsersInGroupChannelResponse, error) {
	res, err := getRequest[dto.ModerationListMutedUsersInGroupChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationListMutedUsersInCustomChannel retrieves a list of muted users in a custom channel.
func (c *Client) ModerationListMutedUsersInCustomChannel(
	ctx context.Context,
	input *dto.ModerationListMutedUsersInCustomChannelRequest,
) (*dto.ModerationListMutedUsersInCustomChannelResponse, error) {
	res, err := getRequest[dto.ModerationListMutedUsersInCustomChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationViewMutedParticipantInOpenChannel retrieves details of a mute imposed on a participant in an open channel.
func (c *Client) ModerationViewMutedParticipantInOpenChannel(
	ctx context.Context,
	input *dto.ModerationViewMutedUserInOpenChannelRequest,
) (*dto.ModerationViewMutedUserInOpenChannelResponse, error) {
	res, err := getRequest[dto.ModerationViewMutedUserInOpenChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationViewMutedMemberInGroupChannel Retrieves details of a mute imposed on a member in a group channel
func (c *Client) ModerationViewMutedMemberInGroupChannel(
	ctx context.Context,
	input *dto.ModerationViewMutedUserInGroupChannelRequest,
) (*dto.ModerationViewMutedUserInGroupChannelResponse, error) {
	res, err := getRequest[dto.ModerationViewMutedUserInGroupChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationMuteParticipantInOpenChannel mutes a participant in an open channel.
func (c *Client) ModerationMuteParticipantInOpenChannel(
	ctx context.Context,
	input *dto.ModerationMuteUserInOpenChannelRequest,
) (*dto.ModerationMuteUserInOpenChannelResponse, error) {
	res, err := postRequest[dto.ModerationMuteUserInOpenChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationMuteMemberInGroupChannel mutes a member in a group channel
func (c *Client) ModerationMuteMemberInGroupChannel(
	ctx context.Context,
	input *dto.ModerationMuteUserInGroupChannelRequest,
) (*dto.EmptyResponse, error) {
	res, err := postRequest[dto.EmptyResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationMuteUserInCustomChannel mutes a user in a custom channel(s) with custom type at once
func (c *Client) ModerationMuteUserInCustomChannel(
	ctx context.Context,
	input *dto.ModerationMuteUserInCustomChannelRequest,
) (*dto.ModerationMuteUserInCustomChannelResponse, error) {
	res, err := postRequest[dto.ModerationMuteUserInCustomChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationUnmuteParticipantInOpenChannel unmutes a participant in an open channel.
func (c *Client) ModerationUnmuteParticipantInOpenChannel(
	ctx context.Context,
	input *dto.ModerationUnmuteUserInOpenChannelRequest,
) (*dto.ModerationUnmuteUserInOpenChannelResponse, error) {
	res, err := deleteRequest[dto.ModerationUnmuteUserInOpenChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationUnmuteMemberInGroupChannel unmutes a member in a group channel
func (c *Client) ModerationUnmuteMemberInGroupChannel(
	ctx context.Context,
	input *dto.ModerationUnmuteUserInGroupChannelRequest,
) (*dto.ModerationUnmuteUserInGroupChannelResponse, error) {
	res, err := deleteRequest[dto.ModerationUnmuteUserInGroupChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationUnmuteUsersInCustomChannel unmutes users in channels with a custom type at once
func (c *Client) ModerationUnmuteUsersInCustomChannel(
	ctx context.Context,
	input *dto.ModerationUnmuteUsersInCustomChannelRequest,
) (*dto.ModerationUnmuteUsersInCustomChannelResponse, error) {
	res, err := deleteRequest[dto.ModerationUnmuteUsersInCustomChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationListBannedParticipantsInOpenChannel list all banned users / participants in an open channel
func (c *Client) ModerationListBannedParticipantsInOpenChannel(
	ctx context.Context,
	input *dto.ModerationListBannedUsersInOpenChannelRequest,
) (*dto.ModerationListBannedUsersInOpenChannelResponse, error) {
	res, err := getRequest[dto.ModerationListBannedUsersInOpenChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationListBannedMembersInGroupChannel list all banned users / Members in an Group channel
func (c *Client) ModerationListBannedMembersInGroupChannel(
	ctx context.Context,
	input *dto.ModerationListBannedUsersInGroupChannelRequest,
) (*dto.ModerationListBannedUsersInGroupChannelResponse, error) {
	res, err := getRequest[dto.ModerationListBannedUsersInGroupChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationListBannedUsersInCustomChannel list all banned users / Users in an Custom channel
func (c *Client) ModerationListBannedUsersInCustomChannel(
	ctx context.Context,
	input *dto.ModerationListBannedUsersInCustomChannelRequest,
) (*dto.ModerationListBannedUsersInCustomChannelResponse, error) {
	res, err := getRequest[dto.ModerationListBannedUsersInCustomChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationViewBannedParticipantInOpenChannel retrieves details of a mute imposed on a participant in an open channel.
func (c *Client) ModerationViewBannedParticipantInOpenChannel(
	ctx context.Context,
	input *dto.ModerationViewBannedUserInOpenChannelRequest,
) (*dto.ModerationViewBannedUserInOpenChannelResponse, error) {
	res, err := getRequest[dto.ModerationViewBannedUserInOpenChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationViewBannedMemberInGroupChannel Retrieves details of a mute imposed on a member in a group channel
func (c *Client) ModerationViewBannedMemberInGroupChannel(
	ctx context.Context,
	input *dto.ModerationViewBannedUserInGroupChannelRequest,
) (*dto.ModerationViewBannedUserInGroupChannelResponse, error) {
	res, err := getRequest[dto.ModerationViewBannedUserInGroupChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationUserListBannedChannels lists open and group channels where a user is banned
func (c *Client) ModerationUserListBannedChannels(
	ctx context.Context,
	input *dto.ModerationListBannedUsersRequest,
) (*dto.ModerationListBannedUsersResponse, error) {
	res, err := getRequest[dto.ModerationListBannedUsersResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationBanParticipantFromOpenChannel Bans a participant from an open channel.
func (c *Client) ModerationBanParticipantFromOpenChannel(
	ctx context.Context,
	input *dto.ModerationBanUserFromOpenChannelRequest,
) (*dto.ModerationBanUserFromOpenChannelResponse, error) {
	res, err := postRequest[dto.ModerationBanUserFromOpenChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationBanMemberFromGroupChannel ban member from a group channel
func (c *Client) ModerationBanMemberFromGroupChannel(
	ctx context.Context,
	input *dto.ModerationBanUserFromGroupChannelRequest,
) (*dto.ModerationBanUserFromGroupChannelResponse, error) {
	res, err := postRequest[dto.ModerationBanUserFromGroupChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationBanUserFromCustomChannels bans a user from channels with specified custom channel types.
func (c *Client) ModerationBanUserFromCustomChannels(
	ctx context.Context,
	input *dto.ModerationBanUsersFromCustomChannelRequest,
) (*dto.ModerationBanUserFromCustomChannelResponse, error) {
	res, err := postRequest[dto.ModerationBanUserFromCustomChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationUnbanParticipantFromOpenChannel Unbans a participant From an open channel.
func (c *Client) ModerationUnbanParticipantFromOpenChannel(
	ctx context.Context,
	input *dto.ModerationUnbanUserFromOpenChannelRequest,
) (*dto.ModerationUnbanUserFromOpenChannelResponse, error) {
	res, err := deleteRequest[dto.ModerationUnbanUserFromOpenChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationUnbanMemberFromGroupChannel Unbans a member From a group channel
func (c *Client) ModerationUnbanMemberFromGroupChannel(
	ctx context.Context,
	input *dto.ModerationUnbanUserFromGroupChannelRequest,
) (*dto.ModerationUnbanUserFromGroupChannelResponse, error) {
	res, err := deleteRequest[dto.ModerationUnbanUserFromGroupChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationUnbanMemberFromCustomChannel Unbans a member From a Custom channel
func (c *Client) ModerationUnbanMemberFromCustomChannel(
	ctx context.Context,
	input *dto.ModerationUnbanUserFromCustomChannelRequest,
) (*dto.ModerationUnbanUserFromCustomChannelResponse, error) {
	res, err := deleteRequest[dto.ModerationUnbanUserFromCustomChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationUpdateStatusBannedParticipantInOpenChannel Updates information about a banned participant of an open channel.
// You can adjust the duration of ban or update the description about the ban with this API.
func (c *Client) ModerationUpdateStatusBannedParticipantInOpenChannel(
	ctx context.Context,
	input *dto.ModerationUpdateBannedUserStatusInOpenChannelRequest,
) (*dto.ModerationUpdateBannedUserStatusInOpenChannelResponse, error) {
	res, err := putRequest[dto.ModerationUpdateBannedUserStatusInOpenChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationUpdateStatusBannedMemberInGroupChannel Updates information about a banned Member of an Group channel.
// You can adjust the duration of ban or update the description about the ban with this API.
func (c *Client) ModerationUpdateStatusBannedMemberInGroupChannel(
	ctx context.Context,
	input *dto.ModerationUpdateBannedUserStatusInGroupChannelRequest,
) (*dto.ModerationUpdateBannedUserStatusInGroupChannelResponse, error) {
	res, err := putRequest[dto.ModerationUpdateBannedUserStatusInGroupChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationFreezeOpenChannel freezes or unfreezes an open channel.
func (c *Client) ModerationFreezeOpenChannel(
	ctx context.Context,
	input *dto.ModerationFreezeOpenChannelRequest,
) (*dto.ModerationFreezeOpenChannelResponse, error) {
	res, err := putRequest[dto.ModerationFreezeOpenChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// ModerationFreezeGroupChannel freezes or unfreezes an Group channel.
func (c *Client) ModerationFreezeGroupChannel(
	ctx context.Context,
	input *dto.ModerationFreezeGroupChannelRequest,
) (*dto.ModerationFreezeGroupChannelResponse, error) {
	res, err := putRequest[dto.ModerationFreezeGroupChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}
