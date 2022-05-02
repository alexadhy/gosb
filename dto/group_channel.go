package dto

import "fmt"

// GroupChannelListOperatorsRequest is the DTO struct for listing group channel operators
type GroupChannelListOperatorsRequest struct {
	ChannelURL string `qs:"-"`
	Token      string `qs:"token"`
	Limit      int    `qs:"limit"`
}

// URL creates group channel list operator URL string
func (o GroupChannelListOperatorsRequest) URL(baseURL string) string {
	base := fmt.Sprintf(
		"%s/%s/%s/operators",
		baseURL,
		groupChannelEndpoint,
		o.ChannelURL,
	)
	s, err := encodeQS[GroupChannelListOperatorsRequest](base, o)
	if err != nil {
		return ""
	}
	return s
}

// GroupChannelListOperatorsResponse is the return response for listing operator endpoint
type GroupChannelListOperatorsResponse struct {
	Operators     []UserResponse `json:"operators,omitempty"`
	NextPageToken string         `json:"next,omitempty"`
	Error         Error          `json:"error,omitempty"`
}

// GroupChannelRegisterOperatorRequest is the request payload for registering operator to an group channel
type GroupChannelRegisterOperatorRequest struct {
	ChannelURL  string   `json:"-"`
	OperatorIDs []string `json:"operator_ids"`
}

// URL creates group channel register operator URL string
func (o GroupChannelRegisterOperatorRequest) URL(baseURL string) string {
	return fmt.Sprintf(
		"%s/%s/%s/operators",
		baseURL,
		groupChannelEndpoint,
		o.ChannelURL,
	)
}

// GroupChannelUnregisterOperatorRequest creates a request query string to unregister group channel operator endpoint.
type GroupChannelUnregisterOperatorRequest struct {
	ChannelURL string   `qs:"-"`
	OpIDs      []string `qs:"operator_ids,comma"`
}

// URL is the URL for GroupChannelUnregisterOperator
func (o GroupChannelUnregisterOperatorRequest) URL(baseURL string) string {
	base := fmt.Sprintf(
		"%s/%s/%s/operators",
		baseURL,
		groupChannelEndpoint,
		o.ChannelURL,
	)
	s, err := encodeQS[GroupChannelUnregisterOperatorRequest](base, o)
	if err != nil {
		return ""
	}
	return s
}

// GroupChannelResponse is the response payload to any group channel endpoint
type GroupChannelResponse struct {
	Name                 *string        `json:"name,omitempty"`
	ChannelURL           *string        `json:"channel_url,omitempty"`
	CoverURL             *string        `json:"cover_url,omitempty"`
	CustomType           *string        `json:"custom_type,omitempty"`
	UnreadMessageCount   *int64         `json:"unread_message_count,omitempty"`
	Data                 *string        `json:"data,omitempty"`
	IsDistinct           *bool          `json:"is_distinct,omitempty"`
	IsPublic             *bool          `json:"is_public,omitempty"`
	IsSuper              *bool          `json:"is_super,omitempty"`
	IsEphemeral          *bool          `json:"is_ephemeral,omitempty"`
	IsAccessCodeRequired *bool          `json:"is_access_code_required,omitempty"`
	HiddenState          *string        `json:"hidden_state,omitempty"`
	MemberCount          *int64         `json:"member_count,omitempty"`
	JoinedMemberCount    *int64         `json:"joined_member_count,omitempty"`
	Members              []UserResponse `json:"members,omitempty"`
	Operators            []UserResponse `json:"operators,omitempty"`
	MaxLengthMessage     *int64         `json:"max_length_message,omitempty"`
	LastMessage          Message        `json:"last_message"`
	CreatedAt            *int64         `json:"created_at,omitempty"`
	Freeze               *bool          `json:"freeze,omitempty"`
}

