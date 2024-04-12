/*
 * @Author: nijineko
 * @Date: 2024-04-12 17:37:47
 * @LastEditTime: 2024-04-12 17:42:50
 * @LastEditors: nijineko
 * @Description: catalog生成测试
 * @FilePath: \Miyako\pkg\catalog\generate_test.go
 */
package catalog

import "testing"

func TestGenerate(t *testing.T) {
	CatalogData, err := Generate("../../assets/test")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(CatalogData)
}
