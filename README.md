# GeeGoWeb

> 学习使用go搭建一个类似gin 的web框架 

## day1 基础框架的搭建

 [classroom](https://geektutu.com/post/gee-day1.html)

[code exercise](https://github.com/lijunxin559/GeeGoWeb)

1. 搭建的思路：使用http.HandleFunc太重复->封装成struct engine使用serveHTTP集成方法->为gee添加New,Run,GET,POST方法。

2. 框架雏形

   ```go
   gee/
     |--gee.go
     |--go.mod
   main.go
   go.mod
   ```

3. 使用go.mod解决依赖，[go.mod学习传送](https://www.jianshu.com/p/760c97ff644c)

4. 最终实现了路由映射表，提供了用户注册静态路由的方法，包装了启动服务的函数。



## day2 Context上下文

 [classroom](https://geektutu.com/post/gee-day2.html)

[code exercise](https://github.com/lijunxin559/GeeGoWeb)

1.将`路由(router)`独立出来，方便之后增强

2.设计`上下文(Context)`，封装 Request 和 Response ，提供对 JSON、HTML 等返回类型的支持

形成一个包含: context<-router<-engine（<-代表调用）



## day3 前缀树定义动态路由

 [classroom](https://geektutu.com/post/gee-day3.html)

[code exercise](https://github.com/lijunxin559/GeeGoWeb)

1.主要就是前缀树的设计存储：怎么去匹配，怎么去插入，怎么去记录（感觉这部分还要多看看）

2.在context中需要一个解析param的方法


## day4 分组路由控制Group

 [classroom](https://geektutu.com/post/gee-day4.html)

[code exercise](https://github.com/lijunxin559/GeeGoWeb)


1.使用分组进行嵌套定义路由

2.Engine作为顶层，包含了所有的RouterGroup和router

3.RouterGroup定义分组处理的engine，并为后续做处理

4.注意：其实目前的所有信息都还是写在同一张router里面的，只是根据路由分组进去的接口不一样了