// GroupChannelListRequest is the url query values for List group channels endpoint
type GroupChannelListRequest struct {
	Token                string            `qs:"token,omitempty"`
	Limit                int               `qs:"limit,omitempty"`
	DistinctMode         DistinctMode      `qs:"distinct_mode" default:"all"`
	PublicMode           PublicMode        `qs:"public_mode" default:"all"`
	SuperMode            SuperMode         `qs:"super_mode" default:"all"`
	HiddenMode           HiddenMode        `qs:"hidden_mode" default:"unhidden_status"`
	MemberStateFilter    MemberStateFilter `qs:"member_state_filter" default:"all"`
	UnreadFilter         UnreadFilter      `qs:"unread_filter" default:"all"`
	CreatedAfter         int64             `qs:"created_after,omitempty"`
	CreatedBefore        int64             `qs:"created_before,omitempty"`
	ShowEmpty            bool              `qs:"show_empty" default:"false"`
	ShowFrozen           bool              `qs:"show_frozen" default:"true"`
	ShowMember           bool              `qs:"show_member" default:"false"`
	ShowDeliveryReceipt  bool              `qs:"show_delivery_receipt" default:"false"`
	ShowReadReceipt      bool              `qs:"show_read_receipt" default:"false"`
	Order                Order             `qs:"order,omitempty" default:"chronological"`
	MetadataOrderKey     string            `qs:"metadata_order_key,omitempty"`
	CustomTypes          []string          `qs:"custom_types,comma,omitempty"`
	CustomTypeStartsWith string            `qs:"custom_type_startswith,omitempty"`
	ChannelURLs          []string          `qs:"channel_urls,comma,omitempty"`
	Name                 string            `qs:"name,omitempty"`
	NameContains         string            `qs:"name_contains,omitempty"`
	NameStartsWith       string            `qs:"name_startswith,omitempty"`
	// MembersExactlyIn Searches for group channels that contain the exact members specified in the string. The value should be a comma-separated string with multiple urlencoded user IDs.
	// An example of a string of urlencoded user IDs is ?members_exactly_in=urlencoded_user_id_1, urlencoded_user_id_2.
	MembersExactlyIn        []string  `qs:"members_exactly_in,comma,omitempty"`
	MembersIncludeIn        []string  `qs:"members_include_in,comma,omitempty"`
	QueryType               LogicalOp `qs:"query_type,omitempty"`
	MembersNickname         []string  `qs:"members_nickname,comma,omitempty"`
	MembersNicknameContains string    `qs:"members_nickname_contains,omitempty"`
	SearchQuery             string    `qs:"search_query,omitempty"`
	SearchFields            []string  `qs:"search_fields,comma,omitempty"`
	MetadataKey             string    `qs:"metadata_key,omitempty"`
	MetadataValues          []string  `qs:"metadata_values,comma,omitempty"`
	MetadataValueStartsWith string    `qs:"metadata_value_startswith,omitempty"`
	MetacounterKey          string    `qs:"metacounter_key,omitempty"`
	MetacounterValues       []string  `qs:"metacounter_values,comma,omitempty"`
	MetacounterValueGT      string    `qs:"metacounter_value_gt,omitempty"`
	MetacounterValueGTE     string    `qs:"metacounter_value_gte,omitempty"`
	MetacounterValueLT      string    `qs:"metacounter_value_lt,omitempty"`
	MetacounterValueLTE     string    `qs:"metacounter_value_lte,omitempty"`
}

// URL returns URL string with query values for listing request
func (o GroupChannelListRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s", baseURL, groupChannelEndpoint)
	s, err := encodeQS[GroupChannelListRequest](base, o)
	if err != nil {
		return ""
	}
	return s
}

// GroupChannelListResponse is the response for list group channel endpoint
type GroupChannelListResponse struct {
	Channels      []GroupChannelResponse `json:"channels,omitempty"`
	NextPageToken string                 `json:"next,omitempty"`
	Error         Error                  `json:"error,omitempty"`
}

// GroupChannelGetRequest is there to construct endpoint for get group channel
type GroupChannelGetRequest struct {
	ChannelURL string `json:"channel_url"`
}

// URL creates URL string for get channel request
func (o GroupChannelGetRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s", baseURL, groupChannelEndpoint, o.ChannelURL)
}

