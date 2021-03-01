# beego框架学习 1 简介

## 项目路由

在bee工具生成的项目文件夹下的主文件为`main.go`，其内容为：

```go
package main

import (
	_ "web_v1/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}


```

![image-20210224134654835](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210224134654835.png)

import 下划线（如：import _ hello/imp）的作用：当导入一个包时，该包下的文件里所有init()函数都会被执行，然而，有些时候我们并不需要把整个包都导入进来，仅仅是是希望它执行init()函数而已。这个时候就可以使用 import _ 引用该包。即使用【import _ 包路径】只是引用该包，仅仅是为了调用init()函数，所以无法通过包名来调用包中的其他函数。

这里在主程序中引用了routers，但是只是执行了其中的init函数，routers中的代码如下：

```go
package routers

import (
   "web_v1/controllers"
   beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
```

路由包里面我们看到执行了路由注册 `web.Router`, 这个函数的功能是映射 URL 到 controller，第一个参数是 URL (用户请求的地址)，这里我们注册的是 `/`，也就是我们访问的不带任何参数的 URL，第二个参数是对应的 Controller，也就是我们即将把请求分发到那个控制器来执行相应的逻辑，

> &是取地址符；*为指针变量的声明，可以理解为使一个地址值生效，和c语言中类似。在golang-100-days中的day12中有定义

我们可以执行类似的方式注册如下路由：

```go
web.Router("/user", &controllers.UserController{})
```

