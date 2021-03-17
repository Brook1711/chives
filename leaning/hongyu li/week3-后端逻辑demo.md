# 网站的后端初探

根据网站的查找和发送弹幕的功能，后端的逻辑基本可以分为收发两个模块，需要和前端统一一下接口API的参数：

## 网站后端具体函数逻辑

1. 模块一，收：
  - 功能：由前端传递字符串，后端在数据库中查找弹幕并返回
  - 伪代码：
    ```go
    fun(str 前端传递的电影名称):
        return 返回Json格式的弹幕
    ```
2. 模块二，发：
  - 功能：将前端传递来的弹幕插入到数据库中
  - 伪代码
    ```go
    fun(json 前端传递的json格式的弹幕):
        return 是否成功插入的消息
    ```
3. 补充：
  - 未来在收模块中查找弹幕方面会应用hash匹配等技术实现精确匹配，这里就先采用字符串匹配做demo

## 网站后端数据库

1. 数据库选型
   - 数据库发展一览
     ![数据库发展](https://raw.githubusercontent.com/Richardhongyu/pic/main/20210315183926.png)
     - 视频链接在[这里](https://university.pingcap.com/courses/PCTA/chapter/%E6%A6%82%E8%BF%B0/lesson/%E5%88%86%E5%B8%83%E5%BC%8F%E6%95%B0%E6%8D%AE%E5%BA%93%E5%8E%86%E5%8F%B2%E5%92%8C%E5%8F%91%E5%B1%95%E8%B6%8B%E5%8A%BF)。归纳一下视频的内容，数据库从开始发展到现在，大体上经历了三代：
     - 以关系型数据库一统江湖为开始，mysql是典型的代表，后面跟着的还有：postgressql，sql server，oracle，IBM......
       - 优点：通用，广泛，可靠
       - 缺点：在面对互联网多样化的存储内容，丰富的可拓展性，大规模的数据时力不从心
     - 然后nosql数据库横空出世
       - 优点：存储能力，性能，可拓展性
       - 缺点：处理复杂业务能力不强，表征数据能力一般，一致性不够好
     - 再接着Newsql初出茅庐，对上面两个杂交了一下。国产tidb是这个时代的领头羊。（三个联合创始人之一还是北邮校友，最近融资2亿美元，已然财富自由。如果我们现在退学加入，没准相当于加入16年的字节[手动狗头]）
       - 优点：杂交就比较耐，结合了双方的优势。
       - 缺点：这个老哥没说。。。推测一下，既然是结合品，那在父母双方擅长的地方这个结合品就有一定的劣势了，说白了还是一个trade-off式的问题。
   - 弹幕如果是以json格式存储，最好是用nosql。因为是键值对存储。
     - 用beego支持的postgres的话
       - 优点：可以满足我们的需求，达到了及格线；官方支持较好
       - 缺点：只是可以满足，但是性能上不是最优解；而且需要从头写代码
     - 但是，鑫鑫目前找到了一个开源的项目，可以解决很多问题。这个方案中的数据库是mongodb。
       - 优点：符合我们对数据库的需求，性能满足需求；代码可以复用，完美解决了上面postgres的问题
       - 缺点：beego的支持性上有点问题，不过我找到了解决方案，后面我具体会看如何连接mongodb。
     - 综上，我觉得我们应该用mongodb
2. beego目前对数据库的支持情况以及拓展解决方案
   - beego orm官方只支持三种数据库:mysql，postgres，sqlite3，不巧，都是关系型数据库。
     - 并且只支持postgreps的json字段（不是mysql不支持json字段，是beego orm不支持建立json字段的mysql表）
      ![图片](https://raw.githubusercontent.com/Richardhongyu/pic/main/20210315183219.png)
   - 如上所述，beego对数据库的支持是有限的，但是我们可以用一些其他的手段来解决beego官方贫瘠的支持情况，方式如下
     - 非ORM
       - orm只是一个中间层，如果不想要（或者不支持对应的数据库），我们就可以抽去这一层。Go 没有内置的驱动支持任何的数据库，但是Go定义了database/sql接口，用户可以基于驱动接口开发相应数据库的驱动。orm封装了这些驱动，我们可以直接用这些driver，而不经过orm。需要注意的是orm本身有一些缺点，可以参考这个[文章](http://www.hydrogen18.com/blog/golang-orms-and-why-im-still-not-using-one.html)
       - 例如，想要连接mongodb的话，可以直接import"gopkg.in/mgo.v2"来操作，这个driver的[官网](http://labix.org/mgo)
     - ORM
       - 如果坚持使用orm有两种思路，一个是尝试新的driver，另一个是选择一个可以支持我们的数据库的orm
       - beego orm官方只支持三种驱动，但是由于beego orm是支持database/sql标准接口的 ORM库，所以理论上来说，只要数据库驱动支持database/sql接口就可以无缝的接入beego orm。我还没来得急测试mongodb driver和beego orm的兼容性。
       - 同时beego并不是一定要用beego orm，所以我们可以用其他支持我们数据库/或者支持某个功能的orm，几个有名的orm的支持情况可以大体参考这个[网站](https://studygolang.com/articles/13563)和[这个](https://www.zhihu.com/question/55072439)，或者直接去官网看。具体的orm有：国产的[gorm](https://github.com/go-gorm/gorm)，这个支持的不多，MySQL, PostgreSQL, SQlite, SQL Server，但是它可以支持json字段的导入；国产的[xorm](https://xorm.io/)，这个orm的源码现在不在github上，在gitea上（这个gitea是go写的类github平台），这个支持的数据库有很多，甚至包括tidb；upper/db，它支持mongodb
3. 连接数据库
   - postgresql，一开始我不知道有mongodb的开源项目，连接了这个数据库，虽然后面不用了，但是记录一下踩过的坑，后面没准有用
     - 安装数据库教程
       - 机器环境：centos  7.8.2003
       - [安装教程1](https://ken.io/note/centos7-postgresql12-install-and-configuration)
       - bug1
         - 报错信息：psql: FATAL: Ident authentication failed for user “postgres”。这个报错有两个版本，一个是这个Ident，还有一个是peer。发生在你连接数据库的时候。解决思路一致。
         - 巨坑来源：/var/lib/pgsql/12/data/pg_hba.conf。如果你按上面的教程找到这个文件的话，你就能发现这个问题的原因。
         - 解决方案
           - [官方救命法宝](https://www.postgresql.org/docs/9.5/static/client-authentication.html)
           - [stackoverflow救命法宝](https://stackoverflow.com/questions/2942485/psql-fatal-ident-authentication-failed-for-user-postgres)
           - 简单来说，这个问题就是因为数据库的权限问题。如果需要彻底根治且不留安全隐患，需要读懂那个配置文件。我主要是卡在local和host在哪种情况下会生效了就采用了简单的方法：把数据库访问方式全部设置为MD5。简单粗暴且有效。
       - 连接数据库
         - 这里直接贴代码就好了
         ```go
         orm.RegisterDriver("postgres", orm.DRPostgres)
         orm.RegisterDataBase("default", "postgres","postgres://postgres:123456@localhost:5432/orm_test_2?sslmode=disable")```
        - 这里的问题是不能加utf-8编码设置，以及sslmode需要设置为disable。另外这个地方的设置可以用”A=B C=D“的格式来写，不一定要用str的格式。
   - mongodb数据库
     - still working

## 部分参考资料

1. [Go语言电子书](https://learnku.com/docs/the-way-to-go/basic-types-and-operators/3586)
2. [Go web开发电子书](https://learnku.com/docs/build-web-application-with-golang/072-json-processing/3196)
3. [官方关于json的介绍内容](https://blog.golang.org/json)


   
