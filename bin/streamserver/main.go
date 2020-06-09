/**
 *  @author: yanKoo
 *  @Date: 2019/4/1 21:53
 *  @Description:
 */
package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/yankooo/video_server/streamserver"
	"log"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *streamserver.ConnLimiter
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = streamserver.NewConnLimiter(cc)
	return m
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/videos/:vid-id", streamserver.StreamHandler)

	router.POST("/upload/:vid-id", streamserver.UploadHandler)

	router.GET("/testpage", streamserver.TestPageHandler)

	return router
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !m.l.GetConn() {
		streamserver.SendErrorResponse(w, http.StatusTooManyRequests, "Too many requests")
		return
	}

	m.r.ServeHTTP(w, r)
	defer m.l.ReleaseConn()
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r, 2)
	if err := http.ListenAndServe(":9000", mh); err != nil {
		log.Println("监听失败")
	}
}