当然，UserController必须在`web_v1/controllers/default.go`中定义。这样用户就可以通过访问 `/user` 去执行 `UserController` 的逻辑。这就是我们所谓的路由，更多更复杂的路由规则请查询 [web 的路由设置](https://beego.me/docs/mvc/controller/router.md)

## web.run()

在import部分执行完路由配置之后，在主程序中会执行`web.run()`函数。该函数完成以下功能：

- 解析配置文件

  beego 会自动解析在 conf 目录下面的配置文件 `app.conf`，通过修改配置文件相关的属性，我们可以定义：开启的端口，是否开启 session，应用名称等信息。

- 执行用户的 hookfunc

  beego 会执行用户注册的 hookfunc，默认的已经存在了注册 mime，用户可以通过函数 `AddAPPStartHook` 注册自己的启动函数。

- 是否开启 session

  会根据上面配置文件的分析之后判断是否开启 session，如果开启的话就初始化全局的 session。

- 是否编译模板

  beego 会在启动的时候根据配置把 views 目录下的所有模板进行预编译，然后存在 map 里面，这样可以有效的提高模板运行的效率，无需进行多次编译。

- 是否开启文档功能

  根据 EnableDocs 配置判断是否开启内置的文档路由功能

- 是否启动管理模块

  beego 目前做了一个很酷的模块，应用内[监控模块](https://beego.me/docs/advantage/monitor.md)，会**在 8088 端口做一个内部监听**，我们可以通过这个端口查询到 QPS、CPU、内存、GC、goroutine、thread 等统计信息。

- 监听服务端口

  这是最后一步也就是我们看到的访问 8080 看到的网页端口，内部其实调用了 `ListenAndServe`，充分利用了 goroutine 的优势

- 这是最后一步也就是我们看到的访问 8080 看到的网页端口，内部其实调用了 `ListenAndServe`，充分利用了 goroutine 的优势

一旦 run 起来之后，我们的服务就监听在两个端口了，一个服务端口 8080 作为对外服务，另一个 8088 端口实行对内监控。

通过这个代码的分析我们了解了 beego 运行起来的过程，以及内部的一些机制。接下来让我们去剥离 Controller 如何来处理逻辑的。

## 状态监控

在8088端口可以查询状态信息，但是默认关闭，可以在`/conf/app.conf`中输入：

```go
EnableAdmin = true
```

开启状态监控

```go
AdminAddr = "localhost"
AdminPort = 8088
```

## controller运行机制

`/routers/default.go`文件中定义了页面路由，将到达网站的请求分发到对应的控制器`controller`进行处理。controller在`/contollers/default.go`中进行定义，其源码：

```go
package controllers

import (
   beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
   beego.Controller
}

func (c *MainController) Get() {
   c.Data["Website"] = "beego.me"
   c.Data["Email"] = "astaxie@gmail.com"
   c.TplName = "index.html"
}
```

上面的代码显示首先我们声明了一个控制器 `MainController`

> go语言中的结构体声明在golang-100-days中的day13中有定义。Go中的函数定义规则在day14中

在该代码中，首先定义了一个`MainController`结构体，

```go
type MainController struct {
   beego.Controller
}
```

这个控制器里面==**组合**==了 `web.Controller`，这就是 Go 的组合方式，也就是 `MainController` 自动拥有了所有 `web.Controller` 的方法。

而 `web.Controller` 拥有很多方法，其中包括 `Init`、`Prepare`、`Post`、`Get`、`Delete`、`Head` 等方法。我们可以==通过重写的方式来实现这些方法==，而我们上面的代码就是重写了 `Get` 方法。

其实，可以把`beego.Controller`看作是一个控制器模版，我们通过重新定义结构体（组合）的方式并自定义其中的某些属性来自定义页面。

我们先前介绍过 beego 是一个 RESTful 的框架，所以我们的请求默认是执行对应 `req.Method` 的方法。例如浏览器的是 `GET` 请求，那么默认就会执行 `MainController` 下的 `Get` 方法。这样我们上面的 Get 方法就会被执行到，这样就进入了我们的逻辑处理。（用户可以改变这个行为，通过注册自定义的函数名，更加详细的请参考[路由设置](https://beego.me/docs/mvc/controller/router.md#自定义方法及-restful-规则)）

> ## RESTful架构：
>
> Fielding将他对互联网软件的架构原则，定名为REST，即Representational State Transfer的缩写。我对这个词组的翻译是"表现层状态转化"。
>
> ### 1、资源（**Resources**）
>
> REST的名称"表现层状态转化"中，省略了主语。"表现层"其实指的是"资源"（Resources）的"表现层"。
>
> **所谓"资源"，就是网络上的一个实体，或者说是网络上的一个具体信息。**它可以是一段文本、一张图片、一首歌曲、一种服务，总之就是一个具体的实在。你可以用一个URI（统一资源定位符）指向它，每种资源对应一个特定的URI。要获取这个资源，访问它的URI就可以，因此URI就成了每一个资源的地址或独一无二的识别符。
>
> 所谓"上网"，就是与互联网上一系列的"资源"互动，调用它的URI。
>
> ### 2、表现层（**Representation**）
>
> "资源"是一种信息实体，它可以有多种外在表现形式。**我们把"资源"具体呈现出来的形式，叫做它的"表现层"（Representation）。**
>
> 比如，文本可以用txt格式表现，也可以用HTML格式、XML格式、JSON格式表现，甚至可以采用二进制格式；图片可以用JPG格式表现，也可以用PNG格式表现。
>
> URI只代表资源的实体，不代表它的形式。严格地说，有些网址最后的".html"后缀名是不必要的，因为这个后缀名表示格式，属于"表现层"范畴，而URI应该只代表"资源"的位置。它的具体表现形式，应该在HTTP请求的头信息中用Accept和Content-Type字段指定，这两个字段才是对"表现层"的描述。
>
> ### 3、状态转化（**State Transfe**）
>
> 访问一个网站，就代表了客户端和服务器的一个互动过程。在这个过程中，势必涉及到数据和状态的变化。
>
> 互联网通信协议HTTP协议，是一个无状态协议。这意味着，所有的状态都保存在服务器端。因此，**如果客户端想要操作服务器，必须通过某种手段，让服务器端发生"状态转化"（State Transfer）。而这种转化是建立在表现层之上的，所以就是"表现层状态转化"。**
>
> 客户端用到的手段，只能是HTTP协议。具体来说，就是HTTP协议里面，四个表示操作方式的动词：GET、POST、PUT、DELETE。它们分别对应四种基本操作：**GET用来获取资源，POST用来新建资源（也可以用于更新资源），PUT用来更新资源，DELETE用来删除资源。**

里面的代码是需要执行的逻辑，这里只是简单的输出数据，我们可以通过各种方式获取数据，然后赋值到 `c.Data` 中，这是一个用来存储输出数据的 map，

> map的定义在day08中

可以赋值任意类型的值，这里我们只是简单举例输出两个字符串。

最后一个就是需要去渲染的模板，`c.TplName` 就是需要渲染的模板，这里指定了 `index.tpl`，如果用户不设置该参数，那么默认会去到模板目录的 `Controller/<方法名>.tpl` 查找，例如上面的方法会去 `maincontroller/get.tpl` ***(文件、文件夹必须小写)\***。

用户设置了模板之后系统会自动的调用 `Render` 函数（这个函数是在 `web.Controller` 中实现的），所以无需用户自己来调用渲染。

当然也可以不使用模版，直接用 `c.Ctx.WriteString` 输出字符串，如

```go
func (c *MainController) Get() {
        c.Ctx.WriteString("hello")
}
```

## goland中读取.tpl文件（可忽略）

**beego也支持`.html`文件，所以直接改文件后缀就行**

在beego框架中，前端文件不是以`.html`格式存储的，而是存为模版文件`.tpl`并存储在`/views`文件夹下，正常情况下goland不会在`.tpl`文件中有高亮提示，我们需要在IDE中进行设置：

![image-20210225103333270](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210225103333270.png)

之后就可以正常在.tpl中使用代码高亮了

## model分析

我们知道 Web 应用中我们用的最多的就是数据库操作，而 model 层一般用来做这些操作，我们的 `bee new` 例子不存在 Model 的演示，但是 `bee api` 应用中存在 model 的应用。说的简单一点，如果您的应用足够简单，那么 Controller 可以处理一切的逻辑，如果您的逻辑里面存在着可以复用的东西，那么就抽取出来变成一个模块。因此 Model 就是逐步抽象的过程，一般我们会在 Model 里面处理一些数据读取，如下是一个日志分析应用中的代码片段：

```go
package models

import (
    "path/filepath"
    "strings"
)

var (
    NotPV []string = []string{"css", "js", "class", "gif", "jpg", "jpeg", "png", "bmp", "ico", "rss", "xml", "swf"}
)

const big = 0xFFFFFF

func LogPV(urls string) bool {
    ext := filepath.Ext(urls)
    if ext == "" {
        return true
    }
    for _, v := range NotPV {
        if v == strings.ToLower(ext) {
            return false
        }
    }
    return true
}
```

所以如果您的应用足够简单，那么就不需要 Model 了；如果你的模块开始多了，需要复用，需要逻辑分离了，那么 Model 是必不可少的。接下来我们将分析如何编写 View 层的东西。

## view编写

在前面编写 Controller 的时候，我们在 Get 里面写过这样的语句 `this.TplName = "index.tpl"`，设置显示的模板文件，默认支持 `tpl` 和 `html` 的后缀名，如果想设置其他后缀你可以调用 `beego.AddTemplateExt` 接口设置，那么模板如何来显示相应的数据呢？beego 采用了 Go 语言默认的模板引擎，所以和 Go 的模板语法一样，Go 模板的详细使用方法请参考[《Go Web 编程》模板使用指南](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/07.4.md)

我们看看快速入门里面的代码（去掉了 css 样式）：

```html
<!DOCTYPE html>

<html>
    <head>
        <title>Beego</title>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    </head>
    <body>
        <header class="hero-unit" style="background-color:#A9F16C">
            <div class="container">
                <div class="row">
                    <div class="hero-text">
                        <h1>Welcome to Beego!</h1>
                        <p class="description">
                            Beego is a simple & powerful Go web framework which is inspired by tornado and sinatra.
                            <br />
                            Official website: <a href="http://{{.Website}}">{{.Website}}</a>
                            <br />
                            Contact me: {{.Email}}
                        </p>
                    </div>
                </div>
            </div>
        </header>
    </body>
</html>
```

我们在 Controller 里面把数据赋值给了 data（map 类型），然后我们在模板中就直接通过 key 访问 `.Website` 和 `.Email` 。这样就做到了数据的输出。接下来我们讲解如何让静态文件输出。

## 静态文件处理

前面我们介绍了如何输出静态页面，但是我们的网页往往包含了很多的静态文件，包括图片、JS、CSS 等，刚才创建的应用里面就创建了如下目录：

```
├── static
    │   ├── css
    │   ├── img
    │   └── js
```

beego 默认注册了 static 目录为静态处理的目录，注册样式：URL 前缀和映射的目录（在`/main.go`文件中`web.Run()`之前加入）：

==`StaticDir`目前有bug==

```
StaticDir["/static"] = "static"
```

用户可以设置多个静态文件处理目录，例如你有多个文件下载目录 download1、download2，你可以这样映射（在 `/main.go` 文件中 `web.Run()` 之前加入）：

```
web.SetStaticPath("/down1", "download1")
web.SetStaticPath("/down2", "download2")
```

这样用户访问 URL `http://localhost:8080/down1/123.txt` 则会请求 download1 目录下的 123.txt 文件。