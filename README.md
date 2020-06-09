<h2 align="center">

![brand](./static/images/logo_sm.png)

</h2>

<div align="center">

MM-Wiki 是一个轻量级的企业知识分享与团队协同软件，可用于快速构建企业 Wiki 和团队知识分享平台。部署方便，使用简单，帮助团队构建一个信息共享、文档管理的协作环境。
</br>

[![stable](https://img.shields.io/badge/stable-stable-green.svg)](https://github.com/phachon/mm-wiki/) 
[![build](https://img.shields.io/shippable/5444c5ecb904a4b21567b0ff.svg)](https://travis-ci.org/phachon/mm-wiki)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/phachon/mm-wiki/master/LICENSE)
[![platforms](https://img.shields.io/badge/platform-All-yellow.svg?style=flat)]()
[![download_count](https://img.shields.io/github/downloads/phachon/mm-wiki/total.svg?style=plastic)](https://github.com/phachon/mm-wiki/releases) 
[![release](https://img.shields.io/github/release/phachon/mm-wiki.svg?style=flat)](https://github.com/phachon/mm-wiki/releases) 

</div>

# 特点
- 部署方便，基于 golang 编写，只需要下载对于平台下二进制文件执行即可。
- 快速安装程序, 提供方便的安装界面程序，无需任何手动操作。
- 独立的空间，空间是一组文档的集合，一般为公司部门或者团队，空间下的文档相互独立。空间可根据需求设置空间访问级别。
- 支持 markdown 语法写作，支持附件上传。
- 完善的系统权限管理，系统可以自定义角色，并为不同角色授予不同的权限。
- 集成统一登录，本系统支持通过外部系统认证用户, 比如与公司的 LDAP 登录融合。具体请看登录认证功能。
- 邮件通知功能，当开启邮件通知，文档更改会通知所有关注该文档的用户。
- 文档具有分享和下载功能，目前只支持下载 MarkDown 源文件。
- 支持文档全文搜索

# 在线快速了解
- [mm-wiki企业知识分享与团队协同软件](http://wiki.cifaz.com/), 用户名:admin, 密码:mmwiki, 点击不了请复制http://wiki.cifaz.com/

# 安装
## 1. 自助安装

打开 https://github.com/phachon/mm-wiki/releases 找到对应平台的版本下载编译好的压缩包

- Linux 平台

    ```
    # 创建目录
    $ mkdir mm_wiki
    $ cd mm_wiki
    # 以 linux amd64 为例，下载最新版本压缩包
    # https://github.com/phachon/mm-wiki/releases 自行下载 wget http://
    # 解压到当前目录
    $ tar -zxvf mm-wiki-linux-amd64.tar.gz
    # 进入程序安装目录
    $ cd install
    # 执行安装程序，默认端口为 8090，指定其他端口加参数 --port=8087
    $ ./install
    # 浏览器访问 http://ip:8090 进入安装界面，完成安装配置
    # Ctrl + C 停止 install 程序, 启动 MM-Wiki 系统
    $ cd ..
    $ ./mm-wiki --conf conf/mm-wiki.conf
    # 浏览器访问你监听的 ip 和端口
    # 开始 MM-Wiki 的使用之旅吧！
    ```

- Windows 平台

    ```
    # 以 windows amd64 为例，下载最新版本压缩包
    # https://github.com/phachon/mm-wiki/releases 自行下载
    # 手动解压到当前目录
    # 进入 install 目录
    # 双击点开 install.exe 文件
    # 浏览器访问 http://ip:8090 进入安装界面，完成安装配置
    # 关闭刚刚点开的 install 窗口
    # 使用 windows 命令行工具（cmd.exe）进入程序根目录
    $ 执行 mm-wiki.exe --conf conf/mm-wiki.conf
    # 浏览器访问你监听的 ip 和端口
    # 开始 MM-Wiki 的使用之旅吧！
    ```

- Docker 部署
    ```
    # 数据库准备
    # 导入docs/databases/data.sql和docs/databases/table.sql（注：需取消注释data.sql中第一条管理用户插入语句）

    # 两种部署方式可用
    # DockerHub（推荐）
    # 从DockerHub下载v0.1.7版本
    # 新增配置文件，数据存放目录以及Mysql数据库配置在mm-wiki.conf配置文件中设置
    # 挂载配置文件及数据存放目录，启动端口为8080
    # docker run -d -p 8080:8081 -v /data/mm-wiki/conf/:/opt/mm-wiki/conf/ -v /data/mm-wiki/data:/data/mm-wiki/data/ --name mm-wiki eahom/mm-wiki:v0.1.7

    # 本地构建最新代码
    # 构建项目镜像
    # docker build -t mm-wiki-image .
    # 新增配置文件，数据存放目录以及Mysql数据库配置在mm-wiki.conf配置文件中设置
    # 挂载配置文件及数据存放目录，启动端口为8080
    # docker run -d -p 8080:8081 -v /data/mm-wiki/conf/:/opt/mm-wiki/conf/ -v /data/mm-wiki/data/:/data/mm-wiki/data/ --name mm-wiki mm-wiki-image
    ```
## 2. 如果需要，可用 nginx 配置反向代理
```
upstream frontends {
    server 127.0.0.1:8088; # MM-Wiki 监听的ip:port
}
server {
    listen      80;
    server_name wiki.intra.xxxxx.com www.wiki.intra.xxxxx.com;
    location / {
        proxy_pass_header Server;
        proxy_set_header Host $http_host;
        proxy_redirect off;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Scheme $scheme;
        proxy_pass http://frontends;
    }
    # 静态资源交由nginx管理
    location /static {
        root        /www/mm-wiki; # MM-Wiki 的根目录
        expires     1d;
        add_header  Cache-Control public;
        access_log  off;
    }
}
```

# 系统预览

### 1 安装
![install](./static/images/preview/install.png)
### 2 登录
![login](./static/images/preview/login.png)
### 3 系统
![system](./static/images/preview/system.png)
### 4 空间文档
![space](./static/images/preview/space.png)
### 5 编辑文档
![edit](./static/images/preview/edit.png)
### 6 文档分享
![share](./static/images/preview/share.png)

# 使用的一些插件

MM-Wiki 是站在巨人的肩膀上开发的一款软件，是因为系统中使用了非常多优秀的插件，非常感谢这些插件的作者们：

- [bootstrap](https://github.com/twbs/bootstrap)
- [awesome-bootstrap-checkbox](https://github.com/flatlogic/awesome-bootstrap-checkbox)
- [bootstrap-iconpicker](https://victor-valencia.github.com/bootstrap-iconpicker)
- [bootstrap-select](http://silviomoreto.github.io/bootstrap-select)
- [bootstrap-switch](https://bttstrp.github.io/bootstrap-switch)
- [bootstrap-tagsinput](https://github.com/bootstrap-tagsinput/bootstrap-tagsinput)
- [editor.md](https://github.com/pandao/editor.md)
- [layout](http://jquery-dev.com)
- [layer](http://layer.layui.com/)
- [metisMenu](https://github.com/onokumus/metisMenu)
- [morris](http://morrisjs.github.com/morris.js/)
- [popover](https://github.com/sandywalker/webui-popover)
- [scrollup](http://markgoodyear.com/labs/scrollup/)
- [zTreev3](http://treejs.cn/)

# 二次开发

环境要求：go 1.8
```
$ git clone https://github.com/phachon/mm-wiki.git
$ cd mm-wiki
$ go build ./
```

>如果你想为 mm-wiki 贡献代码，请加开发者交流群：922827699

## 贡献者列表
<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://phachon.com"><img src="https://avatars3.githubusercontent.com/u/19726268?v=4" width="100px;" alt=""/><br /><sub><b>phachon</b></sub></a><br /><a href="https://github.com/phachon/mm-wiki/commits?author=phachon" title="Code">💻</a></td>
    <td align="center"><a href="https://github.com/cxgreat2014"><img src="https://avatars2.githubusercontent.com/u/15062548?v=4" width="100px;" alt=""/><br /><sub><b>cxgreat2014</b></sub></a><br /><a href="https://github.com/phachon/mm-wiki/commits?author=cxgreat2014" title="Code">💻</a></td>
    <td align="center"><a href="https://github.com/cifaz"><img src="https://avatars0.githubusercontent.com/u/4531158?v=4" width="100px;" alt=""/><br /><sub><b>ccz</b></sub></a><br /><a href="https://github.com/phachon/mm-wiki/commits?author=cifaz" title="Code">💻</a></td>
    <td align="center"><a href="https://www.linkedin.com/in/wangruoran/"><img src="https://avatars3.githubusercontent.com/u/25990237?v=4" width="100px;" alt=""/><br /><sub><b>Ruoran Wang</b></sub></a><br /><a href="https://github.com/phachon/mm-wiki/commits?author=ruoranw" title="Documentation">📖</a></td>
    <td align="center"><a href="https://github.com/eahomliu"><img src="https://avatars3.githubusercontent.com/u/50134691?v=4" width="100px;" alt=""/><br /><sub><b>eahomliu</b></sub></a><br /><a href="https://github.com/phachon/mm-wiki/commits?author=eahomliu" title="Documentation">📖</a> <a href="https://github.com/phachon/mm-wiki/commits?author=eahomliu" title="Code">💻</a></td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

## 支持
请作者喝杯茶吧！

![wechat](./static/images/preview/wechat_1.png) 

![alipay](./static/images/preview/alipay_2.png)

## 反馈
- 官方 QQ 交流群：853467682
- 如果您喜欢该项目，请 [Star](https://github.com/phachon/mm-wiki/stargazers).
- 如果在使用过程中有任何问题， 请提交 [Issue](https://github.com/phachon/mm-wiki/issues).
- 如果您发现并解决了bug，请提交 [Pull Request](https://github.com/phachon/mm-wiki/pulls).
- 如果您想二次开发，欢迎 [Fork](https://github.com/phachon/mm-wiki/network/members).
- 如果你想交个朋友，欢迎发邮件给 [phachon@163.com](mailto:phachon@163.com).

## License

MIT

谢谢
---

Create By phachon
