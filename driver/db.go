package driver

import (
	"database/sql"
	"os"
	"time"

	"github.com/lib/pq"
)

//InitiateDB return a pointer to db conn
func InitiateDB() (*sql.DB, error) {

	pgURL, err := pq.ParseURL(os.Getenv("DB_URL"))

	if err != nil {
		return nil, err
	}

	db, _ := sql.Open("postgres", pgURL)
	for {
		err := db.Ping()
		if err == nil {
			return db, nil
		}
		time.Sleep(5000 * time.Millisecond)
	}

}
