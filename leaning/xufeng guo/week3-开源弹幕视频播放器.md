# 开源弹幕视频播放器

参考：

​	本地视频网页播放：

​	http://jsfiddle.net/dsbonev/cCCZ2/

​	开源弹幕视频播放器

​	https://github.com/chiruom/DanmuPlayer

将以上两项目融合，在beego项目web_v1中实现利用web调用本地视频并播放弹幕

## 1、添加DanmuPlayer的js文件依赖

在将[src](https://github.com/chiruom/DanmuPlayer/tree/master/src)中的文件全部复制到`web_v1/static`文件夹下

![image-20210311183611104](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210311183611104.png)



在`views`文件夹下新建我们要用到的html文件

![image-20210311183921813](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210311183921813.png)

```go
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>弹幕测试</title>
    <link href="../static/css/scojs.css" rel="stylesheet">
    <link href="../static/css/colpick.css" rel="stylesheet">
    <link href="../static/css/bootstrap.css" rel="stylesheet">
    <link rel="stylesheet" href="../static/css/main.css">
</head>
<body>
<input type="file" accept="video/*" />
<div id="danmup" style="left: 50%;margin-left:-400px;top:100px">
</div>
<div style="display: none">
    <span class="glyphicon" aria-hidden="true">&#xe072</span>
    <span class="glyphicon" aria-hidden="true">&#xe073</span>
    <span class="glyphicon" aria-hidden="true">&#xe242</span>
    <span class="glyphicon" aria-hidden="true">&#xe115</span>
    <span class="glyphicon" aria-hidden="true">&#xe111</span>
    <span class="glyphicon" aria-hidden="true">&#xe096</span>
    <span class="glyphicon" aria-hidden="true">&#xe097</span>
</div>


</body>

<script src="../static/js/jquery-2.1.4.min.js"></script>
<script src="../static/js/jquery.shCircleLoader.js"></script>
<script src="../static/js/sco.tooltip.js"></script>
<script src="../static/js/colpick.js"></script>
<script src="../static/js/jquery.danmu.js"></script>
<script src="../static/js/main.js"></script>
<!--<script src="../dist/js/danmuplayer.min.js"></script>-->
<script>
    (function localFileVideoPlayer() {
        'use strict'
        var URL = window.URL || window.webkitURL
        var displayMessage = function (message, isError) {
            var element = document.querySelector('#danmup')
            element.innerHTML = message
            element.className = isError ? 'error' : 'info'
        }
        var playSelectedFile = function (event) {
            var file = this.files[0]
            var type = file.type
            var videoNode = document.querySelector('#danmup')
            //var canPlay = videoNode.canPlayType(type)
            //if (canPlay === '') canPlay = 'no'
            //var message = 'Can play type "' + type + '": ' + canPlay
            //var isError = canPlay === 'no'
            //displayMessage(message, isError)

            //if (isError) {
            //    return
            //}

            var fileURL = URL.createObjectURL(file)
            videoNode.src = fileURL
            $("#danmup").DanmuPlayer({
                src:fileURL,
                height: "480px", //区域的高度
                width: "800px" //区域的宽度
                ,autostart:""
                ,controls:""
                //,urlToGetDanmu:"query.php"
                //,urlToPostDanmu:"stone.php"
            });
            $("#danmup .danmu-div").danmu("addDanmu",[
                { "text":"这是滚动弹幕" ,color:"white",size:1,position:0,time:2}
                ,{ "text":"我是帽子绿" ,color:"green",size:1,position:0,time:3}
                ,{ "text":"哈哈哈啊哈" ,color:"black",size:1,position:0,time:10}
                ,{ "text":"这是顶部弹幕" ,color:"yellow" ,size:1,position:1,time:3}
                ,{ "text":"这是底部弹幕" , color:"red" ,size:1,position:2,time:9}
                ,{ "text":"大家好，我是伊藤橙" ,color:"orange",size:1,position:1,time:3}

            ])
        }
        var inputNode = document.querySelector('input')
        inputNode.addEventListener('change', playSelectedFile, false)


    })()



</script>
<script type="text/javascript">

</script>
</html>
```

在这些代码中：

![image-20210311184020315](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210311184020315.png)

放在头部标签<head>里面的为所用到的`css`依赖包，这些包在之前已经移到了本地

根据分析，头部依赖应该是前端依赖（`css`代码）

在`body`标签结束后，需要引用弹幕所需的`js`依赖包

![image-20210311184320634](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210311184320634.png)

最后一行注释掉的应该是维持弹幕播放器运行的最小`js`依赖包



## 2、前端逻辑

更改`controllers`指定的主页文件为`danmu.html`

![image-20210311183722268](https://cdn.jsdelivr.net/gh/Brook1711/fig_for_blog/img/image-20210311183722268.png)

```go
package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "chives.me"
	c.Data["Email"] = "brook1711@bupt.edu.cn"
	c.TplName = "danmu.html"
}

```

## 3、JavaScript代码

在文件最后有一段js代码：

```javascript
<script>
    (function localFileVideoPlayer() {
        'use strict'
        var URL = window.URL || window.webkitURL
        var displayMessage = function (message, isError) {
            var element = document.querySelector('#danmup')
            element.innerHTML = message
            element.className = isError ? 'error' : 'info'
        }
        var playSelectedFile = function (event) {
            var file = this.files[0]
            var type = file.type
            var videoNode = document.querySelector('#danmup')
            //var canPlay = videoNode.canPlayType(type)
            //if (canPlay === '') canPlay = 'no'
            //var message = 'Can play type "' + type + '": ' + canPlay
            //var isError = canPlay === 'no'
            //displayMessage(message, isError)

            //if (isError) {
            //    return
            //}

            var fileURL = URL.createObjectURL(file)
            videoNode.src = fileURL
            $("#danmup").DanmuPlayer({
                src:fileURL,
                height: "480px", //区域的高度
                width: "800px" //区域的宽度
                ,autostart:""
                ,controls:""
                //,urlToGetDanmu:"query.php"
                //,urlToPostDanmu:"stone.php"
            });
            $("#danmup .danmu-div").danmu("addDanmu",[
                { "text":"这是滚动弹幕" ,color:"white",size:1,position:0,time:2}
                ,{ "text":"我是帽子绿" ,color:"green",size:1,position:0,time:3}
                ,{ "text":"哈哈哈啊哈" ,color:"black",size:1,position:0,time:10}
                ,{ "text":"这是顶部弹幕" ,color:"yellow" ,size:1,position:1,time:3}
                ,{ "text":"这是底部弹幕" , color:"red" ,size:1,position:2,time:9}
                ,{ "text":"大家好，我是伊藤橙" ,color:"orange",size:1,position:1,time:3}

            ])
        }
        var inputNode = document.querySelector('input')
        inputNode.addEventListener('change', playSelectedFile, false)


    })()
    
</script>
```

这里的代码需要注意的是：

采用了以下格式

```javascript
(function Func() {
       
    })(param)
```

>也就是定义了一函数:
>Func()
>然后我们立即执行了它:Func(param);
>
>参考：请问js里两个括号是什么意思？ - 张雄的回答 - 知乎 https://www.zhihu.com/question/48238548/answer/109922725

这里我们一段一段解释：

