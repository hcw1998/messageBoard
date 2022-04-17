package helper

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	model "messageBoard/gin/Models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("nd93bjo4dk4")

type authClaims struct {
	jwt.StandardClaims
	UserID      uint `json:"user_id"`
	IsSuperuser bool `json:"is_superuser"`
}

func GenerateToken(user *model.User) (string, error) {
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, authClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Account,
			ExpiresAt: expiresAt,
		},
		UserID:      uint(user.Id),
		IsSuperuser: bool(user.IsSuperuser),
	})
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (uint, bool, error) {
	var claims authClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})
	if err != nil {
		return 0, false, err
	}
	if !token.Valid {
		return 0, false, errors.New("invalid token")
	}
	id := claims.UserID
	is_superuser := claims.IsSuperuser
	return id, is_superuser, nil
}

func VerifyToken(c *gin.Context) {
	token, ok := getToken(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": -1})
		return
	}

	id, is_superuser, err := ValidateToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": -2})
		return
	}
	c.Set("id", id)
	c.Set("is_superuser", is_superuser)
	c.Writer.Header().Set("Authorization", "Bearer "+token)
	c.Next()
}

func getToken(c *gin.Context) (string, bool) {
	authValue := c.GetHeader("Authorization")
	arr := strings.Split(authValue, " ")
	if len(arr) != 2 {
		return "", false
	}
	authType := strings.Trim(arr[0], "\n\r\t")
	if strings.ToLower(authType) != strings.ToLower("Bearer") {
		return "", false
	}
	return strings.Trim(arr[1], "\n\t\r"), true
}

func VerifySuperuser(c *gin.Context) {
	isSuperuser, _ := c.Get("is_superuser")
	var flag = isSuperuser.(bool)
	fmt.Printf("%t", flag)
	if !flag {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": -1})
		return
	}
	c.Next()
}
