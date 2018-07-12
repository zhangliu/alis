package main

import (
	"log"
	"os"
	"strings"
	"zl/alis/src"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			log.Print(err)
		}
	}()

	if len(os.Args) < 2 {
		log.Println("请输入参数！")
		return
	}
	argStr := strings.Join(os.Args[1:], " ")
	log.Printf("获取到参数：%s", argStr)
	
	runParams := alis.ParseParams(argStr)
	log.Printf("解析出运行参数: 类型=%s, args=%s", runParams.Type, runParams.Args)

	handler := &alis.Handler{Params: *runParams}
	handler.Run(runParams)
}