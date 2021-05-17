package utils

import (
	"fmt"
	"net/http"
)

func ResErr(res *http.ResponseWriter, code int, err error) {
	payload := fmt.Sprintf("%s", err.Error())
	(*res).WriteHeader(code)
	(*res).Write([]byte(payload))
}
