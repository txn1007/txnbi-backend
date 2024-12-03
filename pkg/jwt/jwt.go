package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

func SignForUser(userID int, username string, signKey string) string {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "txnbi",
		// 有效期为 一周
		"exp":      time.Now().Add(7 * 24 * time.Hour).Unix(),
		"userID":   strconv.Itoa(userID),
		"username": username,
	})
	token, err := t.SignedString([]byte(signKey))
	if err != nil {
		panic(err)
	}
	return token
}

func ParseUserToken(token string, signKey string) (userID int, username string, err error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signKey), nil
	})
	if err != nil {
		return 0, "", err
	}
	// 检查 token 是否合法
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return 0, "", fmt.Errorf("invalid token")
	}
	// 获取用户信息
	userID, err = strconv.Atoi(claims["userID"].(string))
	if err != nil {
		return 0, "", err
	}
	return userID, claims["username"].(string), nil
}
