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

