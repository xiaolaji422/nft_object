export const _asyncRoutes = [
  {
    name: "Dashboard",
    path: "/",
    component: () => import('@/layout/index.vue'),
    redirect: "/Dashboard/Workplace",
    meta: { title: "仪表盘", icon: "el-icon-eleme" },
    children: [
      {
        name: "Workplace",
        path: "/Dashboard/Workplace",
        component: () => import('@/views/Dashboard/home.vue'),
        meta: { title: "首页(正式服v1.0.2)", icon: "el-icon-s-tools" },
      }
    ],
  },{
    path: "/tableDemo",
    name: "table",
    component: () => import('@/layout/index.vue'),
    redirect: "/tableDemo/index",
    meta: { title: "仪表盘", icon: "el-icon-eleme",hidden:true},
    children: [
      {
        path: "/tableDemo/index",
        name: "tableDemo",
        component: () => import('@/views/tableDemo/index.vue'),
        meta: { title: "表单示例", icon: "el-icon-s-tools" },
      }
    ],
  },
 {
    path: '/permission',
    component:()=> import('@/layout/index.vue'),
    redirect: '/permission/role',
    alwaysShow: true, 
    name: 'Permission',
    meta: {title: '后台权限',icon: 'lock', img: 'authActive'},
    children: [
      {
        path: '/permission/role',
        component: () => import('@/views/Permission/role.vue'),
        name: 'PermissionRole',
        meta: {title: '角色权限'},
      },
      {
        path: '/permission/user',
        component: () => import('@/views/Permission/user.vue'),
        name: 'UserPermission',
        meta: {  title: '用户权限'},
      },
      {
        path: '/permission/systemApi',
        component: () => import('@/views/Permission/api.vue'),
        name: 'systemApi',
        meta: { title: '接口管理', noCache: true }
      },{
        path: '/permission/systemApiGroup',
        component: () => import('@/views/Permission/apiGroup.vue'),
        name: 'systemApiGroup',
        meta: { title: '接口分组', noCache: true}
      }
    ],
  }
]