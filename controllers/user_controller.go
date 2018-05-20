package controllers

import (
	"net/http"
	"todone-api/model"
	"encoding/json"
	"todone-api/core/repository/user"
	"todone-api/core/network"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"
)

func CreateUser(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	requestUser := new(model.User)
	decoder := json.NewDecoder(network.GetRequestBodyBuffer(req))
	err := decoder.Decode(&requestUser)

	if err != nil {
		network.WriteErrorResponse(w, err, http.StatusUnprocessableEntity)
	}

	success, err := user.CreateUser(requestUser)

	network.SetJson(w)

	if err != nil || !success {
		network.WriteErrorResponse(w, err, http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func GetUser(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	vars := mux.Vars(req)

	var userId int
	var err error
	if userId, err = strconv.Atoi(vars["userId"]); err != nil {
		network.WriteErrorResponse(w, err, http.StatusBadRequest)
	}

	args := map[string]interface{}{"id": userId}

	user, err := user.GetUser(args)

	if err != nil {
		network.WriteErrorResponse(w, err, http.StatusInternalServerError)
	} else {
		network.WritePayloadResponse(w, user, http.StatusOK)
	}
}

func GetUsers(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	var (
		offset, limit int
		orderBy string
		orderDirection string
		err error
		)

	offset, err = network.GetUrlParamInt(req,"offset")
	limit, err = network.GetUrlParamInt(req,"limit")
	orderBy, err = network.GetUrlParamString(req, "orderBy")
	orderDirection, err = network.GetUrlParamString(req, "orderDirection")
	fmt.Println(err)

	fmt.Println(offset)
	fmt.Println(limit)
	fmt.Println(orderBy)
	fmt.Println(orderDirection)

	users, err := user.GetUsers(offset, limit, orderBy, orderDirection)

	if err != nil {
		network.WriteErrorResponse(w, err, http.StatusInternalServerError)
	} else {
		network.WritePayloadResponse(w, users, http.StatusOK)
	}
}