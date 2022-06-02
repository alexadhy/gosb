package dto

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"strconv"
)

// ThreadInfo is thread info about messages
type ThreadInfo struct {
	ReplyCount    *int64           `json:"reply_count,omitempty"`
	MostReplies   []map[string]any `json:"most_replies,omitempty"`
	LastRepliedAt *int64           `json:"last_replied_at,omitempty"`
	UpdatedAt     *int64           `json:"updated_at,omitempty"`
	Error
}

// Message may contain Text, File, or Admin Message
type Message struct {
	*TextMessage
	*FileMessage
	*AdminMessage
	MessageSurvivalSeconds *int `json:"message_survival_seconds,omitempty"`
	Error
}

// UnmarshalJSON unmarshal Message type to its member type (FileMessage / TextMessage / AdminMessage)
func (m Message) UnmarshalJSON(data []byte) error {
	msi := map[string]any{}
	if err := json.Unmarshal(data, &msi); err != nil {
		return err
	}
	msgType, ok := messageType[msi["type"].(string)]
	if !ok {
		return errors.New("error unmarshaling json to Message format: unknown message type")
	}
	if err := json.Unmarshal(data, &msgType); err != nil {
		return err
	}
	switch msgType.(type) {
	case *FileMessage:
		m.FileMessage = msgType.(*FileMessage)
	case *TextMessage:
		m.TextMessage = msgType.(*TextMessage)
	case *AdminMessage:
		m.AdminMessage = msgType.(*AdminMessage)
	}
	return nil
}

// MarshalJSON marshal Message type to JSON byte
func (m Message) MarshalJSON() ([]byte, error) {
	if m.FileMessage != nil {
		return json.Marshal(m.FileMessage)
	}
	if m.AdminMessage != nil {
		return json.Marshal(m.AdminMessage)
	}
	if m.TextMessage != nil {
		return json.Marshal(m.TextMessage)
	}
	return nil, errors.New("unknown message type")
}

var (
	messageType = map[string]any{
		"FILE": &FileMessage{},
		"MESG": &TextMessage{},
		"ADMM": &AdminMessage{},
	}
)

// TextMessage is message of type text on sendbird
type TextMessage struct {
	CommonMessage
	Message              string      `json:"message"`
	Translations         any         `json:"translations,omitempty"`
	Data                 *string     `json:"data,omitempty"`
	OGTag                *OpenGraph  `json:"og_tag,omitempty"`
	IsAppleCriticalAlert *bool       `json:"is_apple_critical_alert,omitempty"`
	ThreadInfo           *ThreadInfo `json:"thread_info,omitempty"`
	ParentMessageID      *int64      `json:"parent_message_id,omitempty"`
	ParentMessageInfo    *ThreadInfo `json:"parent_message_info,omitempty"`
	IsReplyToChannel     *bool       `json:"is_reply_to_channel,omitempty"`
}

// File contains File request and response payload
type File struct {
	URL            string `json:"url"`
	Name           string `json:"name"`
	MIMEType       string `json:"type"`
	Size           int    `json:"size"`
	AdditionalData string `json:"data"`
}

// CommonMessage is common fields included in each object of type Message
type CommonMessage struct {
	MessageID       int64            `json:"message_id"`
	MessageType     string           `json:"type"`
	CustomType      string           `json:"custom_type"`
	ChannelURL      string           `json:"channel_url"`
	User            UserResponse     `json:"user"`
	MentionType     string           `json:"mention_type"`
	MentionedUsers  []UserResponse   `json:"mentioned_users"`
	SortedMetaArray *SortedMetaArray `json:"sorted_metaarray,omitempty"`
	IsRemoved       *bool            `json:"is_removed"`
	CreatedAt       int64            `json:"created_at"`
	UpdatedAt       int64            `json:"updated_at,omitempty"`
}

// FileMessage is message of type file on sendbird
type FileMessage struct {
	CommonMessage
	File             *File       `json:"file,omitempty"`
	Thumbnails       []any       `json:"thumbnails,omitempty"`
	RequireAuth      *bool       `json:"require_auth,omitempty"`
	ThreadInfo       *ThreadInfo `json:"thread_info,omitempty"`
	ParentMessageID  *int64      `json:"parent_message_id,omitempty"`
	IsReplyToChannel *bool       `json:"is_reply_to_channel,omitempty"`
}

