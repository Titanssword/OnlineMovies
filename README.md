# OnlineMovies
A website can play the movies on your server.
本项目采用了前后端分离的设计思路，目前仅完成了部分基本功能，前端采用了AngularJS，后端使用Go, 中间通过http进行数据交互。视频资源使用nginx进行部署。
同时，本项目旨在利用内部局域网的带宽，提供在线播放高清视频从而达到视频共享的一种方式，需要部署的视频由个人管理，本项目不涉及视频版权问题。


## Angular_front
前端基于AngularJS的heros进行开发，![相关链接](https://angular.io/tutorial)
### 使用
首先安装nodejs,
进入Angular_front文件夹，
执行`npm install`, 安装相应包
之后`npm start`, 启动，使用浏览器访问http://localhost:4200
### 部分解析
与后端进行交互主要在hero.service.ts文件当中，
`private heroesUrl = 'http://10.199.130.73/api/movies';  // URL to web api`
在本例当中，使用Go项目搭建的后台给的REST API为`http://10.199.130.73/api/movies`,
后面根据自己的地址或者后台进行调整。


## GoWeb_back
后端基于Go语言进行开发。使用了`github.com/gorilla/mux`进行路由管理。
### 使用
首先安装Go
进入GoWeb_back
运行`go get github.com/gorilla/mux`获取相关包
运行`go run main.go` 运行web后台。
### 部分解析
该后台提供了三种API访问接口
		"/api",
		"/api/movies",
		"/api/movies/{todoId}",
主要是`/api/movies`获取所有电影信息，`/api/movies/{todoId}`根据电影ID，Select相应的电影信息。

 在Handler 中，主要通过查找电影资源目录下的文件生成相应的List, 本例中，主要属性有id,name,url
 ```
 type Todo struct {
	Id      int     `json:"id"`
	Name 	string   `json:"name"`
	Url     string   `json:"url"`
}
type TodoList []Todo
```

## 相关问题
### 跨域
 由于采用前后端分离的部署方式，端口不一样，会出现跨域的问题。（Solved）
 例如，前端使用80端口，后端使用8080端口。此时需要在80端口所在的web服务上配置反向代理。
 
#### Apache2服务
 ```
ProxyPass /api http://127.0.0.1:8080/api
ProxyPassReverse /api http://127.0.0.1:8080/api
<Proxy *>
	Order deny,allow
	Allow from all
</Proxy>
```
 将上述代码放入Apache2配置文件的`VirtualHost *:80`代码块中
 
#### Nginx服务
```
location /api {  
	proxy_pass   http://127.0.0.1:8080/api;  

}
```
将上述代码放入Nginx配置文件的`server`代码块中
### 视频格式的适配问题
 对于部分mkv的视频，存在无法播放声音的问题。
### 功能单一
 前端增加分类，分页，查询，等功能。
 数据增加电影封面，评分等相关信息。

## 最后
如果有任何相关问题, 欢迎提Issue, 提Pr.
