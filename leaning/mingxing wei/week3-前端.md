# Dplayer

## 关于Dplayer

Dplayer是一个浏览器上的支持HTML5的弹幕视频播放器，比较符合项目的需求，并且能够支持vue.js生态，是目前较为流行的浏览器播放器。

## Dplayer基本用法

如果没有接触过前端的话，建议先看一下前端的入门视频，推荐这个**[视频](https://www.bilibili.com/video/BV1Jt411D7j6)**，讲的较为精炼。

可以直接参考**[官方教程](http://dplayer.js.org/zh/guide.html#special-thanks)**，内容比较详细。

```
yarn add dplayer # 安装Dplayer
```

​	1. 打开Vscode编辑器，将下载好的DPlayer.min.js与DPlayer.min.js.map文件放入文件夹中，并引入一个demo.mp4的文件。

![image-20210314111855140](C:\Users\mxxxnick\AppData\Roaming\Typora\typora-user-images\image-20210314111855140.png)

​	2. 新建index.html文件，直接在最前面输一个！就会生成模板，在body里面引入两个脚本

```javascript
 <script src="DPlayer.min.js"></script>
    <script>
        const dp = new DPlayer({
            container: document.getElementById('dplayer'),
            video: {
                url: 'demo.mp4',
            },
        });
    </script>
```

​	3. 按F5运行html，使用chrome:preview模式即可打开页面进行相应的调试，可以看到打开了本地的视频

![image-20210314112008067](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210314112008067.png)

4. 其他功能引入

Dplayer给了许多API，我们使用它的时候只需要在新建new Dplayer时将其引入即可，注意分号与逗号的使用

```javascript
    <script>
        const dp = new DPlayer({
            container: document.getElementById('dplayer'),
            video: {
                url: 'demo.mp4',
            },
            新属性 ：{
            
        	}，
        });
    </script>
```

​	![image-20210314113052517](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210314113052517.png)

6. 弹幕的使用

​	官方给了一个Bilibili的**[API](https://api.prprpr.me/dplayer/)**接口，但是已经年久失修，截至到我测试的时候，是不好用的。但文档给了一种搭建**[后台弹幕库](https://github.com/MoePlayer/DPlayer-node)**的方法，下面将就这个来搭建。参考**[博客](https://www.moerats.com/archives/838/)**来搭建

​	博客使用的是**[docker](https://github.com/MoePlayer/DPlayer-node/blob/master/Dockerfile)**方法来搭建，并提供了dockerfile来进行搭建，是相当nice的，也可以比较好与项目兼容。

### Docker安装 引用自**[博客](https://www.moerats.com/archives/838/)**

**1、安装Docker**

```
#CentOS 6
rpm -iUvh http://dl.fedoraproject.org/pub/epel/6/x86_64/epel-release-6-8.noarch.rpm
yum update -y
yum -y install docker-io
service docker start
chkconfig docker on

#CentOS 7、Debian、Ubuntu
curl -sSL https://get.docker.com/ | sh
systemctl start docker
systemctl enable docker
```

**2、安装Docker Compose**

```
curl -L https://github.com/docker/compose/releases/download/1.17.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
```

**3、运行镜像**
安装`git`：

```
#Debian、Ubuntu系统
apt install git -y

#CentOS系统
yum install -y git
```

再使用命令：

```
i#拉取源码
git clone https://github.com/MoePlayer/DPlayer-node.git
cd DPlayer-node
#新建镜像
docker-compose build
#拉取其它镜像并后台运行
docker-compose up -d
```

此时`api`地址为`http://ip:1207`，数据和日志存放在`/root/dplayer`文件夹。

当然如果你想其它端口，或者修改存放文件夹路径，那你在上面的新建镜像之前，作出如下操作：

```
#编辑DPlayer-node文件夹的docker-compose.yml文件，部分修改如下
mongo:
  volumes:
      - ~/dplayer/db:/data/db  #数据库存放文件夹，~/dplayer/db为映射在外部的路径，自行修改，
web:
  ports:
    - 1207:1207  #api映射到外部的端口，将前面的1207修改成你想要的即可
  volumes:
    - ~/dplayer/logs:/usr/src/app/logs  #同数据库操作
    - ~/dplayer/pm2logs:/root/.pm2/logs  #同上
```

改完后再新建镜像即可，如果你已经新建镜像了，但想改，那就清空之前的镜像再修改，方法参考→[传送门](https://www.moerats.com/archives/161/)。

在过程中会出现很多warning，但是问题不大，**只要不出error就行**，可以看到这里看到在1207端口已经被映射，安装成功

![image-20210314211008582](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210314211008582.png)

我们这里用的是服务器，因此可以直接在互联网上看到相应的界面

访问网址http://161.35.234.230:1207/

​	                                          ![image-20210314211238939](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210314211238939.png)

这里是Not Found，**这是正确的表现！！！！**我在这里卡了好久，总以为是自己做错了，翻来覆去弄了好久，最后在**[另一篇博客](https://my.oschina.net/u/4333022/blog/3324054)**上找到了相应的示例，在后面要加上相应的aid与cid才能找到字幕文件。如：bilibili?aid=80266688&cid=137358410'

![image-20210314212819547](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210314212819547.png)

这才能够在自己的浏览器找到了字幕：

![image-20210314212941703](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210314212941703.png)

## 任意视频加载弹幕

做一个demo来运行任意视频中加载上面的这个弹幕，写一个html页面加载播放器与弹幕库

```javascript
  const dp = new DPlayer({
        container: document.getElementById('dplayer'),
        video: {
            url: url,
            type: 'normal',
        },
        danmaku: {
            id: vid,
            api: 'http://161.35.234.230:1207/',
            addition: ['http://161.35.234.230:1207/v3/bilibili?aid=80266688&cid=137358410']
        },
});
```

这里出现个问题，vscode中的debug for chrome这个插件要经过chrome软件，会出现跨域的安全问题，无法加载视频，出现错误：

```
from origin 'null' has been blocked by CORS policy: Cross origin requests are only supported for protocol schemes: http, data, chrome, chrome-extension, chrome-untrusted, https.
```

这里最好的解决方法就是换个插件😊（千万别和shit mountain斗争），vscode中有一个live serve插件巨好用！！！***强推***

直接右键就可以进行网页端的调试，so nice~就😁

![image-20210314215529084](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210314215529084.png)

可以看到直接加载了digital ocean作为服务器提供的字幕，**此处弹幕与视频无关**。

## 发送的弹幕去哪了

在这里我发了个弹幕 a danmu from mxxxnick，可以在后台调试器中看到我发的这个弹幕

![image-20210314221046950](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210314221046950.png)

并且是一个post的方法，传给客户端

![image-20210314221014797](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210314221014797.png)



虽然不知道为什么，但是当我重新打开这个文件并调试的时候，我能够看到自己发的弹幕，并在下面的文件中找到了我发的弹幕，但是独立于其他的弹幕而存在，当我打开[弹幕的网址](http://161.35.234.230:1207/v3/bilibili?aid=80266688&cid=137358410')，并没有发现我的弹幕。**根据后面的实践**，此处应该的网址应该是一个爬虫出来的基本库，并没有包含到新发送的弹幕。*尚有疑问*

![	](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210314221356413.png)

​    										![image-20210314221724211](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210314221724211.png)

​	后续工作：

 	将发送的弹幕存储到数据库当中。
 	将该文件跑在项目的服务器上。
 	打开本地文件（或许可以参考xufeng guo找到的那个脚本）
 	2021.3.14（浪漫的日子献给代码）
## 将项目打包到服务器上

首先在windows本地进行实验，将自己写的index.html改名为index.tpl放在webApp/views文件夹下，将demo.mp4视频放在webApp下的static文件夹下。

运行webApp.exe可以在本地8080端口看到页面及视频的出现。

![image-20210315095920243](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210315095920243.png)

在服务器端安装同样的方法将文件放进去就行，或者重新打包一下，我按照第一种方法，在服务器端成功打包好了项目。

![image-20210315100525850](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210315100525850.png)

并能够看到了自己以前发的弹幕，看来发的弹幕属实存在。

## 弹幕の发现与备份导出

上面说到已经可以看到自己发的弹幕了，说明在自己搭建的后台数据库中已经储存到了这个弹幕数据，但是并没有在root/dplayer/db文件夹下找到弹幕的json文件。查阅搭建dplayer-node的网址，作者给出了导入及导出json文件的方法。

### 导入dans.json文件

```
mv dans.json ~/dplayer/db/backup/dans.json
docker exec dplayernode_mongo_1 mongoimport -d danmaku -c dans --file /data/db/backup/dans.json
```

### 导出备份

```
docker exec dplayernode_mongo_1 mongoexport -d danmaku -c dans -o /data/db/backup/dans.json
cat ~/dplayer/db/backup/dans.json
```

使用导出命令后可以看到命令行输出了许多弹幕文件信息（10条）

![image-20210315102250593](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210315102250593.png)

保存的备份文件在root/dplayer/data/db/backup/dans.json

![image-20210315102404604](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210315102404604.png)

复制到本地并打开，文件正常

![image-20210315102447752](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210315102447752.png)

之后我又发了一个弹幕，名为**第11个弹幕**

![image-20210315102539010](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210315102539010.png)

重新执行上面的备份命令，这次输出为11个弹幕。

![image-20210315102658430](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210315102658430.png)

说明确确实实直接保存在了mongo数据库里面，并可以导出为json格式。

## 总结

### 工作进展

1. 本次的前端实践通过使用dplayer，在网站上部署了一个简单的demo，实现了本地视频播放与弹幕的基本功能，最大的坑在于一开始盲目追求用VUE来结合dplayer，导致前期工作十分不顺利。
2. 后台方面结合了dplayer-node这一个工具，在自己的服务器利用docker上部署了基于redis及mongo数据库来进行弹幕后台储存及调用。解决了前端到后台弹幕发送的问题，并可以直接导出json备份

### 改进及下一步工作

1. 前端方面美中不足在于没有使用到VUE这一强大的工具，看到一个vue-dplayer的工具，应该尝试用一下这个东西，并学习一下vue的相关知识。此外，并没有实现打开本地文件这个功能（与xufeng guo沟通后，他说他写了一个demo，可以结合一下看看）。
2. 后台上，如何匹配视频仍是一个大问题，不过在看代码的时候，看到一个dplayer的玩家使用了**[md5脚本](https://cdn.bootcss.com/blueimp-md5/2.12.0/js/md5.min.js)**，生成唯一id，从而确定唯一视频，或许可以试一下。B站的aid(视频标识）及cid（弹幕标识）的找寻仍待解决，上周找到的那个爬虫脚本可以试着结合一下。

### 踩过的坑

本周最大的坑就是那个匹配弹幕时候aid及cid没有加上的情况，明明环境已经搭建成功，但不知道怎么用是个大问题（以后要多看源码2333）。此外就是beego的html文件要改名为tpl文件，还有静态资源放到位置及路径。

最后开学太难顶了 **不想上学**😫 😶😑😐😓😤😖😞😟🥶🚄

