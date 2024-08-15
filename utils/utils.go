package utils

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var OK = http.StatusOK

func Error(c *gin.Context, statusCode int, errorMessage string) {
	log.Printf("Error: %s\n", errorMessage)
	c.Header("Content-Type", "application/json")
	c.JSON(statusCode, gin.H{"error": errorMessage})
} // to return error

func Message(c *gin.Context, message string) {
	c.Header("Content-Type", "application/json")
	c.JSON(OK, gin.H{"message": message})
} // to return messages

func Response(c *gin.Context, data interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(OK, data)
} // to return response

func ValidateEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func ValidateURL(url string) bool {
	urlRegex := `^(https?:\/\/)?([\da-z.-]+)\.([a-z.]{2,6})([/\w .-]*)*\/?$`
	re := regexp.MustCompile(urlRegex)
	return re.MatchString(url)
}

func ValidatePANCard(pan string) bool {
	panRegex := `^[A-Z]{5}[0-9]{4}[A-Z]{1}$`
	re := regexp.MustCompile(panRegex)
	return re.MatchString(pan)
}

func IsNonNegative(n int) bool {
	return n > 0
}

func ValidatePassword(password string) (string, bool) {
	minLength := 6
	maxLength := 10
	hasUppercase := false
	hasLowercase := false
	hasDigit := false
	hasSpecial := false

	if password == "" {
		return "Password is required", false
	}

	if len(password) < minLength || len(password) > maxLength {
		return "Password must be between 6 to 10 characters", false
	}

	for _, char := range password {
		switch {
		case 'a' <= char && char <= 'z':
			hasLowercase = true
		case 'A' <= char && char <= 'Z':
			hasUppercase = true
		case '0' <= char && char <= '9':
			hasDigit = true
		case char == '@' || char == '$' || char == '!' || char == '*' || char == '%' || char == '?' || char == '&':
			hasSpecial = true
		}
	}

	if !hasLowercase {
		return "Password must contain at least one lowercase letter", false
	}

	if !hasUppercase {
		return "Password must contain at least one uppercase letter", false
	}

	if !hasDigit {
		return "Password must contain at least one digit", false
	}

	if !hasSpecial {
		return "Password must contain at least one special character (@$!*%?&)", false
	}

	return "", true // Password passed all validations
}

func CheckLength(num int, length int) bool {
	return len(strconv.Itoa(num)) != length
}

func GetTrimedUrl(url string) string {
	return strings.TrimPrefix(url, "https://storage.googleapis.com/golangwithfirebase.appspot.com/")
}
