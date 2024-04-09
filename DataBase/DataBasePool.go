package DataBase

import (
	"github.com/jmoiron/sqlx"
	"sync"
)

// MySQLDBPool 是一个包含数据库连接的池子
type MySQLDBPool struct {
	pool *sync.Pool
}

// NewMySQLDBPool 返回一个新的 MySQL 数据库连接池
func NewMySQLDBPool() *MySQLDBPool {
	return &MySQLDBPool{
		pool: &sync.Pool{
			New: func() interface{} {
				db, err := sqlx.Open("mysql", "root:1310138359@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True")
				if err != nil {
					panic(err)
				}
				return db
			},
		},
	}
}

// Get 从连接池中获取一个数据库连接
func (p *MySQLDBPool) Get() *sqlx.DB {
	return p.pool.Get().(*sqlx.DB)
}

// Put 将一个数据库连接放回连接池
func (p *MySQLDBPool) Put(db *sqlx.DB) {
	p.pool.Put(db)
}
