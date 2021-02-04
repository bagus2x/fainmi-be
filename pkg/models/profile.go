package models

import "github.com/dgrijalva/jwt-go"

// SignUpReq -
type SignUpReq struct {
	Username string `json:"username" validate:"required,gte=4"`
	Email    string `json:"email" validate:"required,email,gte=5"`
	Password string `json:"password" validate:"required,gte=5"`
}

// SignUpRes -
type SignUpRes struct {
	ProfileID   int    `json:"profileID"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	AccessToken string `json:"accessToken"`
}

//SignInReq -
type SignInReq struct {
	UsernameOrEmail string `json:"username" validate:"required,gte=4"`
	Password        string `json:"password" validate:"required,gte=5"`
}

//SignInRes -
type SignInRes struct {
	ProfileID   int    `json:"profileID"`
	Photo       string `json:"photo"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	AccessToken string `json:"accessToken"`
}

// ProfileUpdateReq -
type ProfileUpdateReq struct {
	Photo    string `json:"photo"`
	Username string `json:"username" validate:"required,gte=4"`
	Email    string `json:"email" validate:"required,email,gte=5"`
	Password string `json:"password" validate:"lte=0|gte=5"`
}

// ProfileUpdateRes -
type ProfileUpdateRes struct {
	ProfileID int    `json:"profileID"`
	Photo     string `json:"photo"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

// AccessClaims -
type AccessClaims struct {
	jwt.StandardClaims
	ProfileID int `json:"profileID"`
}