// AdminMessage is message of type admin on sendbird
type AdminMessage struct {
	CommonMessage
	Message string     `json:"message"`
	Data    *string    `json:"data,omitempty"`
	OGTag   *OpenGraph `json:"og_tag,omitempty"`
}

// CommonMessageRequest is the fields / properties required to construct URL parameters & query value of Message endpoints.
type CommonMessageRequest struct {
	ChannelType ChannelType `json:"-" qs:"-"`
	ChannelURL  string      `json:"-" qs:"-"`
}

// MessageListRequest is the payload to list messages
type MessageListRequest struct {
	CommonMessageRequest
	MessageTS                  *int64          `qs:"message_ts,omitempty"`
	MessageID                  *string         `qs:"message_id,omitempty"`
	PrevLimit                  *int            `qs:"prev_limit,omitempty"`
	NextLimit                  *int            `qs:"next_limit,omitempty"`
	Include                    *bool           `qs:"include,omitempty"`
	Reverse                    *bool           `qs:"reverse,omitempty"`
	SenderID                   *string         `qs:"sender_id,omitempty"`
	SenderIDs                  []string        `qs:"sender_ids,comma,omitempty"`
	OperatorFilter             *OperatorFilter `qs:"operator_filter,omitempty"`
	MessageType                *string         `qs:"message_type,omitempty"`
	CustomTypes                []string        `qs:"custom_type,comma,omitempty"`
	IncludingRemoved           *bool           `qs:"including_removed,omitempty"`
	IncludeParentMessageInfo   *bool           `qs:"include_parent_message_info,omitempty"`
	IncludeThreadInfo          *bool           `qs:"include_thread_info,omitempty"`
	IncludeReplyType           *string         `qs:"include_reply_type,omitempty"` // TODO: create ReplyType
	IncludeReactions           *bool           `qs:"include_reactions,omitempty"`
	WithSortedMetaArray        *bool           `qs:"with_sorted_meta_array,omitempty"`
	ShowSubchannelMessagesOnly *bool           `qs:"show_subchannel_messages_only,omitempty"`
	UserID                     string          `qs:"user_id,omitempty"`
}

// URL returns URL string for list messages endpoint
// GET https://api-{application_id}.sendbird.com/v3/{channel_type}/{channel_url}/messages
func (m MessageListRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/%s", baseURL, m.ChannelType, m.ChannelURL, msgEndpoint)
	s, err := encodeQS[MessageListRequest](base, m)
	if err != nil {
		return ""
	}
	return s
}

// MessageListResponse is the response type for list messages endpoint
type MessageListResponse struct {
	Messages []Message `json:"messages,omitempty"`
	Error
}

// MessageViewRequest is the request parameter & url values to get message endpoint
type MessageViewRequest struct {
	CommonMessageRequest
	MessageID                int64 `qs:"-"`
	IncludeParentMessageInfo *bool `qs:"include_parent_message_info,omitempty"`
	IncludeThreadInfo        *bool `qs:"include_thread_info,omitempty"`
	IncludeReactions         *bool `qs:"include_reactions,omitempty"`
	WithSortedMetaArray      *bool `qs:"with_sorted_meta_array,omitempty"`
}

// URL returns URL string for get message endpoint
// GET https://api-{application_id}.sendbird.com/v3/{channel_type}/{channel_url}/messages/{message_id}
func (m MessageViewRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/%s/%d", baseURL, m.ChannelType, m.ChannelURL, msgEndpoint, m.MessageID)
	s, err := encodeQS[MessageViewRequest](base, m)
	if err != nil {
		return ""
	}
	return s
}

// MessageViewResponse is an alias for Message
type MessageViewResponse Message

// MessageSendRequest is the request payload to send message endpoint
type MessageSendRequest struct {
	*MessageRequest
}

// MessageRequest is the request type to Message
type MessageRequest struct {
	*TextMessageSendRequest
	*FileMessageSendRequest
	*FileMessageSendMultipartRequest
	*AdminMessageSendRequest
}

