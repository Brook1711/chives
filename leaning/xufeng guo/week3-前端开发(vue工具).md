目标：找到dplayer前端发送给后端的数据格式并尝试进行修改。修改内容包括弹幕的作者等。

# 1、前端发送弹幕时

## 1.1、抓包

首先在本地跑一下前端应用，并在浏览器中访问，按下f12打开浏览器调试窗口并打开network标签

![image-20210324095720874](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210324095720874.png)

发送弹幕

```
gxf-test-2021-3-24-9:53
```

在浏览器network标签中可以看到：

![image-20210324095649964](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210324095649964.png)

一共发送了两个请求：

![image-20210324095915689](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210324095915689.png)

v3代表的是访问danmuku后端的router

![image-20210324100057860](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210324100057860.png)

第一个请求没有payload，暂时忽略，我们主要看第二个请求

第二个发送的包中的request中包含弹幕的具体信息

![image-20210324100354463](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210324100354463.png)

## 1.2、payload研究

```json
{"author":"DIYgod",
 "time":0.700854,
 "text":"gxf-test-2021-3-24-9:53",
 "color":16777215,
 "type":0}
```

现在要增加实现的功能是后端可以根据前端发送的payload判断

1、是哪个用户发送的弹幕，

2、并且要验证该用户的身份，

3、要知道这条弹幕是发送到哪个视频之下

在前端代码中增加danmuku属性：

```js
        danmaku:{
          id: 'vid-example',
          user:'gxf',
          token:'token-example',
          api: 'http://161.35.234.230:1207/',
          addition: ['http://161.35.234.230:1207/v3/bilibili?aid=80266688&cid=137358410']

        },
```

之后可以看到发送弹幕之后的payload产生了变化：

```json
{
"token":"token-example",
"author":"gxf",
 "id": "vid-example",
"time":2.502,
"text":"gxf-test-2021-3-24-10:24",
"color":16777215,
"type":0}

```



# 2、后端返回数据

## 2.1 发送单条弹幕时的返回数据

```json
{"code":0,
 "data":{"_id":"605aa2d7cedd3c21319cda7d",
         "author":"gxf",
         "time":2.502,
         "text":"gxf-test-2021-3-24-10:24",
         "color":16777215,
         "player":"vid-example",
         "type":0,
         "ip":"::ffff:59.64.129.94",
         "referer":"http://127.0.0.1:8080/",
         "date":1616552663893,
         "__v":0}}
```

这个是beego后端需要重点模仿的部分。

## 2.2 b站爬取弹幕的返回格式

这个可以先不用管，到时候建立两个服务器，一个作为dplayer-node提供弹幕爬取的功能。



# 3、使用POST MAN模拟

## 3.1 发送弹幕

设置POST请求：

![image-20210324125745379](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210324125745379.png)

设置payload格式：

![image-20210324125819108](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210324125819108.png)

在body中设置payload

![image-20210324125946932](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210324125946932.png)

点击发送即可看到返回的结果：

![image-20210324130038490](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210324130038490.png)

## 3.2 拉取指定弹幕池中的弹幕

刚刚的弹幕发送到了弹幕池为`vid-example`的弹幕池中，

设置一个新的request，方法为GET

![image-20210324130220373](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210324130220373.png)

在params中设置id

![image-20210324130317751](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210324130317751.png)

点击发送，看到response

![image-20210324130411123](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210324130411123.png)

即可看到刚刚push上去的弹幕。



# 4、解决方案

## 4.0 简易demo

参考：

https://beego.me/docs/mvc/controller/params.md

### 4.0.1、处理发送弹幕的POST请求

在app.conf中添加：

```go
copyrequestbody = true
```

在`controller/default.go`中添加处理v3接口的controller

```go
type DanmuController struct {
	beego.Controller
}
```

为该controller重写POST函数：


在router.go中添加新的路由:

```go
	beego.Router("/v3", &controllers.DanmuController{})
```



### 4.0.2、处理拉取弹幕的GET请求

## 4.1 实现登陆

主要解决token问题



## 4.2 实现弹幕池

主要和数据库相关









