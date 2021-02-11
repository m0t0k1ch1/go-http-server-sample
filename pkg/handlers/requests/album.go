package requests

import (
	"encoding/json"
	"fmt"

	ov "github.com/go-ozzo/ozzo-validation/v4"
)

// PatchAlbumRequest represents a request to patch an album.
type PatchAlbumRequest struct {
	Title  json.RawMessage `json:"title"`
	Artist json.RawMessage `json:"artist"`
}

// IsEmpty returns whether the request is empty.
func (req PatchAlbumRequest) IsEmpty() bool {
	return !req.HasTitle() && !req.HasArtist()
}

// HasTitle returns whether the req.Title is specified.
func (req PatchAlbumRequest) HasTitle() bool {
	return len(req.Title) > 0
}

// HasArtist returns whether the req.Artist is specified.
func (req PatchAlbumRequest) HasArtist() bool {
	return len(req.Artist) > 0
}

// GetTitleWithValidation validates the req.Title and returns it.
func (req PatchAlbumRequest) GetTitleWithValidation() (string, error) {
	var title string
	if err := json.Unmarshal(req.Title, &title); err != nil {
		return "", fmt.Errorf("title: %w", err)
	}
	if err := ov.Validate(title, ov.Required); err != nil {
		return "", fmt.Errorf("title: %w", err)
	}

	return title, nil
}

// GetArtistWithValidation validates the req.Artist and returns it.
func (req PatchAlbumRequest) GetArtistWithValidation() (string, error) {
	var artist string
	if err := json.Unmarshal(req.Artist, &artist); err != nil {
		return "", fmt.Errorf("artist: %w", err)
	}
	if err := ov.Validate(artist, ov.Required); err != nil {
		return "", fmt.Errorf("artist: %w", err)
	}

	return artist, nil
}
