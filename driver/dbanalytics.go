package driver

import (
	"database/sql"
	"thelight/models"
)

//DBArticleHitCounter will increment counter every article visit
func DBArticleHitCounter(db *sql.DB, payload *models.AnalyticPayload) (int64, error) {
	// ctx := context.Background()

	var insertedID int64

	// err := db.QueryRowContext(
	// 	ctx,
	// 	"UPSERT analytics SET hit=hit+1 WHERE ARTICLE_REF=$1 RETURNING ID",
	// 	payload.ID,
	// ).Scan(&insertedID)
	// if err != nil {
	// 	return 0, err
	// }

	return insertedID, nil
}