// GroupChannelCreateRequest is the payload to send to group channel create endpoint
type GroupChannelCreateRequest struct {
	Users                   []string                    `json:"users"`
	Name                    string                      `json:"name"`
	ChannelURL              string                      `json:"channel_url"`
	CoverURL                string                      `json:"cover_url,omitempty"`
	CoverFile               File                        `json:"cover_file,omitempty"`
	CustomType              string                      `json:"custom_type,omitempty"`
	Data                    string                      `json:"data,omitempty"`
	IsDistinct              *bool                       `json:"is_distinct,omitempty"`
	IsPublic                *bool                       `json:"is_public,omitempty"`
	IsSuper                 *bool                       `json:"is_super,omitempty"`
	IsEphemeral             *bool                       `json:"is_ephemeral,omitempty"`
	AccessCode              string                      `json:"access_code,omitempty"`
	AccessCodeRequired      *bool                       `json:"access_code_required,omitempty"`
	InviterID               string                      `json:"inviter_id,omitempty"`
	Strict                  *bool                       `json:"strict,omitempty"`
	InvitationStatus        map[string]InvitationStatus `json:"invitation_status,omitempty"`
	HiddenStatus            map[string]HiddenMode       `json:"hidden_status,omitempty"`
	OperatorIDs             []string                    `json:"operator_ids,omitempty"`
	BlockSDKUserChannelJoin *bool                       `json:"block_sdk_user_channel_join,omitempty"`
}

// URL creates URL string for create group channel endpoint
func (o GroupChannelCreateRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s", baseURL, groupChannelEndpoint)
}

// GroupChannelUpdateRequest is the payload to send to group channel create endpoint
type GroupChannelUpdateRequest struct {
	ChannelURL  string   `json:"channel_url"`
	Name        string   `json:"name,omitempty"`
	CoverURL    string   `json:"cover_url,omitempty"`
	CoverFile   File     `json:"cover_file,omitempty"`
	CustomType  string   `json:"custom_type,omitempty"`
	Data        string   `json:"data,omitempty"`
	IsDistinct  *bool    `json:"is_distinct,omitempty"`
	IsPublic    *bool    `json:"is_public,omitempty"`
	IsSuper     *bool    `json:"is_super,omitempty"`
	AccessCode  string   `json:"access_code,omitempty"`
	OperatorIDs []string `json:"operator_ids,omitempty"`
}

// URL returns URL string to update group channel endpoint
func (o GroupChannelUpdateRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s", baseURL, groupChannelEndpoint, o.ChannelURL)
}

// GroupChannelDeleteRequest is the constructor for creating delete group-channel endpoint URL
type GroupChannelDeleteRequest struct {
	ChannelURL string `json:"-"`
}

// URL returns URL string for group channel delete endpoint
func (o GroupChannelDeleteRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s", baseURL, groupChannelEndpoint, o.ChannelURL)
}

// GroupChannelJoinRequest is the request payload to join a public group channel
type GroupChannelJoinRequest struct {
	ChannelURL string `json:"-"`
	UserID     string `json:"user_id"`
	AccessCode string `json:"access_code,omitempty"`
}

// GroupChannelJoinResponse is the response payload to group channel join endpoint
type GroupChannelJoinResponse struct {
	Error Error `json:"error,omitempty"`
}

// URL returns URL string for group channel join endpoint
// PUT https://api-{application_id}.sendbird.com/v3/group_channels/{channel_url}/join
func (o GroupChannelJoinRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/join", baseURL, groupChannelEndpoint, o.ChannelURL)
}

// GroupChannelLeaveRequest is the request payload to leave group channel
type GroupChannelLeaveRequest struct {
	ChannelURL     string   `json:"-"`
	UserIDs        []string `json:"user_ids"`
	ShouldLeaveAll *bool    `json:"should_leave_all,omitempty"`
}

// GroupChannelLeaveResponse is the response payload to group channel leave endpoint
type GroupChannelLeaveResponse struct {
	Error Error `json:"error,omitempty"`
}

// URL returns URL string to group channel leave endpoint
func (o GroupChannelLeaveRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/leave", baseURL, groupChannelEndpoint, o.ChannelURL)
}

