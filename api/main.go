/**
 *  @author: yanKoo
 *  @Date: 2019/3/10 20:42
 *  @Description:
 */
package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

/**
 * 注册路由handler
 * 处理流程：handler->validation(1.request, 2.user)->business logic->response
 *     1. data model
 *     2. error handling
 */
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)

	router.POST("/user/:user_name", Login)

	return router
}

func main() {
	r := RegisterHandlers()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println("监听失败")
	}
}
