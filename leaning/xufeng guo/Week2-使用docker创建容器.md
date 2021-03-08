### 使用docker创建容器

标题1～4为基本操作总结，标题5为镜像建立过程

## 1、安装docker

更新包索引

```
sudo apt-get update
```

为了使apt可以通过https使用Repository，先安装以下包：

```
sudo apt-get install -y apt-transport-https ca-certificates curl software-properties-common

```

添加Docker官方GPG密钥：

```
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

```

检查GPG Key信息是否正确：

```
sudo apt-key fingerprint 0EBFCD88
```

添加源信息：

将源信息直接写入/etc/apt/sources.list

```
sudo add-apt-repository \
    "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
    $(lsb_release -cs) \
    stable"

```

再更新下apt包索引：

```html
sudo apt-get update
```

确认Docker的源信息是否正确, 新的源是否添加成功：

```sql
$ sudo apt-cache madison docker-ce
```

- 如果需要安装某个指定版本的Docker CE，而不是最新版本，可用下面命令列出可用的版本列表信息：

```sql
$ sudo apt-cache madison docker-ce
```

第一列是包名，第二列是版本字符串，第三列是存储库名称，它标识出包来自哪个存储库，以及扩展它的稳定性级别。通一下命令安装指定版本的包：

```sql
$ sudo apt-get install docker-ce=[版本字符串]
```

- 查看Docker安装版本详细信息：

```html
sudo docker version
```

- 查看docker服务是否启动：

```ruby
sudo systemctl status docker
```

- 如果Docker未启动，则启动Docker：

```sql
$ sudo systemctl start docke
```

- 运行Hello World，校验Docker是否安装成功：

```ruby
sudo docker run hello-world
```

## 2、启动与停止

```
　　# 启动docker
　　sudo service docker start

　　# 停止docker
　　sudo service docker stop

　　# 重启docker
　　sudo service docker restart
```



## 3、下载docker镜像

### 什么是Docker镜像

　　Docker 镜像是由文件系统叠加而成(是一种文件的存储形式)。最底端是一个文件引 导系统，即 bootfs，这很像典型的 Linux/Unix 的引导文件系统。Docker 用户几乎永远不会和 引导系统有什么交互。实际上，当一个容器启动后，它将会被移动到内存中，而引导文件系 统则会被卸载，以留出更多的内存供磁盘镜像使用。Docker容器启动是需要一些文件的， 而这些文件就可以称为 Docker

![image-20210301172351206](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210301172351206.png)

**Docker 把应用程序及其依赖，打包在 image 文件里面。**只有通过这个文件，才能生成 Docker 容器。image 文件可以看作是容器的模板。Docker 根据 image 文件生成容器的实例。同一个 image 文件，可以生成多个同时运行的容器实例。

​	image 是二进制文件。实际开发中，一个 image 文件往往通过继承另一个 image 文件，加上一些个性化设置而生成。举例来说，你可以在 Ubuntu 的 image 基础上，往里面加入 Apache 服务器，形成你的 image。

　image 文件是通用的，一台机器的 image 文件拷贝到另一台机器，照样可以使用。一般来说，为了节省时间，我们应该尽量使用别人制作好的 image 文件，而不是自己制作。即使要定制，也应该基于别人的 image 文件进行加工，而不是从零开始制作。

