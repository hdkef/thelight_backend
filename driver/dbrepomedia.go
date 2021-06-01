package driver

import (
	"context"
	"database/sql"
	"errors"
	"thelight/models"
)

//DBMediaGetAll will return all media paginated by LastID
func DBMediaGetAll(payload *models.MediaPayload) ([]models.Media, error) {
	ctx := context.Background()

	var limit int64 = 6

	var medias []models.Media

	rows, err := payload.DB.QueryContext(
		ctx,
		"SELECT medias.ID, medias.ImageURL, medias.USER_REF from medias WHERE USER_REF=$1 AND medias.ID > $2 ORDER BY medias.ID ASC LIMIT $3",
		payload.ID, payload.LastID, limit,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var tmp models.Media
		err = rows.Scan(&tmp.ID, &tmp.ImageURL, &tmp.UserRef) //THIS IS WEIRD, IF NO RESULT WILL NOT RETURN ERROR
		if err != nil {
			return nil, err
		}
		medias = append(medias, tmp)
	}

	if len(medias) == 0 {
		return nil, errors.New("NO RESULT")
	}

	return medias, nil
}

//DBMediaInsert will insert media's ImageURL to media database
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

//DBMediaDelete will delete media imageURL in media database
func DBMediaDelete(db *sql.DB, payload *models.MediaPayload, claims *models.WriterInfo) error {
	ctx := context.Background()

	_, err := db.ExecContext(ctx, "DELETE FROM MEDIAS WHERE MEDIAS.ID=$1 AND USER_REF=$2", payload.ID, claims.ID)
	if err != nil {
		return err
	}

	return nil
}

//DBMediaGetImageURL get imagedir of an image from media database
func DBMediaGetImageURL(db *sql.DB, payload *models.MediaPayload, claims *models.WriterInfo) (string, error) {
	ctx := context.Background()

	var imagedir string

	err := db.QueryRowContext(
		ctx,
		"SELECT medias.IMAGEURL FROM medias WHERE medias.ID=$1 AND USER_REF=$2",
		payload.ID, claims.ID,
	).Scan(&imagedir)
	if err != nil {
		return "", nil
	}
	return imagedir, nil
}
