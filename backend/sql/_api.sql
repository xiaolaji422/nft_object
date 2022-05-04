DROP TABLE IF EXISTS `t_auth_admin_api`;
CREATE TABLE `t_auth_admin_api` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID号',
  `admin_id` int(11) NOT NULL COMMENT '人员id',
  `api_id` int(11) NOT NULL COMMENT '接口id',
  `api_name` varchar(100) NOT NULL COMMENT '接口名称',
  `enabled` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态',
  `create_user` varchar(50) DEFAULT '' COMMENT '创建人',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_user` varchar(50) DEFAULT '' COMMENT '修改人',
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_api` (`admin_id`,`api_id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='人员额外接口权限设计表';

DROP TABLE IF EXISTS `t_auth_admin_role`;
CREATE TABLE `t_auth_admin_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID号',
  `admin_id` int(11) NOT NULL COMMENT '角色id',
  `role_id` int(11) NOT NULL COMMENT '角色id',
  `role_name` varchar(50) NOT NULL COMMENT '角色名称',
  `enabled` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态',
  `create_user` varchar(50) DEFAULT '' COMMENT '创建人',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_user` varchar(50) DEFAULT '' COMMENT '修改人',
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `admin_role` (`admin_id`,`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8 COMMENT='人员角色表';

DROP TABLE IF EXISTS `t_auth_api`;
CREATE TABLE `t_auth_api` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID号',
  `group_id` int(11) NOT NULL DEFAULT '0' COMMENT '接口分组',
  `methods` varchar(10) NOT NULL DEFAULT 'GET' COMMENT '请求方法',
  `name` varchar(50) NOT NULL COMMENT '接口名称',
  `route` varchar(100) NOT NULL COMMENT '路径',
  `enabled` tinyint(1) NOT NULL DEFAULT 1 COMMENT '接口状态',
  `limit` tinyint(2) NOT NULL DEFAULT '0' COMMENT '接口限频/s  0=不限制   最大值10',
  `create_user` varchar(50) DEFAULT '' COMMENT '创建人',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_user` varchar(50) DEFAULT '' COMMENT '修改人',
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=92 DEFAULT CHARSET=utf8 COMMENT='"系统api接口表"';

DROP TABLE IF EXISTS `t_auth_api_group`;
CREATE TABLE `t_auth_api_group` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID号',
  `name` varchar(50) NOT NULL COMMENT '分组名称',
  `enabled` tinyint(1) NOT NULL DEFAULT 1 COMMENT '分组状态',
  `create_user` varchar(50) DEFAULT '' COMMENT '创建人',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_user` varchar(50) DEFAULT '' COMMENT '修改人',
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8 COMMENT='"系统api分组表"';


DROP TABLE IF EXISTS `t_auth_role`;
CREATE TABLE `t_auth_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID号',
  `name` varchar(50) NOT NULL COMMENT '角色名称',
  `enabled` tinyint(1) NOT NULL DEFAULT 1 COMMENT '角色状态',
  `create_user` varchar(50) DEFAULT '' COMMENT '创建人',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_user` varchar(50) DEFAULT '' COMMENT '修改人',
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='角色表';

DROP TABLE IF EXISTS `t_auth_role_api`;
CREATE TABLE `t_auth_role_api` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID号',
  `role_id` int(11) NOT NULL COMMENT '角色id',
  `apis` varchar(1000) NOT NULL COMMENT '路径',
  `enabled` tinyint(1) NOT NULL DEFAULT 1 COMMENT '角色状态',
  `create_user` varchar(50) DEFAULT '' COMMENT '创建人',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_user` varchar(50) DEFAULT '' COMMENT '修改人',
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `roleApi` (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8 COMMENT='系统角色权限表';

INSERT INTO `t_auth_api_group` (`id`,`name`) VALUES (1,"接口管理");
INSERT INTO `t_auth_api_group` (`id`,`name`) VALUES (2,"员工管理");
INSERT INTO `t_auth_api_group` (`id`,`name`) VALUES (3,"角色管理");
INSERT INTO `t_auth_api_group` (`id`,`name`) VALUES (4,"用户权限管理");
INSERT INTO `t_auth_api_group` (`id`,`name`) VALUES (5,"用户角色管理");
INSERT INTO `t_auth_api_group` (`id`,`name`) VALUES (6,"接口分组");


INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (1,"所有接口","/admin/api/all","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (1,"接口列表","/admin/api/items","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (1,"接口详情","/admin/api/info","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (1,"接口修改","/admin/api/edit","POST");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (1,"启用/禁用","/admin/api/enable","POST");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (1,"接口添加","/admin/api/add","POST");

INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (2,"员工列表","/admin/admin/items","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (2,"员工详情","/admin/admin/info","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (2,"员工角色","/admin/admin/roles","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (2,"员工授权","/admin/admin/apis","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (2,"员工权限详情","/admin/admin/apiArray","GET");

INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (3,"角色修改","/admin/role/update","POST");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (3,"角色权限列表","/admin/role/getRoleApis","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (3,"设置角色权限","/admin/role/setRoleApis","POST");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (3,"角色列表","/admin/role/items","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (3,"所有角色","/admin/role/all","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (3,"角色添加","/admin/role/add","POST");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (3,"角色详情","/admin/role/info","GET");

INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (4,"启用/禁用","/admin/adminApi/enable","POST");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (4,"接口授权","/admin/adminApi/addApis","POST");

INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (5,"启用/禁用","/admin/adminRole/enable","POST");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (5,"添加角色","/admin/adminRole/addRole","POST");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (5,"删除角色","/admin/adminRole/delRole","POST");

INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (6,"列表(所有)","/admin/apiGroup/all","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (6,"列表(分页)","/admin/apiGroup/items","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (6,"查看详情","/admin/apiGroup/info","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (6,"添加分组","/admin/apiGroup/add","POST");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (6,"修改分组","/admin/apiGroup/edit","POST");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (6,"启用/禁用","/admin/apiGroup/enabled","POST");



