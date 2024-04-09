package Service

import (
	"database/sql"
	"sqlx/CRUD"
	"sqlx/domain"
)

func InsertService(transaction bool, UserInformation domain.User) (sql.Result, error) {
	res, err := CRUD.Insert(transaction, UserInformation)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func DeleteService(transaction bool, UserInformation domain.User) (int64, error) {
	res, err := CRUD.Delete(transaction, UserInformation)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func UpdateService(transaction bool, UserInformation domain.User) (int64, error) {
	res, err := CRUD.Update(transaction, UserInformation)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func SelectService(transaction bool, UserInformation domain.User) ([]domain.User, error) {
	res, err := CRUD.Select(transaction, UserInformation)
	if err != nil {
		return res, err
	}
	return res, nil
}
func BatchSelectService(transaction bool, SliceUserid []int) ([]domain.User, error) {
	UserSlice, err := CRUD.BatchSelect(transaction, SliceUserid)
	if err != nil {
		return nil, err
	}
	return UserSlice, nil
}

func BatchDeleteService(transaction bool, SliceUserid []int) (int64, error) {
	res, err := CRUD.BatchDelete(transaction, SliceUserid)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
