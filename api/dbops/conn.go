/**
 *  @author: yanKoo
 *  @Date: 2019/3/10 22:32
 *  @Description:
 */
package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

/** 
 * 连接数据库
 */
func init() {
	dbConn, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/video_server")
	if err != nil {
		panic(err.Error())
	}
}
