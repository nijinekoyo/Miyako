/*
 * @Author: nijineko
 * @Date: 2024-04-12 17:11:21
 * @LastEditTime: 2024-04-12 17:41:46
 * @LastEditors: nijineko
 * @Description: 生成catalog
 * @FilePath: \Miyako\pkg\catalog\generate.go
 */
package catalog

import (
	"hash/crc32"
	"os"
	"path"
	"path/filepath"

	"github.com/HyacinthusAcademy/miyako/tools/file"
	"github.com/HyacinthusAcademy/miyako/tools/random"
)

/**
 * @description: 生成catalog
 * @param {string} FolderPath 文件夹路径
 * @return {*Catalog} catalog
 * @return {error} 错误
 */
func Generate(FolderPath string) (*Catalog, error) {
	// 遍历文件夹
	Paths, err := file.GetPaths(FolderPath)
	if err != nil {
		return nil, err
	}

	// 生成catalog
	var CatalogData Catalog

	// 包名字为文件夹名字+随机字符
	CatalogData.PackageFile = path.Base(FolderPath) + "_" + random.String(16, 3)

	// 遍历文件
	for _, FilePath := range Paths {
		// 读取文件
		FileData, err := os.ReadFile(FilePath)
		if err != nil {
			return nil, err
		}

		// 生成文件相对路径
		RelativePath, err := filepath.Rel(FolderPath, FilePath)
		if err != nil {
			return nil, err
		}

		// 计算CRC32
		CRC := crc32.ChecksumIEEE(FileData)

		// 添加到catalog
		CatalogData.Files = append(CatalogData.Files, CatalogFile{
			Path: filepath.ToSlash(RelativePath),
			Size: int64(len(FileData)),
			CRC:  CRC,
		})
	}

	return &CatalogData, nil
}