// MarshalJSON for MessageRequest
func (m MessageRequest) MarshalJSON() ([]byte, error) {
	if m.TextMessageSendRequest != nil {
		return json.Marshal(m.TextMessageSendRequest)
	}
	if m.FileMessageSendRequest != nil {
		return json.Marshal(m.FileMessageSendRequest)
	}
	if m.AdminMessageSendRequest != nil {
		return json.Marshal(m.AdminMessageSendRequest)
	}
	return nil, errors.New("unknown message type")
}

// CommonMessageSendRequest is the common message request to be sent to message payload
type CommonMessageSendRequest struct {
	ChannelType               ChannelType                `json:"-" qs:"-"`
	ChannelURL                string                     `json:"-" qs:"-"`
	MessageType               *string                    `json:"message_type,omitempty"`
	UserID                    *string                    `json:"user_id,omitempty"`
	ApnsBundleID              *string                    `json:"apns_bundle_id,omitempty"`
	AppleCriticalAlertOptions *AppleCriticalAlertOptions `json:"apple_critical_alert_options,omitempty"`
}

// TextMessageSendRequest is the request payload for message send request with type MESG
type TextMessageSendRequest struct {
	CommonMessageSendRequest
	CustomType       *string          `json:"custom_type,omitempty"`
	Message          *string          `json:"message,omitempty"`
	MentionType      *string          `json:"mention_type,omitempty"`
	MentionedUserIDS []string         `json:"mentioned_user_ids,omitempty"`
	SortedMetaarray  []map[string]any `json:"sorted_metaarray,omitempty"`
}

// FileMessageSendMultipartRequest is the request payload for message send request with type FILE
type FileMessageSendMultipartRequest struct {
	CommonMessageSendRequest
	File       *os.File
	Thumbnails map[string]string
}

// DTO returns http.Request and error given a FileMessageSendMultipartRequest
func (fm FileMessageSendMultipartRequest) DTO(baseURL, apiToken string) (*http.Request, error) {
	stat, err := fm.File.Stat()
	if err != nil {
		return nil, err
	}

	hdr := make(textproto.MIMEHeader)
	m := map[string]string{}

	b, err := json.Marshal(&fm.CommonMessageSendRequest)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	if len(fm.Thumbnails) > 0 {
		for k, v := range fm.Thumbnails {
			m[k] = v
		}
	}

	cd := mime.FormatMediaType("form-data", m)
	hdr.Set("Content-Disposition", cd)
	hdr.Set("Content-Type", mime.TypeByExtension(fm.File.Name()))
	hdr.Set("Content-Length", strconv.FormatInt(stat.Size(), 10))

	var buf bytes.Buffer
	mw := multipart.NewWriter(&buf)

	part, err := mw.CreatePart(hdr)
	if err != nil {
		return nil, fmt.Errorf("failed to create new form part: %w", err)
	}

	n, err := io.Copy(part, fm.File)
	if err != nil {
		return nil, fmt.Errorf("failed to write form part: %w", err)
	}

	if n != stat.Size() {
		return nil, fmt.Errorf("file size changed while writing: %s", fm.File.Name())
	}

	err = mw.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to prepare form: %w", err)
	}

	msr := MessageSendRequest{&MessageRequest{FileMessageSendMultipartRequest: &fm}}
	requestURL := msr.URL(baseURL)
	u, err := url.Parse(requestURL)
	if err != nil {
		return nil, err
	}

	httpHeader := make(http.Header)
	httpHeader.Add("Api-Token", apiToken)
	httpHeader.Add("Content-Type", mw.FormDataContentType())

	return &http.Request{
		Method:        http.MethodPost,
		URL:           u,
		Header:        httpHeader,
		Body:          ioutil.NopCloser(&buf),
		ContentLength: int64(buf.Len()),
	}, nil
}

// FileMessageSendRequest is the request payload for message send request with type FILE
type FileMessageSendRequest struct {
	CommonMessageSendRequest
	URL         *string `json:"url,omitempty"`
	FileName    *string `json:"file_name,omitempty"`
	FileType    *string `json:"file_type,omitempty"`
	CustomField *string `json:"custom_field,omitempty"`
}

