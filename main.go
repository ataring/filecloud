package main

import (
	"file-manager/config"
	"file-manager/router"
	"fmt"
	"log"
	
	"os"
)

func main() {
	fileInfo, err := os.Stat(config.RootPath)
	if err != nil || !fileInfo.IsDir() {
		fmt.Println("[ERROR] 文件夹路径不存在")
		panic("文件夹路径不存在")
	}

	r := router.StartRoute()
	r.LoadHTMLGlob("view/**")
	
	//r.StaticFS("/static", http.Dir("./static"))
	if err := r.Run(":3333"); err != nil {
		log.Fatal("服务器启动失败...")
	}

}
