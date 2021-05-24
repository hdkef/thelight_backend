package driver

import (
	"context"
	"database/sql"
	"thelight/models"
)

const (
	NEWUSR_AVATARURL = "https://pbs.twimg.com/profile_images/1363210545118150659/Uo-XiGtv_400x400.jpg"
	NEWUSR_BIO       = "Hi! i am a writer"
)

//DBAuthRegister
func DBAuthRegister(db *sql.DB, payload *models.AuthFromClient) (int64, error) {
	ctx := context.Background()

	var insertedID int64

	err := db.QueryRowContext(
		ctx,
		"INSERT INTO users (Name,Pass,AvatarURL,Bio) VALUES ($1,$2,$3,$4) RETURNING ID",
		payload.Name, payload.Pass, NEWUSR_AVATARURL, NEWUSR_BIO,
	).Scan(&insertedID)

	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

//DBAuthLogin
func DBAuthLogin(db *sql.DB, payload *models.AuthFromClient) (string, models.WriterInfo, error) {
	ctx := context.Background()

	var writerinfo models.WriterInfo
	var hpass string

	err := db.QueryRowContext(
		ctx,
		"SELECT ID, Name, AvatarURL, Bio, Pass FROM users WHERE Name=$1",
		payload.Name,
	).Scan(&writerinfo.ID, &writerinfo.Name, &writerinfo.AvatarURL, &writerinfo.Bio, &hpass)
	if err != nil {
		return "", models.WriterInfo{}, err
	}

	return hpass, writerinfo, nil
}

//DBAuthSettings
func DBAuthSettings(db *sql.DB) (models.AuthFromClient, error) {
	return models.AuthFromClient{}, nil
}
