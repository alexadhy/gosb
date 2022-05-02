package dto

import (
	"fmt"
	"io"
	"time"
)

// UserResponse is returned on any endpoint with User data
type UserResponse struct {
	UserID             *string   `json:"user_id,omitempty"`
	Nickname           *string   `json:"nickname,omitempty"`
	UnreadMessageCount *int64    `json:"unread_message_count,omitempty"`
	ProfileURL         *string   `json:"profile_url,omitempty"`
	AccessToken        *string   `json:"access_token,omitempty"`
	IsOnline           *bool     `json:"is_online,omitempty"`
	IsActive           *bool     `json:"is_active,omitempty"`
	CreatedAt          *int64    `json:"created_at,omitempty"`
	LastSeenAt         *int64    `json:"last_seen_at,omitempty"`
	DiscoveryKeys      []string  `json:"discovery_keys,omitempty"`
	PreferredLanguages []string  `json:"preferred_languages,omitempty"`
	HasEverLoggedIn    *bool     `json:"has_ever_logged_in,omitempty"`
	Metadata           *Metadata `json:"metadata,omitempty"`
	FriendDiscoveryKey []string  `json:"friend_discovery_key,omitempty"`
	State              *string   `json:"state,omitempty"`
	Role               *string   `json:"role,omitempty"`
	Error              Error     `json:"error,omitempty"`
}

// UserGetRequest is the request payload to user GET endpoint
type UserGetRequest struct {
	UserID             string  `json:"user_id" qs:"-"`
	IncludeUnreadCount *bool   `qs:"include_unread_count,omitempty"`
	CustomTypes        *string `qs:"custom_types,omitempty"`
	SuperMode          *string `qs:"super_mode" default:"all"`
}

// URL creates URL string to view user endpoint
func (u UserGetRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s", baseURL, userEndpoint, u.UserID)
	s, err := encodeQS[UserGetRequest](base, u)
	if err != nil {
		return ""
	}
	return s
}

// UserListResponse returned on list users
type UserListResponse struct {
	Users         []UserResponse `json:"users,omitempty"`
	NextPageToken *string        `json:"next,omitempty"`
	Error         *Error         `json:"error,omitempty"`
}

// UserListRequest contains URL params / query string
// needed to filter list user
type UserListRequest struct {
	Token              *string `qs:"token,omitempty"`
	Limit              *int    `qs:"limit,omitempty"`
	ActiveMode         *string `qs:"active_mode,omitempty"`
	ShowBot            *bool   `qs:"show_bot,omitempty"`
	UserIDs            *string `qs:"user_ids,omitempty"`
	NickName           *string `qs:"nickname,omitempty"`
	NickNameStartsWith *string `qs:"nickname_startswith,omitempty"`
	MetadataKey        *string `qs:"metadatakey,omitempty"`
	MetadataValuesIn   *string `qs:"metadatavalues_in,omitempty"`
}

// URL creates URL string to user list request plus encodes query param as well
func (u UserListRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s", baseURL, userEndpoint)
	s, err := encodeQS[UserListRequest](base, u)
	if err != nil {
		return ""
	}
	return s
}

// UserCreateRequest is the request payload to create user endpoint.
type UserCreateRequest struct {
	UserID           string    `json:"user_id"`
	Nickname         string    `json:"nickname"`
	ProfileURL       string    `json:"profile_url"`
	ProfileFile      io.Reader `json:"profile_file,omitempty"`
	IssueAccessToken *bool     `json:"issue_access_token" default:"false"`
	DiscoveryKeys    []string  `json:"discovery_keys"`
	Metadata         *Metadata `json:"metadata,omitempty"`
}

// UserUpdateRequest is the request payload to update user endpoint
type UserUpdateRequest struct {
	UserID                  string     `json:"user_id"`
	Nickname                *string    `json:"nickname,omitempty"`
	ProfileURL              *string    `json:"profile_url,omitempty"`
	ProfileFile             io.Reader  `json:"profile_file,omitempty"`
	IssueAccessToken        *bool      `json:"issue_access_token,omitempty" default:"false"`
	IsActive                *bool      `json:"is_active,omitempty"`
	LastSeenAt              *time.Time `json:"last_seen_at,omitempty"`
	DiscoveryKeys           []string   `json:"discovery_keys"`
	PreferredLanguages      []string   `json:"preferred_languages"`
	LeaveAllWhenDeactivated *bool      `json:"leave_all_when_deactivated,omitempty" default:"true"`
}

// URL creates URL string to update user endpoint
func (u UserUpdateRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s", baseURL, userEndpoint, u.UserID)
}

// UserDeleteRequest is the request payload to delete user endpoint
type UserDeleteRequest struct {
	UserID string `json:"user_id"`
}

// URL creates URL string to delete user endpoint
func (u UserDeleteRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s", baseURL, userEndpoint, u.UserID)
}

// UserCreateSessionTokenRequest creates a new session token for user
type UserCreateSessionTokenRequest struct {
	UserID string `json:"user_id"`
}

// URL creates URL string to user create session token request endpoint
func (u UserCreateSessionTokenRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/token", baseURL, userEndpoint, u.UserID)
}

