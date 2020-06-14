-- --------------------------------
-- MM-Wiki 表结构
-- author: phachon
-- --------------------------------

-- 手动安装时首先需要创建数据库
-- CREATE DATABASE IF NOT EXISTS mm_wiki DEFAULT CHARSET utf8;

-- --------------------------------
-- 用户表
-- --------------------------------
DROP TABLE IF EXISTS `mw_user`;
CREATE TABLE `mw_user` (
  `user_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '用户 id',
  `username` varchar(100) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `given_name` varchar(50) NOT NULL DEFAULT '' COMMENT '姓名',
  `mobile` char(13) NOT NULL DEFAULT '' COMMENT '手机号',
  `phone` char(13) NOT NULL DEFAULT '' COMMENT '电话',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
  `department` char(50) NOT NULL DEFAULT '' COMMENT '部门',
  `position` char(50) NOT NULL DEFAULT '' COMMENT '职位',
  `location` char(50) NOT NULL DEFAULT '' COMMENT '位置',
  `im` char(50) NOT NULL DEFAULT '' COMMENT '即时聊天工具',
  `last_ip` varchar(15) NOT NULL DEFAULT '' COMMENT '最后登录ip',
  `last_time` int(11) NOT NULL DEFAULT '0' COMMENT '最后登录时间',
  `role_id` tinyint(3) NOT NULL DEFAULT '0' COMMENT '角色 id',
  `is_forbidden` tinyint(3) NOT NULL DEFAULT '0' COMMENT '是否屏蔽，0 否 1 是',
  `is_delete` tinyint(3) NOT NULL DEFAULT '0' COMMENT '是否删除，0 否 1 是',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';

-- ---------------------------------------------------------------
-- 系统角色表
-- ---------------------------------------------------------------
DROP TABLE IF EXISTS `mw_role`;
CREATE TABLE `mw_role` (
  `role_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '角色 id',
  `name` char(10) NOT NULL DEFAULT '' COMMENT '角色名称',
  `type` tinyint(3) NOT NULL DEFAULT '0' COMMENT '角色类型 0 自定义角色，1 系统角色',
  `is_delete` tinyint(3) NOT NULL DEFAULT '0' COMMENT '是否删除，0 否 1 是',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='系统角色表';

-- -------------------------------------------------------
-- 系统权限表
-- -------------------------------------------------------
DROP TABLE IF EXISTS `mw_privilege`;
CREATE TABLE `mw_privilege` (
  `privilege_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '权限id',
  `name` char(30) NOT NULL DEFAULT '' COMMENT '权限名',
  `parent_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '上级',
  `type` enum('controller','menu') DEFAULT 'controller' COMMENT '权限类型：控制器、菜单',
  `controller` char(100) NOT NULL DEFAULT '' COMMENT '控制器',
  `action` char(100) NOT NULL DEFAULT '' COMMENT '动作',
  `icon` char(100) NOT NULL DEFAULT '' COMMENT '图标（用于展示)',
  `target` char(200) NOT NULL DEFAULT '' COMMENT '目标地址',
  `is_display` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否显示：0不显示 1显示',
  `sequence` int(10) NOT NULL DEFAULT '0' COMMENT '排序(越小越靠前)',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`privilege_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='系统权限表';

-- ------------------------------------------------------------------
-- 系统角色权限对应关系表
-- ------------------------------------------------------------------
DROP TABLE IF EXISTS `mw_role_privilege`;
CREATE TABLE `mw_role_privilege` (
  `role_privilege_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '角色权限关系 id',
  `role_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '角色id',
  `privilege_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '权限id',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`role_privilege_id`),
  KEY (`role_id`),
  KEY (`privilege_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='系统角色权限对应关系表';

-- --------------------------------
-- 空间表
-- --------------------------------
DROP TABLE IF EXISTS `mw_space`;
CREATE TABLE `mw_space` (
  `space_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '空间 id',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '名称',
  `description` varchar(100) NOT NULL DEFAULT '' COMMENT '描述',
  `tags` varchar(255) NOT NULL DEFAULT '' COMMENT '标签',
  `visit_level` enum('private','public') NOT NULL DEFAULT 'public' COMMENT '访问级别：private,public',
  `is_share` tinyint(3) NOT NULL DEFAULT '1' COMMENT '文档是否允许分享 0 否 1 是',
  `is_export` tinyint(3) NOT NULL DEFAULT '1' COMMENT '文档是否允许导出 0 否 1 是',
  `is_delete` tinyint(3) NOT NULL DEFAULT '0' COMMENT '是否删除 0 否 1 是',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`space_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='空间表';

-- --------------------------------
-- 空间成员表
-- --------------------------------
DROP TABLE IF EXISTS `mw_space_user`;
CREATE TABLE `mw_space_user` (
  `space_user_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '用户空间关系 id',
  `user_id` int(10) NOT NULL DEFAULT '0' COMMENT '用户 id',
  `space_id` int(10) NOT NULL DEFAULT '0' COMMENT '空间 id',
  `privilege` tinyint(3) NOT NULL DEFAULT '0' COMMENT '空间成员操作权限 0 浏览者 1 编辑者 2 管理员',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`space_user_id`),
  UNIQUE KEY (`user_id`, `space_id`),
  KEY (`user_id`),
  KEY (`space_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='空间成员表';

-- --------------------------------
-- 文档表
-- --------------------------------
DROP TABLE IF EXISTS `mw_document`;
CREATE TABLE `mw_document` (
  `document_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '文档 id',
  `parent_id` int(10) NOT NULL DEFAULT '0' COMMENT '文档父 id',
  `space_id` int(10) NOT NULL DEFAULT '0' COMMENT '空间id',
  `name` varchar(150) NOT NULL DEFAULT '' COMMENT '文档名称',
  `type` tinyint(3) NOT NULL DEFAULT '1' COMMENT '文档类型 1 page 2 dir',
  `path` char(30) NOT NULL DEFAULT '0' COMMENT '存储根文档到父文档的 document_id 值, 格式 0,1,2,...',
  `sequence` int(10) NOT NULL DEFAULT '0' COMMENT '排序号(越小越靠前)',
  `create_user_id` int(10) NOT NULL DEFAULT '0' COMMENT '创建用户 id',
  `edit_user_id` int(10) NOT NULL DEFAULT '0' COMMENT '最后修改用户 id',
  `is_delete` tinyint(3) NOT NULL DEFAULT '0' COMMENT '是否删除 0 否 1 是',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`document_id`),
  KEY (`parent_id`),
  KEY (`space_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文档表';

-- --------------------------------
-- 用户收藏表
-- --------------------------------
DROP TABLE IF EXISTS `mw_collection`;
CREATE TABLE `mw_collection` (
  `collection_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '用户收藏关系 id',
  `user_id` int(10) NOT NULL DEFAULT '0' COMMENT '用户id',
  `type` tinyint(3) NOT NULL DEFAULT '1' COMMENT '收藏类型 1 文档 2 空间',
  `resource_id` int(10) NOT NULL DEFAULT '0' COMMENT '收藏资源 id ',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`collection_id`),
  KEY (`user_id`),
  UNIQUE key (`user_id`, `resource_id`, `type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户收藏表';

-- --------------------------------
-- 用户关注表
-- --------------------------------
DROP TABLE IF EXISTS `mw_follow`;
CREATE TABLE `mw_follow` (
  `follow_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '关注 id',
  `user_id` int(10) NOT NULL DEFAULT '0' COMMENT '用户id',
  `type` tinyint(3) NOT NULL DEFAULT '1' COMMENT '关注类型 1 文档 2 用户',
  `object_id` int(10) NOT NULL DEFAULT '0' COMMENT '关注对象 id',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`follow_id`),
  KEY (`user_id`),
  KEY (`object_id`),
  UNIQUE key (`user_id`, `object_id`, `type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户关注表';

-- --------------------------------
-- 文档日志表
-- --------------------------------
DROP TABLE IF EXISTS `mw_log_document`;
CREATE TABLE `mw_log_document` (
  `log_document_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '文档日志 id',
  `document_id` int(10) NOT NULL DEFAULT '0' COMMENT '文档id',
  `space_id` int(10) NOT NULL DEFAULT '0' COMMENT '空间id',
  `user_id` int(10) NOT NULL DEFAULT '0' COMMENT '用户id',
  `action` tinyint(3) NOT NULL DEFAULT '1' COMMENT '动作 1 创建 2 修改 3 删除',
  `comment` varchar(255) NOT NULL DEFAULT '' COMMENT '备注信息',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`log_document_id`),
  KEY (`document_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文档日志表';

-- --------------------------------
-- 系统操作日志表
-- --------------------------------
DROP TABLE IF EXISTS `mw_log`;
CREATE TABLE `mw_log` (
  `log_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '系统操作日志 id',
  `level` tinyint(3) NOT NULL DEFAULT '6' COMMENT '日志级别',
  `path` char(100) NOT NULL DEFAULT '' COMMENT '请求路径',
  `get` text NOT NULL COMMENT 'get参数',
  `post` text NOT NULL COMMENT 'post参数',
  `message` varchar(255) NOT NULL DEFAULT '' COMMENT '信息',
  `ip` char(100) NOT NULL DEFAULT '' COMMENT 'ip地址',
  `user_agent` char(200) NOT NULL DEFAULT '' COMMENT '用户代理',
  `referer` char(100) NOT NULL DEFAULT '' COMMENT 'referer',
  `user_id` int(10) NOT NULL DEFAULT '0' COMMENT '用户id',
  `username` char(100) NOT NULL DEFAULT '' COMMENT '用户名',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`log_id`),
  KEY (`level`, `username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='系统操作日志表';

-- --------------------------------
-- 邮件服务器表
-- --------------------------------
DROP TABLE IF EXISTS `mw_email`;
CREATE TABLE `mw_email` (
  `email_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '邮箱 id',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '邮箱服务器名称',
  `sender_address` varchar(100) NOT NULL DEFAULT '' COMMENT '发件人邮件地址',
  `sender_name` varchar(100) NOT NULL DEFAULT '' COMMENT '发件人显示名',
  `sender_title_prefix` varchar(100) NOT NULL DEFAULT '' COMMENT '发送邮件标题前缀',
  `host` char(100) NOT NULL DEFAULT '' COMMENT '服务器主机名',
  `port` int(5) NOT NULL DEFAULT '25' COMMENT '服务器端口',
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(50) NOT NULL DEFAULT '' COMMENT '密码',
  `is_ssl` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否使用ssl， 0 默认不使用 1 使用',
  `is_used` tinyint(3) NOT NULL DEFAULT '0' COMMENT '是否被使用， 0 默认不使用 1 使用',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`email_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='邮件服务器表';

-- --------------------------------
-- 快捷链接表
-- --------------------------------
DROP TABLE IF EXISTS `mw_link`;
CREATE TABLE `mw_link` (
  `link_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '链接 id',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '链接名称',
  `url` varchar(100) NOT NULL DEFAULT '' COMMENT '链接地址',
  `sequence` int(10) NOT NULL DEFAULT '0' COMMENT '排序号(越小越靠前)',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`link_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='快捷链接表';

-- --------------------------------
-- 统一登录认证表
-- --------------------------------
DROP TABLE IF EXISTS `mw_login_auth`;
CREATE TABLE `mw_login_auth` (
  `login_auth_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '认证表主键ID',
  `name` varchar(30) NOT NULL COMMENT '登录认证名称',
  `username_prefix` varchar(30) NOT NULL COMMENT '用户名前缀',
  `url` varchar(200) NOT NULL COMMENT '认证接口 url',
  `ext_data` char(100) NOT NULL DEFAULT '' COMMENT '额外数据: token=aaa&key=bbb',
  `is_used` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否被使用， 0 默认不使用 1 使用',
  `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除 0 否 1 是',
  `create_time` int(11) NOT NULL COMMENT '创建时间',
  `update_time` int(11) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`login_auth_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='统一登录认证表';

-- --------------------------------
-- 全局配置表
-- --------------------------------
DROP TABLE IF EXISTS `mw_config`;
CREATE TABLE `mw_config` (
  `config_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '配置表主键Id',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '配置名称',
  `key` char(50) NOT NULL DEFAULT '' COMMENT '配置键',
  `value` text NOT NULL COMMENT '配置值',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`config_id`),
  unique KEY (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='全局配置表';

-- --------------------------------
-- 系统联系人表
-- --------------------------------
DROP TABLE IF EXISTS `mw_contact`;
CREATE TABLE `mw_contact` (
  `contact_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '联系人 id',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '联系人名称',
  `mobile` char(13) NOT NULL DEFAULT '' COMMENT '联系电话',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
  `position` varchar(100) NOT NULL DEFAULT '' COMMENT '联系人职位',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`contact_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='联系人表';

-- --------------------------------
-- 附件信息表
-- --------------------------------
DROP TABLE IF EXISTS `mw_attachment`;
CREATE TABLE `mw_attachment` (
  `attachment_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '附件 id',
  `user_id` int(10) NOT NULL DEFAULT '0' COMMENT '创建用户id',
  `document_id` int(10) NOT NULL DEFAULT '0' COMMENT '所属文档id',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '附件名称',
  `path` varchar(100) NOT NULL DEFAULT '' COMMENT '附件路径',
  `source` tinyint(1) NOT NULL DEFAULT '0' COMMENT '附件来源， 0 默认是附件 1 图片',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`attachment_id`),
  KEY (`document_id`, `source`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='附件信息表';