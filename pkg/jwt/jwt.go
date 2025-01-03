package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

func SignForUser(userID int64, userAccount string, signKey string) string {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "txnbi",
		// 有效期为 一周
		"exp":         time.Now().Add(7 * 24 * time.Hour).Unix(),
		"userID":      strconv.FormatInt(userID, 10),
		"userAccount": userAccount,
	})
	token, err := t.SignedString([]byte(signKey))
	if err != nil {
		panic(err)
	}
	return token
}

func ParseUserToken(token string, signKey string) (userID int64, userAccount string, err error) {
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
	userID, err = strconv.ParseInt(claims["userID"].(string), 10, 64)
	if err != nil {
		return 0, "", err
	}
	return userID, claims["userAccount"].(string), nil
}

func SignForChartAccessCode(userID, chartID int64, token, signKey string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "txnbi",
		// 有效期为 3 天
		"exp":       time.Now().Add(3 * 24 * time.Hour).Unix(),
		"userID":    userID,
		"chartID":   chartID,
		"userToken": token,
	})
	token, err := t.SignedString([]byte(signKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseChartAccessCode(token string, signKey string) (userID, chartID int64, userToken string, err error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(signKey), nil
	})
	if err != nil {
		return 0, 0, "", err
	}
	// 检查 token 是否合法
	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return 0, 0, "", fmt.Errorf("invalid token")
	}
	// 获取用户,图表信息
	userID, err = strconv.ParseInt(claims["userID"].(string), 10, 64)
	if err != nil {
		return 0, 0, "", err
	}
	chartID, err = strconv.ParseInt(claims["chartID"].(string), 10, 64)
	if err != nil {
		return 0, 0, "", err
	}
	userToken = claims["chartID"].(string)
	return
}
