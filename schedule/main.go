/**
 *  @author: yanKoo
 *  @Date: 2019/4/14 21:56
 *  @Description:
 */
package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
	"video_server/schedule/taskrunner"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/video-delete-record/:vid-id", vidDelRecHandler)

	return router
}

func main() {

	w := taskrunner.NewWorker(5, nil)

	go w.StartWorker()
	time.Sleep(time.Second * 30)

	go taskrunner.Start()
	r := RegisterHandlers()
	http.ListenAndServe(":9001", r)
}

