# MM-Wiki 前端

MM-Wiki 前端基于 React 18 + TypeScript 开发，使用 Ant Design 5 作为 UI 组件库，采用 Container-Component 模式组织代码。

---

## 🛠️ 技术栈

| 技术 | 版本 | 说明 |
|------|------|------|
| [React](https://react.dev/) | 18.x | 前端框架 |
| [TypeScript](https://www.typescriptlang.org/) | 4.x | 类型安全 |
| [Ant Design](https://ant.design/) | 5.x | UI 组件库 |
| [Ant Design Pro Components](https://procomponents.ant.design/) | 2.x | 高级业务组件 |
| [React Router](https://reactrouter.com/) | 6.x | 客户端路由 |
| [Zustand](https://github.com/pmndrs/zustand) | 4.x | 状态管理 |
| [Axios](https://axios-http.com/) | 0.27.x | HTTP 请求 |
| [Monaco Editor](https://microsoft.github.io/monaco-editor/) | — | 代码编辑器 |
| [Cherry Markdown](https://github.com/Tencent/cherry-markdown) | 0.8.x | Markdown 编辑器 |
| [ECharts](https://echarts.apache.org/) | 4.6 | 图表库 |
| [Mermaid](https://mermaid.js.org/) | 9.4 | 流程图渲染 |
| [CRACO](https://craco.js.org/) | 7.x | CRA 配置覆盖工具 |

---

## 📂 目录结构

```
front/
├── package.json              # 项目依赖与脚本
├── package-lock.json         # 依赖锁定文件
├── tsconfig.json             # TypeScript 编译配置
├── craco.config.js           # Webpack 配置覆盖（别名、代码分割）
├── .eslintrc.js              # ESLint 代码检查规则
├── .prettierrc.js            # Prettier 代码格式化规则
├── Dockerfile                # Docker 多阶段构建（Node 构建 + Nginx 运行）
├── nginx.conf                # 生产环境 Nginx 配置
├── LICENSE                   # MIT 开源协议
│
├── public/                   # 公共静态文件（index.html 等）
│
└── src/
    ├── index.tsx             # 应用入口文件
    │
    ├── pages/                # 页面组件（按功能模块划分）
    │   ├── home/             # 首页模块
    │   │   └── Home/         # 首页容器组件
    │   ├── space/            # 空间模块
    │   │   ├── Space/        # 空间列表页
    │   │   └── Doc/          # 文档页（编辑器 + 预览）
    │   └── system/           # 后台管理模块
    │       ├── Main/         # 管理首页
    │       ├── Login/        # 登录页
    │       ├── Install/      # 安装向导页
    │       ├── Account/      # 账号管理
    │       ├── Role/         # 角色管理
    │       ├── Privilege/    # 权限管理
    │       ├── Department/   # 部门管理
    │       ├── Space/        # 空间管理
    │       ├── Log/          # 日志管理
    │       ├── Notice/       # 公告管理
    │       ├── Email/        # 邮件管理
    │       ├── Plugin/       # 插件管理
    │       ├── Link/         # 链接管理
    │       ├── Contact/      # 联系人管理
    │       ├── Config/       # 系统配置
    │       ├── LoginAuth/    # 登录认证管理
    │       ├── Profile/      # 个人中心
    │       ├── System/       # 系统设置
    │       └── Error/        # 错误页面
    │
    ├── components/           # 通用可复用组件
    │   ├── Layout/           # 布局组件（头部、侧边栏、内容区）
    │   ├── Action/           # 操作组件（按钮、弹窗等）
    │   └── DynamicIcon/      # 动态图标组件
    │
    ├── services/             # API 接口封装（按模块划分）
    │   ├── http.ts           # Axios 实例与拦截器配置
    │   ├── Base.ts           # 基础请求方法
    │   ├── Home.ts           # 首页接口
    │   ├── SystemLogin.ts    # 登录接口
    │   ├── SystemAccount.ts  # 账号管理接口
    │   ├── SystemRole.ts     # 角色管理接口
    │   ├── SystemPrivilege.ts # 权限管理接口
    │   ├── SystemSpace.ts    # 空间管理接口
    │   ├── SystemProfile.ts  # 个人中心接口
    │   ├── SpaceSpace.ts     # 空间操作接口
    │   ├── SpaceDoc.ts       # 文档操作接口
    │   ├── UserInteraction.ts # 用户交互接口（收藏、关注）
    │   └── ...               # 其他模块接口
    │
    ├── stores/               # Zustand 状态管理
    │   ├── index.ts          # Store 统一导出
    │   ├── account.ts        # 用户账号状态
    │   ├── doc.ts            # 文档状态
    │   ├── frame.ts          # 框架状态（侧边栏、菜单等）
    │   ├── local.ts          # 本地状态
    │   └── test.ts           # 测试用 Store
    │
    ├── router/               # 路由配置
    │   ├── index.tsx         # 路由入口与定义
    │   ├── paths.ts          # 路由路径常量
    │   ├── type.ts           # 路由类型定义
    │   ├── interceptor.tsx   # 路由拦截器（权限校验）
    │   └── modules/          # 路由模块
    │
    ├── config/               # 前端配置
    │   ├── url.ts            # API 地址与调试配置
    │   ├── setting.ts        # 系统设置（标题、Logo、版权等）
    │   └── layout.ts         # 布局配置（表单布局、侧边栏宽度）
    │
    ├── types/                # TypeScript 类型定义
    │   ├── baseType.ts       # 基础类型
    │   ├── accountType.ts    # 账号类型
    │   ├── docType.ts        # 文档类型
    │   ├── spaceType.ts      # 空间类型
    │   ├── roleType.ts       # 角色类型
    │   ├── privilegeType.ts  # 权限类型
    │   └── ...               # 其他业务类型
    │
    ├── utils/                # 工具函数
    │   ├── utils.ts          # 通用工具
    │   ├── Token.ts          # Token 管理
    │   └── LocalStorage.ts   # 本地存储封装
    │
    ├── assets/               # 静态资源
    │   ├── images/           # 图片资源
    │   └── styles/           # 全局样式
    │
    ├── theme/                # 主题配置
    │   └── index.ts          # Ant Design 主题定制
    │
    └── setupProxy.js         # 开发代理配置
```

---

## 🧩 Container-Component 模式

前端采用 **Container-Component** 模式组织页面代码：

- **Container（容器组件）**：负责数据获取、状态管理和业务逻辑，位于 `pages/` 目录
- **Component（展示组件）**：负责 UI 渲染，接收 props，位于 `components/` 目录

```
pages/system/Account/     ← 容器组件，处理账号管理的数据和逻辑
components/Layout/        ← 展示组件，纯 UI 渲染
```

**数据流：**

```
页面组件(Container) → 调用 Service 获取数据 → 更新 Zustand Store → UI 渲染
```

---

## 🚀 安装与启动

### 环境要求

- Node.js 18+
- npm 8+

### 安装依赖

```bash
npm install
```

### 开发模式

```bash
npm start
```

启动后访问 [http://localhost:3000](http://localhost:3000)，开发模式下 API 请求会自动代理到后端服务 `http://127.0.0.1:8088`。

### 生产构建

```bash
npm run build
```

构建产物输出到 `build/` 目录，可直接部署到 Nginx 或其他 Web 服务器。

### 运行测试

```bash
npm test
```

---

## ⚙️ 配置文件说明

### `src/config/url.ts` — API 地址配置

```typescript
// 开发环境和生产环境均使用 /api 前缀
const devUrlConfig = { proxyUrl: '/api' }
const proUrlConfig = { proxyUrl: '/api' }
```

同时包含开发调试参数配置，可跳过登录认证和权限校验。

### `src/config/setting.ts` — 系统设置

```typescript
export const SettingConfig: ISetting = {
  loginPageTitle: 'MM-Wiki 系统登录',     // 登录页标题
  loginPageSubTitle: '...',               // 登录页副标题
  loginUseDomainLogin: false,             // 是否使用域账号登录
  copyright: 'Copyright © 2023 ...',      // 版权信息
  frameHeaderTitle: ' MM-Wiki ',          // 后台框架头部标题
  footerShowText: '...',                  // 页脚文案
  fileDomain: 'http://localhost:8088'     // 文件存储域名
}
```

### `src/config/layout.ts` — 布局配置

定义 Ant Design 表单布局参数和侧边栏宽度（默认 240px）。

---

## 🔀 开发代理

`src/setupProxy.js` 配置开发环境代理，将 `/api` 和 `/static/upload` 请求转发到后端：

```javascript
module.exports = function (app) {
  // API 请求代理到后端
  app.use('/api', createProxyMiddleware({
    target: 'http://127.0.0.1:8088',
    changeOrigin: true
  }))
  // 上传文件访问代理
  app.use('/static/upload', createProxyMiddleware({
    target: 'http://127.0.0.1:8088',
    changeOrigin: true
  }))
}
```

---

## 📝 代码规范

### ESLint

使用 `@typescript-eslint/parser` 解析器，配合 `eslint-plugin-react`、`eslint-plugin-prettier` 等插件，主要规则：

- 强制驼峰命名
- 禁止多余分号
- 使用单引号
- 不使用尾逗号
- 集成 Prettier 自动格式化

### Prettier

```javascript
// .prettierrc.js
{
  printWidth: 100,        // 每行最大字符数
  tabWidth: 2,            // 缩进空格数
  semi: false,            // 不使用分号
  singleQuote: true,      // 使用单引号
  trailingComma: 'none',  // 不使用尾逗号
  arrowParens: 'always'   // 箭头函数参数总是带括号
}
```

### Webpack 配置（CRACO）

通过 `craco.config.js` 覆盖 CRA 默认配置：

- **路径别名**：`@` 映射到 `src/` 目录
- **代码分割**：将 `node_modules` 和 `antd` 单独打包为异步 chunk

---

## 🐳 Docker 部署

```bash
# 构建镜像（多阶段构建：Node.js 编译 + Nginx 运行）
docker build -t mm-wiki-frontend .

# 运行容器
docker run -d -p 80:80 mm-wiki-frontend
```

生产环境使用 Nginx 作为 Web 服务器，`nginx.conf` 配置：
- `/api` 路径代理到后端服务 `http://backend:8088`
- `/static/upload` 路径代理到后端静态文件服务
- 其他路径使用 `try_files` 支持前端路由

---

## 📄 开源协议

[MIT License](LICENSE)
