/**
 *  @author: yanKoo
 *  @Date: 2019/4/14 20:24
 *  @Description:
 */
package taskrunner

import (
	"errors"
	"github.com/yankooo/video_server/scheduler/dbops"
	"github.com/yankooo/video_server/scheduler/ossops"
	"log"
	"sync"
)

func deleteVideo(vid string) error {
	//err := os.Remove(VIDEO_PATH + vid)
	//if err !=nil && !os.IsNotExist(err){
	//	log.Printf("Delete video file error: %v", err)
	//}
	//return nil
	if ret := ossops.DeleteObject("videos/" +vid, "yankooo-videos"); !ret {
		log.Printf("Deleting video error, oss operation failed.")
		return errors.New("deleting video error")
	}
	return nil
}

func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		log.Printf("Video clear dispathcher error: %v", err)
		return err
	}

	if len(res) == 0 {
		return errors.New("all tasks finished")
	}

	for _, id := range res {
		dc <- id
	}

	return nil
}

func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error

	forloop:
		for {
			select {
			case vid :=<- dc:
				go func(id interface{}) {
					// 会有重复读写的问题
					if err := deleteVideo(id.(string)); err  != nil {
						errMap.Store(id ,err)
						return
					}
					if err := dbops.DelVideoDeletionRecord(id.(string)); err != nil {
						errMap.Store(id, err)
						return
					}
				}(vid)
			default:
				break forloop
			}
		}
	errMap.Range(func(key, value interface{}) bool {
		err = value.(error)
		if err != nil {
			return false
		}
		return true
	})

	return err
}