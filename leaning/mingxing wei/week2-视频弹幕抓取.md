# 视频弹幕抓取

## bilibili弹幕

### 思路一：油猴脚本

针对B站的弹幕爬取相对较多，在GitHub上发现了一个宝藏油猴脚本 **[bilibili evolve](https://github.com/the1812/Bilibili-Evolved)**

记得hongyu li 说过有想把这个项目做成一个插件的形式，感觉是一个比较好的示例

这个脚本可以直接下载.ass弹幕嵌入视频或者下载.json/.xml文件观感极好，也可以下载视频（无需登录）

![image-20210305114045558](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210305114045558.png)

还可以结合IDM下载

![image-20210305203819415](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210305203819415.png)

下载到本地的弹幕***.ass***文件

![image-20210305114233911](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210305114233911.png)

将***.ass***文件嵌入到视频当中，可以看到弹幕存在，效果不错（来自绵羊大姨的白眼）

![image-20210305114527847](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210305114527847.png)

​													这是本地的potplayer

但问题在于这是个油猴脚本，现在的思路可以认为根据av/bv号寻找到.ass文件，之后再上传到服务器上面。

### 思路二 服务器爬虫返回

对于这种思路找到了一个python的爬虫脚本，可以根据输入的url来获取弹幕，并直接输出.ass文件的形式

![image-20210305173118419](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210305173118419.png)

抓取到的是蓝框里面的字幕，已经是足够用了（官方已经进行了筛选）

![image-20210305173526504](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210305173526504.png)

实测完全是可行的（本地的potplayer）

![image-20210305173811622](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210305173811622.png)

## 一点思考

本周的工作主要是进行了爬取弹幕，针对B站的弹幕爬取问题并不大，***问题存在于怎么根据用户上传的本地文件找到相应的弹幕***（maybe需要一个巨大的数据库？）但这显然不太现实，即从0-1的过程是比较困难，尤其是对于像B站这种弹幕。所以是否只瞄准于某些经典大片（类似字幕组等BT站的电影）相对而言较容易匹配，也容易识别，这是本周的一点思考。

## 出现的问题及解决

用setup.py安装的问题

```
python3 setup.py install
```


会安装到默认的python版本，**[指定安装路径](https://blog.csdn.net/sowhatgavin/article/details/81912541)**

```
python setup.py install --prefix=D:\anaconda\envs\bili-scrapy
```

`注意`此处到虚拟环境的目录下即可，不需要进入到lib下的sitepackages

![image-20210305142552008](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210305142552008.png)

可以看到安装成功



寻找sessdata

![sessdat](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/sessdat.PNG)

