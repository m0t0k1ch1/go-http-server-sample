package models

import (
	ov "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
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
		ov.Field(&album.EAN, ov.Required, ov.Length(13, 13), is.Digit),
		ov.Field(&album.Title, ov.Required),
		ov.Field(&album.Artist, ov.Required),
	)
}
