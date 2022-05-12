package gosb

import (
	"context"
	"github.com/alexadhy/gosb/dto"
)

// MessagesList retrieves a list of past messages of a specific channel.
func (c *Client) MessagesList(
	ctx context.Context,
	input *dto.MessageListRequest,
) (*dto.MessageListResponse, error) {
	res, err := getRequest[dto.MessageListResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// MessageView retrieves information on specific message
func (c *Client) MessageView(
	ctx context.Context,
	input *dto.MessageViewRequest,
) (*dto.MessageViewResponse, error) {
	res, err := getRequest[dto.MessageViewResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// MessageSend sends a message to a specific channel
func (c *Client) MessageSend(
	ctx context.Context,
	input *dto.MessageSendRequest,
) (*dto.MessageSentResponse, error) {
	res, err := postRequest[dto.MessageSentResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// MessageUpdate updates specific info on a message
func (c *Client) MessageUpdate(
	ctx context.Context,
	input *dto.MessageUpdateRequest,
) (*dto.MessageUpdateResponse, error) {
	res, err := putRequest[dto.MessageUpdateResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// MessageDelete deletes a specific message from a channel
func (c *Client) MessageDelete(
	ctx context.Context,
	input *dto.MessageDeleteRequest,
) (*dto.MessageDeleteResponse, error) {
	res, err := deleteRequest[dto.MessageDeleteResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// MessageAddMetadata Adds key-values items to store additional information for a message.
func (c *Client) MessageAddMetadata(
	ctx context.Context,
	input *dto.MessageAddMetaRequest,
) (*dto.MessageAddMetaResponse, error) {
	res, err := postRequest[dto.MessageAddMetaResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// MessageUpdateMetadata updates key-values items to store Updateitional information for a message.
func (c *Client) MessageUpdateMetadata(
	ctx context.Context,
	input *dto.MessageUpdateMetaRequest,
) (*dto.MessageUpdateMetaResponse, error) {
	res, err := putRequest[dto.MessageUpdateMetaResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// MessageDeleteMetadata deletes key-values items to store Deleteitional information for a message.
func (c *Client) MessageDeleteMetadata(
	ctx context.Context,
	input *dto.MessageDeleteMetaRequest,
) (*dto.MessageDeleteMetaResponse, error) {
	res, err := deleteRequest[dto.MessageDeleteMetaResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// MessageViewTotalNumberOfMessagesOnChannel Retrieves the total number of messages in a specific channel.
func (c *Client) MessageViewTotalNumberOfMessagesOnChannel(
	ctx context.Context,
	input *dto.MessageTotalCountInChannelRequest,
) (*dto.MessageTotalCountInChannelResponse, error) {
	res, err := getRequest[dto.MessageTotalCountInChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// MessageThreadListReplies retrieves replies of a specific parent message
func (c *Client) MessageThreadListReplies(
	ctx context.Context,
	input *dto.MessageListThreadedRepliesRequest,
) (*dto.MessageListThreadedRepliesResponse, error) {
	res, err := getRequest[dto.MessageListThreadedRepliesResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// MessageThreadViewReply retrieves a specific reply in a thread.
func (c *Client) MessageThreadViewReply(
	ctx context.Context,
	input *dto.MessageViewReplyRequest,
) (*dto.MessageViewReplyResponse, error) {
	res, err := getRequest[dto.MessageViewReplyResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// MessageThreadViewInfo retrieves thread information on a specific message.
func (c *Client) MessageThreadViewInfo(
	ctx context.Context,
	input *dto.MessageViewThreadInfoRequest,
) (*dto.MessageViewThreadInfoResponse, error) {
	res, err := getRequest[dto.MessageViewThreadInfoResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// MessageReply retrieves thread information on a specific message.
func (c *Client) MessageReply(
	ctx context.Context,
	input *dto.MessageReplyRequest,
) (*dto.MessageReplyResponse, error) {
	res, err := postRequest[dto.MessageReplyResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// MessageDeleteReply deletes a specific reply from a channel.
func (c *Client) MessageDeleteReply(
	ctx context.Context,
	input *dto.MessageDeleteRequest,
) (*dto.EmptyResponse, error) {
	res, err := deleteRequest[dto.EmptyResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// MessageSearch searches messages containing specific term
func (c *Client) MessageSearch(
	ctx context.Context,
	input *dto.MessageSearchRequest,
) (*dto.MessageSearchResponse, error) {
	res, err := getRequest[dto.MessageSearchResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}