　为了方便共享，image 文件制作完成后，可以上传到网上的仓库。Docker 的官方仓库 [Docker Hub](https://hub.docker.com/) 是最重要、最常用的 image 仓库。此外，出售自己制作的 image 文件也是可以的。

可以直接在官网寻找镜像：

https://hub.docker.com/search?q=&type=image

### 搜索镜像

```
docker search xxx
```

### 更新镜像 提交镜像副本

```bash
docker commit -m="message" -a="author" IMAGEID username/imagename:tag
```

### 设置镜像标签

```
docker tag IMAGEID username/imagename:newTag
```



### 列出本地镜像

```
　　docker image ls
```

- REPOSITORY：镜像所在的仓库名称
- TAG：镜像标签
- IMAGEID：镜像ID
- CREATED：镜像的创建日期(不是获取该镜像的日期)
- SIZE：镜像大小

Docker维护了镜像仓库，分为共有和私有两种，共有的官方仓库[Docker Hub(https://hub.docker.com/)](https://hub.docker.com/)是最重要最常用的镜像仓库。私有仓库（Private Registry）是开发者或者企业自建的镜像存储库，通常用来保存企业 内部的 Docker 镜像，用于内部开发流程和产品的发布、版本控制。

要想获取某个镜像，我们可以使用pull命令，从仓库中拉取镜像到本地，如

```shell
　　docker image pull library/hello-world
```

　　上面代码中，`docker image pull`是抓取 image 文件的命令。`library/hello-world`是 image 文件在仓库里面的位置，其中`library`是 image 文件所在的组，`hello-world`是 image 文件的名字。

　　由于 Docker 官方提供的 image 文件，都放在[`library`](https://hub.docker.com/r/library/)组里面，所以它的是默认组，可以省略。因此，上面的命令可以写成下面这样。

```shell
　　docker image pull hello-world
```

### 删除镜像

```shell
　　docker image rm 镜像名或镜像id
```

　　如

```shell
　　docker image rm hello-world
```

## 4、docker容器操作

### 创建容器

```shell
　　docker run [option] 镜像名 [向启动容器中传入的命令]
```

常用可选参数说明：

- -i 表示以“交互模式”运行容器
- -t 表示容器启动后会进入其命令行。加入这两个参数后，容器创建就能登录进去。即 分配一个伪终端。
- --name 为创建的容器命名
- -v 表示目录映射关系(前者是宿主机目录，后者是映射到宿主机上的目录，即 宿主机目录:容器中目录)，可以使 用多个-v 做多个目录或文件映射。注意:最好做目录映射，在宿主机上做修改，然后 共享到容器上。
- -d 在run后面加上-d参数,则会创建一个守护式容器在后台运行(这样创建容器后不 会自动登录容器，如果只加-i -t 两个参数，创建后就会自动进去容器)。
- -p 表示端口映射，前者是宿主机端口，后者是容器内的映射端口。可以使用多个-p 做多个端口映射
- -e 为容器设置环境变量
- --network=host 表示将主机的网络环境映射到容器中，容器的网络与主机相同

#### 交互式容器

　　例如，创建一个交互式容器，并命名为myubuntu

```shell
　　docker run -it --name=myubuntu ubuntu /bin/bash
```

#### 守护式容器

　　创建一个守护式容器:如果对于一个需要长期运行的容器来说，我们可以创建一个守护式容器。在容器内部exit退出时，容器也不会停止。

```shell
　　docker run -dit --name=myubuntu2 ubuntu
```

### 进入已运行的容器

```shell
　　docker exec -it 容器名或容器id 进入后执行的第一个命令
```

　　如

```shell
　　docker exec -it myubuntu2 /bin/bash
```

### 查看容器

```shell
　　# 列出本机正在运行的容器
　　docker container ls

　　# 列出本机所有容器，包括已经终止运行的
　　docker container ls --all
```

### 停止与启动容器

```shell
　　# 停止一个已经在运行的容器
　　docker container stop 容器名或容器id

　　# 启动一个已经停止的容器
　　docker container start 容器名或容器id

　　# kill掉一个已经在运行的容器
　　docker container kill 容器名或容器id
```

### 退出容器命令行

```
exit
```



### 删除容器

```shell
　　docker container rm 容器名或容器id
```

### 容器保存为镜像

```shell
　　docker commit 容器名 镜像名
```

### 镜像迁移与备份

我们可以通过save命令将镜像打包成文件，拷贝给别人使用

```shell
　　docker save -o 保存的文件名 镜像名
```

　　如

```shell
　　docker save -o ./ubuntu.tar ubuntu
```

　　在拿到镜像文件后，可以通过load方法，将镜像加载到本地

```shell
　　docker load -i ./ubuntu.tar
```

## 5、针对chives项目的环境打包

参考：https://zhuanlan.zhihu.com/p/166672337

参考的这一篇是建立一个基于cuda的深度学习环境的容器，根据了解，环境打包不能打包本地已有环境，需要新建container然后在container中操作配置环境，编写dockerfile，最后打包为image镜像文件。

### 1）在本地安装docker软件

略过

### 2）在官网拉取最接近项目环境的image

由于在网上搜索到的beego环境都比较老，这里我们新建container并重新部署环境。

#### 拉取镜像

注意！官方源下载速度奇慢无比，换源：

国内镜像有

- docker中国官方

  ```
  https://registry.docker-cn.com
  ```

- 网易

  ```
  http://hub-mirror.c.163.com
  ```

- USTC

  ```
  http://docker.mirrors.ustc.edu.cn
  ```

- 阿里云　　

  需要注册一个阿里云用户,访问 https://cr.console.aliyun.com/#/accelerator 获取专属Docker加速器地址

  ```
  http://<你的ID>.mirror.aliyuncs.com
  ```

打开docker设置：

![image-20210302193325279](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210302193325279.png)

在配置文件中添加源：

![image-20210302193534970](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210302193534970.png)

注意json文件格式（加逗号）

这里拉取最新的ubuntu镜像

https://hub.docker.com/_/ubuntu

在终端输入

```
docker pull ubuntu
```

下载完成后在当前目录下执行：

```
docker image ls
```

查看本地下载的image

![image-20210302205743991](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210302205743991.png)

#### 新建container

通过刚刚下载的image建立对应的container

```bash
docker run -it --name=ubuntu_bee ubuntu /bin/bash
```

- -i 表示以“交互模式”运行容器
- -t 表示容器启动后会进入其命令行。加入这两个参数后，容器创建就能登录进去。即 分配一个伪终端。

`/bin/bash`表示进入bash 命令行

并使用`cat /etc/lsb-release`查看发行版本信息：

![image-20210302210514388](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210302210514388.png)

使用`exit`退出容器

使用

```
docker container stop ubuntu_bee
```

停止该container运行

使用

```
　　docker container rm ubuntu_bee
```

删除该容器

#### 配置ssh访问

在启动docker时映射ssh的`22`端口为`1234`（之后也可以在container中配置端口）

```
docker run -it -p 1234:22 --name=ubuntu_bee ubuntu /bin/bash
```

##### 在ubuntu_bee中配置ssh：

```
root@be4defdb43d7:/# apt-get update
root@be4defdb43d7:/# apt-get upgrade
root@be4defdb43d7:/# apt-get install vim
root@be4defdb43d7:/# apt-get install openssh-server
```

- 设置一个root密码，后面登陆会用到
  根据自己的情况设置一个密码。

```
root@be4defdb43d7:/# passwd 
```

修改配置文件
root@be4defdb43d7:/# vim /etc/ssh/sshd_config
1
注释这一行PermitRootLogin prohibit-password
添加一行PermitRootLogin yes

```
PermitRootLogin prohibit-password
PermitRootLogin yes
```

保存退出

重启ssh服务

```
root@be4defdb43d7:/# /etc/init.d/ssh restart
```

Restarting OpenBSD Secure Shell server sshd

在类putty工具中输入`127.0.0.1:1234`作为地址并输入账户密码

![image-20210302213106529](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210302213106529.png)

同样也可登陆sftp服务

![image-20210302213123336](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210302213123336.png)

### 3）配置golang环境

再次确认本ubuntu系统架构：

```
arch
```

![image-20210302214416777](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210302214416777.png)

架构为`aarch64`

参考https://stackoverflow.com/questions/31851611/differences-between-arm64-and-aarch64/47274698#47274698

>AArch64 is the 64-bit state introduced in the Armv8-A architecture (https://en.wikipedia.org/wiki/ARM_architecture#ARMv8-A). The 32-bit state which is backwards compatible with Armv7-A and previous 32-bit Arm architectures is referred to as AArch32. Therefore the GNU triplet for the 64-bit ISA is aarch64. The Linux kernel community chose to call their port of the kernel to this architecture arm64 rather than aarch64, so that's where some of the arm64 usage comes from.

> As far as I know the Apple backend for aarch64 was called arm64 whereas the LLVM community-developed backend was called aarch64 (as it is the canonical name for the 64-bit ISA) and later the two were merged and the backend now is called aarch64.So **AArch64 and ARM64 refer to the same thing.**

所以下载arm架构的golang：`go1.16.linux-arm64.tar.gz`

在ubuntu 中解压：

```
tar -xzf go1.16.linux-arm64.tar.gz  -C /usr/local
```

在`/root/.bashrc`中添加：

```bash
export GOROOT=/usr/local/go
export GOBIN=$GOROOT/bin
export PATH=$PATH:$GOBIN
```

在终端中查看go版本

![image-20210302220857827](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210302220857827.png)

安装成功

### 4）安装beego和bee

在终端输入

```bash
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go get -v github.com/beego/beego/v2
go get -v github.com/beego/bee/v2
```

新建beego项目并进入整理包依赖

```
bee new web
cd web
go mod tidy
```

运行beego

```
bee run
```

### 5）端口映射

此时beego框架运行在容器的8080端口，但是此时的8080端口并没有影射到宿主机的端口上，所以此时在宿主机的8080端口中访问不到任何东西。

```
docker run -it -p 1234:22 -p 8081:8080 --name=ubuntu_bee ubuntu /bin/bash
```



### 6）保存为image

```
docker commit 容器名 镜像名
```



### 7）从image运行container

```
docker run -it -p 1234:22 -p 8081:8080 --name=ubuntu_bee_container ubuntu_bee_image /bin/bash
```

要注意，每次运行起来后不知为何总要重启ssh服务才可以正常连接

```bash
/etc/init.d/ssh restart
```



### 8）docker官方教程image

```bash
docker pull docker/getting-started
```



```bash
docker run -d -p 80:80 docker/getting-started
```

## 6、dockerfile

以上步骤直接对容器本身进行操作，不管是安装还是删除，都会增大镜像的体积，所以为解决这一问题，使用dockerfile解决环境依赖问题

参考https://blog.csdn.net/bbwangj/article/details/82178774?utm_medium=distribute.pc_relevant.none-task-blog-baidujs_baidulandingword-8&spm=1001.2101.3001.4242

Dockerfile是一个包含用于组合映像的命令的文本文档。可以使用在命令行中调用任何命令。 Docker通过读取`Dockerfile`中的指令自动生成映像。

`docker build`命令用于从Dockerfile构建映像。可以在`docker build`命令中使用`-f`标志指向文件系统中任何位置的Dockerfile。

```shell
docker build -f /path/to/a/Dockerfile
```

>--add-host                   -- Add a custom host-to-IP mapping
>--build-arg                  -- Build-time variables
>--cache-from                 -- Images to consider as cache sources
>--cgroup-parent              -- Parent cgroup for the container
>--compress                   -- Compress the build context using gzip
>--cpu-period                 -- Limit the CPU CFS (Completely Fair Scheduler)
>--cpu-quota                  -- Limit the CPU CFS (Completely Fair Scheduler)
>--cpu-rt-period              -- Limit the CPU real-time period
>--cpu-rt-runtime             -- Limit the CPU real-time runtime
>--cpu-shares             -c  -- CPU shares (relative weight)
>--cpuset-cpus                -- CPUs in which to allow execution
>--cpuset-mems                -- MEMs in which to allow execution
>--disable-content-trust      -- Skip image verification
>--file                   -f  -- Name of the Dockerfile
>--force-rm                   -- Always remove intermediate containers
>--help                       -- Print usage
>--isolation                  -- Container isolation technology
>--label                      -- Set metadata for an image
>--memory                 -m  -- Memory limit
>--memory-swap                -- Total memory limit with swap
>--network                    -- Connect a container to a network
>--no-cache                   -- Do not use cache when building the image
>--pull                       -- Attempt to pull a newer version of the image
>--quiet                  -q  -- Suppress verbose build output
>--rm                         -- Remove intermediate containers after a success
>--shm-size                   -- Size of '/dev/shm' (format is '<number><unit>'
>--squash                     -- Squash newly built layers into a single new la
>--tag                    -t  -- Repository, name and tag for the image
>--target                     -- Set the target build stage to build.
>--ulimit                     -- ulimit options
>--userns                     -- Container user namespace

### dockerfile的基本结构

**基础镜像信息**、**维护者信息**、**镜像操作指令**和**容器启动时执行指令**，

’#’ 为 Dockerfile 中的注释。

### dockerfile 文件说明

Docker以从上到下的顺序运行Dockerfile的指令。为了指定基本映像，第一条指令必须是*FROM*。一个声明以`＃`字符开头则被视为注释。可以在Docker文件中使用`RUN`，`CMD`，`FROM`，`EXPOSE`，`ENV`等指令。

**在这里列出了一些常用的指令。**

**FROM：指定基础镜像，必须为第一个命令**



```
格式：
　　FROM <image>
　　FROM <image>:<tag>
　　FROM <image>@<digest>
示例：
　　FROM mysql:5.6
注：
　　tag或digest是可选的，如果不使用这两个值时，会使用latest版本的基础镜像
```



**MAINTAINER: 维护者信息**

```
格式：
    MAINTAINER <name>
示例：
    MAINTAINER Jasper Xu
    MAINTAINER sorex@163.com
    MAINTAINER Jasper Xu <sorex@163.com>
```

**RUN：构建镜像时执行的命令**



```
RUN用于在镜像容器中执行命令，其有以下两种命令执行方式：
shell执行
格式：
    RUN <command>
exec执行
格式：
    RUN ["executable", "param1", "param2"]
示例：
    RUN ["executable", "param1", "param2"]
    RUN apk update
    RUN ["/etc/execfile", "arg1", "arg1"]
注：
　　RUN指令创建的中间镜像会被缓存，并会在下次构建中使用。如果不想使用这些缓存镜像，可以在构建时指定--no-cache参数，如：docker build --no-cache
```



**ADD：将本地文件添加到容器中，tar类型文件会自动解压(网络压缩资源不会被解压)，可以访问网络资源，类似wget**



```
格式：
    ADD <src>... <dest>
    ADD ["<src>",... "<dest>"] 用于支持包含空格的路径
示例：
    ADD hom* /mydir/          # 添加所有以"hom"开头的文件
    ADD hom?.txt /mydir/      # ? 替代一个单字符,例如："home.txt"
    ADD test relativeDir/     # 添加 "test" 到 `WORKDIR`/relativeDir/
    ADD test /absoluteDir/    # 添加 "test" 到 /absoluteDir/
```



**COPY：功能类似ADD，但是是不会自动解压文件，也不能访问网络资源**

**CMD：构建容器后调用，也就是在容器启动时才进行调用。**



```
格式：
    CMD ["executable","param1","param2"] (执行可执行文件，优先)
    CMD ["param1","param2"] (设置了ENTRYPOINT，则直接调用ENTRYPOINT添加参数)
    CMD command param1 param2 (执行shell内部命令)
示例：
    CMD echo "This is a test." | wc -
    CMD ["/usr/bin/wc","--help"]
注：
 　　CMD不同于RUN，CMD用于指定在容器启动时所要执行的命令，而RUN用于指定镜像构建时所要执行的命令。
```



**ENTRYPOINT：配置容器，使其可执行化。配合CMD可省去"application"，只使用参数。**



```
格式：
    ENTRYPOINT ["executable", "param1", "param2"] (可执行文件, 优先)
    ENTRYPOINT command param1 param2 (shell内部命令)
示例：
    FROM ubuntu
    ENTRYPOINT ["top", "-b"]
    CMD ["-c"]
注：
　　　ENTRYPOINT与CMD非常类似，不同的是通过docker run执行的命令不会覆盖ENTRYPOINT，而docker run命令中指定的任何参数，都会被当做参数再次传递给ENTRYPOINT。Dockerfile中只允许有一个ENTRYPOINT命令，多指定时会覆盖前面的设置，而只执行最后的ENTRYPOINT指令。
```



**LABEL：用于为镜像添加元数据**

```
格式：
    LABEL <key>=<value> <key>=<value> <key>=<value> ...
示例：
　　LABEL version="1.0" description="这是一个Web服务器" by="IT笔录"
注：
　　使用LABEL指定元数据时，一条LABEL指定可以指定一或多条元数据，指定多条元数据时不同元数据之间通过空格分隔。推荐将所有的元数据通过一条LABEL指令指定，以免生成过多的中间镜像。
```

**ENV：设置环境变量**



```
格式：
    ENV <key> <value>  #<key>之后的所有内容均会被视为其<value>的组成部分，因此，一次只能设置一个变量
    ENV <key>=<value> ...  #可以设置多个变量，每个变量为一个"<key>=<value>"的键值对，如果<key>中包含空格，可以使用\来进行转义，也可以通过""来进行标示；另外，反斜线也可以用于续行
示例：
    ENV myName John Doe
    ENV myDog Rex The Dog
    ENV myCat=fluffy
```



**EXPOSE：指定于外界交互的端口**



```
格式：
    EXPOSE <port> [<port>...]
示例：
    EXPOSE 80 443
    EXPOSE 8080
    EXPOSE 11211/tcp 11211/udp
注：
　　EXPOSE并不会让容器的端口访问到主机。要使其可访问，需要在docker run运行容器时通过-p来发布这些端口，或通过-P参数来发布EXPOSE导出的所有端口
```



**VOLUME：用于指定持久化目录**



```
格式：
    VOLUME ["/path/to/dir"]
示例：
    VOLUME ["/data"]
    VOLUME ["/var/www", "/var/log/apache2", "/etc/apache2"
注：
　　一个卷可以存在于一个或多个容器的指定目录，该目录可以绕过联合文件系统，并具有以下功能：
1 卷可以容器间共享和重用
2 容器并不一定要和其它容器共享卷
3 修改卷后会立即生效
4 对卷的修改不会对镜像产生影响
5 卷会一直存在，直到没有任何容器在使用它
```



**WORKDIR：工作目录，类似于cd命令**

```
格式：
    WORKDIR /path/to/workdir
示例：
    WORKDIR /a  (这时工作目录为/a)
    WORKDIR b  (这时工作目录为/a/b)
    WORKDIR c  (这时工作目录为/a/b/c)
注：
　　通过WORKDIR设置工作目录后，Dockerfile中其后的命令RUN、CMD、ENTRYPOINT、ADD、COPY等命令都会在该目录下执行。在使用docker run运行容器时，可以通过-w参数覆盖构建时所设置的工作目录。
```



**USER:**

指定运行容器时的用户名或 UID，后续的 RUN 也会使用指定用户。使用USER指定用户时，可以使用用户名、UID或GID，或是两者的组合。当服务不需要管理员权限时，可以通过该命令指定运行用户。并且可以在之前创建所需要的用户**



 格式:
　　USER user
　　USER user:group
　　USER uid
　　USER uid:gid
　　USER user:gid
　　USER uid:group

 示例：
　　USER www

 注：

　　使用USER指定用户后，Dockerfile中其后的命令RUN、CMD、ENTRYPOINT都将使用该用户。镜像构建完成后，通过`docker run`运行容器时，可以通过-u参数来覆盖所指定的用户。



**ARG：用于指定传递给构建运行时的变量**

```
格式：
    ARG <name>[=<default value>]
示例：
    ARG site
    ARG build_user=www
```

**ONBUILD：用于设置镜像触发器**



```
格式：
　　ONBUILD [INSTRUCTION]
示例：
　　ONBUILD ADD . /app/src
　　ONBUILD RUN /usr/local/bin/python-build --dir /app/src
注：
　　当所构建的镜像被用做其它镜像的基础镜像，该镜像中的触发器将会被钥触发
```

### 利用dockerfile构建beego环境并运行web_v1

1、在chives仓库下的`web_v1`新建`dockerfile`文件，并输入：

```dockerfile
FROM golang:latest
MAINTAINER brook1711 "brook1711@163.com"  
WORKDIR $GOPATH/src/web_v1
ADD ./web_v1/ $GOPATH/src/web_v1  

RUN go get github.com/beego/beego/v2 && go get github.com/beego/bee/v2
RUN go mod tidy && bee run

EXPOSE 8080
```

2、在仓库根目录下执行：

```
docker build -t  web:v1 ./
```

注意这里是在`dockerfile`所在的目录下 执行的命令，这里的`web:v1`指的是tag为v1的名叫web的镜像。

` ./ ` 是上下文路径，是指 docker 在构建镜像，有时候想要使用到本机的文件（比如复制），docker build 命令得知这个路径后，会将路径下的所有内容打包。

**解析**：由于 docker 的运行模式是 C/S。我们本机是 C，docker 引擎是 S。实际的构建过程是在 docker 引擎下完成的，所以这个时候无法用到我们本机的文件。这就需要把我们本机的指定目录下的文件一起打包提供给 docker 引擎使用。

如果未说明最后一个参数，那么默认上下文路径就是 Dockerfile 所在的位置。

**注意**：上下文路径下不要放无用的文件，因为会一起打包发送给 docker 引擎，如果文件过多会造成过程缓慢。

之后就可以在docker应用中看到创建的镜像：

![image-20210305222936205](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210305222936205.png)

之后通过镜像进入docker容器：

```
docker run -it -p 1234:22 -p 8081:8080 --name=web_v1 web:v1 /bin/bash
```

进入容器后执行

```
bee run
```

即可运行服务

通过命令打包：

```
bee pack -be GOOS=linux
```

并在后台执行：

```
docker exec web_v1 -i -t web_v1
```

```
nohup ./web_v1 >/dev/null 2>/dev/null&
```

## 7、dockerfile在goland中的进阶用法

### 找到`add configration`选项，找到`dockerfile`

![image-20210308171714676](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210308171714676.png)



![image-20210308172021771](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210308172021771.png)

1、选择dockerfile路径

2、设定保存的镜像的名字和标签

3、设定跑起来的container的名字

4、设置docker运行起来之后会自动执行什么命令

5、设置端口绑定

6、这里是等效命令行操作

### 操作窗口

在点击绿色按钮运行docker后就会看到编译器下方有一系列操作窗口

![image-20210308174855628](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210308174855628.png)

其中在`attached console` 里面可以对容器进行命令行操作![image-20210308174845363](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210308174845363.png)

在`properties`里面可以看到容器信息，点击`save`按钮可以直接保存镜像（刚运行起来之后必须手动save才会在docker desktop中看到对应镜像）

![image-20210308174833938](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210308174833938.png)

`files`是对容器的类似sftp的窗口

![image-20210308174821202](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210308174821202.png)

### 最终效果

此前将容器的8080端口绑定到了主机的8081端口，所以此时访问主机的`127.0.0.1:8081`即可看到网页：
![image-20210308174808677](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210308174808677.png)