package profile

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/bagus2x/fainmi-be/pkg/entities"
	"github.com/bagus2x/fainmi-be/pkg/models"
	"github.com/bagus2x/fainmi-be/pkg/models/errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

// Service interface to access the user repository
type Service interface {
	SignUp(req *models.SignUpReq) (*models.SignUpRes, error)
	SignIn(req *models.SignInReq) (*models.SignInRes, error)
	UpdateProfile(profileID int, req *models.ProfileUpdateReq) (*models.ProfileUpdateRes, error)
	DeleteProfile(profileID int) error
	CreateAccessToken(profileID int) (string, error)
	ParseAccessToken(token string) (string, error)
}

type service struct {
	repo           Repository
	accessTokenKey string
}

// NewService is instance of user service
func NewService(repo Repository, accessTokenKey string) Service {
	return &service{repo: repo, accessTokenKey: accessTokenKey}
}

func (s service) SignUp(req *models.SignUpReq) (*models.SignUpRes, error) {
	err := validator.New().Struct(req)
	if err != nil {
		return nil, errors.ErrorMessage(errors.ErrBadRequest, err.Error())
	}

	cred, err := s.repo.ReadCredentials(req.Username, req.Email)
	if cred != nil {
		if cred.Email == req.Email {
			return nil, errors.ErrEmailAlreadyExist
		}

		return nil, errors.ErrUsernameAlreadyExist
	}

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return nil, errors.ErrFailedToHash
	}

	profile := &entities.Profile{
		Email:     req.Email,
		Username:  req.Username,
		Password:  hashedPassword,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	err = s.repo.Create(profile)
	if err != nil {
		return nil, errors.ErrFailedToCreateUser
	}

	accessToken, err := s.CreateAccessToken(profile.ProfileID)
	if err != nil {
		return nil, err
	}

	res := &models.SignUpRes{
		ProfileID:   profile.ProfileID,
		Email:       profile.Email,
		Username:    profile.Username,
		AccessToken: accessToken,
	}

	return res, nil
}

func (s service) SignIn(req *models.SignInReq) (*models.SignInRes, error) {
	err := validator.New().Struct(req)
	if err != nil {
		return nil, errors.ErrorMessage(errors.ErrBadRequest, err.Error())
	}

	profile, err := s.repo.ReadByUsernameOrEmail(req.UsernameOrEmail)
	if err != nil {
		return nil, errors.ErrUserNotFound
	}

	isMatch := checkPasswordHash(req.Password, profile.Password)
	if !isMatch {
		return nil, errors.ErrIncorrectPassword
	}

	accessToken, err := s.CreateAccessToken(profile.ProfileID)
	if err != nil {
		return nil, err
	}

	res := &models.SignInRes{
		ProfileID:   profile.ProfileID,
		Email:       profile.Email,
		Username:    profile.Username,
		Photo:       profile.Photo.String,
		AccessToken: accessToken,
	}

	return res, nil
}

func (s service) UpdateProfile(profileID int, req *models.ProfileUpdateReq) (*models.ProfileUpdateRes, error) {
	err := validator.New().Struct(req)
	if err != nil {
		return nil, errors.ErrorMessage(errors.ErrBadRequest, err.Error())
	}

	profile, err := s.repo.Read(profileID)
	if err != nil {
		return nil, errors.ErrUserNotFound
	}

	profile.UpdatedAt = time.Now().Unix()
	profile.Photo = sql.NullString{Valid: req.Photo != "", String: req.Photo}
	profile.Username = req.Username
	profile.Email = req.Email
	if req.Password != "" {
		hashedPassword, err := hashPassword(req.Password)
		if err != nil {
			return nil, errors.ErrFailedToHash
		}
		profile.Password = hashedPassword
	}

	isUpdated, err := s.repo.Update(profileID, profile)
	if err != nil {
		return nil, errors.ErrInternalServer
	}
	if !isUpdated {
		return nil, errors.ErrUserNotFound
	}

	res := models.ProfileUpdateRes{
		ProfileID: profile.ProfileID,
		Photo:     profile.Photo.String,
		Username:  profile.Username,
		Email:     profile.Email,
	}

	return &res, nil
}

func (s service) DeleteProfile(profileID int) error {
	isDeleted, err := s.repo.Delete(profileID)
	if err != nil {
		return errors.ErrInternalServer
	}
	if !isDeleted {
		return errors.ErrUserNotFound
	}

	return nil
}

func (s service) CreateAccessToken(profileID int) (string, error) {
	claims := models.AccessClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
		ProfileID: profileID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.accessTokenKey))
}

func (s service) ParseAccessToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &models.AccessClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.accessTokenKey), nil
	})

	if err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			return "", errors.ErrTokenExpired
		}
		return "", errors.ErrInvalidAccessToken
	}

	if claims, ok := token.Claims.(*models.AccessClaims); ok && token.Valid {
		return claims.Subject, nil
	}

	return "", errors.ErrInvalidAccessToken
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
