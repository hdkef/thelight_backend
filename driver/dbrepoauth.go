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

//DBAuthRegister store new user in user database
func DBAuthRegister(db *sql.DB, payload *models.AuthFromClient) (int64, error) {
	ctx := context.Background()

	var insertedID int64

	err := db.QueryRowContext(
		ctx,
		"INSERT INTO users (Name,Pass,AvatarURL,Bio, Email) VALUES ($1,$2,$3,$4,$5) RETURNING ID",
		payload.Name, payload.Pass, NEWUSR_AVATARURL, NEWUSR_BIO, payload.Email,
	).Scan(&insertedID)

	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

//DBAuthLogin return hashed pass of user, and give writerInfo (for creating claims later)
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

//DBAuthSettings will update user's settings, like bio, name etc
func DBAuthSettings(db *sql.DB, payload *models.Settings) error {
	ctx := context.Background()

	tx, err := db.BeginTx(ctx, nil)
	defer tx.Commit()
	if err != nil {
		return err
	}

	if payload.Bio != "" {
		_, err = tx.ExecContext(ctx, "UPDATE users SET Bio=$1 WHERE ID=$2", payload.Bio, payload.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if payload.Name != "" {
		_, err = tx.ExecContext(ctx, "UPDATE users SET Name=$1 WHERE ID=$2", payload.Name, payload.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if payload.AvatarURL != "" {
		_, err = tx.ExecContext(ctx, "UPDATE users SET AvatarURL=$1 WHERE ID=$2", payload.AvatarURL, payload.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return nil
}
