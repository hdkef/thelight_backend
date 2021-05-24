package driver

import (
	"context"
	"database/sql"
	"thelight/models"
)

//DBMediaGetAll
func DBMediaGetAll(db *sql.DB, payload *models.MediaPayload) ([]models.Media, error) {
	ctx := context.Background()

	var limit int64 = 6
	offset := (payload.Page - 1) * limit

	var medias []models.Media

	rows, err := db.QueryContext(
		ctx,
		"SELECT ID, ImageURL, USER_REF from medias WHERE ID=$1 OFFSET $2 LIMIT $3",
		payload.ID, offset, limit,
	)
	if err != nil {
		return []models.Media{}, err
	}

	for rows.Next() {
		var tmp models.Media
		rows.Scan(&tmp.ID, &tmp.ImageURL, &tmp.UserRef)
		medias = append(medias, tmp)
	}

	return medias, nil
}

//DBMediaInsert
func DBMediaInsert(db *sql.DB, imgurl string, claims *models.WriterInfo) (int64, error) {
	ctx := context.Background()

	var insertedID int64

	err := db.QueryRowContext(
		ctx,
		"INSERT INTO medias (ImageURL,User_Ref) VALUES ($1,$2) RETURNING ID",
		imgurl, claims.ID,
	).Scan(&insertedID)
	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

//DBMediaDelete
func DBMediaDelete(db *sql.DB, payload *models.MediaPayload) error {
	ctx := context.Background()

	_, err := db.ExecContext(ctx, "DELETE FROM medias WHERE id=$1", payload.ID)
	if err != nil {
		return err
	}

	return nil
}
