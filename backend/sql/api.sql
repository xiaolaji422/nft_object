--  1.0.0

INSERT INTO `t_auth_api_group` (`id`,`name`) VALUES (1,"接口管理");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (1,"所有接口","/admin/api/all","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (1,"接口列表","/admin/api/items","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (1,"接口详情","/admin/api/info","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (1,"接口修改","/admin/api/edit","POST");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (1,"启用/禁用","/admin/api/enable","POST");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (1,"接口添加","/admin/api/add","POST");


INSERT INTO `t_auth_api_group` (`id`,`name`) VALUES (2,"员工管理");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (2,"员工列表","/admin/admin/items","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (2,"员工详情","/admin/admin/info","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (2,"员工角色","/admin/admin/roles","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (2,"员工授权","/admin/admin/apis","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (2,"员工权限详情","/admin/admin/apiArray","GET");


INSERT INTO `t_auth_api_group` (`id`,`name`) VALUES (3,"角色管理");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (3,"角色修改","/admin/role/update","POST");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (3,"角色权限列表","/admin/role/getRoleApis","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (3,"设置角色权限","/admin/role/setRoleApis","POST");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (3,"角色列表","/admin/role/items","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (3,"所有角色","/admin/role/all","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (3,"角色添加","/admin/role/add","POST");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (3,"角色详情","/admin/role/info","GET");



INSERT INTO `t_auth_api_group` (`id`,`name`) VALUES (4,"用户权限管理");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (4,"启用/禁用","/admin/adminApi/enable","POST");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (4,"接口授权","/admin/adminApi/addApis","POST");



INSERT INTO `t_auth_api_group` (`id`,`name`) VALUES (5,"用户角色管理");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (5,"启用/禁用","/admin/adminRole/enable","POST");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (5,"添加角色","/admin/adminRole/addRole","POST");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (5,"删除角色","/admin/adminRole/delRole","POST");


INSERT INTO `t_auth_api_group` (`id`,`name`) VALUES (6,"接口分组");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (6,"列表(所有)","/admin/apiGroup/all","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (6,"列表(分页)","/admin/apiGroup/items","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (6,"查看详情","/admin/apiGroup/info","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (6,"添加分组","/admin/apiGroup/add","POST");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (6,"修改分组","/admin/apiGroup/edit","POST");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (6,"启用/禁用","/admin/apiGroup/enabled","POST");


INSERT INTO `t_auth_api_group` (`id`,`name`) VALUES (7,"公告管理");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (7,"最新公告","/admin/notice/query_notice","GET");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (7,"历史公告","/admin/notice/query_his_notice","GET");


INSERT INTO `t_auth_api_group` (`id`,`name`) VALUES (8,"锁单管理");
INSERT INTO `t_auth_api` (`group_id`,`name`,`route`,`methods`) VALUES (8,"商品查找","/admin/album/search","GET");