package controllers

import (
	"crypto/subtle"
	"database/sql"
	"encoding/base64"
	"errors"
	"net/http"
	"quiz-sb-pekan3/database"
	"quiz-sb-pekan3/repository"
	"quiz-sb-pekan3/structs"
	"strings"

	"github.com/gin-gonic/gin"
)

func BasicAuthMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Validasi format Basic Auth
		if len(authHeader) < 6 || authHeader[:6] != "Basic " {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		// Ambil kredensial username dan password yang terkode base64
		encodedCredentials := authHeader[6:]
		decoded, err := base64.StdEncoding.DecodeString(encodedCredentials)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid base64 encoding"})
			c.Abort()
			return
		}

		// Pisahkan username dan password
		parts := strings.SplitN(string(decoded), ":", 2)
		if len(parts) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid basic auth format"})
			c.Abort()
			return
		}

		username := parts[0]
		password := parts[1]

		user, err := AuthenticateUser(db, username, password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			c.Abort()
			return
		}

		// Menyimpan informasi user untuk digunakan di request berikutnya
		c.Set("user", user)

		// Lanjutkan ke handler berikutnya
		c.Next()
	}
}

// Fungsi untuk memverifikasi password dengan hash yang ada di database
func ValidateUserPassword(storedPassword, providedPassword string) bool {
	// Pada kasus nyata, password di database harus di-hash menggunakan bcrypt atau teknik hashing lainnya
	return subtle.ConstantTimeCompare([]byte(storedPassword), []byte(providedPassword)) == 1
}

// Fungsi untuk mendapatkan user dan memverifikasi password
func AuthenticateUser(db *sql.DB, username, password string) (*structs.User, error) {
	user, err := repository.GetUserByUsername(database.DbConnection, username)
	if err != nil {
		return nil, err
	}

	if !ValidateUserPassword(user.Password, password) {
		return nil, errors.New("invalid username or password")
	}

	return user, nil
}

func GetUsernameFromAuthHeader(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("Authorization header is required")
	}

	// Verifikasi format Authorization header (Basic)
	if !strings.HasPrefix(authHeader, "Basic ") {
		return "", errors.New("Invalid authorization header format. Must start with Basic")
	}

	// Mengambil nilai kredensial setelah "Basic "
	encodedCredentials := authHeader[6:]
	decoded, err := base64.StdEncoding.DecodeString(encodedCredentials)
	if err != nil {
		return "", errors.New("Failed to decode base64 credentials")
	}

	// Pisahkan username dan password dengan ":"
	parts := strings.SplitN(string(decoded), ":", 2)
	if len(parts) != 2 {
		return "", errors.New("Invalid basic auth format. Must be 'username:password'")
	}

	// Mengembalikan username
	return parts[0], nil
}
