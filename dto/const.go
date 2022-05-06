package dto

const (
	userEndpoint              = "users"
	reportEndpoint            = "report"
	openChannelEndpoint       = "open_channels"
	groupChannelEndpoint      = "group_channels"
	customChannelEndpoint     = "applications/settings_by_channel_custom_type"
	defaultInvitationEndpoint = "applications/default_channel_invitation_preference"
	settingsEndpoint          = "applications/settings"
	msgEndpoint               = "messages"
	announcementEndpoint      = "announcements"
)

// Order is the order url query value
type Order string

const (
	ORDER_CHRONOLOGICAL         = Order("chronological")
	ORDER_LATEST_LAST_MSG       = Order("latest_last_message")
	ORDER_CHANNEL_ALPHABETICAL  = Order("channel_name_alphabetical")
	ORDER_METADATA_ALPHABETICAL = Order("metadata_value_alphabetical")
)

// LogicalOp used to specify logical operator applied to url query value
type LogicalOp string

const (
	LOGICAL_OP_AND = LogicalOp("AND")
	LOGICAL_OP_OR  = LogicalOp("OR")
)

// DistinctMode used to specify distinct_mode url query value
type DistinctMode string

const (
	DISTINCT_ALL      = DistinctMode("all")
	DISTINCT_DISTINCT = DistinctMode("distinct")
	DISTINCT_NONE     = DistinctMode("nondistinct")
)

// PublicMode specifies public_mode url query value
type PublicMode string

const (
	PUBLIC_MODE_ALL     = PublicMode("all")
	PUBLIC_MODE_PUBLIC  = PublicMode("public")
	PUBLIC_MODE_PRIVATE = PublicMode("private")
)

// SuperMode specifies super_mode url query value
type SuperMode string

const (
	SUPER_MODE_ALL   = SuperMode("all")
	SUPER_MODE_SUPER = SuperMode("super")
	SUPER_MODE_NON   = SuperMode("nonsuper")
)

// HiddenMode specifies hidden_mode url query value
type HiddenMode string

const (
	HIDDEN_MODE_UNHIDDEN_ONLY              = UnreadFilter("unhidden_only")
	HIDDEN_MODE_HIDDEN_ONLY                = UnreadFilter("hidden_only")
	HIDDEN_MODE_HIDDEN_ALLOW_AUTO_UNHIDE   = UnreadFilter("hidden_allow_auto_unhide")
	HIDDEN_MODE_HIDDEN_PREVENT_AUTO_UNHIDE = UnreadFilter("hidden_prevent_auto_unhide")
)

// MemberStateFilter specifies member_state_filter url query value
type MemberStateFilter string

const (
	MEMBER_STATE_FILTER_ALL                   = MemberStateFilter("all")
	MEMBER_STATE_FILTER_INVITED_ONLY          = MemberStateFilter("invited_only")
	MEMBER_STATE_FILTER_JOINED_ONLY           = MemberStateFilter("joined_only")
	MEMBER_STATE_FILTER_INVITED_BY_FRIEND     = MemberStateFilter("invited_by_friend")
	MEMBER_STATE_FILTER_INVITED_BY_NON_FRIEND = MemberStateFilter("invited_by_non_friend")
)

// UnreadFilter specifies unread_filter url query value
type UnreadFilter string

const (
	UNREAD_FILTER_ALL            = UnreadFilter("all")
	UNREAD_FILTER_UNREAD_MESSAGE = UnreadFilter("unread_message")
)

// OGDeterminer is OpenGraph's determiner metadata
type OGDeterminer string

const (
	DETERMINER_A     = OGDeterminer("a")
	DETERMINER_AN    = OGDeterminer("an")
	DETERMINER_EMPTY = OGDeterminer("")
	DETERMINER_AUTO  = OGDeterminer("auto")
)

type InvitationStatus string

const (
	INVITATION_STATUS_JOINED                = InvitationStatus("joined")
	INVITIATION_STATUS_INVITED_BY_FRIEND    = InvitationStatus("invited_by_friend")
	INVITATION_STATUS_INVITED_BY_NON_FRIEND = InvitationStatus("invited_by_non_friend")
)

// OperatorFilter is a field to restrict search scope in operator
type OperatorFilter string

const (
	OPERATOR_FILTER_ALL      = OperatorFilter("all")
	OPERATOR_FILTER_OPERATOR = OperatorFilter("operator")
	OPERATOR_FILTER_NON      = OperatorFilter("nonoperator")
)

// MutedMemberFilter is the field to filter muted member
type MutedMemberFilter string

const (
	MUTED_MEMBER_FILTER_ALL     = MutedMemberFilter("all")
	MUTED_MEMBER_FILTER_MUTED   = MutedMemberFilter("muted")
	MUTED_MEMBER_FILTER_UNMUTED = MutedMemberFilter("unmuted")
)

// ChannelType is the field to restrict channel type
type ChannelType string

const (
	CHANNEL_TYPE_OPEN  = ChannelType("open_channels")
	CHANNEL_TYPE_GROUP = ChannelType("group_channels")
)

// MetaMode is the metacounter value mode
type MetaMode string

const (
	META_MODE_INC = MetaMode("increase")
	META_MODE_DEC = MetaMode("decrease")
	META_MODE_SET = MetaMode("set")
)

// ModerationAction is the moderation action field of profanity filter
type ModerationAction string

const (
	MODERATION_ACTION_MUTE = ModerationAction("mute")
	MODERATION_ACTION_KICK = ModerationAction("kick")
	MODERATION_ACTION_BAN  = ModerationAction("ban")
)

// AnnouncementAction is the actions that can be taken to modify announcement status
type AnnouncementAction string

const (
	ANNOUNCEMENT_ACTION_PAUSE  = AnnouncementAction("pause")
	ANNOUNCEMENT_ACTION_STOP   = AnnouncementAction("stop")
	ANNOUNCEMENT_ACTION_RESUME = AnnouncementAction("resume")
	ANNOUNCEMENT_ACTION_CANCEL = AnnouncementAction("cancel")
)
