## 博客简介

基于Golang后端、Vue3前端的前后端分离全栈博客项目。后端技术栈包括Golang、Gin、Gorm、MySQL、Redis等。前端技术栈包括Vue3、TypeScript、Axios、Element Plus等。

## 在线预览

博客链接：[hitori.cn](https://www.hitori.cn)（已配置ssl，可使用https访问）

![image](https://github.com/blockcheDev/blog-web/assets/89156012/afb8b63b-88c9-423a-abfd-27b5a590a7b1)

![image](https://github.com/blockcheDev/blog-web/assets/89156012/2961969d-f0a8-412b-b177-12b3659126ea)

![image](https://github.com/blockcheDev/blog-web/assets/89156012/c607b701-79ff-40ea-9e2d-ad60252c8124)

## 项目介绍

本项目是本人的第一个web项目，项目还在不断迭代中。前端界面主打一个简洁但不简陋，做了一些动效设计，后端在能力范围内尽可能做到高安全、优性能。

项目结构：根目录为golang后端项目，根目录中的web文件夹内为vue3前端项目。

后端：

- 用户身份使用JWT鉴权，使用HS256签名算法，保证token不可被伪造。
- 数据库存储密码前使用bcrypt算法给密码加盐，防止彩虹表破解。
- 可以使用Github授权登录博客。
- 高频调用使用Redis优化性能：缓存文章列表
- 防SQL注入[（原理看底下的开发问题日志）](#SQL)

前端：

- 使用Element Plus组件库
- 文章编辑使用可视化Markdown编辑器
- 界面好看

其他容器：

- Python脚本每天凌晨4:00自动将服务器的MySQL数据备份至腾讯云COS，可随时获取对应日期的博客数据。
- Python脚本每间隔一个小时从后台拉取博客文章列表，更新sitemap.xml(https://www.hitori.cn/sitemap.xml)

### 使用技术

后端技术栈：

- Golang
- Gin
- Gorm
- MySQL
- Redis
- Nginx
- Python
- Docker

前端技术栈：

- Node.js
- Vue3
- TypeScript
- Element Plus
- Axios

## 启动项目

### 现在可以使用docker-compose一键启动了

#### 1. 要先补全缺少的配置（为了保密所以没有放上来）

- deploy/web/hitori.cn_bundle.crt（域名的SSL证书，替换为自己的）
- deploy/web/hitori.cn.key（同上）

- deploy/.env
```
DATA_DIRECTORY=./blockche-blog

REDIS_PORT=6379

MYSQL_ROOT_PASSWORD=xxxx
MYSQL_PORT=3306

SERVER_PORT=xxxx

COS_SECRET_ID=xxxx
COS_SECRET_KEY=xxxx
```

- deploy/mysql/blog.sql（用于数据库初始化）

- server/config/config.yaml
```yaml
github:
  client_id: "xxxx"
  client_secret: "xxxx"

mysql:
  address: "172.17.0.1:3306"
  user: "root"
  password: "xxxx"
  database: "blog"

redis:
  address: "172.17.0.1:6379"
  password: "xxxx"
  database: 0 # use default DB

jwt:
  key: "xxxx"
```

#### 2. 一键启动

```shell
cd deploy
docker-compose up --build
```


### 手动启动

需要安装Golang、Node.js、MySQL、Redis环境

后端项目运行：

```shell
# 项目根目录
go mod tidy
go run main.go
```

前端项目运行：

```shell
# 进入web目录
cd ./web

# 安装依赖
npm install

# 运行
npm run dev
```

### BUG

#### 已修复

- 注销用户时传密码到后端失败：使用axios，delete请求传data的方法和post、put不太一样。



### 开发问题日志

#### 配置ssl后，使用https前端无法访问http后端接口

下载ssl证书，在Nginx的nginx.conf上配置443端口服务后，就可以访问到前端了。

```conf
    # HTTPS server
    server {
       listen       443 ssl;
       server_name  hitori.cn;

       ssl_certificate      /usr/local/nginx/conf/cert/hitori.cn.pem;
       ssl_certificate_key  /usr/local/nginx/conf/cert/hitori.cn.key;

       ssl_session_cache    shared:SSL:1m;
       ssl_session_timeout  5m;

       ssl_ciphers  HIGH:!aNULL:!MD5;
       ssl_prefer_server_ciphers  on;

        # https代理到http后端服务,注意前端访问后端的地址要改为https://www.hitori.cn/api
       location /api {
            proxy_pass http://www.hitori.cn:8080/api;
        }

       location / {
           root   /usr/local/blog-web/dist;
           index  index.html index.htm;
           try_files $uri $uri/ /index.html;
       }
    }
```

需要注意的是，由于前端是https，浏览器是不允许访问后端的http接口的，所以我们还需要用Nginx将https的请求反向代理到后端的http接口，即加上这段配置：

```
        # https代理到http后端服务,注意前端访问后端的地址要改为https://www.hitori.cn/api
       location /api {
            proxy_pass http://www.hitori.cn:8080/api;
        }
```

添加这段配置后，当前端访问 https://www.hitori.cn/api 时，Nginx会将其代理到 http://www.hitori.cn:8080/api ，所以前端axios的后端地址配置也要修改为https方式了。



#### 接入Github OAuth授权登录

1. 向Github发送post请求获取access_token时，需要在header设置"Accept"为"application/json"，这样Github才会返回json格式的数据，否则默认是无格式的字符串。

```go
// 获取access_token
url := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", client_id, client_secret, code)
req, _ := http.NewRequest("POST", url, nil)
req.Header.Set("Accept", "application/json")
res, err := http.DefaultClient.Do(req)
// res.Body内就是json数据
```



#### 关于防SQL注入<a id="SQL"></a>

这是一个小插曲，本来已经做好自己写过滤模块的准备了，结果 bing 一搜发现 Gorm 本身就自带防 SQL 注入哈哈，省事了。

**使用：**

只要使用 Gorm API 自带的占位符拼接 SQL 语句，Gorm 就会使用预编译功能，即先传入带占位符的SQL语句，完成预编译，再传入参数。

**原理：**

数据库会先根据带占位符的SQL语句完成词法、语法、语义分析，生成语法树，此时语法树的结构已经固定了，再传入参数是无法改变语义的，就算参数是`1 or 1=1`也只会被当成普通字符串。
