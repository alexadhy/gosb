package gosb

import (
	"context"

	"github.com/alexadhy/gosb/dto"
)

// CustomChannelListOperators retrieves a list of operators of a custom channel.
func (c *Client) CustomChannelListOperators(
	ctx context.Context,
	input *dto.CustomChannelListOperatorsRequest,
) (*dto.CustomChannelListOperatorsResponse, error) {
	result, err := getRequest[dto.CustomChannelListOperatorsResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, err
	}
	return result, nil
}

// CustomChannelRegisterOperators Registers one or more operators to a custom channel
func (c *Client) CustomChannelRegisterOperators(
	ctx context.Context,
	input *dto.CustomChannelRegisterOperatorRequest,
) (*dto.EmptyResponse, error) {
	result, err := postRequest[dto.EmptyResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}

	if result.Error != nil {
		return nil, err
	}
	return result, nil
}

// CustomChannelUnregisterOperators cancels the registration of operators in a custom channel but leaves them in the channel.
func (c *Client) CustomChannelUnregisterOperators(
	ctx context.Context,
	input *dto.CustomChannelUnregisterOperatorRequest,
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

// CustomChannelListSettings retrieves a list of the settings for all custom channel types.
func (c *Client) CustomChannelListSettings(
	ctx context.Context,
	input *dto.CustomChannelListSettingsRequest,
) (*dto.CustomChannelListSettingsResponse, error) {
	result, err := getRequest[dto.CustomChannelListSettingsResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, err
	}
	return result, nil
}

// CustomChannelGetSettings retrieves a list of the settings for all custom channel types.
func (c *Client) CustomChannelGetSettings(
	ctx context.Context,
	input *dto.CustomChannelGetSettingsRequest,
) (*dto.Settings, error) {
	result, err := getRequest[dto.Settings](ctx, c, input)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, err
	}
	return result, nil
}

// CustomChannelCreateSettings creates the settings for the channels with a custom channel type.
func (c *Client) CustomChannelCreateSettings(
	ctx context.Context,
	input *dto.CustomChannelCreateSettingRequest,
) (*dto.Settings, error) {
	result, err := postRequest[dto.Settings](ctx, c, input)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, err
	}
	return result, nil
}

// CustomChannelUpdateSettings Updates the settings for the channels with a custom channel type.
func (c *Client) CustomChannelUpdateSettings(
	ctx context.Context,
	input *dto.CustomChannelUpdateSettingRequest,
) (*dto.Settings, error) {
	result, err := putRequest[dto.Settings](ctx, c, input)
	if err != nil {
		return nil, err
	}
	if result.Error != nil {
		return nil, err
	}
	return result, nil
}

// CustomChannelDeleteSettings Deletes the settings for the channels with a custom channel type.
func (c *Client) CustomChannelDeleteSettings(
	ctx context.Context,
	input *dto.CustomChannelDeleteSettingRequest,
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
