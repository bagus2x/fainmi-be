package routes

import (
	"errors"

	"github.com/bagus2x/fainmi-be/pkg/models"
	custom "github.com/bagus2x/fainmi-be/pkg/models/errors"
)

// Get http status code
func code(err error) int {
	switch errors.Unwrap(err) {
	case custom.ErrBadRequest:
		return 400
	case custom.ErrUnauthorized:
		return 401
	case custom.ErrNotFound:
		return 404
	case custom.ErrConflict:
		return 409
	case custom.ErrInternalServer:
	}

	return 500
}

type r models.Response
