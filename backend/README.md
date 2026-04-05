# MM-Wiki 后端服务

MM-Wiki 后端基于 Go 语言开发，使用 Gin Web 框架，采用分层架构设计，提供 RESTful API 接口。

---

## 🏗️ 架构概述

后端采用经典的四层架构设计：

```
Entity（实体层）→ DAO（数据访问层）→ Service（业务逻辑层）→ Controller（控制器层）
```

| 层级 | 职责 | 目录 |
|------|------|------|
| **Entity** | 定义数据库表对应的结构体，是与数据库表的一一映射 | `app/entity/` |
| **DAO** | 封装数据库 CRUD 操作，通过 GORM 操作 MySQL | `app/dao/` |
| **Service** | 处理核心业务逻辑，组合调用多个 DAO，事务管理 | `app/service/` |
| **Controller** | 处理 HTTP 请求，参数校验，调用 Service，返回响应 | `app/controller/` |

**请求处理流程：**

```
HTTP Request → Gin Engine → Middleware（Filter）→ Router → Controller → Service → DAO → MySQL
```

中间件（Filter）按顺序处理：
1. **CORS** — 跨域处理
2. **Logger** — 请求日志
3. **Recovery** — 异常恢复
4. **RateLimit** — 限流
5. **RequestParse** — 请求解析
6. **AuthLoginJWT** — JWT 登录认证
7. **PermissionCheck** — 权限校验

---

## 📂 目录结构

```
backend/
├── main.go                   # 程序入口，初始化各模块并启动 HTTP Server
├── version.go                # 版本号定义
├── Makefile                  # 构建与开发命令
├── Dockerfile                # Docker 镜像构建
├── build.sh                  # 构建发布包脚本
├── deploy.sh                 # 部署脚本
├── go.mod                    # Go 模块依赖
├── go.sum                    # 依赖校验
│
├── conf/                     # 配置文件（按环境区分）
│   ├── dev/app.yaml          # 开发环境配置
│   ├── test/app.yaml         # 测试环境配置
│   └── prod/app.yaml         # 生产环境配置
│
├── config/                   # 配置解析模块
│   ├── app.go                # AppConfig 结构体定义与解析
│   └── config.go             # 配置初始化入口
│
├── router/                   # 路由定义
│   └── router.go             # 所有 API 路由注册，路由表配置
│
├── filter/                   # 中间件
│   ├── auth_login.go         # JWT 登录认证中间件
│   ├── cross.go              # 跨域处理（已迁移至 CORS 库）
│   ├── permission.go         # 权限校验中间件
│   ├── rate_limit.go         # 限流中间件
│   └── request.go            # 请求解析中间件
│
├── global/                   # 全局变量与初始化
│   ├── global.go             # 全局 Gin Engine、常量定义
│   ├── flag_env.go           # 命令行参数与环境变量解析
│   ├── context.go            # 请求上下文相关
│   └── indentify.go          # 身份标识相关
│
├── logger/                   # 日志模块
│   └── logger.go             # Zap 日志初始化与封装
│
├── app/                      # 核心业务代码
│   ├── entity/               # 数据实体（22 张表的结构体定义）
│   │   ├── account.go        # 账号实体
│   │   ├── doc.go            # 文档实体
│   │   ├── space.go          # 空间实体
│   │   ├── role.go           # 角色实体
│   │   ├── privilege.go      # 权限实体
│   │   ├── department.go     # 部门实体
│   │   ├── content.go        # 文档内容实体
│   │   ├── content_version.go # 内容版本实体
│   │   ├── collection.go     # 收藏实体
│   │   ├── follow.go         # 关注实体
│   │   ├── attachment.go     # 附件实体
│   │   ├── notice.go         # 公告实体
│   │   ├── link.go           # 链接实体
│   │   ├── contact.go        # 联系人实体
│   │   ├── email.go          # 邮件配置实体
│   │   ├── plugin.go         # 插件实体
│   │   ├── login_auth.go     # 登录认证实体
│   │   ├── config.go         # 系统配置实体
│   │   ├── log.go            # 操作日志实体
│   │   ├── log_doc.go        # 文档操作日志实体
│   │   ├── space_permission.go # 空间权限实体
│   │   └── ...
│   ├── dao/                  # 数据访问层
│   │   ├── dao.go            # DAO 初始化，DB 连接
│   │   ├── base.go           # 基础 CRUD 方法
│   │   └── ...               # 各实体对应的 DAO 文件
│   ├── service/              # 业务逻辑层
│   │   ├── service.go        # Service 初始化
│   │   ├── auth.go           # 认证服务
│   │   ├── captcha.go        # 验证码服务
│   │   ├── permission.go     # 权限服务
│   │   └── ...               # 各业务模块对应的 Service 文件
│   ├── controller/           # 控制器
│   │   ├── controller.go     # Controller 基础定义
│   │   ├── base.go           # 通用响应方法
│   │   ├── system/           # 系统管理 API（/system 路由组）
│   │   ├── space/            # 空间与文档 API（/space 路由组）
│   │   ├── user/             # 用户交互 API（/user 路由组）
│   │   └── home/             # 首页 API（/home 路由组）
│   ├── cache/                # 缓存
│   │   ├── cache.go          # 缓存基础
│   │   └── permission.go     # 权限缓存
│   ├── event/                # 事件
│   │   ├── event.go          # 事件定义
│   │   └── account_action.go # 账号操作事件
│   └── worker/               # 后台任务
│       └── worker.go         # Worker 定义
│
├── gopkg/                    # 内部工具包
│   ├── errors/               # 统一错误处理
│   ├── log/                  # 日志库配置结构
│   ├── mail/                 # 邮件发送工具
│   └── upload/               # 文件上传工具（支持 local/qiniu/aliyun）
│
├── utils/                    # 通用工具函数
│
├── script/                   # 运维脚本
│   ├── start.sh              # 启动服务
│   ├── stop.sh               # 停止服务
│   ├── restart.sh            # 重启服务
│   ├── monitor.sh            # 监控服务
│   ├── pack.sh               # 打包脚本
│   ├── mysql_docker_install.sh # Docker 安装 MySQL
│   └── mysql_dump.sh         # MySQL 数据备份
│
└── docs/                     # 文档
    └── database/
        └── mm_wiki2.sql      # 完整数据库表结构
```

