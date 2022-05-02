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
}

// Message may contain Text, File, or Admin Message
type Message struct {
	*TextMessage
	*FileMessage
	*AdminMessage
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
	SortedMetaArray []map[string]any `json:"sorted_metaarray,omitempty"`
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
	Error    *Error    `json:"error,omitempty"`
}

// MessageGetRequest is the request parameter & url values to get message endpoint
type MessageGetRequest struct {
	CommonMessageRequest
	MessageID                string `qs:"-"`
	IncludeParentMessageInfo *bool  `qs:"include_parent_message_info,omitempty"`
	IncludeThreadInfo        *bool  `qs:"include_thread_info,omitempty"`
	IncludeReactions         *bool  `qs:"include_reactions,omitempty"`
	WithSortedMetaArray      *bool  `qs:"with_sorted_meta_array,omitempty"`
}

// URL returns URL string for get message endpoint
// GET https://api-{application_id}.sendbird.com/v3/{channel_type}/{channel_url}/messages/{message_id}
func (m MessageGetRequest) URL(baseURL string) string {
	base := fmt.Sprintf("%s/%s/%s/%s/%s", baseURL, m.ChannelType, m.ChannelURL, msgEndpoint, m.MessageID)
	s, err := encodeQS[MessageGetRequest](base, m)
	if err != nil {
		return ""
	}
	return s
}

// MessageGetResponse is an alias for Message
type MessageGetResponse Message

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
	return fmt.Sprintf("%s/%s/%s/%s", baseURL, m.ChannelType, m.ChannelURL, msgEndpoint)
}

// MessageSentResponse is the response type to MessageSendRequest request
type MessageSentResponse Message

type 
