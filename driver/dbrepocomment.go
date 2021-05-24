package driver

import (
	"context"
	"database/sql"
	"thelight/models"
)

//DBCommentGetAll
func DBCommentGetAll(db *sql.DB, payload *models.CommentFromClient) ([]models.Comment, error) {
	ctx := context.Background()

	var comments []models.Comment

	rows, err := db.QueryContext(
		ctx,
		"SELECT ID,Name,Text FROM comments WHERE ARTICLE_REF=$1",
		payload.ID,
	)
	if err != nil {
		return comments, err
	}

	for rows.Next() {
		var tmp models.Comment
		rows.Scan(&tmp.ID, &tmp.Name, &tmp.Text)
		comments = append(comments, tmp)
	}

	return comments, nil
}

//DBCommentInsert
func DBCommentInsert(db *sql.DB, payload *models.CommentFromClient) (int64, error) {
	ctx := context.Background()

	var insertedID int64

	err := db.QueryRowContext(
		ctx,
		"INSERT INTO comments (Name,Text,ARTICLE_REF) VALUES ($1,$2,$3) RETURNING ID",
		payload.CommentFromClient.Name, payload.CommentFromClient.Text, payload.CommentFromClient.ID,
	).Scan(&insertedID)
	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

//DBCommentDelete
func DBCommentDelete(db *sql.DB, payload *models.CommentFromClient) error {
	ctx := context.Background()

	_, err := db.ExecContext(ctx, "DELETE FROM comments WHERE id=$1", payload.ID)
	if err != nil {
		return err
	}

	return nil
}
