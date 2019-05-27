/**
 *  @author: yanKoo
 *  @Date: 2019/4/14 21:21
 *  @Description:
 */
package taskrunner

import (
	"errors"
	"log"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	d := func(dc dataChan) error {
		for i := 0; i < 30; i++ {
			dc <- i
			log.Println("gen data: ", i)
		}
		return nil
	}

	e := func(dc dataChan) error {
		//forloop:
		for {
			select {
			case d := <-dc:
				log.Printf("execute data: %d", d.(int))

			default:
				//break forloop
				log.Println("*****************")
			}
		}
		return errors.New("executor")
	}

	r:= NewRunner(30, false, d,e)
	go r.StartAll()
time.Sleep(time.Second*1000000)
}
