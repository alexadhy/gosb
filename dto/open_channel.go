package dto

import "fmt"

// OpenChannelListOperatorsRequest is the DTO struct for listing open channel operators
type OpenChannelListOperatorsRequest struct {
	ChannelURL string  `qs:"-"`
	Token      *string `qs:"token"`
	Limit      *int    `qs:"limit"`
}

// URL creates open channel list operator URL string
func (o OpenChannelListOperatorsRequest) URL(baseURL string) string {
	base := fmt.Sprintf(
		"%s/%s/%s/operators",
		baseURL,
		openChannelEndpoint,
		o.ChannelURL,
	)
	s, err := encodeQS[OpenChannelListOperatorsRequest](base, o)
	if err != nil {
		return ""
	}
	return s
}

// OpenChannelListOperatorsResponse is the return response for listing operator endpoint
type OpenChannelListOperatorsResponse struct {
	Operators     []UserResponse `json:"operators,omitempty"`
	NextPageToken string         `json:"next,omitempty"`
	Error         Error          `json:"error,omitempty"`
}

// OpenChannelRegisterOperatorRequest is the request payload for registering operator to an open channel
type OpenChannelRegisterOperatorRequest struct {
	ChannelURL  string   `json:"-"`
	OperatorIDs []string `json:"operator_ids"`
}

// URL creates open channel list operator URL string
func (o OpenChannelRegisterOperatorRequest) URL(baseURL string) string {
	return fmt.Sprintf(
		"%s/%s/%s/operators",
		baseURL,
		openChannelEndpoint,
		o.ChannelURL,
	)
}

// OpenChannelUnregisterOperatorRequest creates a request query string to unregister open channel operator endpoint.
type OpenChannelUnregisterOperatorRequest struct {
	ChannelURL string   `qs:"-"`
	OpIDs      []string `qs:"operator_ids,comma"`
}

func (o OpenChannelUnregisterOperatorRequest) URL(baseURL string) string {
	base := fmt.Sprintf(
		"%s/%s/%s/operators",
		baseURL,
		openChannelEndpoint,
		o.ChannelURL,
	)
	s, err := encodeQS[OpenChannelUnregisterOperatorRequest](base, o)
	if err != nil {
		return ""
	}
	return s
}

// OpenChannelResponse is the response to open channel view endpoint
type OpenChannelResponse struct {
	Name             *string        `json:"name,omitempty"`
	ChannelURL       *string        `json:"channel_url,omitempty"`
	CustomType       *string        `json:"custom_type,omitempty"`
	IsEphemeral      *bool          `json:"is_ephemeral,omitempty"`
	ParticipantCount *int64         `json:"participant_count,omitempty"`
	MaxLengthMessage *int64         `json:"max_length_message,omitempty"`
	CreatedAt        *int64         `json:"created_at,omitempty"`
	CoverURL         *string        `json:"cover_url,omitempty"`
	Data             *string        `json:"data,omitempty"`
	Operators        []UserResponse `json:"operators,omitempty"`
	Freeze           *bool          `json:"freeze,omitempty"`
}

// OpenChannelListRequest is the url query values for List open channels endpoint
type OpenChannelListRequest struct {
	Token        string   `qs:"token,omitempty"`
	Limit        int      `qs:"limit,omitempty"`
	CustomTypes  []string `qs:"custom_types,comma,omitempty"`
	NameContains string   `qs:"name_contains,omitempty"`
	URLContains  string   `qs:"url_contains,omitempty"`
	ShowFrozen   *bool    `qs:"show_frozen,omitempty"`
	ShowMetadata *bool    `qs:"show_metadata,omitempty"`
}

// URL returns URL string with query values for listing request
func (o OpenChannelListRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s", baseURL, openChannelEndpoint)
	s, err := encodeQS[OpenChannelListRequest](base, o)
	if err != nil {
		return ""
	}
	return s
}

// OpenChannelListResponse is the response for list open channel endpoint
type OpenChannelListResponse struct {
	Channels      []OpenChannelResponse `json:"channels,omitempty"`
	NextPageToken string                `json:"next,omitempty"`
	Error         Error                 `json:"error,omitempty"`
}

// OpenChannelGetRequest is there to construct endpoint for get open channel
type OpenChannelGetRequest struct {
	ChannelURL string `json:"channel_url"`
}

// URL creates URL string for get channel request
func (o OpenChannelGetRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s", baseURL, openChannelEndpoint, o.ChannelURL)
}

// OpenChannelCreateRequest is the payload to send to sendbird's create channel endpoint
type OpenChannelCreateRequest struct {
	ChannelURL           string   `json:"channel_url"`
	Name                 string   `json:"name"`
	CoverURL             string   `json:"cover_url"`
	CoverFile            *File    `json:"cover_file,omitempty"`
	CustomType           *string  `json:"custom_type,omitempty"`
	Data                 string   `json:"data"`
	IsEphemeral          *bool    `json:"is_ephemeral,omitempty"`
	IsDynamicPartitioned *bool    `json:"is_dynamic_partitioned,omitempty"`
	OperatorIDs          []string `json:"operator_ids,omitempty"`
}

// URL returns URL string for creating open channel endpoint
func (o OpenChannelCreateRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s", baseURL, openChannelEndpoint)
}

// OpenChannelUpdateRequest is the payload to send to SendBird's update open channel endpoint
type OpenChannelUpdateRequest struct {
	ChannelURL  string   `json:"channel_url"`
	Name        *string  `json:"name,omitempty"`
	CoverURL    *string  `json:"cover_url,omitempty"`
	CoverFile   *File    `json:"cover_file,omitempty"`
	CustomType  *string  `json:"custom_type,omitempty"`
	Data        *string  `json:"data,omitempty"`
	OperatorIDs []string `json:"operator_ids,omitempty"`
}

// URL returns URL string for open channel update endpoint
func (o OpenChannelUpdateRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s", baseURL, openChannelEndpoint, o.ChannelURL)
}

// OpenChannelDeleteRequest is the constructor for creating delete open-channel endpoint URL
type OpenChannelDeleteRequest struct {
	ChannelURL string `json:"-"`
}

// URL returns URL string for open channel delete endpoint
func (o OpenChannelDeleteRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s", baseURL, openChannelEndpoint, o.ChannelURL)
}

// OpenChannelListParticipantsRequest is the request type to list all participants in open_channel
type OpenChannelListParticipantsRequest struct {
	ChannelURL string `qs:"-"`
	Token      string `qs:"token,omitempty"`
	Limit      int    `qs:"limit,omitempty"`
}

// URL returns URL string to list participants in open_channel endpoint
func (o OpenChannelListParticipantsRequest) URL(baseURL string) string {
	base := fmt.Sprintf(
		"%s/%s/%s/participants",
		baseURL,
		openChannelEndpoint,
		o.ChannelURL,
	)
	s, err := encodeQS[OpenChannelListParticipantsRequest](base, o)
	if err != nil {
		return ""
	}
	return s
}

// OpenChannelListParticipantsResponse is the response payload type on listing open channel participants endpoint.
type OpenChannelListParticipantsResponse struct {
	Participants  []UserResponse `json:"participants,omitempty"`
	NextPageToken *string        `json:"next,omitempty"`
	Error         *Error         `json:"error,omitempty"`
}
