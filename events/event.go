package events

// WebHookEvents is the common interface for all events
type WebHookEvents interface {
	EventCategory() string
}
