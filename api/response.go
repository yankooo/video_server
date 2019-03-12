/**
 *  @author: yanKoo
 *  @Date: 2019/3/10 21:59
 *  @Description:
 */
package main

import (
	"encoding/json"
	"github.com/yankooo/video_server/api/defs"
	"io"
	"log"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, errResp defs.ErrorResponse) {
	w.WriteHeader(errResp.HttpSC)

	respStr, err := json.Marshal(&errResp.Error)
	if err != nil {
		log.Println("json marshal is fail in sendErrorResponse function.")
	}
	if _, err := io.WriteString(w, string(respStr)); err != nil {
		log.Println("sendErrorResponse is fail")
	}
}

func sendNormalResponse(w http.ResponseWriter, respStr string, sc int) {
	w.WriteHeader(sc)
	if _, err := io.WriteString(w, respStr); err != nil {
		log.Println("sendNormalResponse is fail")
	}
}
