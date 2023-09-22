package auth

import (
	"time"

	"github.com/sebastianmacias/starter/internal/models"
)

type AuthenticationResult struct {
	Token      string
	User       *models.User
	ValidUntil time.Time
	Scopes     []string
}
