package auth

import (
	"time"

	"github.com/futurescopex/starter/internal/models"
)

type AuthenticationResult struct {
	Token      string
	User       *models.User
	ValidUntil time.Time
	Scopes     []string
}
