package authorization

import (
	"net/http"
	"todone-api/data"
	"fmt"
)

func CheckResourceAuthorization(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	sharedData := data.SharedData()
	authorized := true

	userId := sharedData.GetUserId()

	fmt.Println(userId)

	if authorized {
		next(rw, req)
	} else {
		rw.WriteHeader(http.StatusUnauthorized)
	}
}