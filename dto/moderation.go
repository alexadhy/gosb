package dto

type Moderation struct{}

type ModerationUserResource struct {
	UserID     string    `json:"user_id"`
	Nickname   string    `json:"nickname"`
	ProfileURL string    `json:"profile_url"`
	IsOnline   bool      `json:"is_online"`
	LastSeenAt int       `json:"last_seen_at"`
	Metadata   *Metadata `json:"metadata"`
}
