package dto

// Statistic is the object from SendBird
type Statistic struct {
	DateRange                  string  `json:"date_range"`
	TotalAnnouncementCount     int     `json:"total_announcement_count"`
	CanceledAnnouncementCount  int     `json:"canceled_announcement_count"`
	StoppedAnnouncementCount   int     `json:"stopped_announcement_count"`
	CompletedAnnouncementCount int     `json:"completed_announcement_count"`
	TargetChannelCount         int     `json:"target_channel_count"`
	TargetUserCount            int     `json:"target_user_count"`
	SentChannelCount           int     `json:"sent_channel_count"`
	SentUserCount              int     `json:"sent_user_count"`
	OpenCount                  int     `json:"open_count"`
	OpenRate                   float64 `json:"open_rate"`
}
