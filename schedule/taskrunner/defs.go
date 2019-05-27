/**
 *  @author: yanKoo
 *  @Date: 2019/4/14 20:24
 *  @Description:
 */
package taskrunner


// 在controlChan中的数据
const (
	READY_TO_DISPATCH = "d"
	READY_TO_EXECUTE = "e"
	CLOSE = "c"

	VIDEO_PATH = "videos/"
)

type controlChan chan string

type dataChan chan interface{}

type fn func(dc dataChan) error