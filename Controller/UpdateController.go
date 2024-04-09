package Controller

import (
	"fmt"
	"sqlx/Service"
	"sqlx/domain"
)

func Update(transaction bool, username string, userid int) {

	person := domain.User{Username: username, UserId: userid}
	AffectedRowNums, err := Service.UpdateService(transaction, person)
	if err != nil {
		fmt.Println("执行update语句出错了:", err)
	} else {
		fmt.Printf("执行Update语句成功，最终改变了%d条\n", AffectedRowNums)
	}
}
