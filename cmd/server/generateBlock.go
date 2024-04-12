/*
 * @Author: nijineko
 * @Date: 2024-04-12 18:07:58
 * @LastEditTime: 2024-04-12 18:21:01
 * @LastEditors: nijineko
 * @Description: 批量创建文件块
 * @FilePath: \Miyako\cmd\server\generateBlock.go
 */
package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/HyacinthusAcademy/miyako/internal/flag"
	"github.com/HyacinthusAcademy/miyako/pkg/block"
)

/**
 * @description: 创建文件块
 * @return {*error} error
 */
func generateBlock() error {
	// 获取资产文件夹路径
	AssetsPath := flag.Get.AssetsFolder

	// 遍历文件夹
	Dirs, err := os.ReadDir(AssetsPath)
	if err != nil {
		return err
	}

	for _, Dir := range Dirs {
		// 检查是否为文件夹
		if Dir.IsDir() {
			DirPath := filepath.Join(AssetsPath, Dir.Name())

			// 检查是否存在catalog.json
			_, err := os.Stat(filepath.Join(DirPath, "catalog.json"))
			if err == nil {
				continue
			}

			log.Default().Printf("Linking files in %s/......", Dir.Name())

			// 生成文件块
			Catalog, err := block.Generate(DirPath)
			if err != nil {
				return err
			}

			// 序列化Catalog
			CatalogData, err := Catalog.Marshal()
			if err != nil {
				return err
			}

			// 写入Catalog到文件夹
			err = os.WriteFile(filepath.Join(DirPath, "catalog.json"), CatalogData, 0644)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
