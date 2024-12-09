package helper

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// CheckUserType validates if the user type matches the expected role
func CheckUserType(c *gin.Context, role string) error {
	userType := c.GetString("user_type")
	if userType == "" {
		return errors.New("user type not found in context")
	}
	if userType != role {
		return errors.New("unauthorized to access this resource")
	}
	return nil
}

// MatchUserTypeToUid ensures the user type and user ID match expected values
func MatchUserTypeToUid(c *gin.Context, userId string) error {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")

	if userType == "" {
		return errors.New("user type not found in context")
	}

	if userType == "USER" && uid != userId {
		return errors.New("unauthorized to access this resource")
	}

	// Validate against an expected role (optional, based on logic)
	return CheckUserType(c, "ADMIN")
}
