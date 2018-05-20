package validation

import (
	"net/http"
	"encoding/json"
	"todone-api/core/network"
	"todone-api/model"
	userRepo "todone-api/core/repository/user"
	"github.com/badoux/checkmail"
	"regexp"
	"errors"
)


var IsAlphaNumeric = regexp.MustCompile("^[a-zA-Z0-9_]+$").MatchString


func ValidateNewUsernameFormat(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	requestUser, err := network.ReadUser(req)

	if err != nil {
		network.WriteErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}

	validFormat := IsAlphaNumeric(requestUser.Username)

	if validFormat {
		next(w, req)
	} else {
		err := errors.New("invalid username format")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func ValidateNewUsernameExists(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	requestUser := new(model.User)
	decoder := json.NewDecoder(network.GetRequestBodyBuffer(req))
	err := decoder.Decode(&requestUser)

	if err != nil {
		network.WriteErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}

	exists, err := userRepo.UsernameExists(requestUser.Username)

	if err == nil && !exists {
		next(w, req)
	} else {
		err := errors.New("username already exists")
		network.WriteErrorResponse(w, err, http.StatusBadRequest)
	}
}


func ValidateNewEmailFormat(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	requestBodyBuffer := network.GetRequestBodyBuffer(req)

	requestUser := new(model.User)
	decoder := json.NewDecoder(requestBodyBuffer)
	err := decoder.Decode(&requestUser)

	if err != nil {
		network.WriteErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}

	validationError := checkmail.ValidateFormat(requestUser.Email)

	if validationError == nil {
		next(w, req)
	} else {
		err := errors.New("invalid email format")
		network.WriteErrorResponse(w, err, http.StatusBadRequest)
	}
}

func ValidateNewEmailExists(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	requestBodyBuffer := network.GetRequestBodyBuffer(req)

	requestUser := new(model.User)
	decoder := json.NewDecoder(requestBodyBuffer)
	err := decoder.Decode(&requestUser)

	if err != nil {
		network.WriteErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}

	exists, err := userRepo.EmailExists(requestUser.Email)

	if err == nil && !exists {
		next(w, req)
	} else {
		err := errors.New("email already exists")
		network.WriteErrorResponse(w, err, http.StatusBadRequest)
	}
}