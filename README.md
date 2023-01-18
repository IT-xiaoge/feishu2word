# feishu2word
飞书文档分为三种大的类型：知识库、我的空间、共享空间。feishu2word用于下载飞书知识库、我的空间、共享空间中所有权是自己的文档。
## 一、快速开始
feishu2word使用Go语言开发，可适用于个人用户、有管理员权限的企业用户、特定场景下的企业用户（可获得管理员授权的）
### 1.1 申请飞书应用，获取API TOKEN
* 创建应用：登陆[飞书开放平台](https://open.feishu.cn/) - 创建应用 - 创建企业自建应用（应用名称和应用描述随意填写就好）  
* 应用发布：完善应用信息 - 应用发布 - 前往发布 - 创建版本 - 保存 - 申请线上发布 
* 权限申请：权限管理 - 权限配置 - 云文档 - 批量申请开通权限
* 启用网页：开发者后台-启用功能-网页，开启网页功能并配置桌面端主页
* 安全设置：开发者后台-安全设置-重定向URL，添加重定向URL
* 获取token：API调试台-user_access_token-刷新-查看（2小时过期）

至此，飞书token获取完毕
### 1.2 安装feishu2word
1）编译安装  
```golang
$ echo $GOPATH  
$ cd $GOPATH  
$ git clone https://github.com/IT-xiaoge/feishu2word.git
$ cd feishu2word  
$ go build  
```
2）下载安装  
在 [Release](https://github.com/IT-xiaoge/feishu2word/releases) 中下载合适的版本  

### 1.3 下载飞书文档
查看帮助文档：
```golang
$ ./feishu2word -h
feishu wiki download

Usage:
  feishu2word [flags]

Examples:
feishu2word -p xxxx -t xxxx

Flags:
  -h, --help               help for feishu2word
  -p, --savepath string    feishu wiki save to path
  -t, --usertoken string   feishu app user_access_token
```
开始下载文档：
```golang
feishu2word -p {user_access_token} -t {savepath}
```
## 二、支持功能  
* 飞书知识空间wiki下载：docx、execl
* 飞书个人空间或共享空间doc下载：docx、execl
* 失败文档信息打印，失败的文档需要手动去操作下载
## Q&A
### 1）为什么会有下载失败的情况？
当前飞书官方仅支持部分文档类型下载，PPT、XMIND等不支持
### 2）为什么文档中代码块展示异常、缺失部分内容
代码块展示异常是因为我没有做这块功能，部分内容缺失部分原因是飞书官方不支持，部分原因是我没有支持






