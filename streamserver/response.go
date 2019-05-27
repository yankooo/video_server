/**
 *  @author: yanKoo
 *  @Date: 2019/4/1 22:00
 *  @Description:
 */
package main

import (
	"io"
	"log"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, sc int, errMsg string){
	w.WriteHeader(sc)
	if _, err := io.WriteString(w, errMsg); err != nil {
		log.Println("io writer string is error : ", err)
	}
}