// AdminMessageSendRequest is the request payload for message send request with type ADMM
type AdminMessageSendRequest struct {
	CommonMessageSendRequest
	Message    *string `json:"message,omitempty"`
	CustomType *string `json:"custom_type,omitempty"`
	IsSilent   *bool   `json:"is_silent,omitempty"`
}

// AppleCriticalAlertOptions is the request payload to adjust apple alert options
type AppleCriticalAlertOptions struct {
	Sound  *string `json:"sound,omitempty"`
	Volume *string `json:"volume,omitempty"`
}

// URL creates URL string to send message endpoint
// POST https://api-{application_id}.sendbird.com/v3/{channel_type}/{channel_url}/messages
func (m MessageSendRequest) URL(baseURL string) string {
	var chanType string
	var chanURL string
	if m.TextMessageSendRequest != nil {
		chanType = string(m.TextMessageSendRequest.ChannelType)
		chanURL = m.TextMessageSendRequest.ChannelURL
	}
	if m.FileMessageSendRequest != nil {
		chanType = string(m.FileMessageSendRequest.ChannelType)
		chanURL = m.FileMessageSendRequest.ChannelURL
	}
	if m.FileMessageSendMultipartRequest != nil {
		chanType = string(m.FileMessageSendMultipartRequest.ChannelType)
		chanURL = m.FileMessageSendMultipartRequest.ChannelURL
	}
	if m.AdminMessageSendRequest != nil {
		chanType = string(m.AdminMessageSendRequest.ChannelType)
		chanURL = m.AdminMessageSendRequest.ChannelURL
	}
	return fmt.Sprintf("%s/%s/%s/%s", baseURL, chanType, chanURL, msgEndpoint)
}

// MessageSentResponse is the response type to MessageSendRequest request
type MessageSentResponse Message

// MessageUpdateRequest is the request payload and URL param to update message
type MessageUpdateRequest struct {
	ChannelType      ChannelType `json:"-" qs:"-"`
	ChannelURL       string      `json:"-" qs:"-"`
	MessageID        int64       `json:"-" qs:"-"`
	MessageType      string      `json:"message_type"`
	Message          *string     `json:"message,omitempty"`
	CustomType       *string     `json:"custom_type,omitempty"`
	Data             *string     `json:"data,omitempty"`
	MentionType      *string     `json:"mention_type,omitempty"`
	MentionedUserIDs []string    `json:"mentioned_user_ids,omitempty"`
}

// URL returns URL string for update message endpoint
// PUT https://api-{application_id}.sendbird.com/v3/{channel_type}/{channel_url}/messages/{message_id}
func (m MessageUpdateRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%d", baseURL, m.ChannelType, m.ChannelURL, msgEndpoint, m.MessageID)
}

// MessageUpdateResponse is the response type to MessageUpdateRequest
type MessageUpdateResponse Message

// MessageDeleteRequest is the request URL parameters / values for delete message endpoint.
type MessageDeleteRequest struct {
	ChannelType ChannelType `json:"-" qs:"-"`
	ChannelURL  string      `json:"-" qs:"-"`
	MessageID   int64       `json:"-" qs:"-"`
}

// URL returns URL string for delete message endpoint
// DELETE https://api-{application_id}.sendbird.com/v3/{channel_type}/{channel_url}/messages/{message_id}
func (m MessageDeleteRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%d", baseURL, m.ChannelType, m.ChannelURL, msgEndpoint, m.MessageID)
}

// MessageDeleteResponse is the response type to MessageDeleteRequest
type MessageDeleteResponse Message

// MessageAddMetaRequest is the request payload to add meta to message
type MessageAddMetaRequest struct {
	ChannelType     ChannelType  `json:"-" qs:"-"`
	ChannelURL      string       `json:"-" qs:"-"`
	MessageID       int64        `json:"-" qs:"-"`
	SortedMetaArray []SortedMeta `json:"sorted_metaarray"`
}

// URL returns URL string to add metadata to message
// POST https://api-{application_id}.sendbird.com/v3/{channel_type}/{channel_url}/messages/{message_id}/sorted_metaarray
func (m MessageAddMetaRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%d/sorted_metaarray", baseURL, m.ChannelType, m.ChannelURL, msgEndpoint, m.MessageID)
}

