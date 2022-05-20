package events

import "github.com/alexadhy/gosb/dto"

type GroupChannelCommon struct {
	Category string                    `json:"category"`
	Members  []dto.UserResponse        `json:"members,omitempty"`
	Channel  *dto.GroupChannelResponse `json:"channel"`
	AppID    string                    `json:"app_id"`
}

// GroupChannelCreate is invoked when a group channel with the invited members
// is created through the Chat API or SDKs.
// Once this webhook is delivered, the group_channel:join webhook is invoked.
type GroupChannelCreate struct {
	GroupChannelCommon
	CreatedAt int64             `json:"created_at"`
	Inviter   *dto.UserResponse `json:"inviter"`
}

func (g GroupChannelCreate) EventCategory() string {
	return g.GroupChannelCommon.Category
}

// GroupChannelChange webhook event is invoked and delivered when one of the following group channel properties
// has been changed: name and cover_url.
// The changes property in the payload contains the old and new values of the changed channel properties.
type GroupChannelChange struct {
	GroupChannelCommon
	ChangedAt int64             `json:"changed_at"`
	Changes   []Change          `json:"changes"`
	Inviter   *dto.UserResponse `json:"inviter"`
}

func (g GroupChannelChange) EventCategory() string {
	return g.GroupChannelCommon.Category
}

// GroupChannelRemove webhook event is invoked when a group channel is removed.
type GroupChannelRemove struct {
	GroupChannelCommon
	RemovedAt int64 `json:"removed_at"`
}

func (g GroupChannelRemove) EventCategory() string {
	return g.GroupChannelCommon.Category
}

// GroupChannelInvite webhook event is invoked when a user invites another user.
// Note: When the members are already invited, the group_channel:invite webhook isn't invoked.
// By default, the members property is excluded from the payload.
// If you want to include the members,
// select the Group channel member on the Sendbird Dashboard under Settings > Chat > Webhooks > Optional information.
type GroupChannelInvite struct {
	GroupChannelCommon
	InvitedAt int64              `json:"invited_at"`
	Inviter   *dto.UserResponse  `json:"inviter"`
	Invitees  []dto.UserResponse `json:"invitees"`
}

func (g GroupChannelInvite) EventCategory() string {
	return g.GroupChannelCommon.Category
}

// GroupChannelDeclineInvite webhook event is invoked when a user declines an invitation.
// The users property indicates the members who have declined invitations and are in pending status
type GroupChannelDeclineInvite struct {
	GroupChannelCommon
	DeclinedInviteAt int64              `json:"declined_invite_at"`
	Users            []dto.UserResponse `json:"users"`
}

func (g GroupChannelDeclineInvite) EventCategory() string {
	return g.GroupChannelCommon.Category
}

// GroupChannelJoin webhook event is invoked when a user joins a group channel.
// The users property indicates the members who have joined channels while the auto-accept is on or accepted invitations
type GroupChannelJoin struct {
	GroupChannelCommon
	JoinedAt int64              `json:"joined_at"`
	Users    []dto.UserResponse `json:"users"`
}

func (g GroupChannelJoin) EventCategory() string {
	return g.GroupChannelCommon.Category
}

// GroupChannelLeave webhook event is invoked when a user leaves a group channel.
// The users property indicates those who have left a group channel,
// whereas the members indicates those who still remain in the channel.
// The channel_unread_message_count is a value for a user's unread message count of the channel when the user left.
type GroupChannelLeave struct {
	GroupChannelCommon
	LeftAt int64              `json:"left_at"`
	Users  []dto.UserResponse `json:"users"`
}

func (g GroupChannelLeave) EventCategory() string {
	return g.GroupChannelCommon.Category
}

// GroupChannelMessageSend webhook event is invoked when a message is sent within a group channel
type GroupChannelMessageSend struct {
	GroupChannelCommon
	Sender         *dto.UserResponse  `json:"sender"`
	Silent         bool               `json:"silent"`
	SenderIPAddr   string             `json:"sender_ip_addr"`
	CustomType     string             `json:"custom_type"`
	MentionType    string             `json:"mention_type"`
	MentionedUsers []dto.UserResponse `json:"mentioned_users"`
	Members        []dto.UserResponse `json:"members,omitempty"`
	MessageType    string             `json:"type"`
	Payload        *dto.Message       `json:"message"`
	Sdk            string             `json:"sdk"`
}

func (g GroupChannelMessageSend) EventCategory() string {
	return g.GroupChannelCommon.Category
}

// GroupChannelMessageRead webhook event is fired when user has no more unread messages in a group channel
type GroupChannelMessageRead struct {
	GroupChannelCommon
	Members     []dto.UserResponse `json:"members,omitempty"`
	ReadUpdates []dto.UserResponse `json:"read_updates"`
}

func (g GroupChannelMessageRead) EventCategory() string {
	return g.GroupChannelCommon.Category
}

// GroupChannelMessageUpdate webhook event is invoked when a message is sent within a group channel
type GroupChannelMessageUpdate struct {
	GroupChannelCommon
	Changes     []Change           `json:"changes"`
	CustomType  string             `json:"custom_type"`
	Members     []dto.UserResponse `json:"members,omitempty"`
	MessageType string             `json:"type"`
	Payload     *dto.Message       `json:"message"`
	Sdk         string             `json:"sdk"`
}

func (g GroupChannelMessageUpdate) EventCategory() string {
	return g.GroupChannelCommon.Category
}

// GroupChannelMessageDelete webhook event is invoked when a message is deleted from group channel
type GroupChannelMessageDelete struct {
	GroupChannelCommon
	Sender      *dto.UserResponse  `json:"sender"`
	Members     []dto.UserResponse `json:"members,omitempty"`
	CustomType  string             `json:"custom_type"`
	MessageType string             `json:"type"`
	Payload     *dto.Message       `json:"payload"`
}

func (g GroupChannelMessageDelete) EventCategory() string {
	return g.GroupChannelCommon.Category
}

// GroupChannelFreezeUnfreeze webhook event is invoked when a channel operator freezes / unfreezes group channel.
type GroupChannelFreezeUnfreeze struct {
	GroupChannelCommon
	Members []dto.UserResponse `json:"members"`
}

func (g GroupChannelFreezeUnfreeze) EventCategory() string {
	return g.GroupChannelCommon.Category
}

// GroupChannelReactionAdd webhook event is invoked when a user adds reactions to a message
type GroupChannelReactionAdd struct {
	GroupChannelCommon
	Reaction string            `json:"reaction"`
	TS       int64             `json:"ts"`
	User     *dto.UserResponse `json:"user"`
	Message  *dto.Message      `json:"message"`
}

func (g GroupChannelReactionAdd) EventCategory() string {
	return g.GroupChannelCommon.Category
}

// GroupChannelReactionDelete webhook event is invoked when a user deletes reactions from a message.
type GroupChannelReactionDelete struct {
	GroupChannelCommon
	Reaction string            `json:"reaction"`
	TS       int64             `json:"ts"`
	User     *dto.UserResponse `json:"user"`
	Message  *dto.Message      `json:"message"`
}

func (g GroupChannelReactionDelete) EventCategory() string {
	return g.GroupChannelCommon.Category
}
