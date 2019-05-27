/**
 *  @author: yanKoo
 *  @Date: 2019/3/10 21:54
 *  @Description: 实现对数据库几张表的crud
 */
package dbops

import "log"

func AddVideoDeletionRecord(vid string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO video_del_rec (video_id) VALUES(?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(vid)
	if err != nil {
		log.Printf("AddVideoDeletionRecord error: %v", err)
		return err
	}

	defer stmtIns.Close()
	return nil
}
