/**
 *  @author: yanKoo
 *  @Date: 2019/4/14 20:24
 *  @Description:
 */
package taskrunner

type Runner struct {
	Controller controlChan
	Error      controlChan
	Data       dataChan
	dataSize   int
	longLived  bool
	Dispatcher fn
	Executor   fn
}

func NewRunner(size int, longlived bool, d fn, e fn) *Runner {
	return &Runner{
		Controller: make(chan string, 1),
		Error:      make(chan string, 1),
		Data:       make(dataChan, size),
		longLived:  longlived,
		dataSize:   size,
		Dispatcher: d,
		Executor:   e,
	}
}


func (r *Runner) startDispatch() {
	defer func() {
		if !r.longLived {
			close(r.Controller)
			close(r.Error)
			close(r.Data)
		}
	}()

	// 常驻任务,所以用for循环
	for {
		select {
		case c :=<- r.Controller:
			if c == READY_TO_DISPATCH {
				err := r.Dispatcher(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_EXECUTE
				}

			}

			if c == READY_TO_EXECUTE {
				err:= r.Executor(r.Data)
				if err != nil {
					r.Error <- CLOSE
				} else {
					r.Controller <- READY_TO_DISPATCH
				}
			}
		case e :=<- r.Error:
			if e == CLOSE {
				return
			}
		//default:
			
		}
	}
}

func (r *Runner) StartAll()  {
	r.Controller <- READY_TO_DISPATCH
	r.startDispatch()
}

