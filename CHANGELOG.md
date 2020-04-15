# 更新日志：

## v0.1.9（2020-04）

### Fix Bug & Add Feature
#### 修复bug
1. 修复markdown序号问题

#### 新增功能
无

### 升级（Upgrade）
1. 下载新版本到部署该项目的根目录
2. 覆盖解压 (tar -zxvf mm-wiki-v0.1.9-linux-amd64.tar.gz)
3. 执行升级命令
```
./mm-wiki --conf conf/mm-wiki.conf --upgrade
```
3. 重新启动
```
./mm-wiki --conf conf/mm-wiki.conf
```

## v0.1.8（2020-04）

### Fix Bug & Add Feature
#### 修复bug
1. 修复版本号问题
2. 首页和登录页面自适应

#### 新增功能
1. 增加在线部署网站
2. 增加文档移动排序支持
3. 增加目录大纲显示
4. 首页最近文档过滤掉无权限的空间文档
5. 一些样式调整
6. 增加 docker 支持

### 升级（Upgrade）
1. 下载新版本到部署该项目的根目录
2. 覆盖解压 (tar -zxvf mm-wiki-v0.1.8-linux-amd64.tar.gz)
3. 执行升级命令
```
./mm-wiki --conf conf/mm-wiki.conf --upgrade
```
3. 重新启动
```
./mm-wiki --conf conf/mm-wiki.conf
```
### 感谢
特别感谢 [@eahomliu](https://github.com/eahomliu) [@cifaz](https://github.com/cifaz) [@cxgreat2014](https://github.com/cxgreat2014) 几位贡献 PR

## v0.1.7（2020-02）

### Fix Bug & Add Feature
#### 修复bug
1. 修复发送邮件路径不存在

#### 新增功能
无

### 升级（Upgrade）
1. 下载新版本到部署该项目的根目录
2. 覆盖解压 (tar -zxvf mm-wiki-v0.1.7-linux-amd64.tar.gz)
3. 执行升级命令
```
./mm-wiki --conf conf/mm-wiki.conf --upgrade
```
3. 重新启动
```
./mm-wiki --conf conf/mm-wiki.conf
```

## v0.1.6（2020-02）

### Fix Bug & Add Feature
#### 修复bug
1. 修复新安装版本号不存在问题

#### 新增功能
1. 完善 pack 打包脚本
2. 修复  Windows 下 build 脚本编译问题
3. 搜索支持全文搜索功能

### 升级（Upgrade）
1. 下载新版本到部署该项目的根目录
2. 覆盖解压 (tar -zxvf mm-wiki-v0.1.5-mac-amd64.tar.gz)
3. 执行升级命令
```
./mm-wiki --conf conf/mm-wiki.conf --upgrade
```
3. 重新启动
```
./mm-wiki --conf conf/mm-wiki.conf
```
4. 配置文件新增搜索相关配置（增加到自己的配置文件中）
```
# 搜索配置
[search]
interval_time=30
```

## v0.1.5（2019-12）

### Fix Bug & Add Feature
#### 修复bug
1. 修复空间修改报错
2. 修复用户管理修改用户bug

#### 新增功能
1. 超级管理员可以重置用户密码

### 升级（Upgrade）
1. 下载新版本到部署该项目的根目录
2. 覆盖解压 (tar -zxvf mm-wiki-v0.1.5-mac-amd64.tar.gz)
3. 执行升级命令
```
./mm-wiki --conf conf/mm-wiki.conf --upgrade
```
3. 重新启动
```
./mm-wiki --conf conf/mm-wiki.conf
```

## v0.1.4（2019-11）

### Fix Bug & Add Feature
#### 修复bug
1. 邮箱配置发送测试失败问题
2. 去掉手机号验证
3. 修复版本号不存在问题
4. 修复空间修改bug

#### 新增功能
1. 项目改成 go mod 部署
2. 代码优化
3. linux 下增加启动脚本 run.sh

### 升级（Upgrade）
1. 下载新版本到部署该项目的根目录
2. 覆盖解压 (tar -zxvf mm-wiki-v0.1.4-mac-amd64.tar.gz)
3. 执行升级命令
```
./mm-wiki --conf conf/mm-wiki.conf --upgrade
```
3. 重新启动
```
./mm-wiki --conf conf/mm-wiki.conf
```

## v0.1.3（2019-05）
### Fix Bug & Add Feature
#### 修复bug
1. 修复启动命令不支持绝对路径问题
2. 安装向导页面优化
3. 修复首页快捷链接没有按排序号排序问题
4. 代码 go fmt
5. 更新 copyright
6. 修改系统权限 #55
7. 修复私有文档未授权访问漏洞 #55
8. 修复文档导出漏洞 #66
9. 新建文档可回车提交 #43
10. 修复文档编辑窗口过短问题 #46
11. 安装的最小环境限制 #45
11. 空间名修改后没有更新 #53

#### 新增功能
1. 新增附件上传，附件列表查看，附件删除功能
2. 图片上传分目录存储，可查看图片列表并删除
3. 分页增加每一页数量控制功能
4. 导出文件同时导出附件和图片
5. 下载文件同时下载所有的图片和附件

### Upgrade
1. 下载新版本到部署该项目的根目录
2. 覆盖解压
3. 执行升级命令
```
./mm-wiki --conf conf/mm-wiki.conf --upgrade
```
3. 重新启动
```
./mm-wiki --conf conf/mm-wiki.conf
```

## v0.1.2 （2018-11）
### Fix Bug & Add Feature
1. 修复 #16 账号密码回车不能登录问题
2. 修复 #22 微信分享标题问题
3. 优化 js
4. 修复文档内容 a 标签跳转问题
5. 修复搜索框回车键不能搜索问题
6. 增加邮件测试发送功能
7. 修复邮件通知错误日志不能输出到数据库问题
8. 修复邮件通知多个发送人失败问题
9. 增加新版本自动升级命令 --upgrade
10. 增加文档本地保存，意外退出后可恢复本地文档

### Upgrade
1. 下载新版本到部署该项目的根目录
2. 覆盖解压
3. 执行升级命令
```
./mm-wiki --conf conf/mm-wiki.conf --upgrade
```
3. 重新启动
```
./mm-wiki --conf conf/mm-wiki.conf
```


## v0.1.1（2018-08-08）
### Fix Bug & Add Feature
1. 添加角色不能删除 #3
2. 权限删除问题
3. 空间不能删除问题 #12
4. js 优化

### Upgrade
1. 下载新版本到部署该项目的根目录
2. 覆盖解压
3. 重新启动

## v0.1 （2018-07-22）
### 预览版发布
