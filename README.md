## 博客简介

基于Golang后端、Vue3前端的前后端分离全栈博客项目。后端技术栈包括Golang、Gin、Gorm、MySQL等。前端技术栈包括Vue3、TypeScript、Axios、Element Plus等。

## 在线预览

博客链接：[hitori.cn](www.hitori.cn)（未适配移动端）

![image](https://github.com/blockcheDev/blog-web/assets/89156012/afb8b63b-88c9-423a-abfd-27b5a590a7b1)

![image](https://github.com/blockcheDev/blog-web/assets/89156012/2961969d-f0a8-412b-b177-12b3659126ea)

![image](https://github.com/blockcheDev/blog-web/assets/89156012/c607b701-79ff-40ea-9e2d-ad60252c8124)

## 项目介绍

本项目是本人的第一个web项目，项目还在不断迭代中。前端界面主打一个简洁但不简陋，做了一些动效设计，后端在能力范围内尽可能做到高安全、优性能。

项目结构：根目录为golang后端项目，根目录中的web文件夹内为vue3前端项目。

后端：

- 使用JWT鉴权
- 拥有管理员、普通用户两种权限
- 评论和评论回复功能（待实现）
- 高频调用使用Redis优化性能（待实现）
- 防SQL注入（待实现）

前端：

- 使用Element Plus组件库
- 文章编辑使用可视化Markdown编辑器
- 界面好看

### 使用技术

后端技术栈：

- Golang
- Gin
- Gorm
- MySQL
- Redis（待实现）

前端技术栈：

- Node.js
- Vue3
- TypeScript
- Element Plus
- Axios
- Nginx

## 启动项目

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







