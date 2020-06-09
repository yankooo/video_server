package logic

import (
	"flag"
	"fmt"
	"github.com/yankooo/video_server/logic"
	"os"
	"runtime"
	"time"
)

var (
	confFile string // 配置文件路径
)

func initArgs() {
	flag.StringVar(&confFile, "conf", "./logic.json", "where logic.json is.")
	flag.Parse()
}

func initEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		err error
	)

	// 初始化环境
	initArgs()
	initEnv()

	if err = logic.InitConfig(confFile); err != nil {
		goto ERR
	}

	if err = logic.InitStats(); err != nil {
		goto ERR
	}

	if err = logic.InitGateConnMgr(); err != nil {
		goto ERR
	}

	if err = logic.InitService(); err != nil {
		goto ERR
	}

	for {
		time.Sleep(1 * time.Second)
	}

	os.Exit(0)

ERR:
	fmt.Fprintln(os.Stderr, err)
	os.Exit(-1)
}
