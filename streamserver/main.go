/**
 *  @author: yanKoo
 *  @Date: 2019/4/1 21:53
 *  @Description:
 */
package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/videos/:vid-id", streamHandler)

	router.POST("/upload/:vid-id", uploadHandler)

	router.GET("/testpage", testPageHandler)

	return router
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		sendErrorResponse(w, http.StatusTooManyRequests, "Too many requests")
		return
	}

	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

func main()  {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r, 2)
	if err := http.ListenAndServe(":9000", mh); err != nil {
		log.Println("监听失败")
	}
}