package store

import "txnbi-backend/internal/model"

func GetUserByAccount(account string) (*model.User, error) {
	var u model.User
	err := DB.Where("userAccount = ?", account).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}
