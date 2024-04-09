package DataBase

import (
	_ "github.com/go-sql-driver/mysql"
)

var DbPool *MySQLDBPool

func init() {

	// 创建一个 MySQL 数据库连接池
	dbPool := NewMySQLDBPool()
	DbPool = dbPool
}
