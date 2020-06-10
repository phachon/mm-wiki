-- --------------------------------------
-- MM-Wiki 安装数据
-- warning：执行前需保证表里没有任何数据
-- author: phachon
-- --------------------------------------


-- --------------------------------------
-- 系统用户（root）, password：123456, 自动安装不需要插入，手动安装时需要插入该数据
-- --------------------------------------
-- INSERT INTO `mw_user` (`user_id`, `username`, `password`, `given_name`, `email`,  `mobile`, `role_id`, `is_delete`, `create_time`, `update_time`)
-- VALUES ('1', 'root', 'e10adc3949ba59abbe56e057f20f883e', 'root', 'root@123456.com', '1102222', '1', '0', unix_timestamp(now()), unix_timestamp(now()));


-- --------------------------------------
-- 系统角色 1 超级管理员，2 管理员，3 普通用户
-- --------------------------------------
INSERT INTO `mw_role` (`role_id`, `name`, `type`, `is_delete`, `create_time`, `update_time`) VALUES ('1', '超级管理员', '1', '0', unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO `mw_role` (`role_id`, `name`, `type`, `is_delete`, `create_time`, `update_time`) VALUES ('2', '管理员', '1', '0', unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO `mw_role` (`role_id`, `name`, `type`, `is_delete`, `create_time`, `update_time`) VALUES ('3', '普通用户', '1', '0', unix_timestamp(now()), unix_timestamp(now()));


-- --------------------------------------
-- 系统权限
-- --------------------------------------
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (1, '个人中心', 0, 'menu', '', '', 'glyphicon-leaf', '', 1, 1, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (2, '个人资料', 1, 'controller', 'profile', 'info', 'glyphicon-list', '', 1, 11, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (3, '修改资料', 1, 'controller', 'profile', 'edit', 'glyphicon-list', '', 0, 12, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (4, '修改资料保存', 1, 'controller', 'profile', 'modify', 'glyphicon-list', '', 0, 13, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (5, '关注用户列表', 1, 'controller', 'profile', 'followUser', 'glyphicon-list', '', 0, 14, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (6, '关注文档列表', 1, 'controller', 'profile', 'followDoc', 'glyphicon-list', '', 0, 15, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (7, '我的活动', 1, 'controller', 'profile', 'activity', 'glyphicon-list', '', 1, 16, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (8, '修改密码', 1, 'controller', 'profile', 'password', 'glyphicon-list', '', 1, 17, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (9, '修改密码保存', 1, 'controller', 'profile', 'savePass', 'glyphicon-list', '', 0, 18, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (10, '用户管理', 1, 'menu', '', '', 'glyphicon-user', '', 1, 2, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (11, '添加用户', 10, 'controller', 'user', 'add', 'glyphicon-list', '', 1, 21, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (12, '添加用户保存', 10, 'controller', 'user', 'save', 'glyphicon-list', '', 0, 22, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (13, '用户列表', 10, 'controller', 'user', 'list', 'glyphicon-list', '', 1, 23, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (14, '修改用户', 10, 'controller', 'user', 'edit', 'glyphicon-list', '', 0, 24, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (15, '修改用户保存', 10, 'controller', 'user', 'modify', 'glyphicon-list', '', 0, 25, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (16, '屏蔽用户', 10, 'controller', 'user', 'forbidden', 'glyphicon-list', '', 0, 26, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (17, '恢复用户', 10, 'controller', 'user', 'recover', 'glyphicon-list', '', 0, 27, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (18, '用户详情', 10, 'controller', 'user', 'info', 'glyphicon-list', '', 0, 28, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (19, '角色管理', 1, 'menu', '', '', 'glyphicon-gift', '', 1, 3, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (20, '添加角色', 19, 'controller', 'role', 'add', 'glyphicon-list', '', 1, 31, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (21, '添加角色保存', 19, 'controller', 'role', 'save', 'glyphicon-list', '', 0, 32, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (22, '角色列表', 19, 'controller', 'role', 'list', 'glyphicon-list', '', 1, 33, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (23, '修改角色', 19, 'controller', 'role', 'edit', 'glyphicon-list', '', 0, 34, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (24, '修改角色保存', 19, 'controller', 'role', 'modify', 'glyphicon-list', '', 0, 35, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (25, '角色用户列表', 19, 'controller', 'role', 'user', 'glyphicon-list', '', 0, 36, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (26, '角色权限', 19, 'controller', 'role', 'privilege', 'glyphicon-list', '', 0, 37, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (27, '角色权限保存', 19, 'controller', 'role', 'grantPrivilege', 'glyphicon-list', '', 0, 38, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (28, '删除角色', 19, 'controller', 'role', 'delete', 'glyphicon-list', '', 0, 29, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (29, '重置用户角色', 19, 'controller', 'role', 'resetUser', 'glyphicon-list', '', 0, 310, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (30, '权限管理', 1, 'menu', '', '', 'glyphicon-lock', '', 1, 4, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (31, '添加权限', 30, 'controller', 'privilege', 'add', 'glyphicon-list', '', 1, 41, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (32, '添加权限保存', 30, 'controller', 'privilege', 'save', 'glyphicon-list', '', 0, 42, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (33, '权限列表', 30, 'controller', 'privilege', 'list', 'glyphicon-list', '', 1, 43, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (34, '修改权限', 30, 'controller', 'privilege', 'edit', 'glyphicon-list', '', 0, 44, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (35, '修改权限保存', 30, 'controller', 'privilege', 'modify', 'glyphicon-list', '', 0, 45, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (36, '删除权限', 30, 'controller', 'privilege', 'delete', 'glyphicon-list', '', 0, 46, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (37, '空间管理', 1, 'menu', '', '', 'glyphicon-th-large', '', 1, 5, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (38, '添加空间', 37, 'controller', 'space', 'add', 'glyphicon-list', '', 1, 51, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (39, '添加空间保存', 37, 'controller', 'space', 'save', 'glyphicon-list', '', 0, 52, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (40, '空间列表', 37, 'controller', 'space', 'list', 'glyphicon-list', '', 1, 53, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (41, '修改空间', 37, 'controller', 'space', 'edit', 'glyphicon-list', '', 0, 54, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (42, '修改空间保存', 37, 'controller', 'space', 'modify', 'glyphicon-list', '', 0, 55, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (43, '空间成员列表', 37, 'controller', 'space', 'member', 'glyphicon-list', '', 0, 56, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (44, '添加空间成员', 37, 'controller', 'space_user', 'save', 'glyphicon-list', '', 0, 57, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (45, '移除空间成员', 37, 'controller', 'space_user', 'remove', 'glyphicon-list', '', 0, 58, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (46, '更新空间成员权限', 37, 'controller', 'space_user', 'modify', 'glyphicon-list', '', 0, 59, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (47, '删除空间', 37, 'controller', 'space', 'delete', 'glyphicon-list', '', 0, 510, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (48, '空间备份', 37, 'controller', 'space', 'download', 'glyphicon-list', '', 0, 512, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (49, '日志管理', 1, 'menu', '', '', 'glyphicon-list-alt', '', 1, 6, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (50, '系统日志', 49, 'controller', 'log', 'system', 'glyphicon-list', '', 1, 61, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (51, '系统日志详情', 49, 'controller', 'log', 'info', 'glyphicon-list', '', 0, 62, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (52, '文档日志', 49, 'controller', 'log', 'document', 'glyphicon-list', '', 1, 63, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (53, '配置管理', 1, 'menu', '', '', 'glyphicon-cog', '', 1, 7, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (54, '全局配置', 53, 'controller', 'config', 'global', 'glyphicon-list', '', 1, 71, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (55, '全局配置保存', 53, 'controller', 'config', 'modify', 'glyphicon-list', '', 0, 72, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (56, '邮箱配置', 53, 'controller', 'email', 'list', 'glyphicon-list', '', 1, 73, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (57, '添加邮件服务器', 53, 'controller', 'email', 'add', 'glyphicon-list', '', 0, 74, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (58, '添加邮件服务器保存', 53, 'controller', 'email', 'save', 'glyphicon-list', '', 0, 75, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (59, '修改邮件服务器', 53, 'controller', 'email', 'edit', 'glyphicon-list', '', 0, 76, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (60, '修改邮件服务器保存', 53, 'controller', 'email', 'modify', 'glyphicon-list', '', 0, 77, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (61, '启用邮件服务器', 53, 'controller', 'email', 'used', 'glyphicon-list', '', 0, 78, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (62, '删除邮件服务器', 53, 'controller', 'email', 'delete', 'glyphicon-list', '', 0, 79, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (63, '登录认证', 53, 'controller', 'auth', 'list', 'glyphicon-list', '', 1, 81, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (64, '添加登录认证', 53, 'controller', 'auth', 'add', 'glyphicon-list', '', 0, 82, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (65, '添加登录认证保存', 53, 'controller', 'auth', 'save', 'glyphicon-list', '', 0, 83, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (66, '修改登录认证', 53, 'controller', 'auth', 'edit', 'glyphicon-list', '', 0, 84, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (67, '修改登录认证保存', 53, 'controller', 'auth', 'modify', 'glyphicon-list', '', 0, 85, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (68, '删除登录认证', 53, 'controller', 'auth', 'delete', 'glyphicon-list', '', 0, 86, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (69, '启用登录认证', 53, 'controller', 'auth', 'used', 'glyphicon-list', '', 0, 87, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (70, '登录认证文档', 53, 'controller', 'auth', 'doc', 'glyphicon-list', '', 0, 88, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (71, '系统管理', 1, 'menu', '', '', 'glyphicon-link', '', 1, 8, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (72, '快捷链接', 71, 'controller', 'link', 'list', 'glyphicon-list', '', 1, 81, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (73, '添加链接', 71, 'controller', 'link', 'add', 'glyphicon-list', '', 0, 82, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (74, '添加链接保存', 71, 'controller', 'link', 'save', 'glyphicon-list', '', 0, 83, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (75, '修改链接', 71, 'controller', 'link', 'edit', 'glyphicon-list', '', 0, 84, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (76, '修改链接保存', 71, 'controller', 'link', 'modify', 'glyphicon-list', '', 0, 85, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (77, '删除链接', 71, 'controller', 'link', 'delete', 'glyphicon-list', '', 0, 86, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (78, '系统联系人', 71, 'controller', 'contact', 'list', 'glyphicon-list', '', 1, 91, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (79, '添加联系人', 71, 'controller', 'contact', 'add', 'glyphicon-list', '', 0, 92, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (80, '添加联系人保存', 71, 'controller', 'contact', 'save', 'glyphicon-list', '', 0, 93, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (81, '修改联系人', 71, 'controller', 'contact', 'edit', 'glyphicon-list', '', 0, 94, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (82, '修改联系人保存', 71, 'controller', 'contact', 'modify', 'glyphicon-list', '', 0, 95, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (83, '删除联系人', 71, 'controller', 'contact', 'delete', 'glyphicon-list', '', 0, 96, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (84, '统计管理', 1, 'menu', '', '', 'glyphicon-signal', '', 1, 9, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (85, '数据统计', 84, 'controller', 'static', 'default', 'glyphicon-list', '', 1, 91, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (86, 'ajax获取空间文档排行', 84, 'controller', 'static', 'spaceDocsRank', 'glyphicon-list', '', 0, 92, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (87, 'ajax获取收藏文档排行', 84, 'controller', 'static', 'collectDocRank', 'glyphicon-list', '', 0, 93, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (88, 'ajax获取文档数量增长趋势', 84, 'controller', 'static', 'docCountByTime', 'glyphicon-list', '', 0, 94, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (89, '系统监控', 84, 'controller', 'static', 'monitor', 'glyphicon-list', '', 1, 95, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (90, 'ajax获取服务器状态', 84, 'controller', 'static', 'serverStatus', 'glyphicon-list', '', 0, 96, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (91, 'ajax获取服务器时间', 84, 'controller', 'static', 'serverTime', 'glyphicon-list', '', 0, 97, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (92, '测试邮件服务器', 53, 'controller', 'email', 'test', 'glyphicon-list', '', 0, 80, unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO mw_privilege (privilege_id, name, parent_id, type, controller, action, icon, target, is_display, sequence, create_time, update_time) VALUES (93, '导入联系人', 71, 'controller', 'contact', 'import', 'glyphicon-list', '', 0, 97, unix_timestamp(now()), unix_timestamp(now()));



-- ---------------------------------------------
-- 系统角色权限关系数据
-- ---------------------------------------------
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (1, 3, 1, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (2, 3, 2, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (3, 3, 3, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (4, 3, 4, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (5, 3, 5, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (6, 3, 6, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (7, 3, 7, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (8, 3, 8, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (9, 3, 9, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (10, 2, 1, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (11, 2, 2, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (12, 2, 3, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (13, 2, 4, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (14, 2, 5, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (15, 2, 6, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (16, 2, 7, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (17, 2, 8, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (18, 2, 9, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (19, 2, 37, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (20, 2, 38, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (21, 2, 39, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (22, 2, 40, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (23, 2, 41, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (24, 2, 42, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (25, 2, 43, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (26, 2, 44, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (27, 2, 45, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (28, 2, 46, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (29, 2, 47, unix_timestamp(now()));
INSERT INTO mw_role_privilege (role_privilege_id, role_id, privilege_id, create_time) VALUES (30, 2, 48, unix_timestamp(now()));


-- -------------------------------------------
-- 系统配置
-- -------------------------------------------
INSERT INTO `mw_config` VALUES ('1', '主页标题', 'main_title', '这里可以填写公司名称，例如：欢迎来到 XXXX 科技公司 wiki 平台！', unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO `mw_config` VALUES ('2', '主页描述', 'main_description', '这是写一些描述：请使用域账号登录，使用中有任何问题请联系管理员 root@xxx.com！', unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO `mw_config` VALUES ('3', '是否开启自动关注', 'auto_follow_doc_open', '', unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO `mw_config` VALUES ('4', '是否开启邮件通知', 'send_email_open', '', unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO `mw_config` VALUES ('5', '是否开启统一登录', 'sso_open', '', unix_timestamp(now()), unix_timestamp(now()));
-- INSERT INTO `mw_config` (name, key, value, create_time, update_time) VALUES ('系统版本号', 'system_version', 'v0.0.0', unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO `mw_config` (name, `key`, value, create_time, update_time) VALUES ('开启全文搜索', 'fulltext_search_open', '1', unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO `mw_config` (name, `key`, value, create_time, update_time) VALUES ('索引更新间隔', 'doc_search_timer', '3600', unix_timestamp(now()), unix_timestamp(now()));
INSERT INTO `mw_config` (name, `key`, value, create_time, update_time) VALUES ('系统名称', 'system_name', 'Markdown Mini Wiki', unix_timestamp(now()), unix_timestamp(now()));
