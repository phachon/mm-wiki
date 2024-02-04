-- ----------------------------------------------
-- mm_wiki2.sql mm-wiki 新版本表结构
-- Author: phachon
-- Date: 2024-02-01 11:00:00
-- ----------------------------------------------

CREATE DATABASE IF NOT EXISTS mm_wiki2 DEFAULT CHARSET utf8mb4;

-- ----------------------------------------------------------
-- mk_account 系统账号表
-- ----------------------------------------------------------
DROP TABLE IF EXISTS `mk_account`;
CREATE TABLE `mk_account` (
  `account_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '账号id',
  `name` varchar(32) NOT NULL DEFAULT '' COMMENT '账号名',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `given_name` varchar(50) NOT NULL DEFAULT '' COMMENT '姓名',
  `mobile` char(13) NOT NULL DEFAULT '' COMMENT '手机号',
  `phone` char(13) NOT NULL DEFAULT '' COMMENT '电话',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
  `department` char(50) NOT NULL DEFAULT '' COMMENT '部门',
  `position` char(50) NOT NULL DEFAULT '' COMMENT '职位',
  `location` char(50) NOT NULL DEFAULT '' COMMENT '办公位',
  `im` char(50) NOT NULL DEFAULT '' COMMENT '即时聊天工具',
  `last_ip` varchar(15) NOT NULL DEFAULT '' COMMENT '最后登录ip',
  `last_time` int(11) NOT NULL DEFAULT '0' COMMENT '最后登录时间',
  `role_id` tinyint(3) NOT NULL DEFAULT '0' COMMENT '角色 id',
  `is_forbidden` tinyint(3) NOT NULL DEFAULT '0' COMMENT '是否屏蔽: 0 否 1 是',
  `status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '状态: 0 正常 -1 删除',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`account_id`),
  UNIQUE KEY `name` (`name`)
) ENGINE = InnoDB AUTO_INCREMENT = 210510 DEFAULT CHARSET = utf8mb4 COMMENT = '系统账号表';

