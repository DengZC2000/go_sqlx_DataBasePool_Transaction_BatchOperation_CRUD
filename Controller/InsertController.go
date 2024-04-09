package Controller

import (
	"fmt"
	"sqlx/Service"
	"sqlx/domain"
)

func Insert(transaction bool, name string, sex string, email string) {
	person := domain.User{Username: name, Sex: sex, Email: email}
	res, err := Service.InsertService(transaction, person)
	if err != nil {
		fmt.Println("执行insert语句出错了:", err)
	} else {
		AffectedRowNums, _ := res.RowsAffected()
		LastInsertID, _ := res.LastInsertId()
		fmt.Printf("执行insert语句成功，最终插入%d条,插入的最后一条ID为：%d\n", AffectedRowNums, LastInsertID)
	}
}
