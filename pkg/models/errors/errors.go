package errors

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

var (
	// ErrInternalServer A generic error message, given when an unexpected condition was encountered
	ErrInternalServer = errors.New("Internal Server Error")
	// ErrNotFound The requested resource could not be found
	ErrNotFound = errors.New("Not Found")
	// ErrConflict Indicates that the request could not be processed because of conflict
	ErrConflict = errors.New("Conflict")
	// ErrBadRequest The server cannot or will not process the request due to an apparent client error
	ErrBadRequest = errors.New("Bad request")
	// ErrUnauthorized The user does not have valid authentication credentials for the target resource
	ErrUnauthorized = errors.New("Unauthorized")
	// ErrUsernameOrEmailConflict -
	ErrUsernameOrEmailConflict = fmt.Errorf("%w: Username or email already exist", ErrConflict)
	// ErrUserNotFound -
	ErrUserNotFound = fmt.Errorf("%w: User does not exist", ErrNotFound)
	// ErrFailedToHash -
	ErrFailedToHash = fmt.Errorf("%w: Failed to hash", ErrInternalServer)
	// ErrInvalidAccessToken -
	ErrInvalidAccessToken = fmt.Errorf("%w: Invalid Acccess Token", ErrUnauthorized)
	// ErrTokenExpired -
	ErrTokenExpired = fmt.Errorf("%w: Token expired", ErrUnauthorized)
	// ErrFailedToCreateUser -
	ErrFailedToCreateUser = fmt.Errorf("%w: Failed to create user", ErrInternalServer)
	// ErrUsernameAlreadyExist -
	ErrUsernameAlreadyExist = fmt.Errorf("%w: Username already exist", ErrConflict)
	// ErrEmailAlreadyExist -
	ErrEmailAlreadyExist = fmt.Errorf("%w: Email already exist", ErrConflict)
	// ErrIncorrectPassword -
	ErrIncorrectPassword = fmt.Errorf("%w: Icorrect Password", ErrBadRequest)
	// ErrStyleNotFound -
	ErrStyleNotFound = fmt.Errorf("%w: Style does not exist", ErrNotFound)
	// ErrLinkNotFound -
	ErrLinkNotFound = fmt.Errorf("%w: Link does not exist", ErrNotFound)
	// ErrBackgroundNotFound -
	ErrBackgroundNotFound = fmt.Errorf("%w: Background does not exist", ErrNotFound)
	// ErrButtonNotFound -
	ErrButtonNotFound = fmt.Errorf("%w: Button does not exist", ErrNotFound)
	// ErrFontNotFound -
	ErrFontNotFound = fmt.Errorf("%w: Font does not exist", ErrNotFound)
	// ErrLikeNotFound -
	ErrLikeNotFound = fmt.Errorf("%w: Like does not exist", ErrNotFound)
	// ErrorMessage for easy unwrapping
	ErrorMessage = func(err error, msg string) error {
		return fmt.Errorf("%w: %s", err, msg)
	}
	// ErrDatabase -
	ErrDatabase = func(err error) error {
		if err, ok := err.(*pq.Error); ok {
			var errMessage string
			switch err.Code {
			case "23502":
				// not-null constraint violation
				errMessage = err.Detail
				break
			case "23503":
				// foreign key violation
				errMessage = err.Detail
				break
			case "23505":
				// unique constraint violation
				errMessage = err.Detail
				break
			case "23514":
				// check constraint violation
				errMessage = err.Detail
				break
			}
			return fmt.Errorf("%w: %s", ErrBadRequest, errMessage)
		}
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("%w: Item does not exist", ErrNotFound)
		}
		return err
	}
)
