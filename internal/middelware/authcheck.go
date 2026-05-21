package middelware

import (
	"fmt"
	"net/http"
	"tipodikayayagoda/internal/utils"
)

func AuthCheck(token *http.Cookie) (string, error) {

	claims, err := utils.ValidateToken(token.Value)
	if err != nil {
		return "", err
	}

	roleVal, ok := claims["role"]
	if !ok {
		return "", fmt.Errorf("role not found in token")
	}

	role, ok := roleVal.(string)
	if !ok {
		return "", fmt.Errorf("role is not string")
	}

	return role, nil
}
