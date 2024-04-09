package Controller

import (
	"fmt"
	"sqlx/Service"
)

func BatchDelete(transaction bool, userid ...int) {
	var slice_userid []int
	for _, val := range userid {
		slice_userid = append(slice_userid, val)
	}

	AffectedRowNums, err := Service.BatchDeleteService(transaction, slice_userid)
	if err != nil {
		fmt.Println("执行batch_delete语句出错了:", err)
	} else {
		fmt.Printf("执行batch_delete语句成功，最终删除了%d条\n", AffectedRowNums)
	}
}
