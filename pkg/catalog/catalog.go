/*
 * @Author: nijineko
 * @Date: 2024-04-12 17:05:21
 * @LastEditTime: 2024-04-12 17:35:53
 * @LastEditors: nijineko
 * @Description: catalog管理
 * @FilePath: \Miyako\pkg\catalog\catalog.go
 */
package catalog

import "encoding/json"

type CatalogFile struct {
	Path string `json:"path"` // 路径
	Size int64  `json:"size"` // 大小 (bytes)
	CRC  uint32 `json:"crc"`  // CRC32
}

type Catalog struct {
	BlockFile string        `json:"block_file"` // 文件块地址
	Files     []CatalogFile `json:"files"`      // 文件列表
}

// 序列化Catalog
func (c *Catalog) Marshal() ([]byte, error) {
	return json.Marshal(c)
}

// 反序列化Catalog
func Unmarshal(Data []byte) (*Catalog, error) {
	CatalogData := &Catalog{}

	err := json.Unmarshal(Data, CatalogData)
	if err != nil {
		return nil, err
	}
	return CatalogData, nil
}
