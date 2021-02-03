package models

// Album represents a row in albums table.
type Album struct {
	EAN    string `json:"ean"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
}
