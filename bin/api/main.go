package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/yankooo/video_server/api"
	"github.com/yankooo/video_server/api/session"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session
	api.ValidateUserSession(r)

	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", api.CreateUser)

	router.POST("/user/:username", api.Login)

	router.GET("/user/:username", api.GetUserInfo)

	router.POST("/user/:username/videos", api.AddNewVideo)

	router.GET("/user/:username/videos", api.ListAllVideos)

	router.DELETE("/user/:username/videos/:vid-id", api.DeleteVideo)

	router.POST("/videos/:vid-id/comments", api.PostComment)

	router.GET("/videos/:vid-id/comments", api.ShowComments)

	return router
}

func Prepare() {
	session.LoadSessionsFromDB()
}

func main() {
	Prepare()
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)
}
