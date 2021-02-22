# beego环境配置，和web_v1运行

### 运行环境

m1 macbook air

![image-20210222120321698](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210222120321698.png)

golang 版本

go1.16.darwin-arm64

运行IDE：goland

beego版本

```go
require github.com/beego/beego/v2 v2.0.1
require github.com/smartystreets/goconvey v1.6.4
```

# 1、下载go语言安装包并安装

![image-20210222120652613](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210222120652613.png)

一定要注意架构，这台mac是使用的m1芯片所以要下载arm版本不是amd版本

在mac上的安装一路默认即可，环境变量配置为：在～目录下的.bash_profile文件中添加：

```bash
export GOROOT=/usr/local/go
export GOPATH=/Users/～/go
export GOBIN=$GOROOT/bin
export PATH=$PATH:$GOBIN
export PATH=$PATH:$GOPATH/bin
```

需要注意的是，bee工具在命令行中的使用需要

```bash
export PATH=$PATH:$GOPATH/bin
```

大部分教程中都没有提到。

# 2、在jetbrains官网下载goland安装包并下载安装

一路默认即可

# #3、beego和bee的使用

mingxing wei 在文章“beego项目部署和域名绑定”中说的已经很清楚，在这里补充一些

beego项目原作者为astaxie，在github上有beego仓库，但是这是v1版本的beego，目前已经放弃维护，新版本为v2版本，目前在github上的beego账号下的beego仓库中。

![image-20210222133221139](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210222133221139.png)

该项目使用v2版本仓库。另外，安装bee工具也应使用v2仓库

安装时使用：

```bash
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.cn,direct
$ go get -v github.com/beego/beego/v2
$ go get -v github.com/beego/bee/v2
```

在创建新项目时使用

```bash
bee new web_v1
```

进入web_v1后执行：

```bash
go mod tidy
```

该命令的含义是，检查该项目中的依赖包是否有重复利用，并且主动安装没用安装的包。该命令来自go语言自带的包管理工具“go module”，使用go mod安装的包在`GOPATH/pkg/mod`目录下，相当于缓存。

每一个文件夹下需要通过

```
go mod init
```

生成go.mod文件，使用bee命令后项目会自动生成该文件，所有的go mod命令均是围绕着这个文件进行维护的，简单记录一下go mod的命令，官网[go module文档](https://blog.golang.org/using-go-modules)和[技术文档](https://golang.org/doc/tutorial/create-module)

​	download  download modules to local cache

​	edit    edit go.mod from tools or scripts

​	graph    print module requirement graph

​	init    initialize new module in current directory

​	tidy    add missing and remove unused modules

​	vendor   make vendored copy of dependencies

​	verify   verify dependencies have expected content

​	why     explain why packages or modules are needed

最后执行：

```bash
bee run
```

