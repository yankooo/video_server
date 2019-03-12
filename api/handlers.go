/**
 *  @author: yanKoo
 *  @Date: 2019/3/10 20:52
 *  @Description:负责处理请求的业务逻辑
 */
package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/yankooo/video_server/api/dbops"
	"github.com/yankooo/video_server/api/defs"
	"github.com/yankooo/video_server/api/session"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func SignedUp(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// 发送POST请求，有body，所以使用ioutil包里面的ReadAll
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}

	// 使用json解析器反序列化，把res转换为ubody
	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	// 添加用户
	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	// 将生成一个session记录的uuid保存在数据库里面，实现注册
	id := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{Success: true, SessionId: id}

	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		log.Println("json marshal fail")
		return
	} else {
		sendNormalResponse(w, string(resp), 201)
	}
}

func SignedIn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	username := params.ByName("user_name") // 这里的参数是依靠前面router注册地方的user_name
	if _, err := io.WriteString(w, username); err != nil {
		log.Println("登录出错")
	}
}
