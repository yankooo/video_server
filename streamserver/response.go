/**
 *  @author: yanKoo
 *  @Date: 2019/4/1 22:00
 *  @Description:
 */
package streamserver

import (
	"io"
	"log"
	"net/http"
)

func SendErrorResponse(w http.ResponseWriter, sc int, errMsg string) {
	w.WriteHeader(sc)
	if _, err := io.WriteString(w, errMsg); err != nil {
		log.Println("io writer string is error : ", err)
	}
}
