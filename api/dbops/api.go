/**
 *  @author: yanKoo
 *  @Date: 2019/3/10 21:54
 *  @Description: 实现对数据库几张表的crud
 */
package dbops

import (
	"database/sql"
	"github.com/yankooo/video_server/api/defs"
	"github.com/yankooo/video_server/api/utils"
	"log"
	"time"
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
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
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

// 获取用户
func GetUser(loginName string) (*defs.User, error) {
	stmtOut, err := dbConn.Prepare("SELECT id, pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	var id int
	var pwd string

	err = stmtOut.QueryRow(loginName).Scan(&id, &pwd)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	res := &defs.User{Id: id, LoginName: loginName, Pwd: pwd}
	defer stmtOut.Close()
	return res, nil
}

// 添加视频的数据记录插入
func AddNewVideo(aId int, name string) (*defs.VideoInfo, error) {
	// create uuid
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}

	t := time.Now()
	ctime := t.Format("2006-1-2 15:04:05")
	stmtIns, err := dbConn.Prepare(`INSERT INTO video_info 
		(id, author_id, name, display_ctime) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}

	if _, err = stmtIns.Exec(vid, aId, name, ctime); err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{Id: vid, AuthorId: aId, Name: name, DisplayCtime: ctime}

	defer stmtIns.Close()
	return res, nil
}

// 获取视频信息
// 输入参数：视频uuid，返回整个视频的data modal
func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM video_info WHERE id = ?")
	if err != nil {
		return nil, err
	}

	var (
		aid  int
		dct  string
		name string
	)

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	res := &defs.VideoInfo{Id: vid, Name: name, DisplayCtime: dct}

	defer stmtOut.Close()
	return res, nil
}

func ListVideoInfo(uname string, from, to int) ([]*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare(`SELECT video_info.id, video_info.author_id,video_info.name, video_info.display_ctime FROM video_info
			INNER JOIN users ON video_info.author_id = users.id
			WHERE user.login_name = ? AND video_info.create_time > FROM_UNIXTIME(?) AND video_info.create_time <= FROM_UNIXTIME(?)
			ORDER BY video_info.create_time DESC`)

	var res []*defs.VideoInfo

	if err != nil {
		return nil, err
	}

	rows, err := stmtOut.Query(uname, from, to)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id, name, ctime string
		var aid int
		if err := rows.Scan(&id, &aid, &name, &ctime); err != nil {
			return res, err
		}

		vi := &defs.VideoInfo{Id: id, AuthorId: aid, Name: name, DisplayCtime: ctime}
		res = append(res, vi)
	}

	defer stmtOut.Close()
	return res, nil
}

// 根据视频的uuid删除相应的记录
func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM videoinfo WHERE id = ?")
	if err != nil {
		return err
	}

	if _, err := stmtDel.Exec(vid); err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}

// 增加评论
func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}

	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	if _, err := stmtIns.Exec(id, vid, aid, content); err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

// 获取一段时间内的评论
func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(`SELECT comments.id, users.login_name, comments.content FROM comments
			INNER JOIN users ON comments.authorid = users_id
			WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)
			ORDER BY comments.time DESC`)

	var res []*defs.Comment

	if err != nil {
		return nil, err
	}

	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}

		c := &defs.Comment{Id: id, Author: name, Content: content}
		res = append(res, c)
	}
	defer stmtOut.Close()
	return res, nil
}
