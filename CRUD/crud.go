package CRUD

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"sqlx/DataBase"
	"sqlx/domain"
)

func Insert(transaction bool, UserInfo domain.User) (sql.Result, error) {

	Exe_DB := DataBase.DbPool.Get()

	defer DataBase.DbPool.Put(Exe_DB)

	Sql := "insert into user (username, sex, email)values(?, ?, ?)"
	if transaction == true {
		conn, err := Exe_DB.Begin()
		if err != nil {
			return nil, err
		}
		res, err := conn.Exec(Sql, UserInfo.Username, UserInfo.Sex, UserInfo.Email)
		if err != nil {
			conn.Rollback()
			return nil, err
		}
		conn.Commit()
		return res, nil
	}
	res, err := Exe_DB.Exec(Sql, UserInfo.Username, UserInfo.Sex, UserInfo.Email)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func Delete(transaction bool, UserInfo domain.User) (sql.Result, error) {
	Exe_DB := DataBase.DbPool.Get()

	defer DataBase.DbPool.Put(Exe_DB)
	Sql := "delete from user where user_id=?"
	if transaction == true {
		conn, err := Exe_DB.Begin()
		if err != nil {
			return nil, err
		}
		res, err := conn.Exec(Sql, UserInfo.UserId)
		if err != nil {
			conn.Rollback()
			return nil, err
		}
		conn.Commit()
		return res, nil
	}
	res, err := Exe_DB.Exec(Sql, UserInfo.UserId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func Update(transaction bool, UserInfo domain.User) (sql.Result, error) {
	Exe_DB := DataBase.DbPool.Get()

	defer DataBase.DbPool.Put(Exe_DB)
	Sql := "update user set username=? where user_id=?"
	if transaction == true {
		conn, err := Exe_DB.Begin()
		if err != nil {
			return nil, err
		}
		res, err := conn.Exec(Sql, UserInfo.Username, UserInfo.UserId)
		if err != nil {
			conn.Rollback()
			return nil, err
		}
		conn.Commit()
		return res, nil
	}
	res, err := Exe_DB.Exec(Sql, UserInfo.Username, UserInfo.UserId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func Select(transaction bool, UserInfo domain.User) ([]domain.User, error) {
	Exe_DB := DataBase.DbPool.Get()

	defer DataBase.DbPool.Put(Exe_DB)
	var user []domain.User
	Sql := "select user_id, username, sex, email from user where user_id=?"
	if transaction == true {
		conn := Exe_DB.MustBegin()

		err := conn.Select(&user, Sql, UserInfo.UserId)
		if err != nil {
			conn.Rollback()
			return nil, err
		}
		conn.Commit()
		return user, nil
	}
	err := Exe_DB.Select(&user, Sql, UserInfo.UserId)
	if err != nil {

		return nil, err
	}
	return user, nil
}

func BatchSelect(transaction bool, SliceUserID []int) ([]domain.User, error) {
	Exe_DB := DataBase.DbPool.Get()

	defer DataBase.DbPool.Put(Exe_DB)
	var UserSlice []domain.User

	Sql := "select * from user where user_id IN (?)"
	query, args, err := sqlx.In(Sql, SliceUserID)
	if err != nil {
		return nil, err
	}
	if transaction == true {
		conn := Exe_DB.MustBegin()

		err := conn.Select(&UserSlice, query, args...)
		if err != nil {
			conn.Rollback()
			return nil, err
		}
		conn.Commit()

		return UserSlice, nil
	}
	err = Exe_DB.Select(&UserSlice, query, args...)
	if err != nil {
		return nil, err
	}
	return UserSlice, nil
}

func BatchDelete(transaction bool, SliceUserID []int) (sql.Result, error) {
	Exe_DB := DataBase.DbPool.Get()

	defer DataBase.DbPool.Put(Exe_DB)
	Sql := "delete from user where user_id IN (?)"
	query, args, err := sqlx.In(Sql, SliceUserID)
	if err != nil {
		return nil, err
	}
	if transaction == true {
		conn, err := Exe_DB.Begin()
		if err != nil {
			return nil, err
		}
		res, err := conn.Exec(query, args...)
		if err != nil {
			conn.Rollback()
			return nil, err
		}
		conn.Commit()
		return res, nil
	}
	res, err := Exe_DB.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return res, nil
}