// MessageAddMetaResponse is an alias to SortedMetaArrayResponse
type MessageAddMetaResponse SortedMetaArrayResponse

// MessageUpdateMetaRequest is the request payload to update sorted_metaarray
type MessageUpdateMetaRequest struct {
	ChannelType     ChannelType  `json:"-" qs:"-"`
	ChannelURL      string       `json:"-" qs:"-"`
	MessageID       int64        `json:"-" qs:"-"`
	SortedMetaArray []SortedMeta `json:"sorted_metaarray"`
	Mode            *string      `json:"mode,omitempty"`
	Upsert          *bool        `json:"upsert,omitempty"`
}

// URL returns URL string to update metadata to message
// PUT https://api-{application_id}.sendbird.com/v3/{channel_type}/{channel_url}/messages/{message_id}/sorted_metaarray
func (m MessageUpdateMetaRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%d/sorted_metaarray", baseURL, m.ChannelType, m.ChannelURL, msgEndpoint, m.MessageID)
}

// MessageUpdateMetaResponse is an alias to SortedMetaArrayResponse
type MessageUpdateMetaResponse SortedMetaArrayResponse

// MessageDeleteMetaRequest is the request URL params and values to delete sorted_metaarray endpoint
type MessageDeleteMetaRequest struct {
	ChannelType ChannelType `json:"-" qs:"-"`
	ChannelURL  string      `json:"-" qs:"-"`
	MessageID   int64       `json:"-" qs:"-"`
	Keys        []string    `qs:"keys,comma"`
}

// URL returns URL string to update metadata to message
// DELETE https://api-{application_id}.sendbird.com/v3/{channel_type}/{channel_url}/messages/{message_id}/sorted_metaarray
func (m MessageDeleteMetaRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/%s/%d/sorted_metaarray", baseURL, m.ChannelType, m.ChannelURL, msgEndpoint, m.MessageID)
	s, _ := encodeQS[MessageDeleteMetaRequest](base, m)
	return s
}

// MessageDeleteMetaResponse is an alias to SortedMetaArrayResponse
type MessageDeleteMetaResponse SortedMetaArrayResponse

// MessageTotalCountInChannelRequest is the request payload to view the number of messages in total on a channel
type MessageTotalCountInChannelRequest struct {
	ChannelType ChannelType `json:"-" qs:"-"`
	ChannelURL  string      `json:"-" qs:"-"`
}

// URL returns URL string to total messages count in a channel endpoint
// GET https://api-{application_id}.sendbird.com/v3/{channel_type}/{channel_url}/messages/total_count
func (m MessageTotalCountInChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/%s/total_count", baseURL, m.ChannelType, m.ChannelURL, msgEndpoint)
}

// MessageTotalCountInChannelResponse is the response type to total number of messages in a channel
type MessageTotalCountInChannelResponse struct {
	Total *int `json:"total,omitempty"`
	Error
}

// MessageListThreadedRepliesRequest is the request URL constructor for list replies in a 1-depth thread
type MessageListThreadedRepliesRequest struct {
	ChannelType              ChannelType     `json:"-" qs:"-"`
	ChannelURL               string          `json:"-" qs:"-"`
	ParentMessageID          int64           `json:"-" qs:"parent_message_id"`
	MessageTS                int64           `qs:"message_ts"`
	IncludeReplyType         *string         `qs:"include_reply_type,omitempty"`
	IncludeParentMessageInfo *bool           `qs:"include_parent_message_info,omitempty"`
	IncludeThreadInfo        *bool           `qs:"include_thread_info,omitempty"`
	PrevLimit                *int            `qs:"prev_limit,omitempty"`
	NextLimit                *int            `qs:"next_limit,omitempty"`
	Include                  *bool           `qs:"include,omitempty"`
	Reverse                  *bool           `qs:"reverse,omitempty"`
	SenderID                 *string         `qs:"sender_id,omitempty"`
	SenderIDs                []string        `qs:"sender_ids,comma,omitempty"`
	OperatorFilter           *OperatorFilter `qs:"operator_filter,omitempty"`
	MessageType              *string         `qs:"message_type,omitempty"`
	CustomType               *string         `qs:"custom_type,omitempty"`
	IncludingRemoved         *bool           `qs:"including_removed,omitempty"`
	IncludeReactions         *bool           `qs:"include_reactions,omitempty"`
	WithSortedMetaArray      *bool           `qs:"with_sorted_meta_array,omitempty"`
}

