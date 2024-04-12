/*
 * @Author: nijineko
 * @Date: 2024-04-12 17:30:41
 * @LastEditTime: 2024-04-12 17:33:00
 * @LastEditors: nijineko
 * @Description: 生成随机字符串
 * @FilePath: \Miyako\tools\random\string.go
 */
package random

import (
	"math/rand"
	"time"
)

/**
 * @description: 生成随机字符串
 * @param {int} Length 长度
 * @param {int} Level 等级
 * @return {*string} 随机字符串
 */
func String(Length int, Level int) string {
	var Charset string = "1234567890"
	switch Level {
	case 2:
		Charset += "abcdefghijklmnopqrstuvwxyz"
	case 3:
		Charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	Rand := rand.New(rand.NewSource(time.Now().UnixNano() + int64(rand.Intn(100))))
	var Result []byte
	for i := 0; i < Length; i++ {
		Result = append(Result, Charset[Rand.Intn(len(Charset))])
	}
	return string(Result)
}