// GroupChannelHideRequest is the request payload to hide group channel from channel list
type GroupChannelHideRequest struct {
	ChannelURL           string `json:"-"`
	UserID               string `json:"user_id"`
	AllowAutoUnhide      *bool  `json:"allow_auto_unhide,omitempty"`
	ShouldHideAll        *bool  `json:"should_hide_all,omitempty"`
	HidePreviousMessages *bool  `json:"hide_previous_messages,omitempty"`
}

// URL returns URL string to group channel hide endpoint
func (o GroupChannelHideRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/hide", baseURL, groupChannelEndpoint, o.ChannelURL)
}

// GroupChannelHideResponse is the response payload to hide group channel endpoint
type GroupChannelHideResponse struct {
	Error *Error `json:"error,omitempty"`
}

// GroupChannelResetHistoryRequest is the payload to group_channel reset history endpoint
type GroupChannelResetHistoryRequest struct {
	ChannelURL string `json:"-"`
	UserID     string `json:"user_id"`
	ResetAll   *bool  `json:"reset_all,omitempty"`
}

// URL returns URL string to reset user history endpoint
// PUT https://api-{application_id}.sendbird.com/v3/group_channels/{channel_url}/reset_user_history
func (o GroupChannelResetHistoryRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/reset_user_history", baseURL, groupChannelEndpoint, o.ChannelURL)
}

// GroupChannelResetHistoryResponse is the response to reset history endpoint
type GroupChannelResetHistoryResponse struct {
	Error *Error `json:"error,omitempty"`
}

// GroupChannelListMembersRequest is the request type to list all Members in group_channel
type GroupChannelListMembersRequest struct {
	ChannelURL          string             `qs:"-"`
	Token               *string            `qs:"token,omitempty"`
	Limit               *int               `qs:"limit,omitempty"`
	ShowDeliveryReceipt *bool              `qs:"show_delivery_receipt,omitempty"`
	ShowReadReceipt     *bool              `qs:"show_read_receipt,omitempty"`
	ShowMemberIsMuted   *bool              `qs:"show_member_is_muted,omitempty"`
	Order               *Order             `qs:"order,omitempty"`
	OperatorFilter      *OperatorFilter    `qs:"operator_filter,omitempty"`
	MemberStateFilter   *MemberStateFilter `qs:"member_state_filter,omitempty"`
	MutedMemberFilter   *MutedMemberFilter `qs:"muted_member_filter,omitempty"`
	NicknameStartsWith  *string            `qs:"nickname_startswith,omitempty"`
}

// URL returns URL string to list Members in group_channel endpoint
func (o GroupChannelListMembersRequest) URL(baseURL string) string {
	base := fmt.Sprintf(
		"%s/%s/%s/members",
		baseURL,
		groupChannelEndpoint,
		o.ChannelURL,
	)
	s, err := encodeQS[GroupChannelListMembersRequest](base, o)
	if err != nil {
		return ""
	}
	return s
}

// GroupChannelListMembersResponse is the response payload type on listing open channel Members endpoint.
type GroupChannelListMembersResponse struct {
	Members       []UserResponse `json:"members,omitempty"`
	NextPageToken *string        `json:"next,omitempty"`
	Error         *Error         `json:"error,omitempty"`
}

// GroupChannelCheckMemberRequest is the request parameter to check user membership in a group
type GroupChannelCheckMemberRequest struct {
	ChannelURL string `json:"-"`
	UserID     string `json:"-"`
}

// URL returns URL string to check user membership in group channel endpoint
// GET https://api-{application_id}.sendbird.com/v3/group_channels/{channel_url}/members/{user_id}
func (o GroupChannelCheckMemberRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/members/%s", baseURL, groupChannelEndpoint, o.ChannelURL, o.UserID)
}

// GroupChannelCheckMemberResponse is the response type / payload to GroupChannelCheckMemberResponse endpoint
type GroupChannelCheckMemberResponse struct {
	IsMember *bool  `json:"is_member,omitempty"`
	Error    *Error `json:"error,omitempty"`
}
