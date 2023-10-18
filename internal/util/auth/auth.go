package auth

import "github.com/sebastianmacias/clientify_livechat/internal/models"

// UserHasScopes checks if a User has all the required scopes.
func UserHasScopes(user *models.User, requiredScopes []string) bool {
	// Create a map for quick lookup of user's scopes
	userScopesMap := make(map[string]bool)
	for _, scope := range user.Scopes {
		userScopesMap[scope] = true
	}

	// Check if the user has each of the required scopes
	for _, requiredScope := range requiredScopes {
		if _, exists := userScopesMap[requiredScope]; !exists {
			return false
		}
	}

	return true
}

// UserHasAnyScopes checks if a User has any of the required scopes.
func UserHasAnyScopes(user *models.User, requiredScopes []string) bool {
	// Create a map for quick lookup of user's scopes
	userScopesMap := make(map[string]bool)
	for _, scope := range user.Scopes {
		userScopesMap[scope] = true
	}

	// Check if the user has any of the required scopes
	for _, requiredScope := range requiredScopes {
		if _, exists := userScopesMap[requiredScope]; exists {
			return true
		}
	}

	return false
}

// UserHasScope checks if a User has the specific scope.
func UserHasScope(user *models.User, requiredScope string) bool {
	for _, scope := range user.Scopes {
		if scope == requiredScope {
			return true
		}
	}
	return false
}
