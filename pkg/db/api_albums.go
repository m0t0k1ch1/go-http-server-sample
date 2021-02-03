package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/m0t0k1ch1/go-http-server-sample/pkg/models"
)

// FetchAlbums fetches all albums.
func FetchAlbums(ctx context.Context, exe Executer) ([]*models.Album, error) {
	rows, err := exe.QueryContext(ctx, `
		SELECT *
		FROM albums
		ORDER BY ean
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the albums: %w", err)
	}

	return scanAlbums(rows)
}

func scanAlbums(rows *sql.Rows) ([]*models.Album, error) {
	defer rows.Close()

	albums := []*models.Album{}
	for rows.Next() {
		album, err := scanAlbum(rows)
		if err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan album rows: %w", err)
	}

	return albums, nil
}

func scanAlbum(s Scanner) (*models.Album, error) {
	var album models.Album

	err := s.Scan(
		&album.EAN,
		&album.Title,
		&album.Artist,
	)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, nil

	case err != nil:
		return nil, fmt.Errorf("failed to scan an album row: %w", err)

	default:
		return &album, nil
	}
}
