/*
 * @Author: nijineko
 * @Date: 2024-04-12 19:49:41
 * @LastEditTime: 2024-04-12 21:39:41
 * @LastEditors: nijineko
 * @Description: 客户端入口
 * @FilePath: \Miyako\cmd\client\main.go
 */
package main

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/nijinekoyo/miyako/pkg/block"
	"github.com/schollz/progressbar/v3"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: client <CatalogAddr> <SavePath>")

		os.Exit(1)
	}

	// 获取第一个命令参数作为Catalog地址
	CatalogAddr := os.Args[1]
	var SavePath string = "./download/"
	if len(os.Args) >= 3 {
		// 获取第二个命令参数作为保存路径
		SavePath = os.Args[2]
	}

	// 获取Catalog
	CatalogData, err := getCatalog(CatalogAddr)
	if err != nil {
		panic(err)
	}

	// 拼接文件块地址
	CatalogURL, err := url.Parse(CatalogAddr)
	if err != nil {
		panic(err)
	}
	BlockURL := fmt.Sprintf("%s://%s%s/%s",
		CatalogURL.Scheme,
		CatalogURL.Host,
		filepath.ToSlash(filepath.Dir(CatalogURL.Path)),
		CatalogData.BlockFile,
	)

	// 创建文件块下载对象
	BlockDownload, err := block.NewDownload(BlockURL, CatalogData, SavePath)
	if err != nil {
		panic(err)
	}
	defer BlockDownload.Close()

	// 创建一个进度条
	Progressbar := progressbar.NewOptions(int(BlockDownload.TotalSize),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionFullWidth(),
		progressbar.OptionSetDescription("Downloading..."),
		progressbar.OptionClearOnFinish(),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[light_blue]=[reset]",
			SaucerHead:    "[light_blue]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	// 开始下载
	err = BlockDownload.Start(func(WrittenSize, _ int64) {
		Progressbar.Add64(WrittenSize)
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Download success")
}
