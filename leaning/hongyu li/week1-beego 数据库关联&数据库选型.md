# beego 数据库关联&数据库选型

## beego开发前的配置工作

1. 开发工具  
  - GoLand
    - GoLand可以远程连接到服务器上，提供本地化的开发体验，非常适合网站开发。
2. 配置GoLand步骤
  - 配置文件传输
    - 打开配置菜单Configuration
      - ![打开配置菜单](https://raw.githubusercontent.com/Richardhongyu/pic/main/20210301160424.png)
    - 选择对应的文件传输协议
      - ![文件传输协议](https://raw.githubusercontent.com/Richardhongyu/pic/main/20210301160728.png)
    - 编辑SSH配置详情
      - ![编辑](https://raw.githubusercontent.com/Richardhongyu/pic/main/20210301160919.png)
    - 配置项目的映射路径，注意deployment的路径是相对路径
      - ![映射路径](https://raw.githubusercontent.com/Richardhongyu/pic/main/20210301161111.png)
    - 成功打开文件系统
      - ![文件系统](https://raw.githubusercontent.com/Richardhongyu/pic/main/20210301161227.png)
    - 小tips:
      - 注意文件系统的协议选择
      - 可以在option选项中修改文件传输中的覆盖，删除等操作
      - 推荐打开automatic upload选项，可以在每次修改后自动执行上传命令
      - 可以设置快捷键打开remote host
      - 配置文件可以保留，取消勾选visible only for this project后，就可以在多个工程中重复使用同一个服务器的配置
  - 命令行
    - ![打开ssh会话](https://raw.githubusercontent.com/Richardhongyu/pic/main/20210301170454.png)
    - 配置SSH快捷键
      - 第一个是启动SSH连接
        - ![启动](https://raw.githubusercontent.com/Richardhongyu/pic/main/20210301183727.png)
      - 第二个是切换主界面到SSH
        - ![切换](https://raw.githubusercontent.com/Richardhongyu/pic/main/20210301183828.png)
        - 使用顺序是先启动配置好的SSH，然后在写代码和运行之间切换
  - 其他
    - 还可以配置remote debug
      - 安装Go的debugger，这里采用的是delve
        - [官方安装教程](https://github.com/go-delve/delve/blob/master/Documentation/installation/README.md)
    - 但是目前没有发现GoLand如何直接连接远端的Go，在goroot上目前只支持本地的文件，只能用SSH运行代码。这个需求有待后续发掘新方法。值得一提的是目前vscode已经支持了，后面再折腾吧

## beego项目开发

### 安装与配置

1. 在mingxin wei的文章“beego项目部署和域名绑定”和xufeng guo“beego环境配置，和web_v1运行”中均介绍了beego的安装  
  - 我这里不再详细介绍，前者介绍了beego在windows和centos系统下的安装，后者介绍了在mac m1系统下的安装
  - 需要补充的是在centos系统下，如果需要全局运行bee命令，必须添加$GOPATH/bin也就是$GOBIN到$PATH环境变量。

### 项目创建

1. 创建项目
  - 在远程创建项目
    ```go
    bee new web
    ```
  - 在本地创建文件名相同的工程，并配置deployedment path为远程工程文件，最后拉取工程文件
2. 运行项目
  - 参见xufeng guo的教程，需要先运行
    ```go
    go mod tidy
    ```
3. 代码编写
  - 学习过程主要参考了xufeng guo的beego框架学习和beego官网的快速入门
  - 目标
    
    - 编写两个API，分别实现对数据库的写入和删除
  - 过程
    - 编写路由
      - 路由实现将对应的请求分发到controller
      - 路由分为三种：固定路由，正则路由，自动路由；我们这里采用固定路由。
      - 在routers/route.go文件中添加路由规则
        ```go
        beego.Router("/write", &controllers.WriteController{})
        beego.Router("/read", &controllers.ReadController{})
        ```
    - controller
      - controller负责具体处理一个逻辑
      - 一个controller中的处理函数分为两部分：结构体定义以及具体的逻辑
      - 在controllers/default.go文件中添加处理逻辑
        - 先定义对应的结构体
          ```go
          type WriteController struct {
            beego.Controller
        }
  
          type ReadController struct {
            beego.Controller
          }
          ```
        - 根据需要的逻辑，重写beego.Controller的方法，我们这里重写的是Post方法
    - model
      
      - beego受到django和SQLAlchemy的启发，采用了ORM的思想，在sql语句和go语言操作数据库的接口中添加了一层，简化了编程难度。