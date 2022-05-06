package dto

import "fmt"

// EmojiCategory corresponds to SendBird's response type to emoji category
type EmojiCategory struct {
	ID            int    `json:"id,omitempty"`
	Name          string `json:"name,omitempty"` // Name of the emoji category
	URL           string `json:"url,omitempty"`
	ApplicationID int    `json:"application_id,omitempty"`
	Error         *Error `json:"error,omitempty"`
}

// Emoji corresponds to the SendBird's response type to emoji
type Emoji struct {
	ID              int    `json:"id,omitempty"`
	Key             string `json:"key,omitempty"`
	URL             string `json:"url,omitempty"`
	ApplicationID   int    `json:"application_id,omitempty"`
	EmojiCategoryID int    `json:"emoji_category_id,omitempty"`
	Error           *Error `json:"error,omitempty"`
}

// EmojiEnableRequest turns on and off message reactions in SendBird app
type EmojiEnableRequest struct {
	Enabled bool `json:"enabled"`
}

// URL returns URL string to enable emoji reaction settings
// PUT https://api-{application_id}.sendbird.com/v3/applications/settings/reactions
func (e EmojiEnableRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/reactions", baseURL, settingsEndpoint)
}

// EmojiEnableResponse is the response type to emoji enablement preference endpoint.
type EmojiEnableResponse struct {
	Reactions *bool  `json:"reactions"`
	Error     *Error `json:"error,omitempty"`
}

// EmojiListReactionsRequest contains URL params, and query values for listing emoji reactions endpoint
type EmojiListReactionsRequest struct {
	ChannelType ChannelType `json:"-" qs:"-"`
	ChannelURL  string      `json:"-" qs:"-"`
	MessageID   int64       `json:"-" qs:"-"`
	ListUsers   *bool       `qs:"list_users"`
}

// URL returns URL string to list emoji reactions in a message
// GET https://api-{application_id}.sendbird.com/v3/{channel_type}/{channel_url}/messages/{message_id}/reactions
func (e EmojiListReactionsRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/%s/%d/reactions", baseURL, e.ChannelType, e.ChannelURL, msgEndpoint, e.MessageID)
	s, _ := encodeQS[EmojiListReactionsRequest](base, e)
	return s
}

// ReactionListUserOrCount is a generic constraint for emoji reaction response
type ReactionListUserOrCount interface {
	[]string | int
}

// ReactionListResponse is the response type to a reaction
type ReactionListResponse[T ReactionListUserOrCount] map[string]T

// EmojiAddReactionRequest is the request params & payload to send to add reaction to a message
type EmojiAddReactionRequest struct {
	ChannelType ChannelType `json:"-" qs:"-"`
	ChannelURL  string      `json:"-" qs:"-"`
	MessageID   int64       `json:"-" qs:"-"`
	UserID      string      `json:"user_id"`
	Reaction    string      `json:"reaction"`
}

// URL returns URL string to add emoji reaction in a message
// POST https://api-{application_id}.sendbird.com/v3/{channel_type}/{channel_url}/messages/{message_id}/reactions
func (e EmojiAddReactionRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%d/reactions", baseURL, e.ChannelType, e.ChannelURL, msgEndpoint, e.MessageID)
}

// EmojiAddReactionResponse is the response type to add emoji reaction to a message
type EmojiAddReactionResponse EmojiReactionEventResponse

// EmojiReactionEventResponse is the response type to any operation on emoji reaction
type EmojiReactionEventResponse struct {
	UserID    *string `json:"user_id,omitempty"`
	Operation *string `json:"operation,omitempty"`
	Success   *bool   `json:"success,omitempty"`
	Reaction  *string `json:"reaction,omitempty"`
	UpdatedAt *int64  `json:"updated_at,omitempty"`
	Error     *Error  `json:"error,omitempty"`
}

