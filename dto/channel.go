package dto

import (
	"encoding/json"
	"errors"
)

// Channel is the common Channel object
type Channel struct {
	*OpenChannelResponse
	*GroupChannelResponse
}

// UnmarshalJSON unmarshal Channel type to its member type (OpenChannelResponse / GroupChannelResponse)
func (c Channel) UnmarshalJSON(data []byte) error {
	msi := map[string]any{}
	if err := json.Unmarshal(data, &msi); err != nil {
		return err
	}
	if _, ok := msi["is_public"]; ok {
		g := GroupChannelResponse{}
		return json.Unmarshal(data, &g)
	}
	oc := OpenChannelResponse{}
	return json.Unmarshal(data, &oc)
}

// MarshalJSON marshal Message type to JSON byte
func (c Channel) MarshalJSON() ([]byte, error) {
	if c.OpenChannelResponse != nil {
		return json.Marshal(c.OpenChannelResponse)
	}
	if c.GroupChannelResponse != nil {
		return json.Marshal(c.GroupChannelResponse)
	}
	return nil, errors.New("unknown message type")
}
