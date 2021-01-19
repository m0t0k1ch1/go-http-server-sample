package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/m0t0k1ch1/go-http-server-sample/pkg/rdb"
)

// Album represents a row of the albums table.
type Album struct {
	EAN    string `json:"ean"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

// FetchAlbums fetches all rows of the albums table.
func FetchAlbums(ctx context.Context, exe rdb.Executer) ([]*Album, error) {
	rows, err := exe.QueryContext(ctx, `
		SELECT *
		FROM albums
		ORDER BY ean
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all rows of the albums table: %w", err)
	}

	return scanAlbums(rows)
}

func scanAlbums(rows *sql.Rows) ([]*Album, error) {
	defer rows.Close()

	albums := []*Album{}
	for rows.Next() {
		album, err := scanAlbum(rows)
		if err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan rows of the albums table: %w", err)
	}

	return albums, nil
}

func scanAlbum(s rdb.Scanner) (*Album, error) {
	var album Album

	err := s.Scan(
		&album.EAN,
		&album.Title,
		&album.Artist,
	)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, nil

	case err != nil:
		return nil, fmt.Errorf("failed to scan a row of the albums table: %w", err)

	default:
		return &album, nil
	}
}
