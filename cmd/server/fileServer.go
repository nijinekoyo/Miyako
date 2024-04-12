/*
 * @Author: nijineko
 * @Date: 2024-04-12 19:44:21
 * @LastEditTime: 2024-04-12 19:44:22
 * @LastEditors: nijineko
 * @Description: 文件服务器
 * @FilePath: \Miyako\cmd\server\fileServer.go
 */
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nijinekoyo/miyako/internal/flag"
)

// 启动文件服务器
func startFileServer() error {
	// 创建文件服务
	FileServer := http.FileServer(http.Dir(flag.Get.AssetsFolder))
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if info, err := os.Stat(flag.Get.AssetsFolder + r.URL.Path); err == nil && info.IsDir() {
			// 如果是目录则返回404
			http.NotFound(w, r)
			return
		}

		// 否则将请求交给文件服务器处理
		FileServer.ServeHTTP(w, r)
	}))

	// 启动 HTTP 服务器
	log.Default().Println("Server started at http://0.0.0.0:" + flag.Get.HTTPPort)
	err := http.ListenAndServe(":"+flag.Get.HTTPPort, nil)
	if err != nil {
		return err
	}

	return nil
}
