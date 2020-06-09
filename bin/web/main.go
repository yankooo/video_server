package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/yankooo/video_server/web"
	"net/http"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", web.HomeHandler)

	router.POST("/", web.HomeHandler)

	router.GET("/userhome", web.UserHomeHandler)

	router.POST("/userhome", web.UserHomeHandler)

	router.POST("/api", web.ApiHandler)

	router.GET("/videos/:vid-id", web.ProxyVideoHandler)

	router.POST("/upload/:vid-id", web.ProxyUploadHandler)

	router.ServeFiles("/statics/*filepath", http.Dir("./templates"))

	return router
}

func main() {
	r := RegisterHandler()
	http.ListenAndServe(":8080", r)
}
