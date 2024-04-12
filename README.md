# Miyako
一个加速大量小文件下载的工具

此项目是源于之前玩的一个游戏，他会在更新版本时下载大量的小文件，但是没有任何优化所以速度非常的慢，故想到有没有办法优化一下大量小文件的下载场景  
简单构思后想到，既然下载速度慢是因为小文件太多，那就把他们都拼成一个大文件就好了，毕竟大文件的下载速度理论上可以跑满带宽！

项目的原理就是将大量的小文件直接链接成一个块文件，然后生成一份catalog，客户端通过catalog下载块文件并依据文件大小和路径在本地还原文件  
并且基于`HTTP Range`的特性，还可以实现断点重连和下载某个文件的功能

## 如何使用
1. 前往[Releases](https://github.com/nijinekoyo/miyako/releases)下载最新的客户端和服务端二进制构建
2. 部署服务端
   1. 将二进制文件放置到任意目录，并在同级目录创建一个`assets`文件夹
   2. 将需要加速下载的小文件放置到`assets`文件夹内的任意文件夹，例如`assets/example`
   3. 启动服务端，服务端启动时会自动按照`assets`目录下的所有子文件夹为单位创建打包文件，并为他们创建`catalog.json`
   4. 服务端提示`Server started at http://0.0.0.0:8080`即为启动完成
3. 通过客户端下载文件
   1. 将二进制文件放置到任意目录
   2. 输入`client <CatalogAddr>`开始下载文件，`<CatalogAddr>`为`catalog.json`的下载地址，服务端启动时会自动生成，假设你的服务端文件路径是`assets/example`，则catalog地址就为`http://localhost:8080/example/catalog.json`
   3. 下载的文件会自动保存在`download`文件夹

## 构建
构建需要Golang >= 1.22.2
1. 安装依赖
``` shell
go mod tidy
```
2. 构建服务端
``` shell
go build ./cmd/server/
```
3. 构建客户端
``` shell
go build ./cmd/client/
```

## 引用
引入包
```
go get -u github.com/nijinekoyo/miyako
```
包使用说明
1. pkg/catalog  
用于创建和解析catalog，API请查看[文档](https://pkg.go.dev/github.com/nijinekoyo/miyako/pkg/catalog)
2. pkg/block  
用于生成文件块和下载文件块，API请查看[文档](https://pkg.go.dev/github.com/nijinekoyo/miyako/pkg/block)

## 许可
本项目基于`MIT License`协议分发