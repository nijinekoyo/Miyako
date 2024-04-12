/*
 * @Author: nijineko
 * @Date: 2024-04-12 17:23:26
 * @LastEditTime: 2024-04-12 17:23:30
 * @LastEditors: nijineko
 * @Description: 文件夹工具
 * @FilePath: \Miyako\tools\file\folder.go
 */
package file

import (
	"os"
	"path/filepath"
)

/**
 * @description: 遍历出文件夹内所有文件路径
 * @param {string} DirPth
 * @return {[]string} 文件路径
 * @return {error} 错误
 */
func GetPaths(DirPth string) ([]string, error) {
	Dir, err := os.ReadDir(filepath.Clean(DirPth))
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)

	var Dirs []string
	var Files []string
	for _, fi := range Dir {
		// 目录, 继续递归遍历
		if fi.IsDir() {
			Dirs = append(Dirs, filepath.Clean(DirPth+PthSep+fi.Name()))
		} else {
			Files = append(Files, filepath.Clean(DirPth+PthSep+fi.Name()))
		}
	}

	// 读取子目录下文件
	for _, Table := range Dirs {
		TempFiles, _ := GetPaths(Table)
		for _, TempFile := range TempFiles {
			Files = append(Files, filepath.Clean(TempFile))
		}
	}

	return Files, nil
}
