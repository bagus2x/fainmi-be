package models

import "github.com/dgrijalva/jwt-go"

// SignUpRequest -
type SignUpRequest struct {
	Username string `json:"username" validate:"required,gte=4"`
	Email    string `json:"email" validate:"required,email,gte=5"`
	Password string `json:"password" validate:"required,gte=5"`
}

// SignUpResponse -
type SignUpResponse struct {
	ProfileID   int    `json:"profileID"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	AccessToken string `json:"accessToken"`
}

//SignInRequest -
type SignInRequest struct {
	UsernameOrEmail string `json:"username" validate:"required,gte=4"`
	Password        string `json:"password" validate:"required,gte=5"`
}

//SignInResponse -
type SignInResponse struct {
	ProfileID   int    `json:"profileID"`
	Photo       string `json:"photo"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	AccessToken string `json:"accessToken"`
}

// ProfileUpdateRequest -
type ProfileUpdateRequest struct {
	Photo    string `json:"photo"`
	Username string `json:"username" validate:"required,gte=4"`
	Email    string `json:"email" validate:"required,email,gte=5"`
	Password string `json:"password" validate:"lte=0|gte=5"`
}

// ProfileUpdateResponse -
type ProfileUpdateResponse struct {
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
