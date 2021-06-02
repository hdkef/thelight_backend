package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/smtp"
	"os"
	"thelight/driver"
	"thelight/models"
	"thelight/utils"
	"time"
)

//HandleEmailVerReq will handle email verification request
func (x *AuthHandler) HandleEmailVerReq() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		var payload models.AuthFromClient
		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		code := fmt.Sprint(rand.Int31())

		insertedID, err := storeToUsersReg(x.db, &payload, code)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		err = sendEmailVer(&payload, code)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		go waitAndDelete(x.db, insertedID)
	}
}

//sendEmailVer will send Email verification and store temporary data on users_reg table
func sendEmailVer(payload *models.AuthFromClient, code string) error {

	smtp_email := os.Getenv("SMTP_EMAIL")
	smtp_pass := os.Getenv("SMTP_PASS")
	smtp_host := os.Getenv("SMTP_HOST")
	smtp_port := os.Getenv("SMTP_PORT")
	addr := fmt.Sprintf("%s:%s", smtp_host, smtp_port)
	auth := smtp.PlainAuth("", smtp_email, smtp_pass, smtp_host)

	subject := "Subject : The Light Email Verification\n"
	body := "Hi this is your email verification code : "

	msg := []byte(subject + body + code)

	err := smtp.SendMail(addr, auth, smtp_email, []string{payload.Email}, msg)
	if err != nil {
		return err
	}
	return nil
}

//storeToUsersReg will store to payload to temporary table
func storeToUsersReg(db *sql.DB, payload *models.AuthFromClient, code string) (int64, error) {

	insertedID, err := driver.StoreUsersReg(db, payload, code)
	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

//waitAndDelete will wait for 5 minutes before temporary data in users_reg get deleted
func waitAndDelete(db *sql.DB, ID int64) {
	time.Sleep(300 * time.Second)
	err := driver.DeleteUsersReg(db, ID)
	if err != nil {
		fmt.Println(err)
		return
	}
}
