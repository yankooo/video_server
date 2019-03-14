/**
* @Author: yanKoo
* @Date: 2019/3/12 15:38
* @Description: 选择sync包下的map来做缓存
 */
package session

import (
	"github.com/yankooo/video_server/api/dbops"
	"github.com/yankooo/video_server/api/defs"
	"github.com/yankooo/video_server/api/utils"
	"log"
	"sync"
	"time"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64 {
	return time.Now().UnixNano() / 1000000 // 纳秒/10 00000
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid) // 先删除存储在缓存中的session
	// 删除数据库中的过期session
	if err := dbops.DeleteSession(sid); err != nil {
		log.Printf("Delete Expired Seesion Fail: %s", err)
	}
}

// 从数据库中加载session，放进内存中的map
func LoadSessionFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}

	r.Range(func(key, value interface{}) bool {
		ss := value.(*defs.SimpleSession)
		sessionMap.Store(key, ss)
		return true
	})
}

func GenerateNewSessionId(un string) string {
	sid, _ := utils.NewUUID()
	ct := nowInMilli()
	ttl := ct + 30*60*1000 // Server side session valid time is 30 min

	ss := &defs.SimpleSession{Username: un, TTL: ttl}
	sessionMap.Store(sid, ss)
	if err := dbops.InsertSeesion(sid, ttl, un); err != nil {
		log.Printf("insert session fail : %s", err)
		return ""
	}
	return sid
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	ct := nowInMilli()
	if ok {
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, false
	} else {
		ss, err := dbops.RetrieveSession(sid)
		if err != nil || ss == nil {
			return "", true
		}

		if ss.TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}

		sessionMap.Store(sid, ss)
		return ss.Username, false
	}
}
