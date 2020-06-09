/**
 *  @author: yanKoo
 *  @Date: 2019/4/14 20:24
 *  @Description:
 */
package taskrunner

import (
	"log"
	"time"
)

type Worker struct {
	ticker *time.Ticker
	runner *Runner
}

func NewWorker(interval time.Duration, r *Runner) *Worker {
	return &Worker{
		ticker: time.NewTicker(interval * time.Second),
		runner: r,
	}
}

func (w *Worker) startWorker() {
	for {
		select {
		case <-w.ticker.C:
			go w.runner.StartAll()
		}
	}
}

func (w *Worker) StartWorker() {
	for {
		select {
		case <-w.ticker.C:
			log.Println("ticker is run")
		}
	}
}

func Start() {
	// Start video file cleaning
	r := NewRunner(3, true, VideoClearDispatcher, VideoClearExecutor)
	w := NewWorker(3, r)
	go w.startWorker()
}
