package data

import (
	"github.com/dgrijalva/jwt-go"
)

type sharedData struct {
	token jwt.Token
	userId uint64
}

func (s *sharedData) SetToken(token *jwt.Token) {
	s.token = *token
}

func (s *sharedData) GetToken() jwt.Token {
	return s.token
}

func (s *sharedData) SetUserId(userId *uint64) {
	s.userId = *userId
}

func (s *sharedData) GetUserId() uint64 {
	return s.userId
}

var (
	s *sharedData
)

func SharedData() *sharedData {
	if s == nil {
		s = &sharedData{
		}
	}
	return s
}