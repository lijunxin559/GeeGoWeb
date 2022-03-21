# GeeGoWeb

> 学习使用go搭建一个类似gin 的web框架 

## day1 基础框架的搭建

 [classroom](https://geektutu.com/post/gee-day1.html)

[code exercise](https://github.com/lijunxin559/GeeGoWeb/tree/main/http-base)

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

[code exercise](https://github.com/lijunxin559/GeeGoWeb/tree/main/context)

1.将`路由(router)`独立出来，方便之后增强

2.设计`上下文(Context)`，封装 Request 和 Response ，提供对 JSON、HTML 等返回类型的支持

形成一个包含: context<-router<-engine（<-代表调用）



## day3 前缀树定义动态路由

 [classroom](https://geektutu.com/post/gee-day3.html)

[code exercise](https://github.com/lijunxin559/GeeGoWeb/tree/main/router)

1.主要就是前缀树的设计存储：怎么去匹配，怎么去插入，怎么去记录（感觉这部分还要多看看）

2.在context中需要一个解析param的方法


## day4 分组路由控制Group

 [classroom](https://geektutu.com/post/gee-day4.html)

[code exercise](https://github.com/lijunxin559/GeeGoWeb/tree/main/group)


1.使用分组进行嵌套定义路由

2.Engine作为顶层，包含了所有的RouterGroup和router

3.RouterGroup定义分组处理的engine，并为后续做处理

4.注意：其实目前的所有信息都还是写在同一张router里面的，只是根据路由分组进去的接口不一样了


## day5 中间件MiddleWare

[classroom](https://geektutu.com/post/gee-day5.html)

[code exercise](https://github.com/lijunxin559/GeeGoWeb/tree/main/middleware)

1.收到请求之后，请求的所有信息保存在Context之中，所有应作用于该路由器的中间件也保存在Context之中，依次进行调用，保存在Context中的目的是为了能在handler处理前后都有机会执行中间件，这个根据用户定义

2.中间件函数的编写分为上下两个部分，即handler的前后

3.c.Next()中每次执行完一个中间件函数，将context权限交给下一个函数，最后执行路由的handler，最后再递归调用回来

4.定义Use函数，将中间件应用到某个Group上

5.得到一个具体的请求之后，去匹配每一个group，将相应的中间件函数赋值给context，执行c.Next()

## day6 HTML模版渲染

[classroom](https://geektutu.com/post/gee-day6.html)

[code exercise](https://github.com/lijunxin559/GeeGoWeb/tree/main/template) 

1.*template.Template 和 template.FuncMap对象，前者将所有的模板加载进内存，后者是所有的自定义模板渲染函数

2.Context中指定成员变量engine，能够通过engine访问HTML模版

3.特别注意当前的目录架构
static下放置所有的资源文件，访问的时候使用/assets/filename进行访问，即服务器和用户路径访问分离

```
---gee/
---static/
   |---css/
        |---geektutu.css
   |---file1.txt
   |---file2.txt
   |---file3.txt 
---templates/
   |---arr.tmpl
   |---css.tmpl
   |---custom_func.tmpl
---main.go

```

## day7 错误恢复

[classroom](https://geektutu.com/post/gee-day7.html)

[code exercise](https://github.com/lijunxin559/GeeGoWeb/tree/main/error-deal) 

1.golang 中panic可以通过defer中的recover恢复，定义Recovery中间件函数，先执行router  deal调用，然后捕捉处理错误，恢复到监听一层

2.定义trace函数，可以使用runtime.Callers跳过几层调用之后，打印函数栈过程


  <em>至此我们就完成了一个类似gin的web框架！完结撒花！～</em>

>>再次感谢原作者的详细讲解，一套下来之后，理解加深很多！