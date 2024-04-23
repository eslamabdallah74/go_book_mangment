package helper

import (
	"fmt"
	"net/http"
)

func MakeContentTypeApplicationJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func CustomError(err any) {
	fmt.Printf("error... : \n %v", err)
}
