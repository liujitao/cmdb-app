/*
CREATE DATABASE `cmdb` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
GRANT ALL PRIVILEGES ON cmdb.* to 'cmdb'@'127.0.0.1' IDENTIFIED BY 'cmdb';
FLUSH PRIVILEGES;
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` varchar(20) NOT NULL COMMENT '用户ID',
  `avatar` varchar(255) NOT NULL COMMENT '头像',
  `mobile` varchar(50) NOT NULL COMMENT '手机号码',
  `email` varchar(50) NOT NULL COMMENT '邮箱',
  `user_name` varchar(50) NOT NULL COMMENT '用户名称',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `gender` tinyint(4) NOT NULL COMMENT '性别[ 0.女  1.男  2.未知]',
  `department_id` varchar(20) NOT NULL COMMENT '部门ID',
  `status` tinyint(1) NOT NULL COMMENT '状态 [ 0.禁用 1.正常 ]',
  `admin_flag` tinyint(1) NOT NULL COMMENT '0. 普通用户  1. 管理员',
  `create_at` datetime NOT NULL COMMENT '创建时间',
  `create_user` varchar(20) NOT NULL COMMENT '创建人',
  `update_at` datetime NOT NULL COMMENT '最后一次修改时间',
  `update_user` varchar(20) NOT NULL COMMENT '最后一次修改人',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `UK_ulo5s2i7qoksp54tgwl_mobile` (`mobile`) USING BTREE,
  UNIQUE KEY `UK_6ixlo2i7qoksp54tgwl_username` (`user_name`) USING BTREE,
  UNIQUE KEY `UK_6i5ixxs2i7984erhgwl_email` (`email`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='[ 权限 ] 用户表';

INSERT INTO `sys_user` VALUES ('5748354935248', 'https://oss.mhtled.com/avatar/6bcf17f12e524820abb9d983b6384fdc.png', '13910514434', 'admin@abc.com', '管理员', '$2a$10$Hxx27PDN.Eu3nWhB07Dk3eHbLqT5/0zCKR/h3FtYpne6KA9Qhj0uO', 2, '10001', 1, 1, '2021-05-14 11:17:33', '5748354935248', '2021-12-17 10:12:24', '5748354935248');
INSERT INTO `sys_user` VALUES ('5748354935249', 'https://oss.mhtled.com/avatar/69356c4b58014aa5b157229a8953695c.jpg', '15071525231', 'test1@mhtled.com', '测试用户2', '$2a$10$egh4K7nA02uTqzS.wX1rDeiB4NBWgyFqE86WGb7b08d4Vr3IVzSZq', 2, '10005', 0, 0, '2021-05-14 11:17:33', '5748354935248', '2021-12-17 10:12:56', '5748354935248');
INSERT INTO `sys_user` VALUES ('5748354935250', 'https://oss.mhtled.com/avatar/e28d9e7463a54374905ec297d923f990.jpg', '15071525235', 'doudou@mhtled.com', '王陈陈', '$2a$10$Hxx27PDN.Eu3nWhB07Dk3eHbLqT5/0zCKR/h3FtYpne6KA9Qhj0uO', 0, '10005', 1, 0, '2021-05-14 11:17:33', '5748354935248', '2022-02-24 11:58:57', '5748354935248');
INSERT INTO `sys_user` VALUES ('5748354935251', 'https://oss.mhtled.com/avatar/de422f1053f042d3bbb5f36480095baf.jpg', '18711336565', '1234567890@mhtled.com', 'koushuiwa', '$2a$10$RBhJGpKGthZGb8OZovRd1.14z0XToeyHqOqrvy21DmIhbbEG8U.Ey', 0, '10011', 1, 0, '2021-10-09 09:38:09', '5748354935248', '2021-12-17 10:13:26', '5748354935248');
INSERT INTO `sys_user` VALUES ('5748354935252', 'https://oss.mhtled.com/avatar/9e29bef889db433f91f9d7edf26c7aa8.jpg', '18682035178', '42199386@mhtled.com', 'Yang Jian Ming', '$2a$10$HOyIJ.RQlQk4srEO62H/xuFLccMKpogBrIvlf7zXoCFh0eeULp586', 0, '10011', 1, 0, '2021-11-01 09:44:01', '5748354935248', '2021-12-17 10:13:31', '5748354935248');
INSERT INTO `sys_user` VALUES ('5748354935253', 'https://oss.mhtled.com/image/mhtled_logo.png', '15222222222', 'zhangyao@mhtled.com', '张瑶', '$2a$10$Hxx27PDN.Eu3nWhB07Dk3eHbLqT5/0zCKR/h3FtYpne6KA9Qhj0uO', 1, '10011', 1, 0, '2021-11-23 15:03:13', '5748354935248', '2022-03-25 14:26:09', '5748354935248');
INSERT INTO `sys_user` VALUES ('5749496392112', 'https://images.xaaef.com/5096297972b0469f80467771ba9c0fda.jpg', '15071525332', 'nihaoya123@qq.com', '你好呀', '$2a$10$5Vxu.lkpQT9awGBm/Uh3UOFs7DlMp9qL8b1zO.lfiJN2DbYn.eD5S', 1, '10010', 1, 0, '2022-03-25 11:21:39', '5748354935248', '2022-03-25 11:21:39', '5748354935248');

DROP TABLE IF EXISTS `sys_department`;
CREATE TABLE `sys_department` (
  `id` varchar(20) NOT NULL COMMENT '部门ID',
  `department_name` varchar(50) NOT NULL COMMENT '部门名称',
  `parent_id` varchar(20) NOT NULL DEFAULT 0 COMMENT '父主键',
  `leader_id` varchar(50) NOT NULL COMMENT '领导名称',
  `sort_id` tinyint(4) NOT NULL COMMENT '排序',
  `description` varchar(255) NOT NULL COMMENT '部门描述',
  `create_at` datetime NOT NULL COMMENT '创建时间',
  `create_user` varchar(20) NOT NULL COMMENT '创建人',
  `update_at` datetime NOT NULL COMMENT '最后一次修改时间',
  `update_user` varchar(20) NOT NULL COMMENT '最后一次修改人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10100 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='[ 权限 ] 部门表';

INSERT INTO `sys_department` VALUES ('10001', 'ABC技术有限公司', 0, '5748354935248', 1, 'ABC技术有限公司', '2021-07-05 15:20:44', '5748354935248', '2021-07-05 15:20:47', '5748354935248');
INSERT INTO `sys_department` VALUES ('10005', '研发部', '10001', '5748354935250', 3, '研发部', '2021-07-19 10:55:32', '5748354935248', '2022-03-24 16:25:46', '5748354935248');
INSERT INTO `sys_department` VALUES ('10010', '市场部', '10001', '5748354935252', 2, '市场部', '2021-07-29 10:22:18', '5748354935248', '2021-07-29 10:25:42', '5748354935248');
INSERT INTO `sys_department` VALUES ('10011', '软件组', '10005', '5748354935253', 1, '软件组', '2021-07-29 10:58:19', '5748354935248', '2021-07-29 14:36:51', '5748354935248');

DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` varchar(20) NOT NULL COMMENT '角色ID',
  `role_name` varchar(255) NOT NULL COMMENT '角色名称',
  `description` varchar(255) NOT NULL COMMENT '角色描述',
  `create_at` datetime NOT NULL COMMENT '创建时间',
  `create_user` varchar(20) NOT NULL COMMENT '创建人',
  `update_at` datetime NOT NULL COMMENT '最后一次修改时间',
  `update_user` varchar(20) NOT NULL COMMENT '最后一次修改人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='[ 权限 ] 角色表';

INSERT INTO `sys_role` VALUES ('10001', '系统管理员', '系统管理员', '2021-07-05 15:02:03', '5748354935248', '2021-12-17 09:50:21', '5748354935248');
INSERT INTO `sys_role` VALUES ('10038', '业务人员', '业务人员', '2021-12-17 09:51:27', '5748354935248', '2022-03-25 14:13:50', '5748354935248');
INSERT INTO `sys_role` VALUES ('10039', '测试人员', '测试人员', '2021-12-17 09:51:59', '5748354935248', '2022-03-25 14:13:58', '5748354935248');
INSERT INTO `sys_role` VALUES ('10040', '开发人员', '开发人员', '2021-12-17 09:58:41', '5748354935248', '2022-03-24 17:51:20', '5748354935248');

DROP TABLE IF EXISTS `sys_permission`;
CREATE TABLE `sys_permission` (
  `id` varchar(20) NOT NULL COMMENT '菜单ID',
  `title` varchar(100) NOT NULL COMMENT '菜单名称',
  `parent_id` varchar(20) NOT NULL DEFAULT 0 COMMENT '父主键',
  `name` varchar(255) NOT NULL COMMENT '名称',
  `path` varchar(255) NOT NULL COMMENT '路径',
  `component` varchar(255) NOT NULL COMMENT '组件',
  `redirect` varchar(255) NOT NULL COMMENT '重定向',
  `icon` varchar(255) NOT NULL DEFAULT '#' COMMENT '菜单图标',
  `sort_id` tinyint(4) NOT NULL COMMENT '排序',
  `permission_type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '菜单类型 [ 0:目录 1:菜单 2:功能/按钮/操作 ]',
  `create_at` datetime NOT NULL COMMENT '创建时间',
  `create_user` varchar(20) NOT NULL COMMENT '创建人',
  `update_at` datetime NOT NULL COMMENT '最后一次修改时间',
  `update_user` varchar(20) NOT NULL COMMENT '最后一次修改人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='[ 权限 ] 菜单权限表';

INSERT INTO `sys_permission` VALUES ('10002', '文档管理', '', 'Article', '/article', 'Layout', '/article/list', 'el-icon-document', 2, 0, '2021-07-09 12:30:00', '5748354935248', '2021-11-18 09:08:22', '5748354935248');
INSERT INTO `sys_permission` VALUES ('10052', '文章管理', '10002', 'ArticleList', '/article/list', 'views/article/index', '', '', 1, 1, '2021-07-09 12:30:00', '5748354935248', '2021-11-18 09:08:22', '5748354935248');
INSERT INTO `sys_permission` VALUES ('10053', '文章分类', '10002', 'ArticleCategory', '/article/category', 'views/article/category', '', '', 2, 1, '2021-07-09 12:30:00', '5748354935248', '2021-11-18 09:08:22', '5748354935248');

INSERT INTO `sys_permission` VALUES ('10001', '系统管理', '', 'Setting', '/setting', 'Layout', '/setting/user', 'el-icon-setting', 9, 0, '2021-07-09 12:30:00', '5748354935248', '2021-11-18 09:08:28', '5748354935248');
INSERT INTO `sys_permission` VALUES ('10004', '用户管理', '10001', 'SettingUser', '/setting/user', 'views/user/index', '', '', 1, 1, '2021-07-09 12:30:00', '5748354935248', '2022-03-24 16:34:56', '5748354935248');
INSERT INTO `sys_permission` VALUES ('10005', '角色管理', '10001', 'SettingRole', '/settting/role', 'views/role/index', '', '', 2, 1, '2021-07-09 12:30:00', '5748354935248', '2022-03-24 16:35:30', '5748354935248');
INSERT INTO `sys_permission` VALUES ('10007', '部门管理', '10001', 'SysDepartment', '/setting/department', 'views/department/index', '', '', 3, 1, '2021-07-09 12:30:00', '5748354935248', '2021-07-27 16:46:09', '5748354935248');
INSERT INTO `sys_permission` VALUES ('10003', '权限管理', '10001', 'SettingMenu', '/setting/permission', 'views/permission/index', '', '', 4, 1, '2021-07-09 12:30:00', '5748354935248', '2022-03-24 16:35:18', '5748354935248');

INSERT INTO `sys_permission` VALUES ('10082', '视图', '10003', '', '/setting/permission/get', '', '', '', 1, 2, '2022-03-25 14:05:50', '5748354935248', '2022-03-25 14:09:31', '5748354935248');
INSERT INTO `sys_permission` VALUES ('10066', '新增', '10003', '', '/setting/permission/post', '', '', '', 2, 2, '2022-03-25 13:50:51', '5748354935248', '2022-03-25 14:09:47', '5748354935248');
INSERT INTO `sys_permission` VALUES ('10067', '修改', '10003', '', '/setting/permission/patch', '', '', '', 3, 2, '2022-03-25 13:51:35', '5748354935248', '2022-03-25 13:52:39', '5748354935248');
INSERT INTO `sys_permission` VALUES ('10068', '删除', '10003', '', '/setting/permission/delete', '', '', '', 4, 2, '2022-03-25 13:51:56', '5748354935248', '2022-03-25 13:52:46', '5748354935248');

INSERT INTO `sys_permission` VALUES ('10083', '视图', '10004', '', '/setting/user/get', '', '', '', 1, 2, '2022-03-25 14:06:20', '5748354935248', '2022-03-25 14:09:57', '5748354935248');
INSERT INTO `sys_permission` VALUES ('10069', '新增', '10004', '', '/setting/user/post', '', '', '', 2, 2, '2022-03-25 13:53:38', '5748354935248', '2022-03-25 14:10:06', '5748354935248');
INSERT INTO `sys_permission` VALUES ('10071', '修改', '10004', '', '/setting/user/patch', '', '', '', 3, 2, '2022-03-25 13:54:40', '5748354935248', '2022-03-25 13:55:43', '5748354935248');
INSERT INTO `sys_permission` VALUES ('10072', '删除', '10004', '', '/setting/user/delete', '', '', '', 4, 2, '2022-03-25 13:54:58', '5748354935248', '2022-03-25 13:55:57', '5748354935248');
INSERT INTO `sys_permission` VALUES ('10073', '重置密码', '10004', '', '/setting/user/password', '', '', '', 6, 2, '2022-03-25 13:55:20', '5748354935248', '2022-03-25 13:55:20', '5748354935248');

INSERT INTO `sys_permission` VALUES ('10084', '视图', '10005', '', '/setting/role/get', '', '', '', 1, 2, '2022-03-25 14:06:44', '5748354935248', '2022-03-25 14:10:15', '5748354935248');
INSERT INTO `sys_permission` VALUES ('10074', '新增', '10005', '', '/setting/role/post', '', '', '', 2, 2, '2022-03-25 13:56:51', '5748354935248', '2022-03-25 14:10:33', '5748354935248');
INSERT INTO `sys_permission` VALUES ('10075', '修改', '10005', '', '/setting/role/patch', '', '', '', 3, 2, '2022-03-25 13:57:25', '5748354935248', '2022-03-25 13:57:25', '5748354935248');
INSERT INTO `sys_permission` VALUES ('10076', '删除', '10005', '', '/setting/role/delete', '', '', '', 4, 2, '2022-03-25 13:57:43', '5748354935248', '2022-03-25 13:57:43', '5748354935248');
INSERT INTO `sys_permission` VALUES ('10077', '修改权限', '10005', '', '/setting/role/permission', '', '', '', 5, 2, '2022-03-25 13:58:03', '5748354935248', '2022-03-25 13:58:03', '5748354935248');

INSERT INTO `sys_permission` VALUES ('10085', '视图', '10007', '', '/setting/department/get', '', '', '', 1, 2, '2022-03-25 14:07:13', '5748354935248', '2022-03-25 14:10:40', '5748354935248');
INSERT INTO `sys_permission` VALUES ('10078', '新增', '10007', '', '/setting/department/post', '', '', '', 2, 2, '2022-03-25 13:58:58', '5748354935248', '2022-03-25 14:13:01', '5748354935248');
INSERT INTO `sys_permission` VALUES ('10079', '修改', '10007', '', '/setting/department/patch', '', '', '', 3, 2, '2022-03-25 13:59:08', '5748354935248', '2022-03-25 13:59:08', '5748354935248');
INSERT INTO `sys_permission` VALUES ('10080', '删除', '10007', '', '/setting/department/delete', '', '', '', 4, 2, '2022-03-25 13:59:22', '5748354935248', '2022-03-25 13:59:22', '5748354935248');
INSERT INTO `sys_permission` VALUES ('10081', '修改权限', '10007','', '/setting/department/permission', '', '', '', 5, 2, '2022-03-25 14:00:05', '5748354935248', '2022-03-25 14:00:05', '5748354935248');

DROP TABLE IF EXISTS `sys_department_permission`;
CREATE TABLE `sys_department_permission` (
  `department_id` varchar(20) NOT NULL COMMENT '部门ID',
  `permission_id` varchar(20) NOT NULL COMMENT '菜单ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='[ 权限 ] 部门和菜单权限表';

DROP TABLE IF EXISTS `sys_role_permission`;
CREATE TABLE `sys_role_permission` (
  `role_id` varchar(20) NOT NULL COMMENT '角色ID',
  `permission_id` varchar(20) NOT NULL COMMENT '菜单ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='[ 权限 ] 角色和菜单权限表';


INSERT INTO `sys_role_permission` VALUES ('10001', '10001');
INSERT INTO `sys_role_permission` VALUES ('10001', '10066');
INSERT INTO `sys_role_permission` VALUES ('10001', '10002');
INSERT INTO `sys_role_permission` VALUES ('10001', '10052');
INSERT INTO `sys_role_permission` VALUES ('10001', '10053');
INSERT INTO `sys_role_permission` VALUES ('10001', '10003');
INSERT INTO `sys_role_permission` VALUES ('10001', '10067');
INSERT INTO `sys_role_permission` VALUES ('10001', '10068');
INSERT INTO `sys_role_permission` VALUES ('10001', '10004');
INSERT INTO `sys_role_permission` VALUES ('10001', '10069');
INSERT INTO `sys_role_permission` VALUES ('10001', '10005');
INSERT INTO `sys_role_permission` VALUES ('10001', '10071');
INSERT INTO `sys_role_permission` VALUES ('10001', '10007');
INSERT INTO `sys_role_permission` VALUES ('10001', '10072');
INSERT INTO `sys_role_permission` VALUES ('10001', '10073');
INSERT INTO `sys_role_permission` VALUES ('10001', '10074');
INSERT INTO `sys_role_permission` VALUES ('10001', '10075');
INSERT INTO `sys_role_permission` VALUES ('10001', '10076');
INSERT INTO `sys_role_permission` VALUES ('10001', '10077');
INSERT INTO `sys_role_permission` VALUES ('10001', '10078');
INSERT INTO `sys_role_permission` VALUES ('10001', '10079');
INSERT INTO `sys_role_permission` VALUES ('10001', '10080');
INSERT INTO `sys_role_permission` VALUES ('10001', '10081');
INSERT INTO `sys_role_permission` VALUES ('10001', '10082');
INSERT INTO `sys_role_permission` VALUES ('10001', '10083');
INSERT INTO `sys_role_permission` VALUES ('10001', '10084');
INSERT INTO `sys_role_permission` VALUES ('10001', '10085');
INSERT INTO `sys_role_permission` VALUES ('10038', '10001');
INSERT INTO `sys_role_permission` VALUES ('10038', '10082');
INSERT INTO `sys_role_permission` VALUES ('10038', '10083');
INSERT INTO `sys_role_permission` VALUES ('10038', '10003');
INSERT INTO `sys_role_permission` VALUES ('10038', '10084');
INSERT INTO `sys_role_permission` VALUES ('10038', '10004');
INSERT INTO `sys_role_permission` VALUES ('10038', '10085');
INSERT INTO `sys_role_permission` VALUES ('10038', '10005');
INSERT INTO `sys_role_permission` VALUES ('10038', '10007');

DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role` (
  `user_id` varchar(20) NOT NULL COMMENT '用户ID',
  `role_id` varchar(20) NOT NULL COMMENT '角色ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='[ 权限 ] 用户和角色表';

INSERT INTO `sys_user_role` VALUES ('5748354935248', '10001');
INSERT INTO `sys_user_role` VALUES ('5749496392112', '10001');
INSERT INTO `sys_user_role` VALUES ('5748354935250', '10038');
INSERT INTO `sys_user_role` VALUES ('5748354935250', '10039');
INSERT INTO `sys_user_role` VALUES ('5748354935253', '10039');

DROP TABLE IF EXISTS `sys_user_department`;
CREATE TABLE `sys_user_department` (
  `user_id` varchar(20) NOT NULL COMMENT '用户ID',
  `department_id` varchar(20) NOT NULL COMMENT '部门ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='[ 权限 ] 用户和部门表';

INSERT INTO `sys_user_department` VALUES ('5748354935248', '10001');
INSERT INTO `sys_user_department` VALUES ('5748354935250', '10005');
INSERT INTO `sys_user_department` VALUES ('5748354935253', '10010');
INSERT INTO `sys_user_department` VALUES ('5748354935252', '10011');
