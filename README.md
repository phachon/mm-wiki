# MM-Wiki v2

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go)](https://go.dev/)
[![React](https://img.shields.io/badge/React-18-61DAFB?logo=react)](https://react.dev/)

**MM-Wiki** 是一个轻量级的企业知识分享与团队协同软件，适用于企业内部 Wiki、知识库、文档管理等场景。

v2 版本基于 Go + React 全新重写，具备更现代化的技术栈和更灵活的架构设计。

---

## ✨ 功能特性

- 📁 **空间与文档管理** — 支持多空间、树形文档结构，灵活组织知识体系
- ✏️ **Markdown 编辑** — 集成 Cherry Markdown 和 Monaco Editor，支持实时预览
- 👥 **用户与权限管理** — 完善的用户、角色、权限三层模型，支持 RBAC 权限控制
- 🏢 **部门管理** — 支持组织架构管理，按部门分配人员
- 📧 **邮件服务配置** — 可配置邮件服务器，支持通知推送
- 🔗 **快捷链接与联系人** — 便捷的常用链接和联系人管理
- 📢 **系统公告** — 支持公告发布与管理
- 🔐 **登录认证** — 支持本地登录和外部 SSO 认证
- 📝 **文档版本控制** — 完整的文档历史版本管理，支持版本对比与恢复
- ⭐ **收藏与关注** — 用户可收藏文档和空间，关注感兴趣的内容
- 📎 **文件附件** — 支持文档附件上传与管理
- 🔌 **插件系统** — 可扩展的插件机制
- 📊 **操作日志** — 完整的操作审计日志
- 🛠️ **安装向导** — 首次使用提供交互式安装引导

---

## 🖼️ 截图预览

> 截图待补充

---

## 🛠️ 技术栈

### 后端

| 技术 | 说明 |
|------|------|
| [Go 1.20](https://go.dev/) | 后端开发语言 |
| [Gin](https://gin-gonic.com/) | HTTP Web 框架 |
| [GORM](https://gorm.io/) | ORM 框架 |
| [MySQL 8.0](https://www.mysql.com/) | 关系型数据库 |
| [JWT](https://github.com/golang-jwt/jwt) | 登录认证 |
| [Zap](https://github.com/uber-go/zap) | 结构化日志 |
| [Lumberjack](https://github.com/natefinch/lumberjack) | 日志文件轮转 |

### 前端

| 技术 | 说明 |
|------|------|
| [React 18](https://react.dev/) | 前端框架 |
| [TypeScript](https://www.typescriptlang.org/) | 类型安全的 JavaScript |
| [Ant Design 5](https://ant.design/) | UI 组件库 |
| [Zustand](https://github.com/pmndrs/zustand) | 轻量级状态管理 |
| [Monaco Editor](https://microsoft.github.io/monaco-editor/) | 代码编辑器 |
| [Cherry Markdown](https://github.com/Tencent/cherry-markdown) | Markdown 编辑器 |
| [React Router v6](https://reactrouter.com/) | 前端路由 |
| [Axios](https://axios-http.com/) | HTTP 请求 |

---

## 🚀 快速开始

### 环境要求

- **Go** 1.20+
- **Node.js** 18+（推荐使用 LTS 版本）
- **MySQL** 8.0+

### 1. 克隆项目

```bash
git clone https://github.com/phachon/mm-wiki.git
cd mm-wiki
```

### 2. 数据库准备

**方式一：使用 Docker Compose（推荐）**

```bash
# 启动 MySQL 容器，自动创建数据库和导入表结构与初始数据
make docker-up
```

Docker Compose 会自动完成以下操作：
- 创建 `mm_wiki2` 数据库
- 导入表结构（`backend/docs/database/mm_wiki2.sql`）
- 导入初始数据（`docs/database/data.sql.txt`）

**方式二：手动安装 MySQL**

```bash
# 1. 创建数据库
mysql -u root -p -e "CREATE DATABASE IF NOT EXISTS mm_wiki2 DEFAULT CHARSET utf8mb4;"

# 2. 导入表结构
mysql -u root -p mm_wiki2 < backend/docs/database/mm_wiki2.sql

# 3. 导入初始化数据
mysql -u root -p mm_wiki2 < docs/database/data.sql.txt
```

### 3. 启动后端服务

```bash
cd backend
make dev
```

后端默认监听 `http://127.0.0.1:8088`，使用 `conf/dev/app.yaml` 配置文件。

### 4. 启动前端服务

```bash
cd front
npm install
npm start
```

前端默认运行在 `http://localhost:3000`，开发模式下自动代理 API 请求到后端。

### 5. 访问系统

打开浏览器访问 [http://localhost:3000](http://localhost:3000)

默认管理员账号：
- 用户名：`root`
- 密码：`111111`

---

## 📂 项目结构

```
mm-wiki/
├── Makefile                  # 根目录 Make 命令（开发、构建、Docker）
├── docker-compose.yml        # Docker Compose 配置（MySQL 服务）
├── README.md                 # 项目说明文档
├── backend/                  # 后端服务（Go）
│   ├── main.go               # 程序入口
│   ├── version.go            # 版本号
│   ├── Makefile              # 后端 Make 命令
│   ├── Dockerfile            # 后端 Docker 镜像
│   ├── build.sh              # 构建脚本
│   ├── deploy.sh             # 部署脚本
│   ├── conf/                 # 配置文件目录
│   │   ├── dev/app.yaml      # 开发环境配置
│   │   ├── test/app.yaml     # 测试环境配置
│   │   └── prod/app.yaml     # 生产环境配置
│   ├── config/               # 配置解析
│   ├── router/               # 路由定义
│   ├── filter/               # 中间件（认证、权限、限流、跨域）
│   ├── global/               # 全局变量与命令行参数
│   ├── logger/               # 日志初始化
│   ├── app/                  # 业务逻辑
│   │   ├── entity/           # 数据实体（表结构映射）
│   │   ├── dao/              # 数据访问层
│   │   ├── service/          # 业务服务层
│   │   ├── controller/       # 控制器（API 处理）
│   │   ├── cache/            # 缓存
│   │   ├── event/            # 事件处理
│   │   └── worker/           # 后台任务
│   ├── gopkg/                # 内部工具包
│   │   ├── errors/           # 错误处理
│   │   ├── log/              # 日志工具
│   │   ├── mail/             # 邮件发送
│   │   └── upload/           # 文件上传
│   ├── utils/                # 工具函数
│   ├── script/               # 运维脚本
│   └── docs/                 # 文档
│       └── database/         # 数据库表结构 SQL
├── front/                    # 前端应用（React）
│   ├── package.json          # 依赖与脚本
│   ├── craco.config.js       # Webpack 配置覆盖
│   ├── tsconfig.json         # TypeScript 配置
│   ├── Dockerfile            # 前端 Docker 镜像
│   ├── nginx.conf            # Nginx 配置（生产部署）
│   ├── src/
│   │   ├── index.tsx         # 应用入口
│   │   ├── pages/            # 页面组件
│   │   ├── components/       # 通用组件
│   │   ├── services/         # API 接口封装
│   │   ├── stores/           # Zustand 状态管理
│   │   ├── router/           # 路由配置
│   │   ├── config/           # 前端配置
│   │   ├── types/            # TypeScript 类型定义
│   │   ├── utils/            # 工具函数
│   │   ├── assets/           # 静态资源
│   │   ├── theme/            # 主题配置
│   │   └── setupProxy.js     # 开发代理配置
│   └── public/               # 公共静态文件
└── docs/                     # 项目文档
    └── database/             # 数据库初始化数据
```

---

## 💻 开发指南

### Make 命令

项目根目录提供了统一的 Make 命令：

```bash
make help               # 查看所有可用命令
make dev                # 提示同时启动前后端（需两个终端）
make dev-backend        # 启动后端开发服务
make dev-frontend       # 启动前端开发服务
make build              # 构建前后端
make build-backend      # 仅构建后端
make build-frontend     # 仅构建前端
make docker-up          # 启动 MySQL Docker 容器
make docker-down        # 停止 MySQL Docker 容器
make lint               # 代码检查
make test               # 运行后端测试
make clean              # 清理构建产物
```

### 后端 Make 命令

```bash
cd backend
make dev                # 使用 go run 启动（开发模式）
make build              # 编译为二进制文件
make run                # 编译并运行
make test               # 运行所有测试
make lint               # 代码检查（golangci-lint 或 go vet）
make linux              # 交叉编译 Linux amd64
make windows            # 交叉编译 Windows amd64
make darwin             # 交叉编译 macOS amd64
```

通过 `ENV` 变量切换环境配置：

```bash
make dev ENV=dev        # 开发环境（默认）
make dev ENV=test       # 测试环境
make dev ENV=prod       # 生产环境
```

---

## 🐳 Docker 部署

### 后端镜像

```bash
cd backend
docker build -t mm-wiki-backend .
docker run -d -p 8088:8088 mm-wiki-backend
```

### 前端镜像

```bash
cd front
docker build -t mm-wiki-frontend .
docker run -d -p 80:80 mm-wiki-frontend
```

### 构建脚本

后端提供了 `build.sh` 和 `deploy.sh` 脚本用于打包和部署：

```bash
cd backend

# 构建发布包（默认 prod 环境）
sh build.sh

# 指定环境构建
sh build.sh dev

# 构建产物目录结构：
# release/
#   ├── bin/mm-wiki     # 可执行文件
#   ├── conf/           # 配置文件
#   ├── logs/           # 日志目录
#   └── script/         # 运维脚本
```

### 运维脚本

`backend/script/` 目录包含常用运维脚本：

| 脚本 | 说明 |
|------|------|
| `start.sh` | 启动服务 |
| `stop.sh` | 停止服务 |
| `restart.sh` | 重启服务 |
| `monitor.sh` | 监控服务 |
| `pack.sh` | 打包脚本 |
| `mysql_docker_install.sh` | Docker 安装 MySQL |
| `mysql_dump.sh` | MySQL 备份脚本 |

---

## ⚙️ 配置参考

后端配置文件位于 `backend/conf/{env}/app.yaml`，支持 `dev`、`test`、`prod` 三种环境。

```yaml
# 全局配置
Global:
  env: dev                    # 运行环境
  gin_mode: debug             # Gin 模式: debug | release | test

# 服务配置
Server:
  ip: 127.0.0.1              # 监听 IP
  port: 8088                  # 监听端口
  read_timeout: 2000          # 读超时（毫秒）
  write_timeout: 2000         # 写超时（毫秒）

# 数据库配置
Database:
  mm_wiki:
    host: 127.0.0.1           # MySQL 地址
    port: 3306                # MySQL 端口
    name: mm_wiki2            # 数据库名
    user: root                # 用户名
    pass: "123456"            # 密码
    table_prefix: "mk_"      # 表前缀
    conn_max_idle: 20         # 最大空闲连接数
    conn_max_connection: 200  # 最大连接数
    conn_max_lifetime: 200    # 连接最大存活时间（秒）

# 文件上传配置
Upload:
  doc_file:
    type: local               # 上传方式: local | qiniu | aliyun
    local_dir: ./static/upload # 本地存储目录
    domain: http://localhost:8088 # 访问域名

# 日志配置
Logger:
  default:
    - writer: console         # 控制台输出
      level: debug
    - writer: file            # 文件输出
      level: info
      formatter: json
      writer_config:
        filename: ./logs/access.log
        max_age: 7            # 保留天数
        max_backups: 10       # 最大备份数
        compress: false       # 是否压缩
        max_size: 10          # 单文件最大大小（MB）

# 认证配置（可选）
Auth:
  jwt_secret: ""              # JWT 密钥（为空使用默认值）
  expire_hours: 3             # Token 过期时间（小时）

# 跨域配置（可选）
CORS:
  allow_origins: []           # 允许的来源，为空则允许所有
  allow_methods: []           # 允许的 HTTP 方法
  allow_headers: []           # 允许的请求头
```

命令行参数：

```bash
./mm-wiki -conf ./conf/dev/app.yaml   # 指定配置文件路径
./mm-wiki -env dev                     # 指定运行环境
```

也可通过环境变量 `KAPP_ENV_TYPE` 设置运行环境。

---

## 🤝 参与贡献

欢迎提交 Issue 和 Pull Request 参与项目贡献！

1. Fork 本仓库
2. 创建你的功能分支（`git checkout -b feature/amazing-feature`）
3. 提交你的修改（`git commit -m 'feat: add amazing feature'`）
4. 推送到分支（`git push origin feature/amazing-feature`）
5. 创建 Pull Request

---

## 📄 开源协议

本项目基于 [MIT License](front/LICENSE) 开源。

Copyright © 2023 [phachon](https://github.com/phachon)
