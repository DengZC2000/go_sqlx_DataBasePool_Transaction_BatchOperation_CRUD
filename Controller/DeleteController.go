package Controller

import (
	"fmt"
	"sqlx/Service"
	"sqlx/domain"
)

func Delete(transaction bool, userid int) {
	person := domain.User{UserId: userid}
	AffectedRowNums, err := Service.DeleteService(transaction, person)
	if err != nil {
		fmt.Println("执行delete语句出错了:", err)
	} else {
		fmt.Printf("执行delete语句成功，最终删除了%d条\n", AffectedRowNums)
	}
}