// URL returns URL string to list replies in a message thread
// GET https://api-{application_id}.sendbird.com/v3/{channel_type}/{channel_url}/messages
func (m MessageListThreadedRepliesRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/%s", baseURL, m.ChannelType, m.ChannelURL, msgEndpoint)
	s, _ := encodeQS[MessageListThreadedRepliesRequest](base, m)
	return s
}

// MessageListThreadedRepliesResponse is an alias to MessageListResponse
type MessageListThreadedRepliesResponse MessageListResponse

// MessageViewReplyRequest is the request payload to view a specific reply in a thread
type MessageViewReplyRequest struct {
	ChannelType              ChannelType `json:"-" qs:"-"`
	ChannelURL               string      `json:"-" qs:"-"`
	MessageID                int64       `json:"-" qs:"message_id"`
	IncludeReplyType         *string     `qs:"include_reply_type,omitempty"`
	IncludeParentMessageInfo *bool       `qs:"include_parent_message_info,omitempty"`
	IncludeThreadInfo        *bool       `qs:"include_thread_info,omitempty"`
}

// URL returns URL string to view specific reply to a message endpoint.
// GET https://api-{application_id}.sendbird.com/v3/{channel_type}/{channel_url}/messages/{message_id}
func (m MessageViewReplyRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/%s/%d", baseURL, m.ChannelType, m.ChannelURL, msgEndpoint, m.MessageID)
	s, _ := encodeQS[MessageViewReplyRequest](base, m)
	return s
}

// MessageViewReplyResponse is an alias for type Message
type MessageViewReplyResponse Message

// MessageViewThreadInfoRequest is the request URL params to view thread info
type MessageViewThreadInfoRequest struct {
	ChannelType     ChannelType `json:"-" qs:"-"`
	ChannelURL      string      `json:"-" qs:"-"`
	ParentMessageID int64       `json:"-" qs:"parent_message_id"`
}

// URL returns URL string to be used for MessageViewThreadInfo endpoint
// GET https://api-{application_id}.sendbird.com/v3/{channel_type}/{channel_url}/messages/thread_info
func (m MessageViewThreadInfoRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/%s/%s", baseURL, m.ChannelType, m.ChannelURL, msgEndpoint, "thread_info")
	s, _ := encodeQS[MessageViewThreadInfoRequest](base, m)
	return s
}

// MessageViewThreadInfoResponse is an alias for ThreadInfo
type MessageViewThreadInfoResponse ThreadInfo

// MessageReplyRequest is the param and request payload for sending reply to message in a thread
type MessageReplyRequest struct {
	ChannelType     ChannelType `json:"-" qs:"-"`
	ChannelURL      string      `json:"-" qs:"-"`
	ParentMessageID int64       `json:"parent_message_id"`
	// common properties
	SendPush       *bool   `json:"send_push,omitempty"`
	MarkAsRead     *bool   `json:"mark_as_read,omitempty"`
	CreatedAt      *int64  `json:"created_at,omitempty"`
	DedupID        *string `json:"dedup_id,omitempty"`
	ReplyToChannel *bool   `json:"reply_to_channel,omitempty"`
	*TextMessageReplyRequest
	*FileMessageReplyRequest
	*FileMessageMultipartReplyRequest
}

// MarshalJSON for MessageRequest
func (m MessageReplyRequest) MarshalJSON() ([]byte, error) {
	if m.TextMessageReplyRequest != nil {
		return json.Marshal(m.TextMessageReplyRequest)
	}
	if m.FileMessageReplyRequest != nil {
		return json.Marshal(m.FileMessageReplyRequest)
	}
	return nil, errors.New("unknown message type")
}

// TextMessageReplyRequest is the request payload to reply message on thread endpoint
type TextMessageReplyRequest struct {
	TextMessageSendRequest
}

