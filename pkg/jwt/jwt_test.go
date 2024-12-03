package jwt

import (
	"testing"
)

func TestJWT(t *testing.T) {
	signKey := "123"

	username := "test"
	userID := 1

	//生成token
	token := SignForUser(userID, username, signKey)

	//解析 token，并判断能否正确获取 token 中的用户信息
	id, u, err := ParseUserToken(token, signKey)
	if err != nil {
		t.Error(err)
	}
	if id != userID {
		t.Errorf("expect user id %d, got %d", userID, id)
	}
	if u != username {
		t.Errorf("expect username %s, got %s", u, username)
	}

}
