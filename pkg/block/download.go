/*
 * @Author: nijineko
 * @Date: 2024-04-12 19:58:14
 * @LastEditTime: 2024-04-12 22:27:45
 * @LastEditors: nijineko
 * @Description: 下载文件块
 * @FilePath: \Miyako\pkg\block\download.go
 */
package block

import (
	"errors"
	"hash/crc32"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/nijinekoyo/miyako/pkg/catalog"
)

type Download struct {
	BlockResponse *http.Response
	CatalogData   *catalog.Catalog
	SavePath      string
	TotalSize     int64
}

/**
 * @description: 创建一个新的块下载
 * @param {string} BlockURL 文件块地址
 * @param {*catalog.Catalog} CatalogData Catalog数据
 * @param {string} SavePath 保存路径
 * @return {*Download} 下载对象
 */
func NewDownload(BlockURL string, CatalogData *catalog.Catalog, SavePath string) (*Download, error) {
	// 创建HTTP链接
	Client := &http.Client{
		Timeout: 0,
	}

	BlockRequest, err := http.NewRequest(http.MethodGet, BlockURL, nil)
	if err != nil {
		return nil, err
	}

	BlockResponse, err := Client.Do(BlockRequest)
	if err != nil {
		return nil, err
	}

	if BlockResponse.StatusCode != http.StatusOK {
		return nil, errors.New("Download block failed")
	}

	return &Download{
		BlockResponse: BlockResponse,
		CatalogData:   CatalogData,
		SavePath:      SavePath,
		TotalSize:     BlockResponse.ContentLength,
	}, nil
}

/**
 * @description: 开始下载文件块
 * @return {*error} 错误
 */
func (d *Download) Start(Progress func(WrittenSize int64, TotalSize int64)) error {
	// 遍历Catalog文件列表
	for _, FileInfo := range d.CatalogData.Files {
		FilePath := filepath.Join(d.SavePath + FileInfo.Path)

		// 创建文件夹
		err := os.MkdirAll(filepath.Dir(FilePath), os.ModePerm)
		if err != nil {
			return err
		}

		// 按照Size读取文件
		FileData := make([]byte, FileInfo.Size)
		LimitFileData := io.LimitReader(d.BlockResponse.Body, int64(FileInfo.Size))

		_, err = io.ReadFull(LimitFileData, FileData)
		if err != nil {
			return err
		}

		// 检查CRC32
		CRC := crc32.ChecksumIEEE(FileData)
		if CRC != FileInfo.CRC {
			return errors.New(FileInfo.Path + " CRC32 not match")
		}

		// 写入文件
		err = os.WriteFile(FilePath, FileData, 0644)
		if err != nil {
			return err
		}

		// 进度回调
		Progress(FileInfo.Size, d.TotalSize)
	}

	return nil
}

/**
 * @description: 关闭块下载
 */
func (d *Download) Close() {
	d.BlockResponse.Body.Close()
}
