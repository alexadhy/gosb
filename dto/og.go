package dto

// OpenGraph provides open graph objects
type OpenGraph struct {
	Title           *string       `json:"og:title,omitempty"`
	Type            *string       `json:"og:type,omitempty"`
	Image           *string       `json:"og:image,omitempty"`
	URL             *string       `json:"og:url,omitempty"`
	Audio           *string       `json:"og:audio,omitempty"`
	Description     *string       `json:"og:description,omitempty"`
	Determiner      *OGDeterminer `json:"og:determiner,omitempty"`
	Locale          *string       `json:"og:locale,omitempty"`
	LocaleAlternate []string      `json:"og:locale:alternate,omitempty"`
	SiteName        *string       `json:"og:site_name,omitempty"`
	Video           *string       `json:"og:video,omitempty"`
}
