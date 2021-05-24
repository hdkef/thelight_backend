package driver

import (
	"context"
	"database/sql"
	"thelight/models"
)

//DBMediaGetAll
func DBMediaGetAll(db *sql.DB) (models.MediaPayload, error) {
	return models.MediaPayload{}, nil
}

//DBMediaInsert
func DBMediaInsert(db *sql.DB, imgurl string, claims *models.WriterInfo) (int64, error) {
	ctx := context.Background()

	var insertedID int64

	err := db.QueryRowContext(
		ctx,
		"INSERT INTO medias (ImageURL,User_Ref) RETURNING ID",
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
