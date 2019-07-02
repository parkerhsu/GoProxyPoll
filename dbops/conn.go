package dbops

import (
	"database/sql"
	"github.com/gomodule/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
)

var (
	dbConn *sql.DB
	err error
)

func init() {
	dbConn, err := redis.Dial("tcp", "127.0.0.1")
	if err != nil {
		panic(err.Error())
	}

}