// FileMessageReplyRequest is the request payload to reply message on thread endpoint
type FileMessageReplyRequest struct {
	FileMessageSendRequest
}

// FileMessageMultipartReplyRequest is the request payload to reply message on thread endpoint
type FileMessageMultipartReplyRequest struct {
	FileMessageSendMultipartRequest
}

// DTO returns request body and error to FileMessageMultipartReplyRequest payload
func (fm FileMessageMultipartReplyRequest) DTO(baseURL, apiToken string) (*http.Request, error) {
	return fm.FileMessageSendMultipartRequest.DTO(baseURL, apiToken)
}

// URL returns URL string to message reply endpoint
// POST https://api-{application_id}.sendbird.com/v3/{channel_type}/{channel_url}/messages
func (m MessageReplyRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/%s", baseURL, m.ChannelType, m.ChannelURL, msgEndpoint)
}

// MessageReplyResponse is an alias to Message
type MessageReplyResponse Message

// MessageDeleteReplyRequest is the URL params to delete reply endpoint
type MessageDeleteReplyRequest struct {
	ChannelType ChannelType `json:"-" qs:"-"`
	ChannelURL  string      `json:"-" qs:"-"`
	MessageID   int64       `json:"message_id"`
}

// URL returns URL string to message delete reply endpoint
// DELETE https://api-{application_id}.sendbird.com/v3/{channel_type}/{channel_url}/messages/{message_id}
func (m MessageDeleteReplyRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%d", baseURL, m.ChannelType, m.ChannelURL, msgEndpoint, m.MessageID)
}

// MessageSearchRequest is the message search URL params & query values
type MessageSearchRequest struct {
	ChannelURL            string   `json:"-" qs:"channel_url"`
	UserID                string   `qs:"user_id"`
	Query                 string   `qs:"query"`
	CustomType            *string  `qs:"custom_type,omitempty"`
	Limit                 *int     `qs:"limit,omitempty"`
	ExactMatch            *bool    `qs:"exact_match,omitempty"`
	AdvancedQuery         *bool    `qs:"advanced_query,omitempty"`
	Synonym               *bool    `qs:"synonym,omitempty"`
	SortField             *string  `qs:"sort_field,omitempty"`
	Score                 *bool    `qs:"score,omitempty"`
	MessageTSFrom         *int64   `qs:"message_ts_from,omitempty"`
	MessageTSTo           *int64   `qs:"message_ts_to,omitempty"`
	After                 *string  `qs:"after,omitempty"`
	Before                *string  `qs:"before,omitempty"`
	Reverse               *bool    `qs:"reverse,omitempty"`
	Token                 *string  `qs:"token,omitempty"`
	TargetUserIDs         []string `qs:"target_user_ids,comma,omitempty"`
	WithChannelInResponse *bool    `qs:"with_channel_in_response,omitempty"`
}

// MessageSearchResponse is the response payload to search request endpoint
type MessageSearchResponse struct {
	TotalCount    *int      `json:"total_count,omitempty"`
	Results       []Message `json:"results,omitempty"`
	StartCursor   *string   `json:"start_cursor,omitempty"`
	EndCursor     *string   `json:"end_cursor,omitempty"`
	HasPrev       *bool     `json:"has_prev,omitempty"`
	HasNext       *bool     `json:"has_next,omitempty"`
	NextPageToken string    `json:"next,omitempty"`
	Error
}

// URL returns URL string for SendBird's search message endpoint
// GET https://api-{application_id}.sendbird.com/v3/search/messages
func (m MessageSearchRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/search/%s", baseURL, msgEndpoint)
	s, _ := encodeQS[MessageSearchRequest](base, m)
	return s
}

// MessageMigrateRequest is the request type to migrate messages to other channel
type MessageMigrateRequest struct {
	TargetChannelURL string    `json:"-"`
	Messages         []Message `json:"messages"`
	UpdateReadTS     bool      `json:"update_read_ts"`
	RewindReadTS     *bool     `json:"rewind_read_ts,omitempty"`
}

// MessageMigrateResponse will return empty body and/or Error if any
type MessageMigrateResponse EmptyResponse

