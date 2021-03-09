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
  - 学习过程主要参考了xufeng guo的beego框架学习和beego官网的快速入门以及MVC架构
  - 目标
    - 编写两个API，分别实现对数据库的写入和删除
  - 过程(MVC)
    - 编写路由
      - 路由实现将对应的请求分发到controller
      - 路由分为三种：固定路由，正则路由，自动路由；我们这里采用固定路由。
      - 在routers/route.go文件中添加路由规则
        ```go
        beego.Router("/create", &controllers.CreateController{})//C
        beego.Router("/delete", &controllers.DeleteController{})//D
        beego.Router("/update", &controllers.UpdateController{})//U
        beego.Router("/read", &controllers.ReadController{})//R
        ```
    - controller（C）
      - controller负责具体处理一个逻辑
      - 一个controller中的处理函数分为两部分：结构体定义以及具体的逻辑
      - 在controllers/default.go文件中添加处理逻辑
        - 先定义对应的结构体
          ```go
          type CreateController struct {
            beego.Controller
          }

          type DeleteController struct {
            beego.Controller
          }

          type UpdateController struct {
            beego.Controller
          }

          type ReadController struct {
            beego.Controller
          }
          ``` 
        - 根据需要的逻辑，重写beego.Controller的方法，我们这里重写的是Post方法
    - model(M)
      - beego受到django和SQLAlchemy的启发，采用了ORM的思想，在sql语句和go语言操作数据库的接口中添加了一层，简化了编程难度。
      - model分为连接数据库，增删改查
      - 连接数据库
        - 要想成功的连接数据库，就需要搞懂beego的运行顺序
          - 首先执行import中的内容，如果import前面有_的话，就只执行对应package的init函数          
          - 依次递归调用；
          - 我用fmt打印了hello world的顺序，可以帮助理解；
        - 然后开始连接
          - 这里直接按官方的代码是跑不通的，建议按下面的教程走
          - beego目前支持以下的数据库
            - MySQL：github.com/go-sql-driver/mysql
            - PostgreSQL：github.com/lib/pq
            - Sqlite3：github.com/mattn/go-sqlite3
          - 需要改动的代码有两部分models/model.go和main.go：
          - main.go
            - 这个比较简单，就是引入models.go中的package models，和数据库建立连接并建表	
            ```go
            _ "web/models"
             ```
          - models/model.go
            - 首先在models/model.go中import引入两个包
              ```go
              "github.com/beego/beego/v2/client/orm"
  	          _ "github.com/go-sql-driver/mysql"
              ```       
              - 前者是orm框架的包
              - 后者是mysql的驱动
            - 然后开始构建一些数据的结构体
              ```go
              type User struct {
                Id          int
                Name        string
                Profile     *Profile   `orm:"rel(one)"` // OneToOne relation
                Post        []*Post `orm:"reverse(many)"` // 设置一对多的反向关系
              }

              type Profile struct {
                Id          int
                Age         int16
                User        *User   `orm:"reverse(one)"` // 设置一对一反向关系(可选)
              }

              type Post struct {
                Id    int
                Title string
                User  *User  `orm:"rel(fk)"`    //设置一对多关系
                Tags  []*Tag `orm:"rel(m2m)"`
              }

              type Tag struct {
                Id    int
                Name  string
                Posts []*Post `orm:"reverse(many)"` //设置多对多反向关系
              }
              ``` 
            - 接着写init函数
              ```go
              func init() {
                  //
                  fmt.Println("hello world0")
                  // 需要在init中注册定义的model
                  orm.RegisterModel(new(User), new(Post), new(Profile), new(Tag))
                  orm.RegisterDriver("mysql", orm.DRMySQL)
                  orm.RegisterDataBase("default", "mysql", "root:admin@tcp(localhost:3306)/orm_test3?charset=utf8")
                  orm.RunSyncdb("default", false, true)
                  //默认使用 default，你可以指定为其他数据库
                  o := orm.NewOrm()

                  profile := new(Profile)
                  profile.Age = 30

                  user := new(User)
                  user.Profile = profile
                  user.Name = "slene"

                  fmt.Println(o.Insert(profile))
                  fmt.Println(o.Insert(user))
                  fmt.Println("hello world1")
              }
              ```
              - 这里的问题就比较多了
              - 首先需要把orm.RegisterModel放在orm.RegisterDataBase的前面，这个语句代表用上面的结构体来替代表的结构
              - 然后用我们的mysql驱动来作为数据库的驱动，另外安装在centos上安装mysql也有一波坑，可以结合搜索引擎利用递归检索法解决问题；
              - 最后开始连接，'default'参数是不能修改的：因为默认需要一个default参数的数据库，mysql代表数据库的类型，接着一长串是数据库的用户名，密码，连接协议，ip，端口，以及数据库名，和编码格式
              - 最后要在数据库里面建表，但是这一步只能运行一次，之后需要注释掉检测到重复会报错，而在开发过程中，beego是一直运行的，这里是个坑；
              - 最后测试一下：向数据库中写两条数据，然后在mysql中查看一下；
        - 开始CRUD的不归路
          - 前言：[为啥增删改查对应CRUD](https://www.v2ex.com/t/658554)
          - 直接讲CRUD没啥意思，结合contorller来看比较有意义；另外control中的增删改查貌似有更高级的实现方法，时间原因这里说点简单的；
          - 首先引入一些包
            ```go
              "web/models"
              "github.com/beego/beego/v2/client/orm"
            ``` 
            - 前者是要读取models中的结构
            - 后者是要读取orm结构
          - 接着我们要修改go语言中的struct，然后利用orm将对struct的操作转化成对mysql中的table的操作
            - 增 C
              ```go
              func (c *CreateController) Get() {
                  c.TplName = "index.tpl"

                  o := orm.NewOrm()
                  var user models.User
                  profile := new(models.Profile)
                  profile.Age = 18
                  user.Name = "wei"
                  user.Profile = profile
                  _, errPro := o.Insert(profile)
                  id, err := o.Insert(&user)
                  if errPro != nil {
                    fmt.Println(id)
                    fmt.Println("写入失败", err)
                    fmt.Println("写入失败", errPro)
                    c.Ctx.WriteString("写入失败")
                    return
                  }
                  fmt.Println("写入成功")
                  }
              ``` 
              - 首先，controller里面的函数可以绑定一个模板，在后面的views里面会用到
              - 然后要声明一个数据库，这个数据库可以引入配置，进而操作不同的数据库，这里我们就用default数据库
              - 然后声明一个struct结构，注意这里有不同的声明形式，和最后插入的方式相对应
              - 给这个结构体赋值，因为user表中的profile字段不能为空，所以还要声明一个profile结构体
              - 最后插入数据库
              - nil类似于null，但是有更灵活的变化
              - fmt在运行窗口中打印，c.Ctx.WriteString在网页中写入
            - 删除 D
              ```go
              func (c *DeleteController) Get() {
                //c.TplName = "index.tpl"

                o := orm.NewOrm()
                if num, err := o.Delete(&models.User{Id: 8}); err == nil {
                  fmt.Println(num)
                  c.Ctx.WriteString("删除成功")
                  return
                }
                fmt.Println("删除失败")
                c.Ctx.WriteString("删除失败")
              }
              ```
              - 删除的内容大同小异，但是我在官方教程中看到一种炫技的写法
              - 就是在if的判断条件前用;隔开逻辑语句
              - 然后o.Delete会返回两个值前者是影响了多少行，后者是执行过程中得到的错误
              - 声明一个models.User，并选择ID为8的那一条
              - 后面就没啥说的，唯一注意的是加一个retuen
            - 改 U
              ```go
              func (c *UpdateController) Get() {
                  //c.TplName = "index.tpl"

                  o := orm.NewOrm()
                  user := models.User{}
                  user.Id = 5
                  err := o.Read(&user)
                  if err == nil {
                    user.Name = "weiweiwei"
                    _ , err = o.Update(&user)
                    if err != nil{
                      c.Ctx.WriteString("写入失败")
                      fmt.Println("写入失败")
                    }
                    fmt.Println("查找成功")
                  }
                  fmt.Println("查找成功")
              }
              ``` 
              - 改的话需要先查找到这个数据，查找的条件可以写在外面
              - 然后利用Read读取，最后在查找到的情况下更改那个数据
            - 查 R
              ```go
              func (c *ReadController) Get() {
                  //c.TplName = "index.tpl"

                  o := orm.NewOrm()
                  user := models.User{}
                  user.Id = 1
                  err := o.Read(&user, "id")
                  if err != nil{
                    //beego.ControllerInfo{更新失败}
                    fmt.Println("查找失败")
                    return
                  }
                  fmt.Println("查找成功", user)
                  c.Ctx.WriteString(user.Name)
              }
              ```
              - 查比较简单，前面用过好多次了，就不详细说了
    - View(V)
      - 模板渲染
      - 逻辑函数
      - 静态文件
        - 静态文件放置
          - StaticDir["/static"] = "static"
          - go web.SetStaticPath("/static","public")





