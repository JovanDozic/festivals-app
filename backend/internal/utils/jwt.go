package utils

import (
	"context"
	"errors"
	"log"
	"time"

	modelsUser "backend/internal/models/user"

	"github.com/golang-jwt/jwt"
)

type JWTUtil struct {
	Secret []byte
}

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

type contextKey string

const (
	UserKey     contextKey = "user"
	usernameKey contextKey = "username"
	roleKey     contextKey = "role"
)

func NewJWTUtil(secret string) *JWTUtil {
	return &JWTUtil{Secret: []byte(secret)}
}

func (j *JWTUtil) GenerateToken(username string, role string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		string(usernameKey): username,
		string(roleKey):     role,
		"exp":               time.Now().Add(time.Hour * 24).Unix(), // ? lower this down?
	})

	tokenString, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}

	log.Println("token generated successfully")
	return tokenString, nil
}

func (j *JWTUtil) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.Secret), nil
	})
	if err != nil {
		log.Println("error validating token:", err)
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		log.Println("provided token is invalid")
		return nil, errors.New("invalid token")
	}
	if time.Now().Unix() > claims.ExpiresAt {
		log.Println("provided token is expired")
		return nil, errors.New("token expired")
	}
	if claims.Role == "" || (claims.Role != "ADMIN" && claims.Role != "EMPLOYEE" && claims.Role != "ATTENDEE" && claims.Role != "ORGANIZER") {
		log.Println("role not found in claims")
		return nil, errors.New("role not found in claims")
	}

	log.Println("token validated successfully")
	return claims, nil
}

func AuthAdmin(c context.Context) bool {
	claims, ok := c.Value(UserKey).(*Claims)
	if !ok {
		return false
	}
	return claims.Role == string(modelsUser.RoleAdmin)
}

func AuthOrganizer(c context.Context) bool {
	claims, ok := c.Value(UserKey).(*Claims)
	if !ok {
		return false
	}
	return claims.Role == string(modelsUser.RoleOrganizer)
}

func AuthAttendeeRole(c context.Context) bool {
	claims, ok := c.Value(UserKey).(*Claims)
	if !ok {
		return false
	}
	return claims.Role == string(modelsUser.RoleAttendee)
}

func AuthEmployee(c context.Context) bool {
	claims, ok := c.Value(UserKey).(*Claims)
	if !ok {
		return false
	}
	return claims.Role == string(modelsUser.RoleEmployee)
}

func GetUsername(c context.Context) string {
	claims, ok := c.Value(UserKey).(*Claims)
	if !ok {
		return ""
	}
	return claims.Username
}

func Auth(c context.Context) bool {
	_, ok := c.Value(UserKey).(*Claims)
	return ok
}
