/*
CREATE DATABASE `cmdb` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
GRANT ALL PRIVILEGES ON cmdb.* to 'cmdb'@'127.0.0.1' IDENTIFIED BY 'cmdb';
FLUSH PRIVILEGES;
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id`  varchar(20) NOT NULL NOT NULL COMMENT '用户ID',
  `avatar` varchar(100) NOT NULL COMMENT '头像',
  `mobile` varchar(20) NOT NULL COMMENT '手机号码',
  `email` varchar(20) NOT NULL COMMENT '邮箱',
  `user_name` varchar(20) NOT NULL COMMENT '用户名称',
  `password` varchar(100) NOT NULL COMMENT '密码',
  `gender` tinyint(4) NOT NULL COMMENT '性别[ 0.女  1.男  2.未知]',
  `status` tinyint(1) NOT NULL COMMENT '状态 [ 0.禁用 1.正常 ]',
  `create_at` datetime NOT NULL COMMENT '创建时间',
  `create_user` varchar(20) NOT NULL COMMENT '创建人',
  `update_at` datetime COMMENT '最后一次修改时间',
  `update_user` varchar(20) COMMENT '最后一次修改人',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `UK_ulo5s2i7qoksp54tgwl_mobile` (`mobile`) USING BTREE,
  UNIQUE KEY `UK_6ixlo2i7qoksp54tgwl_username` (`user_name`) USING BTREE,
  UNIQUE KEY `UK_6i5ixxs2i7984erhgwl_email` (`email`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='[ 权限 ] 用户表';

INSERT INTO `sys_user` VALUES ('000000000', 'https://oss.mhtled.com/image/mhtled_logo.png', '13910514434', 'admin@abc.com', '管理员', '$2a$10$Hxx27PDN.Eu3nWhB07Dk3eHbLqT5/0zCKR/h3FtYpne6KA9Qhj0uO', 2, 1, '2021-05-14 11:17:33', '管理员', '2021-12-17 10:12:24', '管理员');
INSERT INTO `sys_user` VALUES ('000000001', 'https://oss.mhtled.com/image/mhtled_logo.png', '13900000001', 'user1@abc.com', '孙悟空', '$2a$10$Hxx27PDN.Eu3nWhB07Dk3eHbLqT5/0zCKR/h3FtYpne6KA9Qhj0uO', 1, 1, '2021-05-14 11:17:33', '管理员', '2021-12-17 10:12:56', '管理员');
INSERT INTO `sys_user` VALUES ('000000002', 'https://oss.mhtled.com/image/mhtled_logo.png', '13900000002', 'user2@abc.com', '天津饭', '$2a$10$Hxx27PDN.Eu3nWhB07Dk3eHbLqT5/0zCKR/h3FtYpne6KA9Qhj0uO', 1, 1, '2021-05-14 11:17:33', '管理员', '2022-02-24 11:58:57', '管理员');
INSERT INTO `sys_user` VALUES ('000000003', 'https://oss.mhtled.com/image/mhtled_logo.png', '13900000003', 'user3@abc.com', '贝吉塔', '$2a$10$Hxx27PDN.Eu3nWhB07Dk3eHbLqT5/0zCKR/h3FtYpne6KA9Qhj0uO', 1, 1, '2021-10-09 09:38:09', '管理员', '2021-12-17 10:13:26', '管理员');
INSERT INTO `sys_user` VALUES ('000000004', 'https://oss.mhtled.com/image/mhtled_logo.png', '13900000004', 'user4@abc.com', '龟仙人', '$2a$10$Hxx27PDN.Eu3nWhB07Dk3eHbLqT5/0zCKR/h3FtYpne6KA9Qhj0uO', 1, 1, '2021-11-01 09:44:01', '管理员', '2021-12-17 10:13:31', '管理员');
INSERT INTO `sys_user` VALUES ('000000005', 'https://oss.mhtled.com/image/mhtled_logo.png', '13900000005', 'user5@abc.com', '弗利萨', '$2a$10$Hxx27PDN.Eu3nWhB07Dk3eHbLqT5/0zCKR/h3FtYpne6KA9Qhj0uO', 1, 1, '2021-11-23 15:03:13', '管理员', '2022-03-25 14:26:09', '管理员');
INSERT INTO `sys_user` VALUES ('000000006', 'https://oss.mhtled.com/image/mhtled_logo.png', '13900000006', 'user6@abc.com', '布尔玛', '$2a$10$Hxx27PDN.Eu3nWhB07Dk3eHbLqT5/0zCKR/h3FtYpne6KA9Qhj0uO', 0, 0, '2022-03-25 11:21:39', '管理员', '2022-03-25 11:21:39', '管理员');

DROP TABLE IF EXISTS `sys_department`;
CREATE TABLE `sys_department` (
  `id`varchar(20) NOT NULL COMMENT '部门ID',
  `department_name` varchar(20) NOT NULL COMMENT '部门名称',
  `parent_id` varchar(20) NOT NULL DEFAULT 0 COMMENT '父主键',
  `sort_id` tinyint(4) NOT NULL COMMENT '排序',
  `description` varchar(100) NOT NULL COMMENT '部门描述',
  `create_at` datetime NOT NULL COMMENT '创建时间',
  `create_user` varchar(20) NOT NULL COMMENT '创建人',
  `update_at` datetime COMMENT '最后一次修改时间',
  `update_user` varchar(20) COMMENT '最后一次修改人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='[ 权限 ] 部门表';

INSERT INTO `sys_department` VALUES ('100000', '集团公司', '', 1, '集团公司', '2021-07-05 15:20:44', '管理员', '2021-07-05 15:20:47', '管理员');
INSERT INTO `sys_department` VALUES ('100001', '研发部', '100000', 1, '研发部', '2021-07-19 10:55:32', '管理员', '2022-03-24 16:25:46', '管理员');
INSERT INTO `sys_department` VALUES ('100002', '工程部', '100000', 2, '工程部', '2021-07-19 10:55:32', '管理员', '2022-03-24 16:25:46', '管理员');
INSERT INTO `sys_department` VALUES ('100003', '市场部', '100000', 3, '市场部', '2021-07-29 10:22:18', '管理员', '2021-07-29 10:25:42', '管理员');
INSERT INTO `sys_department` VALUES ('100004', '软件组', '100001', 1, '研发部-软件组', '2021-07-29 10:58:19', '管理员', '2021-07-29 14:36:51', '管理员');
INSERT INTO `sys_department` VALUES ('100005', '硬件组', '100001', 2, '研发部-硬件组', '2021-07-29 10:58:19', '管理员', '2021-07-29 14:36:51', '管理员');
INSERT INTO `sys_department` VALUES ('100006', '测试组', '100001', 3, '研发部-测试组', '2021-07-29 10:58:19', '管理员', '2021-07-29 14:36:51', '管理员');
INSERT INTO `sys_department` VALUES ('100007', '硬件1组', '100005', 1, '研发部-硬件1组', '2021-07-29 10:58:19', '管理员', '2021-07-29 14:36:51', '管理员');
INSERT INTO `sys_department` VALUES ('100008', '硬件2组', '100005', 2, '研发部-硬件2组', '2021-07-29 10:58:19', '管理员', '2021-07-29 14:36:51', '管理员');
INSERT INTO `sys_department` VALUES ('100009', '东北区', '100002', 1, '工程部-东北区', '2021-07-19 10:55:32', '管理员', '2022-03-24 16:25:46', '管理员');
INSERT INTO `sys_department` VALUES ('100010', '西北区', '100002', 2, '工程部-西北区', '2021-07-19 10:55:32', '管理员', '2022-03-24 16:25:46', '管理员');
INSERT INTO `sys_department` VALUES ('100011', '华南区', '100002', 3, '工程部-华南区', '2021-07-19 10:55:32', '管理员', '2022-03-24 16:25:46', '管理员');
INSERT INTO `sys_department` VALUES ('100012', '华南区1组', '100011', 1, '工程部-华南区-1组', '2021-07-19 10:55:32', '管理员', '2022-03-24 16:25:46', '管理员');
INSERT INTO `sys_department` VALUES ('100013', '华南区2组', '100011', 2, '工程部-华南区-2组', '2021-07-19 10:55:32', '管理员', '2022-03-24 16:25:46', '管理员');
INSERT INTO `sys_department` VALUES ('100014', '华南区3组', '100011', 3, '工程部-华南区-3组', '2021-07-19 10:55:32', '管理员', '2022-03-24 16:25:46', '管理员');

DROP TABLE IF EXISTS `sys_permission`;
CREATE TABLE `sys_permission` (
  `id` varchar(20) NOT NULL COMMENT '菜单ID',
  `parent_id` varchar(20) NOT NULL DEFAULT 0 COMMENT '父主键',
  `title` varchar(100) NOT NULL COMMENT '菜单名称',
  `name` varchar(100) NOT NULL COMMENT '名称',
  `path` varchar(100) NOT NULL COMMENT '路径',
  `component` varchar(100) NOT NULL COMMENT '组件',
  `redirect` varchar(100) NOT NULL COMMENT '重定向',
  `icon` varchar(100) NOT NULL DEFAULT '#' COMMENT '菜单图标',
  `permission_type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '菜单类型 [ 0:目录 1:菜单 2:功能/按钮/操作 ]',
  `sort_id` tinyint(4) NOT NULL DEFAULT 0 COMMENT '排序',
  `create_at` datetime NOT NULL COMMENT '创建时间',
  `create_user` varchar(20) NOT NULL COMMENT '创建人',
  `update_at` datetime COMMENT '最后一次修改时间',
  `update_user` varchar(20) COMMENT '最后一次修改人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='[ 权限 ] 菜单权限表';

INSERT INTO `sys_permission` VALUES ('120001', '', '资产管理', 'Asset', '/asset', 'Layout', '/asset/server', 'el-icon-s-grid', 0, 1, '2021-07-09 12:30:00', '管理员', '2021-11-18 09:08:22', '管理员');
INSERT INTO `sys_permission` VALUES ('120002', '', '系统管理', 'Setting', '/setting', 'Layout', '/setting/user', 'el-icon-setting', 0, 9, '2021-07-09 12:30:00', '管理员', '2021-11-18 09:08:28', '管理员');
INSERT INTO `sys_permission` VALUES ('120003', '120001', '服务器', 'AssetServer', '/asset/server', 'views/server/index', '', '', 1, 11, '2021-07-09 12:30:00', '管理员', '2021-11-18 09:08:22', '管理员');
INSERT INTO `sys_permission` VALUES ('120004', '120001', '网络设备', 'AssetNetwork', '/asset/network', 'views/network/index', '', '', 1, 12, '2021-07-09 12:30:00', '管理员', '2021-11-18 09:08:22', '管理员');
INSERT INTO `sys_permission` VALUES ('120005', '120001', '存储设备', 'AssetStorage', '/asset/storage', 'views/storage/index', '', '', 1, 13, '2021-07-09 12:30:00', '管理员', '2021-11-18 09:08:22', '管理员');
INSERT INTO `sys_permission` VALUES ('120006', '120001', '备件', 'ArticlSpare', '/asset/spare', 'views/spare/index', '', '', 1, 14, '2021-07-09 12:30:00', '管理员', '2021-11-18 09:08:22', '管理员');

INSERT INTO `sys_permission` VALUES ('120007', '120002', '用户管理', 'SettingUser', '/setting/user', 'views/user/index', '', '', 1, 11, '2021-07-09 12:30:00', '管理员', '2022-03-24 16:34:56', '管理员');
INSERT INTO `sys_permission` VALUES ('120008', '120002', '角色管理', 'SettingRole', '/settting/role', 'views/role/index', '', '', 1, 12, '2021-07-09 12:30:00', '管理员', '2022-03-24 16:35:30', '管理员');
INSERT INTO `sys_permission` VALUES ('120009', '120002', '权限管理', 'SettingMenu', '/setting/permission', 'views/permission/index', '', '', 1, 13, '2021-07-09 12:30:00', '管理员', '2022-03-24 16:35:18', '管理员');
INSERT INTO `sys_permission` VALUES ('120010', '120002', '部门管理', 'SettingDepartment', '/setting/department', 'views/department/index', '', '', 1, 14, '2021-07-09 12:30:00', '管理员', '2021-07-27 16:46:09', '管理员');

INSERT INTO `sys_permission` VALUES ('120011', '120007', '视图', '', '/setting/user/get', '', '', '', 2, 101, '2022-03-25 14:06:20', '管理员', '2022-03-25 14:09:57', '管理员');
INSERT INTO `sys_permission` VALUES ('120012', '120007', '新增',  '', '/setting/user/post', '', '', '', 2, 102, '2022-03-25 13:53:38', '管理员', '2022-03-25 14:10:06', '管理员');
INSERT INTO `sys_permission` VALUES ('120013', '120007', '修改', '', '/setting/user/patch', '', '', '', 2, 103, '2022-03-25 13:54:40', '管理员', '2022-03-25 13:55:43', '管理员');
INSERT INTO `sys_permission` VALUES ('120014', '120007', '删除', '', '/setting/user/delete', '', '', '', 2, 104, '2022-03-25 13:54:58', '管理员', '2022-03-25 13:55:57', '管理员');
INSERT INTO `sys_permission` VALUES ('120015', '120007', '重置密码', '', '/setting/user/password', '', '', '', 2, 105, '2022-03-25 13:55:20', '管理员', '2022-03-25 13:55:20', '管理员');

INSERT INTO `sys_permission` VALUES ('120016', '120008', '视图', '', '/setting/role/get', '', '', '', 2, 101, '2022-03-25 14:06:44', '管理员', '2022-03-25 14:10:15', '管理员');
INSERT INTO `sys_permission` VALUES ('120017', '120008', '新增', '', '/setting/role/post', '', '', '', 2, 102, '2022-03-25 13:56:51', '管理员', '2022-03-25 14:10:33', '管理员');
INSERT INTO `sys_permission` VALUES ('120018', '120008', '修改', '', '/setting/role/patch', '', '', '', 2, 103, '2022-03-25 13:57:25', '管理员', '2022-03-25 13:57:25', '管理员');
INSERT INTO `sys_permission` VALUES ('120019', '120008', '删除', '', '/setting/role/delete', '', '', '', 2, 104, '2022-03-25 13:57:43', '管理员', '2022-03-25 13:57:43', '管理员');

INSERT INTO `sys_permission` VALUES ('120020', '120009', '视图', '', '/setting/permission/get', '', '', '', 2, 101, '2022-03-25 14:05:50', '管理员', '2022-03-25 14:09:31', '管理员');
INSERT INTO `sys_permission` VALUES ('120021', '120009', '新增', '', '/setting/permission/post', '', '', '', 2, 102, '2022-03-25 13:50:51', '管理员', '2022-03-25 14:09:47', '管理员');
INSERT INTO `sys_permission` VALUES ('120022', '120009', '修改', '', '/setting/permission/patch', '', '', '', 2, 103, '2022-03-25 13:51:35', '管理员', '2022-03-25 13:52:39', '管理员');
INSERT INTO `sys_permission` VALUES ('120023', '120009', '删除', '', '/setting/permission/delete', '', '', '', 2, 104, '2022-03-25 13:51:56', '管理员', '2022-03-25 13:52:46', '管理员');

INSERT INTO `sys_permission` VALUES ('120024', '120010', '视图', '', '/setting/department/get', '', '', '', 2, 101, '2022-03-25 14:07:13', '管理员', '2022-03-25 14:10:40', '管理员');
INSERT INTO `sys_permission` VALUES ('120025', '120010', '新增', '', '/setting/department/post', '', '', '', 2, 102, '2022-03-25 13:58:58', '管理员', '2022-03-25 14:13:01', '管理员');
INSERT INTO `sys_permission` VALUES ('120026', '120010', '修改', '', '/setting/department/patch', '', '', '', 2, 103, '2022-03-25 13:59:08', '管理员', '2022-03-25 13:59:08', '管理员');
INSERT INTO `sys_permission` VALUES ('120027', '120010', '删除', '', '/setting/department/delete', '', '', '', 2, 104, '2022-03-25 13:59:22', '管理员', '2022-03-25 13:59:22', '管理员');

DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` varchar(20) NOT NULL COMMENT '角色ID',
  `role_name` varchar(20) NOT NULL COMMENT '角色名称',
  `description` varchar(100) NOT NULL COMMENT '角色描述',
  `create_at` datetime NOT NULL COMMENT '创建时间',
  `create_user` varchar(20) NOT NULL COMMENT '创建人',
  `update_at` datetime COMMENT '最后一次修改时间',
  `update_user` varchar(20) COMMENT '最后一次修改人',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='[ 权限 ] 角色表';

INSERT INTO `sys_role` VALUES ('110000', '管理员', '管理员', '2021-07-05 15:02:03', '管理员', '2021-12-17 09:50:21', '管理员');
INSERT INTO `sys_role` VALUES ('110001', '研发人员', '研发人员', '2021-12-17 09:51:27', '管理员', '2022-03-25 14:13:50', '管理员');
INSERT INTO `sys_role` VALUES ('110002', '工程人员', '工程人员', '2021-12-17 09:51:59', '管理员', '2022-03-25 14:13:58', '管理员');
INSERT INTO `sys_role` VALUES ('110003', '销售人员', '销售人员', '2021-12-17 09:58:41', '管理员', '2022-03-24 17:51:20', '管理员');

DROP TABLE IF EXISTS `sys_role_permission`;
CREATE TABLE `sys_role_permission` (
  `role_id` varchar(20) NOT NULL COMMENT '角色ID',
  `permission_id` varchar(20) NOT NULL COMMENT '菜单ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='[ 权限 ] 角色和菜单权限表';

INSERT INTO `sys_role_permission` VALUES ('110000', '120001');
INSERT INTO `sys_role_permission` VALUES ('110000', '120002');
INSERT INTO `sys_role_permission` VALUES ('110000', '120003');
INSERT INTO `sys_role_permission` VALUES ('110000', '120004');
INSERT INTO `sys_role_permission` VALUES ('110000', '120005');
INSERT INTO `sys_role_permission` VALUES ('110000', '120006');
INSERT INTO `sys_role_permission` VALUES ('110000', '120007');
INSERT INTO `sys_role_permission` VALUES ('110000', '120008');
INSERT INTO `sys_role_permission` VALUES ('110000', '120009');
INSERT INTO `sys_role_permission` VALUES ('110000', '120010');
INSERT INTO `sys_role_permission` VALUES ('110000', '120011');
INSERT INTO `sys_role_permission` VALUES ('110000', '120012');
INSERT INTO `sys_role_permission` VALUES ('110000', '120013');
INSERT INTO `sys_role_permission` VALUES ('110000', '120014');
INSERT INTO `sys_role_permission` VALUES ('110000', '120015');
INSERT INTO `sys_role_permission` VALUES ('110000', '120016');
INSERT INTO `sys_role_permission` VALUES ('110000', '120017');
INSERT INTO `sys_role_permission` VALUES ('110000', '120018');
INSERT INTO `sys_role_permission` VALUES ('110000', '120019');
INSERT INTO `sys_role_permission` VALUES ('110000', '120020');
INSERT INTO `sys_role_permission` VALUES ('110000', '120021');
INSERT INTO `sys_role_permission` VALUES ('110000', '120022');
INSERT INTO `sys_role_permission` VALUES ('110000', '120023');
INSERT INTO `sys_role_permission` VALUES ('110000', '120024');
INSERT INTO `sys_role_permission` VALUES ('110000', '120025');
INSERT INTO `sys_role_permission` VALUES ('110000', '120026');
INSERT INTO `sys_role_permission` VALUES ('110000', '120027');

INSERT INTO `sys_role_permission` VALUES ('110001', '120001');
INSERT INTO `sys_role_permission` VALUES ('110001', '120002');
INSERT INTO `sys_role_permission` VALUES ('110001', '120003');
INSERT INTO `sys_role_permission` VALUES ('110001', '120004');
INSERT INTO `sys_role_permission` VALUES ('110001', '120005');
INSERT INTO `sys_role_permission` VALUES ('110001', '120006');
INSERT INTO `sys_role_permission` VALUES ('110001', '120007');
INSERT INTO `sys_role_permission` VALUES ('110001', '120008');
INSERT INTO `sys_role_permission` VALUES ('110001', '120009');
INSERT INTO `sys_role_permission` VALUES ('110001', '120010');
INSERT INTO `sys_role_permission` VALUES ('110001', '120011');
INSERT INTO `sys_role_permission` VALUES ('110001', '120016');
INSERT INTO `sys_role_permission` VALUES ('110001', '120020');
INSERT INTO `sys_role_permission` VALUES ('110001', '120024');

INSERT INTO `sys_role_permission` VALUES ('110002', '120001');
INSERT INTO `sys_role_permission` VALUES ('110002', '120003');
INSERT INTO `sys_role_permission` VALUES ('110002', '120004');
INSERT INTO `sys_role_permission` VALUES ('110002', '120005');
INSERT INTO `sys_role_permission` VALUES ('110002', '120006');

DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role` (
  `user_id` varchar(20) NOT NULL COMMENT '用户ID',
  `role_id` varchar(20) NOT NULL COMMENT '角色ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='[ 权限 ] 用户和角色表';

INSERT INTO `sys_user_role` VALUES ('000000000', '110000');
INSERT INTO `sys_user_role` VALUES ('000000001', '110001');
INSERT INTO `sys_user_role` VALUES ('000000002', '110001');
INSERT INTO `sys_user_role` VALUES ('000000003', '110002');
INSERT INTO `sys_user_role` VALUES ('000000004', '110002');
INSERT INTO `sys_user_role` VALUES ('000000005', '110002');
INSERT INTO `sys_user_role` VALUES ('000000006', '110003');

DROP TABLE IF EXISTS `sys_user_department`;
CREATE TABLE `sys_user_department` (
  `user_id` varchar(20) NOT NULL COMMENT '用户ID',
  `department_id` varchar(20) NOT NULL COMMENT '部门ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='[ 权限 ] 用户和部门表';

INSERT INTO `sys_user_department` VALUES ('000000000', '100000');
INSERT INTO `sys_user_department` VALUES ('000000001', '100004');
INSERT INTO `sys_user_department` VALUES ('000000002', '100006');
INSERT INTO `sys_user_department` VALUES ('000000003', '100007');
INSERT INTO `sys_user_department` VALUES ('000000004', '100012');
INSERT INTO `sys_user_department` VALUES ('000000005', '100013');
INSERT INTO `sys_user_department` VALUES ('000000006', '100003');

DROP TABLE IF EXISTS `sys_department_permission`;
CREATE TABLE `sys_department_permission` (
  `department_id` varchar(20) NOT NULL COMMENT '部门ID',
  `permission_id` varchar(20) NOT NULL COMMENT '菜单ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='[ 权限 ] 部门和菜单权限表';