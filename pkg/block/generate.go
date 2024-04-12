/*
 * @Author: nijineko
 * @Date: 2024-04-12 17:53:08
 * @LastEditTime: 2024-04-12 20:15:01
 * @LastEditors: nijineko
 * @Description: 创建文件块
 * @FilePath: \Miyako\pkg\block\generate.go
 */
package block

import (
	"os"
	"path/filepath"

	"github.com/nijinekoyo/miyako/pkg/catalog"
)

/**
 * @description: 创建一个文件块
 * @param {string} FolderPath 文件夹路径
 * @return {*Catalog} catalog
 * @return {error} 错误
 */
func Generate(FolderPath string) (*catalog.Catalog, error) {
	// 为文件夹创建catalog
	CatalogData, err := catalog.Generate(FolderPath)
	if err != nil {
		return nil, err
	}

	// 生成文件块
	BlockFile, err := os.OpenFile(filepath.Join(FolderPath, CatalogData.BlockFile), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	// 写入文件块
	for _, File := range CatalogData.Files {
		// 读取文件
		FileData, err := os.ReadFile(filepath.Join(FolderPath, File.Path))
		if err != nil {
			return nil, err
		}

		// 写入文件块
		Size, err := BlockFile.Write(FileData)
		if err != nil {
			return nil, err
		}

		// 核对文件大小
		if int64(Size) != File.Size {
			return nil, os.ErrInvalid
		}
	}

	// 关闭文件块
	err = BlockFile.Close()
	if err != nil {
		return nil, err
	}

	return CatalogData, nil
}
