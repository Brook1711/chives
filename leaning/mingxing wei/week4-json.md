# 弹幕json的传递发送

这周的主要工作是对弹幕发送给后台的json文件进行追踪，并找到发送给后端的路径与方法。

## 浏览器抓包

![image-20210319163849269](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210319163849269.png)

这里我使用了代理，因此远程地址的端口是10808，

观察其堆栈调用情况：

![image-20210321144012607](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210321144012607.png)

发现其是在danmaku.js这个文件下执行的相应的操作，在提示位置的135行设置一个断点： 		![image-20210321144138866](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210321144138866.png)

然后发送一个弹幕进行测试，此处发送的是demo1这个弹幕，发送如下：

![image-20210321145521508](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210321145521508.png)

danmaku.js将相应的弹幕数据danmakudata进行html加密``此处的html加密没啥用``等方法重新定义新变量danmaku，并且通过send的方法传给后台。

![image-20210321150752705](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210321150752705.png)

```javascript
{
    key: 'send',
        value: function send(dan, callback) {
            var _this3 = this;	// 此处
            var danmakuData = {
                token: this.options.api.token,
                id: this.options.api.id,
                author: this.options.api.user,
                time: this.options.time(),
                text: dan.text,
                color: dan.color,
                type: dan.type
            }; // 此处是将弹幕文件打包成json
            this.options.apiBackend.send({
                url: this.options.api.address + 'v3/',
                data: danmakuData,
                success: callback,
                error: function error(msg) {
                    _this3.options.error(msg || _this3.options.tran('Danmaku send failed'));
                }
            });
            this.dan.splice(this.danIndex, 0, danmakuData);
            this.danIndex++;
            var danmaku = {  //实际发送的变量
                text: this.htmlEncode(danmakuData.text), //htmlencode加密，但没什么卵用，还是明文
                color: danmakuData.color,
                type: danmakuData.type,
                border: '2px solid ' + this.options.borderColor
            };
            this.draw(danmaku); //在播放器实时弹出弹幕
            this.events && this.events.trigger('danmaku_send', danmakuData);
        }//绑定发送按钮，发送danmamuData
},
```

向服务器发送的参数：

![image-20210321154551561](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210321154551561.png)

dplayer-node使用了**[xhr.js](https://developer.mozilla.org/zh-CN/docs/Web/API/XMLHttpRequest)**与**[Axios.js](http://www.axios-js.com/zh-cn/docs/index.html)**两个库来处理前后端交互。

此处代码可以**[参考](https://zhuanlan.zhihu.com/p/37962469)**

![image-20210321155615991](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210321155615991.png)

将上传弹幕识别为post方法，并根据url进行上传。

Axois.js里的config

![image-20210321160049042](https://cdn.jsdelivr.net/gh/nickmxxx/pic-bed/img/image-20210321160049042.png)

使用Axios将config解析为**[promise](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Promise)**

```javascript
var promise = Promise.resolve(config);
```

`Promise.resolve(value)`方法返回一个以给定值解析后的[`Promise`](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Promise) 对象。如果这个值是一个 promise ，那么将返回这个 promise ；如果这个值是thenable（即带有[`then `](https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Promise/then)方法），返回的promise会“跟随”这个thenable的对象，采用它的最终状态；否则返回的promise将以此值完成。此函数将类promise对象的多层嵌套展平。

前端的部分就是通过上述的内容将其传到服务器后台。

## 服务器的接收

我看了一下dplayer-node的源码，其使用的是以node.js为基础的**[KOA](https://blog.csdn.net/qq_43389371/article/details/113614934)**框架

