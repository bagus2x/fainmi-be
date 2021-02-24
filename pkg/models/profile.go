package models

import "github.com/dgrijalva/jwt-go"

// Profile -
type Profile struct {
	ProfileID int    `json:"profileID"`
	Photo     string `json:"photo"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

// Token -
type Token struct {
	AccessToken string `json:"accessToken"`
}

// SignUpRequest -
type SignUpRequest struct {
	Username string `json:"username" validate:"required,gte=4"`
	Email    string `json:"email" validate:"required,email,gte=5"`
	Password string `json:"password" validate:"required,gte=5"`
}

// SignUpResponse -
type SignUpResponse struct {
	Profile Profile `json:"profile"`
	Token   Token   `json:"token"`
}

//SignInRequest -
type SignInRequest struct {
	UsernameOrEmail string `json:"username" validate:"required,gte=4"`
	Password        string `json:"password" validate:"required,gte=5"`
}

//SignInResponse -
type SignInResponse struct {
	Profile Profile `json:"profile"`
	Token   Token   `json:"token"`
}

// ProfileUpdateRequest -
type ProfileUpdateRequest struct {
	Photo    string `json:"photo"`
	Username string `json:"username" validate:"required,gte=4"`
	Email    string `json:"email" validate:"required,email,gte=5"`
	Password string `json:"password" validate:"lte=0|gte=5"`
}

// ProfileUpdateResponse -
type ProfileUpdateResponse Profile

// GetProfileResponse -
type GetProfileResponse Profile

// AccessClaims -
type AccessClaims struct {
	jwt.StandardClaims
	ProfileID int `json:"profileID"`
}
