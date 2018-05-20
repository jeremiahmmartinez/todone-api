package network

import (
	"net/http"
	"io/ioutil"
	"bytes"
	"todone-api/model"
	"encoding/json"
	"fmt"
	"strconv"
	"errors"
)

func ReadUser(req *http.Request) (*model.User, error) {
	requestUser := new(model.User)
	decoder := json.NewDecoder(GetRequestBodyBuffer(req))

	if err := decoder.Decode(&requestUser); err != nil {
		fmt.Println(err)
		return requestUser, err
	}

	return requestUser, nil
}

func ReadRequestBody(r *http.Request) []byte {
	requestBodyBytes, _ := ioutil.ReadAll(r.Body)
	defer func() {
		r.Body.Close()
		r.Body = ioutil.NopCloser(bytes.NewBuffer(requestBodyBytes))
	}()

	return requestBodyBytes
}

func GetRequestBodyBuffer(r *http.Request) *bytes.Buffer {
	return bytes.NewBuffer(ReadRequestBody(r))
}

func GetUrlParamString(req *http.Request, key string) (string, error) {
	var (
		targetParam string
		err error
	)

	params, ok := req.URL.Query()[key]

	if !ok || len(params) < 1 {
		err = errors.New("missing parameter: '" + targetParam + "'")
	} else {
		targetParam = params[0]
	}

	return targetParam, err
}

func GetUrlParamInt(req *http.Request, key string) (int, error) {
	var (
		targetParam string
		targetInt int
		err error
	)

	params, ok := req.URL.Query()[key]

	if !ok || len(params) < 1 {
		return targetInt, errors.New("missing parameter")
	}

	targetParam = params[0]

	if targetInt, err = strconv.Atoi(targetParam); err != nil {
		targetInt = 0
	}

	return targetInt, err
}