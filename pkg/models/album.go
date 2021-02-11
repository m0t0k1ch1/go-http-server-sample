package models

import (
	ov "github.com/go-ozzo/ozzo-validation/v4"

	v "github.com/m0t0k1ch1/go-http-server-sample/pkg/validation"
)

// Album represents a row in albums table.
type Album struct {
	EAN    string `json:"ean"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

// Validate validates the album.
func (album Album) Validate() error {
	return ov.ValidateStruct(&album,
		ov.Field(&album.EAN, ov.Required, ov.By(v.ValidateEAN)),
		ov.Field(&album.Title, ov.Required),
		ov.Field(&album.Artist, ov.Required),
	)
}
