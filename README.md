# chives

 go web project

> 贡献者：
>
> hongyu li
>
> mingxin wei
>
> xufeng guo

# 项目readme书写规范

我们定义一下书写`点子`的格式，这样我们以后书写和回顾的时候轻松一点。

推荐使用`typora`编辑器，因为某些高亮和其他格式在vscode显示不出来。注意⚠️，**需要对typora进行一下设置才可以显示高亮！！！**

![image-20210218135957865](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210218135957865.png)

## 想法类：

不限于任何形式的内容，天马行空、任意想象。位置可以放在本文档相关性最强的内容之下。

用引用格式并配上姓名拼音mingxin wei、hongyu li、xufeng guo

例如：

> mingxin wei: 是否可以将其做成外加字幕的形式，像是在本地视频中嵌入$.srt$字幕文件一样，对于我们而言，进行的工作是将弹幕是可以进行实时补充，不断增加新用户的新弹幕。

## 技术类：

涉及具体代码实现和技术的内容。位置可以放在本文档相关性最强的内容之下。

配合`代码`、$L_aT_eX$公式、**==加粗和高亮==** 文本格式和

```markdown
代码块段落
```

进行说明。

例如：

`web`后端采用==**golang**==，

for循环代码实现为：

```go
package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop")
}
```



## 内容类：

就是把思路说清楚就行，咱也不是大创，不需要写太多，尽量说明清楚，作为readme骨干；想法类和技术类作为血肉。

# 项目代码管理规范

推荐使用`github desktop`，方便简单

## 分支问题

在技术方案变更和有重大bug时，经微信群组讨论添加分支，其余commit全部在main分支里，==一定记得经常fetch==

## commit信息填写

commit时规范标题和详情

![image-20210218140944210](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210218140944210.png)

1、标题中首先使用英文中括号标识，有`success`和`debug`两种，`success`表示当前功能编译通过，可以正常运行；`debug`表示不能正常运行，但是写到一半，当成存档点。

2、自己更改的哪个模块，就要和之前的模块版本对应+1

3、如果代码里有别人的代码，就要加上他的github 账号，以后有bug方便追责（滑稽）

# 项目内容

## 项目中主营业务

一个基于golang的web业务。主营业务为第三方托管的视频弹幕、评论和长评（可拓展，比如字幕业务等）。主要亮点为绕过视频版权方授权，用户自主获取视频源，网站只提供弹幕和评论。
>mingxin wei: 是否可以将其做成外加字幕的形式，像是在本地视频中嵌入$.srt$字幕文件一样，对于我们而言，进行的工作是将弹幕是可以进行实时补充，不断增加新用户的新弹幕。

如果该方案较为离谱可以转型为视频教育平台，主营第三方学习资料。
## 用户体验

和哔哩哔哩的使用体验一样，只不过用户需要先下载自己想看的片源，之后用户登陆到我们的网站之后在网页端选取和播放下载的视频，并可以在该视频上发表弹幕和评论。也就是说，在整个过程中只有视频是离线的，弹幕和评论都是托管在我们的平台上。
>mingxin wei: 关于主要应用场景的瞄准问题，在寒假在家的时间，我一直有一个困扰，就是无论我将手机上的视频投屏到电视上的时候，或者是应用机顶盒自带的软件市场下载的视频软件播放视频，都存在着无法食用弹幕的情况，我在想我们这个平台是否可以瞄准这个场景，作为机顶盒子或者手机投屏的弹幕的补充。

> mingxin wei: 刚开始爬取别的网站弹幕；区分服务内容，分级服务；srt内容更新较慢。

> hongyu li: 做成不同接口，.srt和实时弹幕


想象一下，如果实现的话就是一个**没有审核制度的b站**（手动滑稽）

# 技术实现

（目前只瞄准web就行了，不一定获得过这个阶段）：

全部功能内容在web端实现，不考虑app等移动端。

## 服务器后台

使用**==golang==**，原因是**==golang==**并发性好，性能高速度快而且上手快~~，而且岗位薪水高~~。

微软服务器一年免费试用

https://mp.weixin.qq.com/s/SXOM76Lh8dnM2ETW--pwPg

#### ==开源web框架==

https://github.com/hoisie/web

​	看上去挺简单的，但是已经有5年没有更新了

https://github.com/labstack/echo

​	这个看上去靠谱多了

![image-20210218225305814](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210218225305814.png)

找到一篇知乎文章，，里面有个大**==goweb框架==**对比

![img](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/v2-49443c9154b18ab93979f8867a481e87_b.jpg)

经过一番xjb对比，最后决定使用beego

https://beego.me/docs/intro/

![image-20210218231319554](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210218231319554.png)

https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210218231248758.png

## 数据库

应该是要采用独立数据库的，==**mysql**==？这一块我也不太懂，后面慢慢摸索吧

## 前端

使用==**JavaScript**==？这一块我不太懂，刚开始前端可以不太漂亮，先做一个`demo`出来。

# 推荐教程

目前主要是先把后台建起来，api接口搞起来，先学一下golang。找到一个比较新的和全面的教程：

https://github.com/rubyhan1314/Golang-100-Days.git

比较新，可以学一下

https://github.com/astaxie/build-web-application-with-golang

找到了一个比较新的用golang构建web应用的教程

关于web前端的教程可以看这个

https://github.com/qianguyihao/Web

然后还有我自己写的博客，内部特供(滑稽)

Brook1711.com

关于前端和其他的我暂时没思路。

# 最后

新年快乐🎉🧨