# beego项目通过域名访问实现

## mingxinwei from bupt Email:mxwei@bupt.edu.cn

本文旨在实现在digital ocean平台下，centos7.6环境搭建的beego的demo实现，可以通过域名访问到该beego项目，类似个人博客的搭建。用到了Cloudflare作为DNS解析服务器，并在freenom上注册了域名。

## Windows下的本地实现

### 踩坑：beego官网上的

```
go get github.com/beego/beego/v2@v2.0.0
在powershell报错： go: cannot use path@version syntax in GOPATH mode
```
###  原因:

go get 命令在Go1.11后，可以用此命令来获取依赖的特定版本，可以用来升级和降级依赖。会自动修改 go.mod 文件，而且依赖的依赖版本号也可能会变。在 go.mod 中使用 exclude 排除的包，不能 go get 下来。

与以前不同的是，新版 go get 可以在末尾加 @ 符号，用来指定版本。go get 命令需在go.mod同级目录下执行，否则会报出错误go: cannot use path@version syntax in GOPATH mode。而且在使用go get下载依赖时，要求仓库必须用 vX.Y.Z 格式打 tag，以下是简单罗列的匹配规则。
```
go get github.com/gorilla/mux    # 匹配最新的一个 tag
go get github.com/gorilla/mux@latest    # 和上面一样
go get github.com/gorilla/mux@v1.6.2    # 匹配 v1.6.2
go get github.com/gorilla/mux@e3702bed2 # 匹配 v1.6.2
go get github.com/gorilla/mux@c856192   # 匹配 c85619274f5d
go get github.com/gorilla/mux@master    # 匹配 master 分支
```
来源:<https://www.cnblogs.com/smallleiit/p/12493404.html>

### 安装bee的方法：

找到一个go的``中文镜像``通过设置进行换源类似操作，一切都丝滑了起来！！！

