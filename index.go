package main

import (
	"log"
	"os"
	"net/http"
	"time"
	"zl/alis/src"
	"zl/alis/src/utils"
	"strings"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			log.Print(err)
		}
	}()

	// if len(os.Args) < 2 {
	// 	log.Println("请输入参数！")
	// 	return
	// }
	// argStr := strings.Join(os.Args[1:], " ")
	// log.Printf("获取到参数：%s", argStr)
	
	file := "./cmd.wav"
	if !utils.IsFileExist(file) {
		log.Print("voice file not found!")
		return
	}

	time.Sleep(10)

	runParams := alis.ParseParams(argStr)
	log.Printf("解析出运行参数: 类型=%s, args=%s", runParams.Type, runParams.Args)

	handler := &alis.Handler{Params: *runParams}
	handler.Run(runParams)
}

func getText(file string) string {
	if isExist, _ := utils.IsFileExist(file); isExist {
		log.Print("voice file not found!")
		return ""
	}
	url := "http://api.xfyun.cn/v1/service/v1/iat"
	contentType := "application/x-www-form-urlencoded; charset=utf-8"
	body := ""
	req, _ := http.Post(url, contentType, strings.NewReader(body))
	client := &http.Client{}
	client.Do(req)
}