// URL returns URL string for SendBird's migrate messages endpoint
// POST https://api-{application_id}.sendbird.com/v3/migration/{target_channel_url}
func (m MessageMigrateRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/migration/%s", baseURL, m.TargetChannelURL)
}

// MessageMarkAllAsDeliveredRequest mark all as delivered request payload
type MessageMarkAllAsDeliveredRequest struct {
	ChannelURL string `json:"-"`
	UserID     string `json:"user_id"`
}

// URL for MessageMarkAllAsDeliveredRequest request endpoint
// PUT https://api-{application_id}.sendbird.com/v3/group_channels/{channel_url}/messages/mark_as_delivered
func (m MessageMarkAllAsDeliveredRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/%s/mark_as_delivered", baseURL, groupChannelEndpoint, m.ChannelURL, msgEndpoint)
}

// MessageMarkAllAsDeliveredResponse is the response type to mark all as delivered request endpoint.
type MessageMarkAllAsDeliveredResponse struct {
	TS *int64 `json:"ts,omitempty"`
	Error
}

// MessageMarkAllAsReadRequest is the request payload and URL params to mark all messages as read on group channel.
type MessageMarkAllAsReadRequest struct {
	ChannelURL string `json:"-"`
	UserID     string `json:"user_id"`
	Timestamp  int64  `json:"timestamp,omitempty"`
}

// URL provides URL string to mark all messages in a group channel as read for a specific user endpoint.
// PUT /group_channels/{channel_url}/messages/mark_as_read
func (m MessageMarkAllAsReadRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/%s/mark_as_read", baseURL, groupChannelEndpoint, m.ChannelURL, msgEndpoint)
}

// MessageMarkAllAsReadResponse is the response payload to mark all messages as read on group channel endpoint
type MessageMarkAllAsReadResponse EmptyResponse

// MessageMarkAllAsReadAllJoinedGroupChannelRequest is the request payload and URL param to mark all of user's
// unread messages as read in specified group channels
type MessageMarkAllAsReadAllJoinedGroupChannelRequest struct {
	UserID      string   `json:"-"`
	ChannelURLs []string `json:"channel_urls"`
}

// URL returns URL string to mark all messages as read on group channel endpoint
// PUT /users/{user_id}/mark_as_read_all
func (m MessageMarkAllAsReadAllJoinedGroupChannelRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/mark_as_read_all", baseURL, m.UserID)
}

// MessageMarkAllAsReadAllJoinedGroupChannelResponse is the response type to marks all of a user’s unread messages
// as read in group channels.
type MessageMarkAllAsReadAllJoinedGroupChannelResponse EmptyResponse

// MessageViewUnreadCountRequest is the request parameters
type MessageViewUnreadCountRequest struct {
	ChannelURL string   `qs:"-"`
	UserIDs    []string `qs:"user_ids,comma,omitempty"`
}

// URL returns URL string for message view unread count request endpoint.
// GET https://api-{application_id}.sendbird.com/v3/group_channels/{channel_url}/messages/unread_count
func (m MessageViewUnreadCountRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/%s/unread_count", baseURL, groupChannelEndpoint, m.ChannelURL, msgEndpoint)
	s, _ := encodeQS[MessageViewUnreadCountRequest](base, m)
	return s
}

// MessageViewUnreadCountResponse is the response type / payload to message view unread count URL
type MessageViewUnreadCountResponse struct {
	Unread map[string]any `json:"unread,omitempty"`
	Error
}

// MessageTranslateRequest is the request payload to send to translate message endpoint
type MessageTranslateRequest struct {
	CommonMessageRequest
	MessageID   int64    `json:"-"`
	TargetLangs []string `json:"target_langs"`
}

// URL returns URL string for message translation request endpo.nt.
// POST https://api-{application_id}.sendbird.com/v3/{channel_type}/{channel_url}/messages/{message_id}/translation
func (m MessageTranslateRequest) URL(baseURL string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%d/translation", baseURL, m.ChannelType, m.ChannelURL, msgEndpoint, m.MessageID)
}

// MessageTranslateResponse is an alias for Message
type MessageTranslateResponse Message