-- ----------------------------------------------------------
-- mk_role 系统角色表
-- ----------------------------------------------------------
DROP TABLE IF EXISTS `mk_role`;
CREATE TABLE `mk_role` (
  `role_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '角色id',
  `name` char(10) NOT NULL DEFAULT '' COMMENT '角色名称',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '角色备注',
  `role_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '角色类型：0 自定义角色 1 默认角色',
  `privilege_ids` varchar(500) NOT NULL DEFAULT '' COMMENT '角色下的权限ID，逗号隔开',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态 0 正常 -1 删除',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`role_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '系统角色表';

-- ----------------------------------------------------------
-- mk_account_role 账号角色关系表
-- ----------------------------------------------------------
DROP TABLE IF EXISTS `mk_account_role`;
CREATE TABLE `mk_account_role` (
  `account_role_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '账号角色id',
  `role_id` int(10) NOT NULL DEFAULT '0' COMMENT '角色id',
  `account_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '账号id',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`account_role_id`),
  KEY `role_id` (`role_id`),
  KEY `account_id` (`account_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '账号角色关系表';

-- ----------------------------------------------------------
-- mk_privilege 权限表
-- ----------------------------------------------------------
DROP TABLE IF EXISTS `mk_privilege`;
CREATE TABLE `mk_privilege` (
  `privilege_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '权限id',
  `identify` char(125) NOT NULL DEFAULT '' COMMENT '权限唯一标识',
  `name` char(125) NOT NULL DEFAULT '' COMMENT '权限名',
  `parent_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '上级id',
  `parent_ids` varchar(250) NOT NULL DEFAULT '' COMMENT '所有上级id,隔开',
  `privilege_type` tinyint(1) DEFAULT '1' COMMENT '权限类型: 1导航 2菜单 3操作',
  `page_router` char(200) NOT NULL DEFAULT '' COMMENT '权限路由',
  `api_marks` varchar(500) NOT NULL DEFAULT '' COMMENT '接口标识（多个标识逗号隔开）',
  `icon` char(100) NOT NULL DEFAULT '' COMMENT '图标（用于展示)',
  `is_display` tinyint(1) NOT NULL DEFAULT '0' COMMENT '菜单是否显示: 0不显示 1显示',
  `sequence` int(10) NOT NULL DEFAULT '1' COMMENT '排序(越小越靠前)',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`privilege_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '系统权限表';

-- ----------------------------------------------------------
-- mk_log 操作日志表
-- ----------------------------------------------------------
DROP TABLE IF EXISTS `mk_log`;
CREATE TABLE `mk_log` (
  `log_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '日志id',
  `uri` varchar(256) NOT NULL DEFAULT '' COMMENT '接口uri',
  `get` text NOT NULL COMMENT 'get参数',
  `post` text NOT NULL COMMENT 'post参数',
  `message` varchar(255) NOT NULL DEFAULT '' COMMENT '信息',
  `level` tinyint(1) NOT NULL DEFAULT '0' COMMENT '级别',
  `file` varchar(256) NOT NULL DEFAULT '' COMMENT '文件',
  `line` int(10) NOT NULL DEFAULT '0' COMMENT '行数',
  `ip` char(100) NOT NULL DEFAULT '' COMMENT 'ip地址',
  `account_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '帐号id',
  `account_name` char(100) NOT NULL DEFAULT '' COMMENT '帐号名',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`log_id`),
  KEY (`level`),
  KEY (`account_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '操作日志表';

-- ----------------------------------------------------------
-- mk_notice 系统公告表
-- ----------------------------------------------------------
DROP TABLE IF EXISTS `mk_notice`;
CREATE TABLE `mk_notice` (
  `notice_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '公告id',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '公告标题',
  `content` varchar(500) NOT NULL DEFAULT '' COMMENT '公告内容',
  `account_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '账号id',
  `account_name` varchar(125) NOT NULL DEFAULT '' COMMENT '发布账号名',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '公告状态 0 正常 -1 删除',
  `publish_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '发布状态 0 待发布 1 已发布',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
  `start_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '开始时间',
  `end_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '结束时间',
  PRIMARY KEY (`notice_id`),
  KEY (`status`),
  KEY (`account_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '系统公告表';


-- --------------------------------
-- mk_space 空间表
-- --------------------------------
DROP TABLE IF EXISTS `mk_space`;
CREATE TABLE `mk_space` (
  `space_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '空间 id',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '名称',
  `description` varchar(100) NOT NULL DEFAULT '' COMMENT '描述',
  `tags` varchar(255) NOT NULL DEFAULT '' COMMENT '标签',
  `visit_level` tinyint(3) NOT NULL DEFAULT '0' COMMENT '访问级别: 0 公开 1 私有',
  `is_share` tinyint(3) NOT NULL DEFAULT '1' COMMENT '文档是否允许分享 0 否 1 是',
  `is_export` tinyint(3) NOT NULL DEFAULT '1' COMMENT '文档是否允许导出 0 否 1 是',
  `status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '状态: 0 正常 -1 删除',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`space_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='空间表';

-- --------------------------------
-- mk_space_account 空间成员表
-- --------------------------------
DROP TABLE IF EXISTS `mk_space_account`;
CREATE TABLE `mk_space_account` (
  `space_account_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '空间账号关系 id',
  `account_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '账号 id',
  `space_id` int(10) NOT NULL DEFAULT '0' COMMENT '空间 id',
  `privilege` tinyint(3) NOT NULL DEFAULT '0' COMMENT '空间成员操作权限 0 浏览者 1 编辑者 2 管理员',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`space_account_id`),
  UNIQUE KEY (`account_id`, `space_id`),
  KEY (`account_id`),
  KEY (`space_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='空间成员表';

-- --------------------------------
-- mk_doc 文档表
-- --------------------------------
DROP TABLE IF EXISTS `mk_doc`;
CREATE TABLE `mk_doc` (
  `doc_id` varchar(125) NOT NULL DEFAULT '' COMMENT '文档唯一ID',
  `parent_id` varchar(125) NOT NULL DEFAULT '' COMMENT '文档父 id',
  `space_id` int(10) NOT NULL DEFAULT '0' COMMENT '空间id',
  `name` varchar(150) NOT NULL DEFAULT '' COMMENT '文档标题',
  `type` tinyint(3) NOT NULL DEFAULT '1' COMMENT '文档类型 1 page 2 dir',
  `path` varchar(1000) NOT NULL DEFAULT '0' COMMENT '存储根文档到父文档的 doc_id 值, 格式 0,1,2,...',
  `sequence` int(10) NOT NULL DEFAULT '0' COMMENT '排序号(越小越靠前)',
  `create_account_id` int(10) NOT NULL DEFAULT '0' COMMENT '创建账号 id',
  `edit_account_id` int(10) NOT NULL DEFAULT '0' COMMENT '最后修改账号 id',
  `status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '状态: 0 正常 -1 删除',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`doc_id`),
  KEY (`parent_id`),
  KEY (`space_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文档表';

-- --------------------------------
-- mk_collection 账号收藏表
-- --------------------------------
DROP TABLE IF EXISTS `mk_collection`;
CREATE TABLE `mk_collection` (
  `collection_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '账号收藏关系 id',
  `account_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '账号id',
  `collection_type` tinyint(3) NOT NULL DEFAULT '1' COMMENT '收藏类型 1 文档 2 空间',
  `resource_id` varchar(225) NOT NULL DEFAULT '' COMMENT '收藏资源 id ',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`collection_id`),
  KEY (`account_id`),
  UNIQUE key (`account_id`, `resource_id`, `collection_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='账号收藏表';

-- --------------------------------
-- mk_follow 账号关注表
-- --------------------------------
DROP TABLE IF EXISTS `mk_follow`;
CREATE TABLE `mk_follow` (
  `follow_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '关注 id',
  `account_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '账号id',
  `follow_type` tinyint(3) NOT NULL DEFAULT '1' COMMENT '关注类型 1 文档 2 用户',
  `object_id` varchar(225) NOT NULL DEFAULT '' COMMENT '关注对象 id',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`follow_id`),
  KEY (`account_id`),
  KEY (`object_id`),
  UNIQUE key (`account_id`, `object_id`, `follow_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='账号关注表';

-- --------------------------------
-- mk_log_doc 文档操作日志表
-- --------------------------------
DROP TABLE IF EXISTS `mk_log_doc`;
CREATE TABLE `mk_log_doc` (
  `log_doc_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '文档日志 id',
  `doc_id` varchar(125) NOT NULL DEFAULT '0' COMMENT '文档id',
  `space_id` int(10) NOT NULL DEFAULT '0' COMMENT '空间id',
  `account_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '账号id',
  `action` tinyint(3) NOT NULL DEFAULT '1' COMMENT '动作 1 创建 2 修改 3 删除',
  `comment` varchar(255) NOT NULL DEFAULT '' COMMENT '备注信息',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`log_doc_id`),
  KEY (`doc_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文档日志表';

-- --------------------------------
-- 邮件服务器表
-- --------------------------------
DROP TABLE IF EXISTS `mk_email`;
CREATE TABLE `mk_email` (
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='邮件服务器表';

-- --------------------------------
-- mk_link 快捷链接表
-- --------------------------------
DROP TABLE IF EXISTS `mk_link`;
CREATE TABLE `mk_link` (
  `link_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '链接 id',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '链接名称',
  `url` varchar(100) NOT NULL DEFAULT '' COMMENT '链接地址',
  `sequence` int(10) NOT NULL DEFAULT '0' COMMENT '排序号(越小越靠前)',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`link_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='快捷链接表';

-- --------------------------------
-- mk_login_auth 统一登录认证表
-- --------------------------------
DROP TABLE IF EXISTS `mk_login_auth`;
CREATE TABLE `mk_login_auth` (
  `login_auth_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '认证表主键ID',
  `name` varchar(30) NOT NULL COMMENT '登录认证名称',
  `account_prefix` varchar(30) NOT NULL COMMENT '账号登录前缀',
  `url` varchar(200) NOT NULL COMMENT '认证接口 url',
  `ext_data` varchar(500) NOT NULL DEFAULT '' COMMENT '额外数据: token=aaa&key=bbb',
  `is_used` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否被使用， 0 默认不使用 1 使用',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态 0 正常 -1 删除',
  `create_time` int(11) NOT NULL COMMENT '创建时间',
  `update_time` int(11) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`login_auth_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='统一登录认证表';

-- --------------------------------
-- mk_config 全局配置表
-- --------------------------------
DROP TABLE IF EXISTS `mk_config`;
CREATE TABLE `mk_config` (
  `config_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '配置表主键Id',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '配置名称',
  `key` char(50) NOT NULL DEFAULT '' COMMENT '配置键',
  `value` text NOT NULL COMMENT '配置值',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`config_id`),
  unique KEY (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='全局配置表';

-- --------------------------------
-- mk_contact 系统联系人表
-- --------------------------------
DROP TABLE IF EXISTS `mk_contact`;
CREATE TABLE `mk_contact` (
  `contact_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '联系人 id',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '联系人名称',
  `mobile` char(13) NOT NULL DEFAULT '' COMMENT '联系电话',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
  `position` varchar(100) NOT NULL DEFAULT '' COMMENT '联系人职位',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`contact_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='联系人表';

-- --------------------------------
-- mk_attachment 附件信息表
-- --------------------------------
DROP TABLE IF EXISTS `mk_attachment`;
CREATE TABLE `mk_attachment` (
  `attachment_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '附件 id',
  `account_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建用户id',
  `doc_id` varchar(125) NOT NULL DEFAULT '0' COMMENT '所属文档id',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '附件名称',
  `path` varchar(100) NOT NULL DEFAULT '' COMMENT '附件路径',
  `source` tinyint(1) NOT NULL DEFAULT '0' COMMENT '附件来源， 0 默认是附件 1 图片',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`attachment_id`),
  KEY (`doc_id`, `source`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='附件信息表';