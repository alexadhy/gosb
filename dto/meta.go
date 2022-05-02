package dto

import "fmt"

// Metadata is returned on any object that has metadata
type Metadata map[string]any

// Metacounter is returned on any object that has metacounter
type Metacounter map[string]any

type MetaValue interface {
	*string | *int
}

// MetaInputRequest contains Metacounter and Metadata, this structure is used as payload to Input both Metadata / Metacounter.
type MetaInputRequest[T MetaValue] struct {
	ChannelType string       `json:"-" qs:"-"`
	ChannelURL  string       `json:"-" qs:"-"`
	Metacounter *Metacounter `json:"metacounter,omitempty"`
	Metadata    *Metadata    `json:"metadata,omitempty"`
	MetaValue   T            `json:"value,omitempty"`
	MetaMode    *MetaMode    `json:"mode,omitempty"`
	Upsert      *bool        `json:"upsert,omitempty"`
}

// URL Inputs URL string for both Metacounter and Metadata endpoints
func (m MetaInputRequest[T]) URL(baseURL string, key string) string {
	base := fmt.Sprintf("%s/%s/%s", baseURL, m.ChannelType, m.ChannelURL)
	baseData := fmt.Sprintf("%s/metadata", base)
	baseCounter := fmt.Sprintf("%s/metacounter", base)

	if key != "" {
		if m.Metadata != nil {
			return fmt.Sprintf("%s/%s", baseData, key)
		}
		if m.Metacounter != nil {
			return fmt.Sprintf("%s/%s", baseCounter, key)
		}
	}
	if m.Metadata != nil {
		return baseData
	}
	return baseCounter
}
