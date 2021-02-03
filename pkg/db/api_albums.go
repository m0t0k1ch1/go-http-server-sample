package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/m0t0k1ch1/go-http-server-sample/pkg/models"
)

// CreateAlbum creates a new album.
func CreateAlbum(ctx context.Context, exe Executer, album *models.Album) error {
	if _, err := exe.ExecContext(ctx, `
		INSERT INTO albums (ean, title, artist)
		VALUES (?, ?, ?)
	`, album.EAN, album.Title, album.Artist); err != nil {
		return fmt.Errorf("failed to create an album: %w", err)
	}

	return nil
}

// FetchAlbums fetches all albums.
func FetchAlbums(ctx context.Context, exe Executer) ([]*models.Album, error) {
	rows, err := exe.QueryContext(ctx, `
		SELECT *
		FROM albums
		ORDER BY ean
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all albums: %w", err)
	}

	return scanAlbums(rows)
}

// FetchAlbum fetches an album by specifying EAN.
func FetchAlbum(ctx context.Context, exe Executer, ean string) (*models.Album, error) {
	album, err := scanAlbum(exe.QueryRowContext(ctx, `
		SELECT *
		FROM albums
		WHERE ean = ?
	`, ean))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch an album by specifying EAN: %w", err)
	}

	return album, nil
}

// FetchAlbumForUpdate fetches an album for update by specifying EAN.
func FetchAlbumForUpdate(ctx context.Context, exe Executer, ean string) (*models.Album, error) {
	album, err := scanAlbum(exe.QueryRowContext(ctx, `
		SELECT *
		FROM albums
		WHERE ean = ?
		FOR UPDATE
	`, ean))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch an album for update by specifying EAN: %w", err)
	}

	return album, nil
}

// DeleteAlbum deletes an album by specifying EAN.
func DeleteAlbum(ctx context.Context, exe Executer, ean string) error {
	if _, err := exe.ExecContext(ctx, `
		DELETE FROM albums
		WHERE ean = ?
	`, ean); err != nil {
		return fmt.Errorf("failed to delete an album by specifying EAN: %w", err)
	}

	return nil
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
