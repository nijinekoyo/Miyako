/*
 * @Author: nijineko
 * @Date: 2024-04-12 17:18:47
 * @LastEditTime: 2024-04-12 19:38:42
 * @LastEditors: nijineko
 * @Description: flag封装
 * @FilePath: \Miyako\internal\flag\flag.go
 */
package flag

import "flag"

type Flag struct {
	AssetsFolder string // 资产文件夹路径
	HTTPPort     string // HTTP端口
}

var Get Flag // 全局参数变量

/**
 * @description: 初始化参数
 * @return {error} 错误信息
 */
func Init() error {
	// 参数解析
	AssetsFolder := flag.String("assets_folder", "./assets", "资产文件夹路径")
	HTTPPort := flag.String("http_port", "8080", "HTTP端口")
	flag.Parse()

	// 赋值全局参数
	Get.AssetsFolder = *AssetsFolder
	Get.HTTPPort = *HTTPPort

	return nil
}
