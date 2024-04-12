/*
 * @Author: nijineko
 * @Date: 2024-04-12 19:53:04
 * @LastEditTime: 2024-04-12 19:55:11
 * @LastEditors: nijineko
 * @Description: 获取Catalog并解析
 * @FilePath: \Miyako\cmd\client\catalog.go
 */
package main

import (
	"io"
	"net/http"

	"github.com/nijinekoyo/miyako/pkg/catalog"
)

/**
 * @description: 获取Catalog
 * @param {string} Address Catalog地址
 * @return {*catalog.Catalog} Catalog
 */
func getCatalog(Address string) (*catalog.Catalog, error) {
	// 下载Catalog
	CatalogResponse, err := http.Get(Address)
	if err != nil {
		return nil, err
	}

	// 读取返回数据
	CatalogData, err := io.ReadAll(CatalogResponse.Body)
	if err != nil {
		return nil, err
	}

	// 解析Catalog
	Catalog, err := catalog.Unmarshal(CatalogData)
	if err != nil {
		return nil, err
	}

	return Catalog, nil
}
