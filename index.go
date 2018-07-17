package main

import (
	"log"
	"os"
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
	args := os.Args[1:]
	
	runParams := alis.ParseParams(args)
	log.Printf("解析出运行参数: type=%s, args=%s", runParams.Type, runParams.Args)

	handler := &alis.Handler{Params: *runParams}
	handler.Run(runParams)
}