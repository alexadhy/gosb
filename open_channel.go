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
	u, err := makeRequestURL(input, c)
	if err != nil {
		return nil, err
	}

	result, err := getRequest[dto.OpenChannelListOperatorsResponse](ctx, c.hc, u)
	if err != nil {
		return nil, err
	}

	return result, result.Error
}

// OpenChannelRegisterOperators Registers one or more operators to an open channel
func (c *Client) OpenChannelRegisterOperators(
	ctx context.Context,
	input *dto.OpenChannelRegisterOperatorRequest,
) (*dto.OpenChannelResponse, error) {
	u, err := makeRequestURL(input, c)
	if err != nil {
		return nil, err
	}

	result, err := postRequest[dto.OpenChannelResponse](ctx, c.hc, u, input)
	if err != nil {
		return nil, err
	}

	return result, result.Error
}

// OpenChannelUnregisterOperators cancels the registration of operators in an open channel but leaves them as participants
func (c *Client) OpenChannelUnregisterOperators(
	ctx context.Context,
	input *dto.OpenChannelUnregisterOperatorRequest,
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

// OpenChannelList list all open channels
func (c *Client) OpenChannelList(ctx context.Context, input *dto.OpenChannelListRequest) (*dto.OpenChannelListResponse, error) {
	u, err := makeRequestURL(input, c)
	if err != nil {
		return nil, err
	}

	result, err := getRequest[dto.OpenChannelListResponse](ctx, c.hc, u)
	if err != nil {
		return nil, err
	}

	return result, result.Error
}

// OpenChannelGet retrieves information on an open channel.
func (c *Client) OpenChannelGet(ctx context.Context, input *dto.OpenChannelGetRequest) (*dto.OpenChannelResponse, error) {
	u, err := makeRequestURL(input, c)
	if err != nil {
		return nil, err
	}

	result, err := getRequest[dto.OpenChannelResponse](ctx, c.hc, u)
	if err != nil {
		return nil, err
	}

	return result, result.Error
}

// OpenChannelCreate creates an open channel.
func (c *Client) OpenChannelCreate(ctx context.Context, input *dto.OpenChannelCreateRequest) (*dto.OpenChannelResponse, error) {
	u, err := makeRequestURL(input, c)
	if err != nil {
		return nil, err
	}

	result, err := postRequest[dto.OpenChannelResponse](ctx, c.hc, u, input)
	if err != nil {
		return nil, err
	}

	return result, result.Error
}

// OpenChannelUpdate Updates information on a specific open channel.
func (c *Client) OpenChannelUpdate(ctx context.Context, input *dto.OpenChannelUpdateRequest) (*dto.OpenChannelResponse, error) {
	u, err := makeRequestURL(input, c)
	if err != nil {
		return nil, err
	}

	result, err := putRequest[dto.OpenChannelResponse](ctx, c.hc, u, input)
	if err != nil {
		return nil, err
	}

	return result, result.Error
}

// OpenChannelDelete deletes a specific open channel.
func (c *Client) OpenChannelDelete(ctx context.Context, input *dto.OpenChannelDeleteRequest) (*dto.EmptyResponse, error) {
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
