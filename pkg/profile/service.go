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
	SignUp(req *models.SignUpRequest) (*models.SignUpResponse, error)
	SignIn(req *models.SignInRequest) (*models.SignInResponse, error)
	GetProfile(profileID int) (*models.GetProfileResponse, error)
	UpdateProfile(profileID int, req *models.ProfileUpdateRequest) (*models.ProfileUpdateResponse, error)
	DeleteProfile(profileID int) error
	CreateAccessToken(profileID int) (string, error)
	ParseAccessToken(token string) (*models.AccessClaims, error)
}

type service struct {
	repo           Repository
	accessTokenKey string
}

// NewService is instance of user service
func NewService(repo Repository, accessTokenKey string) Service {
	return &service{repo: repo, accessTokenKey: accessTokenKey}
}

func (s service) SignUp(req *models.SignUpRequest) (*models.SignUpResponse, error) {
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

	res := &models.SignUpResponse{
		Profile: models.Profile{
			ProfileID: profile.ProfileID,
			Email:     profile.Email,
			Username:  profile.Username,
		},
		Token: models.Token{
			AccessToken: accessToken,
		},
	}

	return res, nil
}

func (s service) SignIn(req *models.SignInRequest) (*models.SignInResponse, error) {
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

	res := &models.SignInResponse{
		Profile: models.Profile{
			ProfileID: profile.ProfileID,
			Email:     profile.Email,
			Username:  profile.Username,
			Photo:     profile.Photo.String,
		},
		Token: models.Token{
			AccessToken: accessToken,
		},
	}

	return res, nil
}

func (s service) GetProfile(profileID int) (*models.GetProfileResponse, error) {
	profile, err := s.repo.Read(profileID)
	if err != nil {
		return nil, errors.ErrUserNotFound
	}

	res := &models.GetProfileResponse{
		ProfileID: profile.ProfileID,
		Photo:     profile.Photo.String,
		Username:  profile.Username,
		Email:     profile.Email,
	}

	return res, nil
}

func (s service) UpdateProfile(profileID int, req *models.ProfileUpdateRequest) (*models.ProfileUpdateResponse, error) {
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

	updated, err := s.repo.Update(profileID, profile)
	if err != nil {
		return nil, errors.ErrDatabase(err)
	}
	if !updated {
		return nil, errors.ErrUserNotFound
	}

	res := models.ProfileUpdateResponse{
		ProfileID: profile.ProfileID,
		Photo:     profile.Photo.String,
		Username:  profile.Username,
		Email:     profile.Email,
	}

	return &res, nil
}

func (s service) DeleteProfile(profileID int) error {
	deleted, err := s.repo.Delete(profileID)
	if err != nil {
		return errors.ErrInternalServer
	}
	if !deleted {
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

func (s service) ParseAccessToken(accessToken string) (*models.AccessClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &models.AccessClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.accessTokenKey), nil
	})

	if err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			return nil, errors.ErrTokenExpired
		}
		return nil, errors.ErrInvalidAccessToken
	}

	if claims, ok := token.Claims.(*models.AccessClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.ErrInvalidAccessToken
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