// UserCreateSessionTokenResponse is the response to session token issuance
type UserCreateSessionTokenResponse struct {
	UserID    string `json:"user_id"`
	Token     string `json:"token"`
	ExpiresAt *int64 `json:"expires_at"`
	Error     *Error `json:"error,omitempty"`
}

// UserRevokeAllSessionTokenRequest is the request DTO to revoke all session token endpoint
type UserRevokeAllSessionTokenRequest struct {
	UserID string
}

func (u UserRevokeAllSessionTokenRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/token", baseURL, userEndpoint, u.UserID)
}

// UserListGroupChannelRequest is the request query string to
type UserListGroupChannelRequest struct {
	UserID               string            `qs:"-"`
	Token                *string           `qs:"token,omitempty"`
	Limit                *int              `qs:"limit,omitempty"`
	DistinctMode         DistinctMode      `qs:"distinct_mode" default:"all"`
	PublicMode           PublicMode        `qs:"public_mode" default:"all"`
	SuperMode            SuperMode         `qs:"super_mode" default:"all"`
	HiddenMode           HiddenMode        `qs:"hidden_mode" default:"unhidden_status"`
	MemberStateFilter    MemberStateFilter `qs:"member_state_filter" default:"all"`
	UnreadFilter         UnreadFilter      `qs:"unread_filter" default:"all"`
	CreatedAfter         *int64            `qs:"created_after,omitempty"`
	CreatedBefore        *int64            `qs:"created_before,omitempty"`
	ShowEmpty            bool              `qs:"show_empty" default:"false"`
	ShowFrozen           bool              `qs:"show_frozen" default:"true"`
	ShowMember           bool              `qs:"show_member" default:"false"`
	ShowDeliveryReceipt  bool              `qs:"show_delivery_receipt" default:"false"`
	ShowReadReceipt      bool              `qs:"show_read_receipt" default:"false"`
	Order                Order             `qs:"order,omitempty" default:"chronological"`
	MetadataOrderKey     *string           `qs:"metadata_order_key,omitempty"`
	CustomTypes          []string          `qs:"custom_types,comma,omitempty"`
	CustomTypeStartsWith *string           `qs:"custom_type_startswith,omitempty"`
	ChannelURLs          []string          `qs:"channel_urls,comma,omitempty"`
	Name                 *string           `qs:"name,omitempty"`
	NameContains         *string           `qs:"name_contains,omitempty"`
	NameStartsWith       *string           `qs:"name_startswith,omitempty"`
	// MembersExactlyIn Searches for group channels that contain the exact members specified in the string. The value should be a comma-separated string with multiple urlencoded user IDs.
	// An example of a string of urlencoded user IDs is ?members_exactly_in=urlencoded_user_id_1, urlencoded_user_id_2.
	MembersExactlyIn        []string   `qs:"members_exactly_in,comma,omitempty"`
	MembersIncludeIn        []string   `qs:"members_include_in,comma,omitempty"`
	QueryType               *LogicalOp `qs:"query_type,omitempty"`
	MembersNickname         []string   `qs:"members_nickname,comma,omitempty"`
	MembersNicknameContains *string    `qs:"members_nickname_contains,omitempty"`
	SearchQuery             *string    `qs:"search_query,omitempty"`
	SearchFields            []string   `qs:"search_fields,comma,omitempty"`
	MetadataKey             *string    `qs:"metadata_key,omitempty"`
	MetadataValues          []string   `qs:"metadata_values,comma,omitempty"`
	MetadataValueStartsWith *string    `qs:"metadata_value_startswith,omitempty"`
	MetacounterKey          *string    `qs:"metacounter_key,omitempty"`
	MetacounterValues       []string   `qs:"metacounter_values,comma,omitempty"`
	MetacounterValueGT      *string    `qs:"metacounter_value_gt,omitempty"`
	MetacounterValueGTE     *string    `qs:"metacounter_value_gte,omitempty"`
	MetacounterValueLT      *string    `qs:"metacounter_value_lt,omitempty"`
	MetacounterValueLTE     *string    `qs:"metacounter_value_lte,omitempty"`
}

// URL returns URL for sendbird list group channel by user endpoint
// GET https://api-{application_id}.sendbird.com/v3/users/{user_id}/my_group_channels
func (u UserListGroupChannelRequest) URL(baseURL string) string {
	base := fmt.Sprintf(
		"%s/%s/%s/my_group_channels",
		baseURL,
		userEndpoint,
		u.UserID,
	)
	s, err := encodeQS[UserListGroupChannelRequest](base, u)
	if err != nil {
		return ""
	}
	return s
}

type UserListGroupChannelResponse struct {
	GroupChannels []GroupChannelResponse `json:"channels,omitempty"`
	NextPageToken *string                `json:"next,omitempty"`
	Error         *Error                 `json:"error,omitempty"`
}

// UserLeaveAllGroupChannelRequest is a payload for leaving all group_channel endpoint
type UserLeaveAllGroupChannelRequest struct {
	UserID     string  `json:"-"`
	CustomType *string `json:"custom_type,omitempty"`
}

// URL returns URL string for sendbird leave all user's group channel
// PUT https://api-{application_id}.sendbird.com/v3/users/{user_id}/my_group_channels/leave
func (u UserLeaveAllGroupChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf(
		"%s/%s/%s/my_group_channels",
		baseURL,
		userEndpoint,
		u.UserID,
	)
}
