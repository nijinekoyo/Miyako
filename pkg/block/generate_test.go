/*
 * @Author: nijineko
 * @Date: 2024-04-12 18:01:54
 * @LastEditTime: 2024-04-12 18:01:55
 * @LastEditors: nijineko
 * @Description: 文件块生成测试
 * @FilePath: \Miyako\pkg\block\generate_test.go
 */
package block

import "testing"

func TestGenerate(t *testing.T) {
	CatalogData, err := Generate("../../assets/test")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(CatalogData)
}
