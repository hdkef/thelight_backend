package driver

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//InitiateDB return a pointer to db conn
func InitiateDB() *gorm.DB {

	host := os.Getenv("PGHOST")
	user := os.Getenv("PGUSER")
	pass := os.Getenv("PGPASS")
	dbname := os.Getenv("PGDBNAME")
	port := os.Getenv("PGPORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable Timezone=Asia/Makassar", host, user, pass, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	db.AutoMigrate(&Article{}, &User{}, &Comment{}, &Media{}, &Draft{})

	return db

}
