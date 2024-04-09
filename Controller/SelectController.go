package Controller

import (
	"fmt"
	"sqlx/Service"
	"sqlx/domain"
)

func Select(transaction bool, userid int) []domain.User {
	person := domain.User{UserId: userid}
	UserSlice, err := Service.SelectService(transaction, person)
	if err != nil {
		fmt.Println("查找中出错了:", err)
		return nil
	}
	return UserSlice
}
