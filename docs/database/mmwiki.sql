-- --------------------------------
-- database: mm-wiki
-- author: phachon
-- --------------------------------

-- --------------------------------
-- 用户表
-- --------------------------------
DROP TABLE IF EXISTS `mw_user`;
CREATE TABLE `mw_user` (
  `user_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `username` varchar(100) NOT NULL DEFAULT '' COMMENT '用户名',
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
  `role` tinyint(3) NOT NULL DEFAULT '0' COMMENT '1 普通用户 2 管理员;3超级管理员',
  `is_delete` tinyint(3) NOT NULL DEFAULT '0' COMMENT '是否删除，0 否 1 是',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';

-- --------------------------------
-- 空间表
-- --------------------------------
DROP TABLE IF EXISTS `mw_space`;
CREATE TABLE `mw_space` (
  `space_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '空间表主键id',
  `category_id` int(10) NOT NULL DEFAULT '0' COMMENT '空间分类id',
  `description` varchar(100) NOT NULL DEFAULT '' COMMENT '描述',
  `is_delete` tinyint(3) NOT NULL DEFAULT '0' COMMENT '是否删除 0 否 1 是',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`space_id`),
  KEY (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='空间表';

-- --------------------------------
-- 分类表
-- --------------------------------
DROP TABLE IF EXISTS `mw_category`;
CREATE TABLE `mw_category` (
  `category_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '分类表主键id',
  `name` varchar(32) NOT NULL DEFAULT '' COMMENT '分类名',
  `is_delete` tinyint(3) NOT NULL DEFAULT '0' COMMENT '是否删除 0 否 1 是',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='分类表';

-- --------------------------------
-- 文档表
-- --------------------------------
DROP TABLE IF EXISTS `mw_page`;
CREATE TABLE `mw_page` (
  `page_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '文档表主键id',
  `parent_id` int(10) NOT NULL DEFAULT '0' COMMENT '文档父id',
  `space_id` int(10) NOT NULL DEFAULT '0' COMMENT '空间id',
  `type` tinyint(3) NOT NULL DEFAULT '' COMMENT '文档类型 1 目录 2 页面',
  `title` varchar(150) NOT NULL DEFAULT '' COMMENT '文档标题',
  `path` varchar(100) NOT NULL DEFAULT '' COMMENT 'markdown 文件路径',
  `create_user_id` int(10) NOT NULL DEFAULT '0' COMMENT '创建用户id',
  `edit_user_id` int(10) NOT NULL DEFAULT '0' COMMENT '最后修改用户id',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`page_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文档表';

-- --------------------------------
-- 用户空间关系表
-- --------------------------------
DROP TABLE IF EXISTS `mw_user_space`;
CREATE TABLE `mw_user_space` (
  `user_space_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '用户项目关系表 id',
  `user_id` int(10) NOT NULL DEFAULT '0' COMMENT '用户 id',
  `space_id` int(10) NOT NULL DEFAULT '0' COMMENT '项目 id',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`user_space_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户空间关系表';

-- --------------------------------
-- 用户收藏表
-- --------------------------------
DROP TABLE IF EXISTS `mw_collection`;
CREATE TABLE `mw_collection` (
  `collection_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '用户收藏关系表 id',
  `user_id` int(10) NOT NULL DEFAULT '0' COMMENT '用户id',
  `username` varchar(100) NOT NULL DEFAULT '' COMMENT '用户名',
  `type` tinyint(3) NOT NULL DEFAULT '1' COMMENT '收藏类型 1 文档 2 空间',
  `resource_id` int(10) NOT NULL DEFAULT '0' COMMENT '收藏资源 id ',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`collection_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户收藏表';

-- --------------------------------
-- 用户关注表
-- --------------------------------
DROP TABLE IF EXISTS `mw_follow`;
CREATE TABLE `mw_follow` (
  `follow_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '关注表id',
  `user_id` int(10) NOT NULL DEFAULT '0' COMMENT '用户id',
  `follow_user_id` int(10) NOT NULL DEFAULT '0' COMMENT '被关注用户 id',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`follow_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户关注表';

-- --------------------------------
-- 操作日志表
-- --------------------------------
DROP TABLE IF EXISTS `mw_log`;
CREATE TABLE `mw_log` (
  `log_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '日志id',
  `level` tinyint(3) NOT NULL DEFAULT '6' COMMENT '日志级别',
  `controller` char(100) NOT NULL DEFAULT '' COMMENT '控制器',
  `action` char(100) NOT NULL DEFAULT '' COMMENT '动作',
  `get` text NOT NULL COMMENT 'get参数',
  `post` text NOT NULL COMMENT 'post参数',
  `message` varchar(255) NOT NULL DEFAULT '' COMMENT '信息',
  `ip` char(100) NOT NULL DEFAULT '' COMMENT 'ip地址',
  `user_agent` char(200) NOT NULL DEFAULT '' COMMENT '用户代理',
  `referer` char(100) NOT NULL DEFAULT '' COMMENT 'referer',
  `user_id` int(10) NOT NULL DEFAULT '0' COMMENT '帐号id',
  `username` char(100) NOT NULL DEFAULT '' COMMENT '帐号名',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  PRIMARY KEY (`log_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='操作日志表';


-- --------------------------------
-- 联系人表
-- --------------------------------
DROP TABLE IF EXISTS `mw_contact`;
CREATE TABLE `mw_contact` (
  `contact_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '联系人表主键ID',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '联系人名称',
  `telephone` char(13) NOT NULL DEFAULT '' COMMENT '联系人座机电话',
  `mobile` char(13) NOT NULL DEFAULT '' COMMENT '联系人手机',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
  `position` varchar(100) NOT NULL DEFAULT '' COMMENT '联系人职位',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`contact_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='联系人表';