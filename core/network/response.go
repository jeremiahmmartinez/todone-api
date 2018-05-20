package network

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type ErrorResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

type PayloadResponse struct {
	Code int `json:"code"`
	Payload interface{} `json:"payload"`
}

type MessageResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

func (err ErrorResponse) Error() string {
	return fmt.Sprintf("[%d] %s", err.Code, err.Message)
}

func WriteErrorResponse(writer http.ResponseWriter, err error, errorCode int) {
	writer.WriteHeader(errorCode)
	SetJson(writer)

	validationError := ErrorResponse{Message:err.Error(), Code:errorCode}
	if err := json.NewEncoder(writer).Encode(validationError); err != nil {
		panic(err)
	}
}

func WritePayloadResponse(writer http.ResponseWriter, payload interface{}, code int) {
	writer.WriteHeader(code)
	SetJson(writer)

	response := PayloadResponse{Code:code, Payload:payload}
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		panic(err)
	}
}

func WriteMessageResponse(writer http.ResponseWriter, message string, code int) {
	writer.WriteHeader(code)
	SetJson(writer)

	response := MessageResponse{Code:code, Message:message}
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		panic(err)
	}
}

func SetJson(writer http.ResponseWriter)  {
	writer.Header().Set("Content-Type", "application/json")
}

func Write(writer http.ResponseWriter, output string) {
	writer.Write([]byte(output))
}