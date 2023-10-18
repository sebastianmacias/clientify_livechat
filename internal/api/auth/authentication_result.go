package auth

import (
	"time"

	"github.com/sebastianmacias/clientify_livechat/internal/models"
)

type AuthenticationResult struct {
	Token      string
	User       *models.User
	ValidUntil time.Time
	Scopes     []string
}
