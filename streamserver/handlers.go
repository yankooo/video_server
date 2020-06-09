/**
 *  @author: yanKoo
 *  @Date: 2019/4/1 21:53
 *  @Description:
 */
package streamserver

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func TestPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("videos/upload.html")

	_ = t.Execute(w, nil)
}

func StreamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	//vl := VIDEO_DIR + vid + ".mp4"
	//video, err := os.Open(vl)
	//if err != nil {
	//	log.Printf("Error when try to open file: %v", err)
	//	sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
	//	return
	//}
	//defer func() {
	//	if err := video.Close(); err != nil {
	//		log.Println("Video close error")
	//	}
	//}()
	//
	//w.Header().Set("Content-Type", "video/mp4")
	//http.ServeContent(w, r, "", time.Now(), video)

	targetUrl := "http://yankooo-videos.oss-cn.aliyuncs.com/videos/" + vid
	http.Redirect(w, r, targetUrl, http.StatusMovedPermanently)
}

func UploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is too big")
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error when try to get file: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
	}

	fn := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR+fn+".mp4", data, 0777)
	if err != nil {
		log.Printf("Write file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	ossFn := "videos/" + fn
	path := "./videos/" + fn
	bn := "yankooo-videos"
	ret := UploadToOss(ossFn, path, bn)
	if !ret {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}

	os.Remove(path)
	w.WriteHeader(http.StatusCreated)
	_, _ = io.WriteString(w, "Uploaded successfully")
}
