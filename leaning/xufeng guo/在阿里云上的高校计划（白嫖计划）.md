---
title: 在阿里云上的高校计划（白嫖计划）
date: 2021-02-24 11:30:22
tags: [golang, beginners]
categories:
    - coding
    - golang
index_img: https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210211140605581.png
banner_img: https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/golang-hacker (1).jpg
---

# 在阿里云上的高校计划（白嫖计划）部署beego

阿里云高校计划地址

https://developer.aliyun.com/adc/student/

首先需要进行实名认证和学生认证

和别家不同的是，阿里云的学生服务器在到期前一个月可以通过学习课程免费续费。

接下来在服务器上部署一个简单的beego项目

## 1、登陆服务器

过程略

## 2、配置golang环境

```
 wget https://golang.google.cn/dl/go1.16.linux-amd64.tar.gz
 # 解压文件 
 tar xfz go1.16.linux-amd64.tar.gz -C /usr/local
```

配置环境变量

```
#修改~/.bashrc
vim ~/.bashrc
#添加Gopath路径
export GOROOT=/usr/local/go  
export GOPATH=$PATH:$GOROOT/bin 
export PATH=$GOPATH/bin:$PATH
export PATH=$GOROOT/bin:$PATH
# 激活配置
source ~/.bashrc
```

安装beego

```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go get -v github.com/beego/beego/v2
go get -v github.com/beego/bee/v2
```

使用bee工具生成beego项目文件

```bash
bee new web_test
```

在web_test下利用go mod 整理依赖

```
go mod tidy
```

在web_test下运行命令

```
bee run
```

但是此时并不能通过ip加端口的方式直接访问，需要配置一下安全组：

![image-20210224125900147](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210224125900147.png)

点击当前策略组：

![image-20210224125955441](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210224125955441.png)

添加如下策略组即可正常访问

![image-20210224130017506](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210224130017506.png)

