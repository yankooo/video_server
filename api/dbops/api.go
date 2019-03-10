/**
 *  @author: yanKoo
 *  @Date: 2019/3/10 21:54
 *  @Description: 实现对数据库几张表的crud
 */
package dbops

import "database/sql"

/**
 * 连接数据库
 */
func openConn() *sql.DB {
	return nil
}

func AddUserCredential(loginName, pwd string) error {
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	return nil
}