```
$ go env -w GO111MODULE=on
$ go env -w GOPROXY=https://goproxy.cn,direct
$ go get -v github.com/beego/bee
```
[来源1](https://goproxy.cn/)
       [来源2](https://github.com/goproxy/goproxy.cn/blob/master/README.zh-CN.md)

### 安装beego/bee v2方法:
```
go get -u github.com/beego/beego/v2
go get -u github.com/beego/bee/v2
```
### 使用bee创建新项目
进入E:\GoPATH\src下，执行bee new webApp
![image-20210220235949192](https://raw.githubusercontent.com/nickmxxx/pic-bed/main/img/image-20210220235949192.png)

进入到webapp文件夹

![image-20210221000603079](https://raw.githubusercontent.com/nickmxxx/pic-bed/main/img/image-20210221000603079.png)

执行结果：

![image-20210221000403491](https://raw.githubusercontent.com/nickmxxx/pic-bed/main/img/image-20210221000403491.png)


### 使用域名访问beego项目
在本地执行bee run操作（上述解决方案）
	   或者使用pack方法将其打包成exe文件，使其在后台运行

```
bee pack -be GOOS=windows
```

![image-20210221224028639](https://raw.githubusercontent.com/nickmxxx/pic-bed/main/img/image-20210221224028639.png)
###  安装nginx 进行反向代理实现域名访问beego项目
[官方下载地址](http://nginx.org/en/download.html)
       下载后，将压缩包解压到 C:\Tools\Nginx Nginx目录所在的路径中不要有中文注意字符，也不建议有空格。
		启动nginx进行反向代理

```
C: && cd C:\tools\Nginx
start nginx
```
打开浏览器进入输入localhost出现nginx界面即为安装成功

![image-20210221224407527](https://raw.githubusercontent.com/nickmxxx/pic-bed/main/img/image-20210221224407527.png)

记录一下nginx常用命令

| nginx -h        | 查看帮助信息                                  |
| --------------- | --------------------------------------------- |
| nginx -v        | 查看Nginx版本                                 |
| nginx -s stop   | 停用Nginx                                     |
| nginx -s quit   | 优雅的停用Nginx（处理完正在进行中请求后停用） |
| nginx -s reload | 重新加载配置，并优雅的重启进程                |
| nginx -s reopen | 重启日志文件                                  |

### 编辑本地hosts文件 进行DNS域名解析
使127.0.0.1解析到mxxchive.tk这个域名上
```
127.0.0.1 mxxchive.tk
```

### 编辑nginx.conf 添加server

    server {
        listen       80;
        server_name  demo1.test; # 改为你自己的域名
     
        charset utf-8;
     
        location /(css|js|fonts|img)/ {
            access_log off;
            expires 1d;
     
            root "E:/Gopath/src/webApp/static"; # 自己的静态文件夹
            try_files $uri @backend;
        }
     
        location / {
            try_files /_not_exists_ @backend;
        }
     
        location @backend {
            proxy_set_header X-Forwarded-For $remote_addr;
            proxy_set_header Host            $http_host;
     
            proxy_pass http://127.0.0.1:8080;
        }
    }


### 最终结果

通过浏览器访问mxxchive.tk 成功在本地浏览器``通过nginx反代``访问到了beego项目，但在ipad及手机上无法访问，原因是在域名解析的时候只是针对了本地的127.0.0.1，并不涉及到外网的操作，因此无法从外网访问到项目。

![image-20210222205217045](https://raw.githubusercontent.com/nickmxxx/pic-bed/main/img/image-20210222205217045.png)

### Centos下环境配置
将beego项目打包到digitalocean云服务器上（因为还有一部分费用没花完2333333），并通过cloudflare进行DNS解析，实现外网访问到beego项目

在digital ocean上租用centos7服务器，并首先搭建环境安装go、beego、nginx

### centos安装go环境

进入到centos，首先看有没有wget命令
```
rpm -qa|grep wget  # 有的话会出现版本
yum install -y wget # 安装wget
wget https://dl.google.com/go/go1.16.linux-amd64.tar.gz # wget执行下载
```
执行ls命令可以看到下载的go安装包，并执行安装命令

```
tar -C /usr/local -zxvf go1.16.linux-amd64.tar.gz
```

![image-20210222084602548](https://raw.githubusercontent.com/nickmxxx/pic-bed/main/img/image-20210222084602548.png)

配置环境变量，进入profile中编辑
```
vim /etc/profile # 在最末尾添加
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
export GOPATH=/data/gowork
source /etc/profile # 刷新环境变量

go version # 打印版本信息
```

![image-20210222085125456](https://raw.githubusercontent.com/nickmxxx/pic-bed/main/img/image-20210222085125456.png)

至此，centos下的go配置完成

### 安装v2版本的beego及bee工具

```
go get -u github.com/beego/beego/v2
go get -u github.com/beego/bee/v2
```

此处不再赘述过程，但要注意此处我们用的是digitalocean的国外服务器，因此不需要配置goproxy，国外下载东西就是快，丝滑！！！

### 安装nginx

需要首先安装依赖软件。

```
yum install -y gcc-c++ # 安装 GCC
yum install -y pcre pcre-devel # 安装 PCRE
yum install -y zlib zlib-devel # 安装 zlib
yum install -y openssl openssl-devel # 安装 OpenSSL
```

如果懒得每一步进行确认 在install 后面加个-y即可

下载最新版的nginx

```
wget http://nginx.org/download/nginx-1.19.7.tar.gz
tar -xvf nginx-1.19.7.tar.gz # 解压文件，与go安装类似
cd nginx-1.19.7 # 进入文件夹，从源码安装
./configure # 生成makefile 执行make命令
make
make install
```

将目录切换到：/usr/local/nginx/sbin 执行

```
./nginx # 启动nginx
./nginx -s stop # 关闭退出 Nginx
./nginx -s quit
./nginx -s reload # 重启 Nginx
ps aux|grep nginx # 查看 Nginx 进程
```

此时在浏览器输入我们的digitalocean的ip地址即能够看到nginx界面，即nginx配置成功。

![image-20210222091721295](https://raw.githubusercontent.com/nickmxxx/pic-bed/main/img/image-20210222091721295.png)

为了方便以后的操作（不用每一次都进入nginx的文件夹），我们配置一下nginx的环境变量。

```
vim /etc/profile # 配置环境变量
export NGINX_HOME=/usr/local/nginx # 这个是自己的nginx的文件地址
export PATH=$PATH:$NGINX_HOME/sbin
source /etc/profile # 重新加载环境
```

输入nginx -v，可以看到nginx的版本为1.19.7，即成功

![image-20210222093030324](https://raw.githubusercontent.com/nickmxxx/pic-bed/main/img/image-20210222093030324.png)

## 将项目移植到服务器的centos上

### 项目迁移

将demo项目打包到Linux服务器平台,生成了webApp.tar.gz的文件

```
bee pack -be GOOS=linux
```

将文件传到 Centos 系统，解压传到服务器的二进制包

```
mkdir /home/webapp # 新建webapp文件夹并解压
tar -xvf webApp.tar.gz -C /home/webapp # 解压到此文件夹
```

给 webapp 这个目录权限

```
chmod -R 777 /home/webapp
chmod +x /home/webapp
```

切换到解压的应用根目录并运行

```
cd /home/webapp
./webapp 
```

![image-20210222110703474](https://raw.githubusercontent.com/nickmxxx/pic-bed/main/img/image-20210222110703474.png)

可以看到，后台已经在执行webApp的应用，并关联在8080端口

```
netstat -ntlp
```

![image-20210222110759067](https://raw.githubusercontent.com/nickmxxx/pic-bed/main/img/image-20210222110759067.png)

同时可以看到，在浏览器输入服务器ip地址:8080端口，可以看到beego项目

![image-20210222111353073](https://raw.githubusercontent.com/nickmxxx/pic-bed/main/img/image-20210222111353073.png)

可以看到终端输出的连接建立

![image-20210222111507986](https://raw.githubusercontent.com/nickmxxx/pic-bed/main/img/image-20210222111507986.png)

Unix/Linux下一般比如想让某个程序在后台运行，很多都是使用 & 在程序结尾来让程序自动运行。

如要运行webapp在后台：

```
./webApp &
```

但是如果终端关闭，那么程序也会被关闭。但是为了能够后台运行，那么我们就可以使用nohup这个命令。

若提示：

```
[~]$ appending output to nohup.out
```

证明运行成功，同时把程序运行的输出信息放到当前目录的 nohup.out 文件中去。

###  配置nohup
#### 后台执行
出现问题：没有nohup命令

```
nohup: command not found
```

[解决](https://blog.csdn.net/leisure_life/article/details/80533492)

```
which nohup # 查看本地的nohup
/usr/bin/nohup # 我的nohup位置
Cd # 切换到根目录
vim .bash_profile # 配置环境文件 在后边加上:/usr/bin，保存，退出
source ~/.bash_profile # 刷新环境变量
nohup --version # 查看版本 有显示即为正确
```

![image-20210222103503461](https://raw.githubusercontent.com/nickmxxx/pic-bed/main/img/image-20210222103503461.png)

进入webapp文件夹下并执行nohup命令

```
cd /home/webapp
nohup ./webApp &
```

出现问题：

```
nohup: ignoring input and appending output to ‘nohup.out’
```

[解决方案](https://blog.csdn.net/jiangyu1013/article/details/81476184)

```
nohup ./webApp >/dev/null 2>/dev/null&
nohup ./webApp >/path/to/file
```

0、1和2分别表示标准输入、标准输出和标准错误信息输出，可以用来指定需要重定向的标准输入或输出。

在一般使用时，默认的是标准输出，既1。当我们需要特殊用途时，可以使用其他标号。例如，将某个程序的错误信息输出到log文件中：./program 2>log。这样标准输出还是在屏幕上，但是错误信息会输出到log文件中。

#### 结束nohup后台任务

```
netstat -tpnul # 查看有哪些进程正在执行
```

![image-20210222144324004](https://raw.githubusercontent.com/nickmxxx/pic-bed/main/img/image-20210222144324004.png)

```
ps aux |grep webApp # 加载某个进程查看pid
kill -9 2681 # 逐个杀死webApp进程
kill -9 2712
```

![image-20210222144527053](https://raw.githubusercontent.com/nickmxxx/pic-bed/main/img/image-20210222144527053.png)

![image-20210222145618784](https://raw.githubusercontent.com/nickmxxx/pic-bed/main/img/image-20210222145618784.png)

可以看到后台进程webApp已经被killed

## DNS解析及域名绑定

### Cloudflare绑定域名服务器，实现DNS地址解析

进入DNS管理，我这里使用的是test开头，其他的名称也都可以，但注意要把代理状态设置为``仅限DNS，即为灰色``。

![image-20210222164518053](https://raw.githubusercontent.com/nickmxxx/pic-bed/main/img/image-20210222164518053.png)

项目用的是freenom注册的域名，其自带DNS的解析，要将其换成Cloudflare的，具体操作见该[博客](https://shikey.com/2020/05/14/cloudfare-freenom.html)写的很清楚：

但注意Cloudflare提供的名称服务器与博客有出入，以Cloudflare的为准。

![image-20210222154920692](https://raw.githubusercontent.com/nickmxxx/pic-bed/main/img/image-20210222154920692.png)

是否成功可以观察Cloudflare的状态，如果显示有效即为成功，如下面的是成功的(配置完成后可能需要1小时或以上的时间)

![image-20210222155155168](https://raw.githubusercontent.com/nickmxxx/pic-bed/main/img/image-20210222155155168.png)

![image-20210222155448927](https://raw.githubusercontent.com/nickmxxx/pic-bed/main/img/image-20210222155448927.png)

另外可以用站长工具:  <https://ping.chinaz.com> 来进行ping该域名 demo.mxxchive.tk，如果可以看到ping出来相应的ip地址即为成功，注意此处前面的test一定要加上去，否则就解析不出来。

![image-20210222210612864](https://raw.githubusercontent.com/nickmxxx/pic-bed/main/img/image-20210222210612864.png)

### 配置nginx

```
chmod -R 777 /usr/local/nginx/conf/ # 给予权限
chmod +x /usr/local/nginx/conf/
vim /usr/local/nginx/conf/nginx.conf # nginx.conf为配置目录
```

在nginx.conf中添加server配置

```
server {
    listen       80;
    server_name  demo.mxxchive.tk; # 改为你自己的域名
 
    charset utf-8;
 
    location /(css|js|fonts|img)/ {
        access_log off;
        expires 1d;
 
        root "/home/webapp/static";
        try_files $uri @backend;
    }
 
    location / {
        try_files /_not_exists_ @backend;
    }
 
    location @backend {
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header Host            $http_host;
        proxy_pass http://161.35.234.230:8080; # 千万注意这里是服务器的ip地址而非本地地址
    }
}

```

一定要在每次执行前检测运行nginx是否通过，如果通过再reload重新加载

```
systemctl start nginx.service # 启动nginx
systemctl enable nginx.service  # 设为开机启动
nginx -t   #检查配置文件的语法
nginx -s reload    #重新加载配置文件
```



![image-20210222160725231](https://raw.githubusercontent.com/nickmxxx/pic-bed/main/img/image-20210222160725231.png)

可以看到语法无误，启动nginx

启动nginx出现：Failed to start nginx.service: Unit not found.

[解决方法](https://www.jianshu.com/p/1ca5a62df1a9?from=singlemessage)

### 最终效果

通过demo.mxxchive.tk可以访问到当前的beego的demo项目

![image-20210222211609494](https://raw.githubusercontent.com/nickmxxx/pic-bed/main/img/image-20210222211609494.png)

并可以在后台看到记录

![image-20210222211713691](https://raw.githubusercontent.com/nickmxxx/pic-bed/main/img/image-20210222211713691.png)

## 总结

本次实现了将beego项目demo进行域名解析与绑定，首先再windows平台下进行了模拟实现，随后将其打包好的文件迁移到Centos环境。对我来说最难的一点在于nginx的server配置，总是使用本地的127.0.0.1进行相应的匹配。但后来查阅相关知识发现应该将其设置为服务器的ip地址，这样才是连接到外网，```proxy_pass http://161.35.234.230:8080;```才是正确写法并能够通过Cloudflare做DNS解析，并最终实现了整个流程。