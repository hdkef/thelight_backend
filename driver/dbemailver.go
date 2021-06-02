package driver

import (
	"context"
	"database/sql"
	"thelight/models"
)

//StoreUsersReg will store user's data temporarily before
func StoreUsersReg(db *sql.DB, payload *models.AuthFromClient, code string) (int64, error) {
	ctx := context.Background()

	var insertedID int64

	err := db.QueryRowContext(
		ctx,
		"INSERT INTO users_reg (Name,Pass,Email,Code) VALUES ($1,$2,$3,$4) RETURNING ID",
		payload.Name, payload.Pass, payload.Email, code,
	).Scan(&insertedID)
	if err != nil {
		return insertedID, nil
	}

	return insertedID, nil
}

//SelectUsersReg will select all column filtered by Email in users_reg table
func SelectUsersReg(db *sql.DB, Email string) (models.AuthFromClient, error) {
	ctx := context.Background()

	var users models.AuthFromClient

	err := db.QueryRowContext(ctx, "SELECT Name,Pass,Email,Code FROM users_reg WHERE Email=$1", Email).Scan(
		&users.Name, &users.Pass, &users.Email, &users.Code,
	)
	if err != nil {
		return models.AuthFromClient{}, err
	}
	return users, nil
}

//DeleteUsersReg will delete one item in delete users reg
func DeleteUsersReg(db *sql.DB, ID int64) error {
	ctx := context.Background()

	_, err := db.ExecContext(ctx, "DELETE FROM users_reg WHERE ID=$1", ID)
	if err != nil {
		return err
	}
	return nil
}