// EmojiDeleteReactionRequest is the URL params & query value(s) for deleting emoji reaction to a message
type EmojiDeleteReactionRequest struct {
	ChannelType ChannelType `json:"-" qs:"-"`
	ChannelURL  string      `json:"-" qs:"-"`
	MessageID   int64       `json:"-" qs:"-"`
	UserID      string      `qs:"user_id"`
	Reaction    string      `qs:"reaction"`
}

// URL returns URL string to delete emoji reaction in a message
// DELETE https://api-{application_id}.sendbird.com/v3/{channel_type}/{channel_url}/messages/{message_id}/reactions
func (e EmojiDeleteReactionRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/%s/%d/reactions", baseURL, e.ChannelType, e.ChannelURL, msgEndpoint, e.MessageID)
	s, _ := encodeQS[EmojiDeleteReactionRequest](base, e)
	return s
}

// EmojiDeleteReactionResponse is the response type to Delete emoji reaction to a message
type EmojiDeleteReactionResponse EmojiReactionEventResponse

// EmojiUseDefaultRequest is the request params, url values and payload to adjust default emoji reaction settings
type EmojiUseDefaultRequest struct {
	UseDefaultEmoji bool `json:"use_default_emoji"`
}

// URL returns URL string to adjust preferences of using standard / default 7 emoji reactions in settings
// PUT https://api-{application_id}.sendbird.com/v3/applications/settings/use_default_emoji
func (e EmojiUseDefaultRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/use_default_emoji", baseURL, settingsEndpoint)
}

// EmojiListAllRequest is the request param / query values / payload for listing all emojis available
type EmojiListAllRequest struct{}

// EmojiListAllResponse lists all emoji resources
type EmojiListAllResponse struct {
	Emojis []Emoji `json:"emojis,omitempty"`
	Error  *Error  `json:"error,omitempty"`
}

// URL returns URL string to list all available emojis in an application
// GET https://api-{application_id}.sendbird.com/v3/emojis
func (e EmojiListAllRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/emojis", baseURL)
}

// EmojiViewRequest is the request type to retrieve a specific emoji
type EmojiViewRequest struct {
	EmojiKey string `json:"emoji_key"`
}

// URL returns URL string to view single specific emoji in an app
// GET https://api-{application_id}.sendbird.com/v3/emojis/{emoji_key}
func (e EmojiViewRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/emojis/%s", baseURL, e.EmojiKey)
}

// EmojiAddRequest is the request payload to add emojis to the application
type EmojiAddRequest struct {
	CategoryID int     `json:"emoji_category_id"`
	Emojis     []Emoji `json:"emojis"` // only key and URL is needed here
}

// URL returns URL string to add emojis to the application
// POST https://api-{application_id}.sendbird.com/v3/emojis
func (e EmojiAddRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/emojis", baseURL)
}

// EmojiAddResponse returns list of newly added []Emoji
type EmojiAddResponse EmojiListAllResponse

// EmojiCategoryUpdateRequest contains URL params and request body to update emoji in an app
type EmojiCategoryUpdateRequest struct {
	CategoryID  int    `json:"-"`
	CategoryURL string `json:"url"`
}

// URL returns URL string to update emoji category endpoint
// PUT https://api-{application_id}.sendbird.com/v3/emoji_categories/{emoji_category_id}
func (e EmojiCategoryUpdateRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/emoji_categories/%d", baseURL, e.CategoryID)
}

// EmojiCategoryUpdateResponse is an alias for Emoji
type EmojiCategoryUpdateResponse Emoji

// EmojiCategoryDeleteRequest is the request URL params for deleting emoji category
type EmojiCategoryDeleteRequest struct {
	CategoryID int `json:"-"`
}

// URL returns URL string to EmojiCategory delete endpoint
func (e EmojiCategoryDeleteRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/emoji_categories/%d", baseURL, e.CategoryID)
}

// EmojiCategoryDeleteResponse is the response type to EmojiCategory delete endpoint
type EmojiCategoryDeleteResponse EmptyResponse