---

## 🚀 构建与运行

### 环境要求

- Go 1.20+
- MySQL 8.0+

### 开发模式

```bash
# 使用 go run 启动（开发环境）
make dev

# 指定环境
make dev ENV=test
make dev ENV=prod
```

### 编译运行

```bash
# 编译
make build

# 编译并运行
make run

# 交叉编译
make linux      # Linux amd64
make windows    # Windows amd64
make darwin     # macOS amd64
```

### 测试

```bash
make test       # 运行所有测试（含覆盖率）
make lint       # 代码检查
```

---

## ⚙️ 配置参考

配置文件：`conf/{env}/app.yaml`

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
    host: 127.0.0.1           # MySQL 主机
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
    local_dir: ./static/upload # 本地存储路径
    domain: http://localhost:8088 # 文件访问域名

# 日志配置
Logger:
  default:
    - writer: console
      level: debug
    - writer: file
      level: info
      formatter: json
      writer_config:
        filename: ./logs/access.log
        max_age: 7
        max_backups: 10
        compress: false
        max_size: 10
    - writer: file
      level: error
      formatter: json
      writer_config:
        filename: ./logs/error.log
        max_age: 7
        max_backups: 10
        compress: false
        max_size: 10

# 认证配置
Auth:
  jwt_secret: ""              # JWT 密钥（为空使用默认值）
  expire_hours: 3             # Token 过期时间（小时）

# 跨域配置
CORS:
  allow_origins: []           # 允许的来源列表
  allow_methods: []           # 允许的 HTTP 方法
  allow_headers: []           # 允许的请求头
```

### 命令行参数

| 参数 | 说明 | 默认值 |
|------|------|--------|
| `-conf` | 配置文件路径 | `./app.yaml` |
| `-env` | 运行环境（`dev`/`test`/`prod`） | 空（使用环境变量） |

也可通过环境变量 `KAPP_ENV_TYPE` 设置运行环境。

---

## 🔌 API 路由

所有 API 按功能模块分为 4 个路由组：

### `/system` — 系统管理

| 模块 | 路径前缀 | 说明 |
|------|----------|------|
| 认证 | `/system/auth/*` | 登录、验证码 |
| 个人中心 | `/system/profile/*` | 个人信息、修改密码、关注列表、活动 |
| 账号管理 | `/system/account/*` | 用户增删改查、状态管理 |
| 角色管理 | `/system/role/*` | 角色增删改查、角色用户、角色权限 |
| 权限管理 | `/system/privilege/*` | 权限增删改查 |
| 日志管理 | `/system/log/*` | 操作日志查询 |
| 公告管理 | `/system/notice/*` | 公告增删改查、发布 |
| 部门管理 | `/system/department/*` | 部门增删改查 |
| 邮件管理 | `/system/email/*` | 邮件服务器配置、发送测试 |
| 插件管理 | `/system/plugin/*` | 插件增删改查、启停 |
| 安装向导 | `/system/install/*` | 系统安装、初始化数据 |
| 链接管理 | `/system/link/*` | 快捷链接增删改查 |
| 联系人管理 | `/system/contact/*` | 联系人增删改查 |
| 系统配置 | `/system/config/*` | 系统配置查看与修改 |
| 登录认证 | `/system/login_auth/*` | 外部登录认证管理（SSO） |
| 空间管理 | `/system/space/*` | 空间管理与管理员分配 |

### `/space` — 空间与文档

| 模块 | 路径前缀 | 说明 |
|------|----------|------|
| 空间 | `/space/space/*` | 获取所有空间、空间文档列表 |
| 空间设置 | `/space/setting/*` | 空间基本设置、权限管理 |
| 文档 | `/space/doc/*` | 文档创建、编辑、删除、移动、排序、搜索、历史版本 |
| 附件 | `/space/attachment/*` | 附件列表、上传、删除 |

### `/user` — 用户交互

| 模块 | 路径前缀 | 说明 |
|------|----------|------|
| 互动 | `/user/interaction/*` | 收藏、关注操作 |
| 用户 | `/user/account/*` | 用户列表查询 |
| 部门 | `/user/department/*` | 部门列表查询 |

### `/home` — 首页

| 模块 | 路径前缀 | 说明 |
|------|----------|------|
| 首页 | `/home/*` | 我的空间、收藏空间、收藏文档 |

---

## 🗄️ 数据库

### 表结构

数据库使用 MySQL 8.0，数据库名 `mm_wiki2`，共 22 张表，表前缀 `mk_`。

完整建表 SQL：`docs/database/mm_wiki2.sql`

### 初始化

```bash
# Docker 方式（在项目根目录）
docker compose up -d

# 手动方式
mysql -u root -p mm_wiki2 < docs/database/mm_wiki2.sql
mysql -u root -p mm_wiki2 < ../docs/database/data.sql.txt
```

---

## 📄 开源协议

[MIT License](../front/LICENSE)
