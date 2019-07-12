# xcmdblog
#
go开发的markdown笔记系统

使用方法
1. 克隆项目到$GOPATH/src 目录
```shell
https://github.com/mkyue/xcmdblog.git
```
2. 安装dep依赖管理，确认$GOPATH/bin 目录加入到环境path
```shell
go get -u -v github.com/golang/dep/cmd/dep
```
3. 更新下载依赖
```
dep ensure -v
```
4. 基于beego开发， 测试通过 bee run 即可运行
