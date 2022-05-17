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

// MessageMigrate migrates messages from another system into Sendbird's system
func (c *Client) MessageMigrate(
	ctx context.Context,
	input *dto.MessageMigrateRequest,
) (*dto.MessageMigrateResponse, error) {
	res, err := postRequest[dto.MessageMigrateResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// MessageMarkAllDeliveredGroupChannel marks all messages in a group channel as delivered for a specific user.
func (c *Client) MessageMarkAllDeliveredGroupChannel(
	ctx context.Context,
	input *dto.MessageMarkAllAsDeliveredRequest,
) (*dto.MessageMarkAllAsDeliveredResponse, error) {
	res, err := putRequest[dto.MessageMarkAllAsDeliveredResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// MessageMarkAllAsReadGroupChannel marks all messages in a group channel as read for a specific user.
func (c *Client) MessageMarkAllAsReadGroupChannel(
	ctx context.Context,
	input *dto.MessageMarkAllAsReadRequest,
) (*dto.MessageMarkAllAsReadResponse, error) {
	res, err := putRequest[dto.MessageMarkAllAsReadResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// MessageMarkAllAsReadAllJoinedGroupChannel marks all of a user's unread messages as read in group channels.
func (c *Client) MessageMarkAllAsReadAllJoinedGroupChannel(
	ctx context.Context,
	input *dto.MessageMarkAllAsReadAllJoinedGroupChannelRequest,
) (*dto.MessageMarkAllAsReadAllJoinedGroupChannelResponse, error) {
	res, err := putRequest[dto.MessageMarkAllAsReadAllJoinedGroupChannelResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// MessageViewUnreadCountMessagesPerMember retrieves the total number of each
// member's unread messages in a specific group channel.
func (c *Client) MessageViewUnreadCountMessagesPerMember(
	ctx context.Context,
	input *dto.MessageViewUnreadCountRequest,
) (*dto.MessageViewUnreadCountResponse, error) {
	res, err := getRequest[dto.MessageViewUnreadCountResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// MessageEnableReactions turns on and off message reactions in a Sendbird application.
func (c *Client) MessageEnableReactions(
	ctx context.Context,
	input *dto.EmojiEnableRequest,
) (*dto.EmojiEnableResponse, error) {
	res, err := putRequest[dto.EmojiEnableResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// MessageListReaction retrieves a list of reactions made to a specific message.
func (c *Client) MessageListReaction(
	ctx context.Context,
	input *dto.EmojiListReactionsRequest,
) (*dto.EmojiListReactionsResponse, error) {
	res, err := getRequest[dto.EmojiListReactionsResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MessageAddReaction adds a specific reaction to a message.
func (c *Client) MessageAddReaction(
	ctx context.Context,
	input *dto.EmojiAddReactionRequest,
) (*dto.EmojiAddReactionResponse, error) {
	res, err := postRequest[dto.EmojiAddReactionResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// MessageRemoveReaction removes a specific reaction from a message.
func (c *Client) MessageRemoveReaction(
	ctx context.Context,
	input *dto.EmojiDeleteReactionRequest,
) (*dto.EmojiDeleteReactionResponse, error) {
	res, err := deleteRequest[dto.EmojiDeleteReactionResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// EmojiUseDefault determines whether to use the seven default emojis initially provided by Sendbird.
func (c *Client) EmojiUseDefault(
	ctx context.Context,
	input *dto.EmojiUseDefaultRequest,
) (*dto.EmptyResponse, error) {
	res, err := putRequest[dto.EmptyResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// EmojiListAll retrieves an array of all emojis registered to the application.
func (c *Client) EmojiListAll(
	ctx context.Context,
	input *dto.EmojiListAllRequest,
) (*dto.EmojiListAllResponse, error) {
	res, err := getRequest[dto.EmojiListAllResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// EmojiView retrieves a specific emoji.
func (c *Client) EmojiView(
	ctx context.Context,
	input *dto.EmojiViewRequest,
) (*dto.Emoji, error) {
	res, err := getRequest[dto.Emoji](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// EmojiAdd adds an array of new emojis to the application
func (c *Client) EmojiAdd(
	ctx context.Context,
	input *dto.EmojiAddRequest,
) (*dto.EmojiAddResponse, error) {
	res, err := postRequest[dto.EmojiAddResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// EmojiUpdateURL updates the image URL of a specific emoji
func (c *Client) EmojiUpdateURL(
	ctx context.Context,
	input *dto.EmojiUpdateURLRequest,
) (*dto.EmojiUpdateURLResponse, error) {
	res, err := putRequest[dto.EmojiUpdateURLResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// EmojiDelete deletes a specific emoji from the application
func (c *Client) EmojiDelete(
	ctx context.Context,
	input *dto.EmojiDeleteRequest,
) (*dto.EmojiDeleteResponse, error) {
	res, err := deleteRequest[dto.EmojiDeleteResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// EmojiListAllCategories retrieves an array of all emoji categories with their emojis
func (c *Client) EmojiListAllCategories(
	ctx context.Context,
	input *dto.EmojiListAllCategoriesRequest,
) (*dto.EmojiListAllCategoriesResponse, error) {
	res, err := getRequest[dto.EmojiListAllCategoriesResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// EmojiViewCategory retrieves a specific emoji category including its emoji
func (c *Client) EmojiViewCategory(
	ctx context.Context,
	input *dto.EmojiViewCategoryRequest,
) (*dto.EmojiViewCategoryResponse, error) {
	res, err := getRequest[dto.EmojiViewCategoryResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// EmojiAddCategory adds an array of new emoji categories to the app
func (c *Client) EmojiAddCategory(
	ctx context.Context,
	input *dto.EmojiAddCategoryRequest,
) (*dto.EmojiAddCategoryResponse, error) {
	res, err := postRequest[dto.EmojiAddCategoryResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// EmojiUpdateCategory updates an emoji category URL
func (c *Client) EmojiUpdateCategory(
	ctx context.Context,
	input *dto.EmojiCategoryUpdateRequest,
) (*dto.EmojiCategoryUpdateResponse, error) {
	res, err := putRequest[dto.EmojiCategoryUpdateResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// EmojiDeleteCategory deletes specific emoji category
func (c *Client) EmojiDeleteCategory(
	ctx context.Context,
	input *dto.EmojiCategoryDeleteRequest,
) (*dto.EmojiCategoryDeleteResponse, error) {
	res, err := deleteRequest[dto.EmojiCategoryDeleteResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// MessageTranslate translates a message into specific language
func (c *Client) MessageTranslate(
	ctx context.Context,
	input *dto.MessageTranslateRequest,
) (*dto.MessageTranslateResponse, error) {
	res, err := postRequest[dto.MessageTranslateResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// AnnouncementList lists all announcements
func (c *Client) AnnouncementList(
	ctx context.Context,
	input *dto.AnnouncementListRequest,
) (*dto.AnnouncementListResponse, error) {
	res, err := getRequest[dto.AnnouncementListResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// AnnouncementView retrieves information on specific announcement
func (c *Client) AnnouncementView(
	ctx context.Context,
	input *dto.AnnouncementViewRequest,
) (*dto.AnnouncementViewResponse, error) {
	res, err := getRequest[dto.AnnouncementViewResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// AnnouncementCreate creates an announcement
func (c *Client) AnnouncementCreate(
	ctx context.Context,
	input *dto.AnnouncementCreateRequest,
) (*dto.AnnouncementCreateResponse, error) {
	res, err := postRequest[dto.AnnouncementCreateResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// AnnouncementUpdate Updates an announcement
func (c *Client) AnnouncementUpdate(
	ctx context.Context,
	input *dto.AnnouncementUpdateRequest,
) (*dto.EmptyResponse, error) {
	res, err := putRequest[dto.EmptyResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}

// AnnouncementViewStats retrieves information on specific announcement statistics
func (c *Client) AnnouncementViewStats(
	ctx context.Context,
	input *dto.AnnouncementViewStatisticsRequest,
) (*dto.AnnouncementViewStatisticsResponse, error) {
	res, err := getRequest[dto.AnnouncementViewStatisticsResponse](ctx, c, input)
	if err != nil {
		return nil, err
	}
	return res, res.Error
}
