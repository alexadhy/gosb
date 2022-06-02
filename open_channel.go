package gosb

import (
	"context"

	"github.com/alexadhy/gosb/dto"
)

// OpenChannelListOperators retrieves a list of operators of an open channel.
func (c *Client) OpenChannelListOperators(
	ctx context.Context,
	input *dto.OpenChannelListOperatorsRequest,
) (*dto.OpenChannelListOperatorsResponse, error) {

	result, err := getRequest[dto.OpenChannelListOperatorsResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, err
	}
	return result, nil
}

// OpenChannelRegisterOperators Registers one or more operators to an open channel
func (c *Client) OpenChannelRegisterOperators(
	ctx context.Context,
	input *dto.OpenChannelRegisterOperatorRequest,
) (*dto.OpenChannelResponse, error) {

	result, err := postRequest[dto.OpenChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, err
	}
	return result, nil
}

// OpenChannelUnregisterOperators cancels the registration of operators in an open channel but leaves them as participants
func (c *Client) OpenChannelUnregisterOperators(
	ctx context.Context,
	input *dto.OpenChannelUnregisterOperatorRequest,
) (*dto.EmptyResponse, error) {

	result, err := deleteRequest[dto.EmptyResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, err
	}
	return result, nil
}

// OpenChannelList list all open channels
func (c *Client) OpenChannelList(ctx context.Context, input *dto.OpenChannelListRequest) (*dto.OpenChannelListResponse, error) {

	result, err := getRequest[dto.OpenChannelListResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, err
	}
	return result, nil
}

// OpenChannelGet retrieves information on an open channel.
func (c *Client) OpenChannelGet(ctx context.Context, input *dto.OpenChannelGetRequest) (*dto.OpenChannelResponse, error) {

	result, err := getRequest[dto.OpenChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, err
	}
	return result, nil
}

// OpenChannelCreate creates an open channel.
func (c *Client) OpenChannelCreate(ctx context.Context, input *dto.OpenChannelCreateRequest) (*dto.OpenChannelResponse, error) {

	result, err := postRequest[dto.OpenChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, err
	}
	return result, nil
}

// OpenChannelUpdate Updates information on a specific open channel.
func (c *Client) OpenChannelUpdate(ctx context.Context, input *dto.OpenChannelUpdateRequest) (*dto.OpenChannelResponse, error) {

	result, err := putRequest[dto.OpenChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, err
	}
	return result, nil
}

// OpenChannelDelete deletes a specific open channel.
func (c *Client) OpenChannelDelete(ctx context.Context, input *dto.OpenChannelDeleteRequest) (*dto.EmptyResponse, error) {

	result, err := deleteRequest[dto.EmptyResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
 return nil, err
 }
 return result, nil
}
