package Controller

import (
	"fmt"
	"sqlx/Service"
	"sqlx/domain"
)

func BatchSelect(transaction bool, userid ...int) []domain.User {
	var UserIDSlice []int
	for _, val := range userid {
		UserIDSlice = append(UserIDSlice, val)
	}
	UserSlice, err := Service.BatchSelectService(transaction, UserIDSlice)
	if err != nil {
		fmt.Println("批量查找中出错了:", err)
		return nil
	}
	return UserSlice
}
