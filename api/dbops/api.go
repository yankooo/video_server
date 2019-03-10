/**
 *  @author: yanKoo
 *  @Date: 2019/3/10 21:54
 *  @Description: 实现对数据库几张表的crud
 */
package dbops

import (
	"database/sql"
	"log"
)

func AddUserCredential(loginName, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")
	if err != nil {
		return err
	}

	if _, err := stmtIns.Exec(loginName, pwd); err != nil {
		log.Println("insert user fail")
	}

	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM user WHERE login_name = ?")
	if err != nil {
		log.Println("select user prepare fail")
	}

	var pwd string
	if err := stmtOut.QueryRow(loginName).Scan(&pwd); err != nil && err != sql.ErrNoRows {
		log.Println("query pwd fail")
	}
	defer stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?")
	if err != nil {
		log.Printf("DeleteUser error: %s", err)
		return err
	}

	if _, err := stmtDel.Exec(loginName, pwd); err != nil {
		log.Println("delete user fail")
	}

	if err := stmtDel.Close(); err != nil {
		log.Println("mysql connection close fail")
	}
	return nil
}
