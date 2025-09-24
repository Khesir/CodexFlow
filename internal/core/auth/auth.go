package auth

import (
	models "app/model"
	"crypto/sha256"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("supersecret")

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userID int) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func GenerateRefreshToken() (string, error) {
	// refresh token is just a random JWT without user data
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr, &Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("bad claims")
	}
	return claims, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashRefreshToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return string(sum[:])
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
			return
		}

		// Extract token
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := ParseToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// Put user_id into gin.Context
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}

// func Login(c *gin.Context, db *sqlx.DB) {
// 	var req struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
// 		return
// 	}

// 	var user models.User
// 	err := db.Get(&user, "SELECT * FROM users WHERE email = $1 AND deleted_at IS NULL", req.Email)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
// 		return
// 	}

// 	// Compare password properly (bcrypt recommended)
// 	if user.Password != req.Password {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
// 		return
// 	}

// 	// Generate tokens
// 	accessToken, _ := GenerateAccessToken(user.ID)
// 	refreshToken, _ := GenerateRefreshToken()

// 	// Hash & store refresh token
// 	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(refreshToken)))
// 	_, _ = db.Exec("UPDATE users SET refresh_token=$1, last_login_at=NOW() WHERE id=$2", hash, user.ID)

// 	c.JSON(http.StatusOK, gin.H{
// 		"accessToken":  accessToken,
// 		"refreshToken": refreshToken,
// 	})
// }

func RefreshToken(c *gin.Context, db *sqlx.DB) {
	var req struct {
		RefreshToken string `json:"refreshToken"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Hash input token
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(req.RefreshToken)))

	var user models.User
	err := db.Get(&user, "SELECT * FROM users WHERE refresh_token=$1 AND deleted_at IS NULL", hash)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	// Validate token expiration
	_, err = jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Expired or invalid refresh token"})
		return
	}

	// Issue new access token
	accessToken, _ := GenerateAccessToken(user.ID)

	c.JSON(http.StatusOK, gin.H{
		"accessToken": accessToken,
	})
}

func Signout(c *gin.Context, db *sqlx.DB) {
	userID := c.GetInt("user_id") // extracted from AuthGuard

	_, _ = db.Exec("UPDATE users SET refresh_token=NULL WHERE id=$1", userID)

	c.JSON(http.StatusOK, gin.H{"message": "Signed out"})
}

func SignIn(c *gin.Context, db *sqlx.DB) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var user models.User
	err := db.Get(&user, "SELECT * FROM users WHERE email=$1 AND deleted_at IS NULL", req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if !CheckPasswordHash(req.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate tokens
	accessToken, err := GenerateAccessToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	refreshToken, err := GenerateRefreshToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}

	// Hash & store refresh token
	hash := HashRefreshToken(refreshToken)
	_, _ = db.Exec("UPDATE users SET refresh_token=$1, last_login_at=NOW() WHERE id=$2", hash, user.ID)

	c.JSON(http.StatusOK, gin.H{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
		"user": gin.H{
			"id":      user.ID,
			"email":   user.Email,
			"isAdmin": user.IsAdmin,
		},
	})
}
