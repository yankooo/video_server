/**
 *  @author: yanKoo
 *  @Date: 2019/3/10 20:52
 *  @Description:负责处理请求的业务逻辑
 */
package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if _, err := io.WriteString(w, "create User"); err != nil {
		log.Println("create io error")
	}
}

func Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	username := params.ByName("user_name") // 这里的参数是依靠前面router注册地方的user_name
	if _, err := io.WriteString(w, username); err != nil {
		log.Println("登录出错")
	}
}
