/*
 * @Author: nijineko
 * @Date: 2024-04-12 17:02:05
 * @LastEditTime: 2024-04-12 19:42:49
 * @LastEditors: nijineko
 * @Description: 服务端入口
 * @FilePath: \Miyako\cmd\server\main.go
 */
package main

import (
	"log"

	"github.com/nijinekoyo/miyako/internal/flag"
)

func main() {
	// 初始化Flag
	err := flag.Init()
	if err != nil {
		log.Default().Panic(err)
	}

	// 为资产文件夹生成文件块
	err = generateBlock()
	if err != nil {
		log.Default().Panic(err)
	}

	// 启动文件服务器
	if err = startFileServer(); err != nil {
		log.Default().Panic(err)
	}
}
