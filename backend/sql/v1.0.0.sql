DROP TABLE IF EXISTS `t_admin`;
CREATE TABLE `t_admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `login_name` varchar(40) DEFAULT '' COMMENT "登录账户",
  `role_name` varchar(50) DEFAULT '' COMMENT '用户角色',
  `password` varchar(50) DEFAULT "" COMMENT "密码",
  `enabled` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态',
  `level` tinyint(1) DEFAULT 1 COMMENT "层级",
  `is_admin` tinyint(1) DEFAULT 0 COMMENT '是否是超级管理员',
  `modified_user` varchar(50) DEFAULT 'admin' COMMENT '配置人',
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `login_name` (`login_name`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='"人员账户表"';


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
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='人员额外接口权限设计表';

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
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='人员角色表';

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
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_route` (`route`),
  KEY `idx_group` (`group_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='"系统api接口表"';

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
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='"系统api分组表"';


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
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='角色表';

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
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='系统角色权限表';