export const _asyncRoutes = [
  // {
  //   name: "Dashboard",
  //   path: "/",
  //   component: () => import('@/layout/index.vue'),
  //   redirect: "/Dashboard/Workplace",
  //   meta: { title: "仪表盘", icon: "el-icon-eleme" },
  //   children: [
  //     {
  //       name: "Workplace",
  //       path: "/Dashboard/Workplace",
  //       component: () => import('@/views/Dashboard/home.vue'),
  //       meta: { title: "首页(正式服v1.0.2)", icon: "el-icon-s-tools" },
  //     }
  //   ],
  // },
  {
    path: "/",
    name: "Notice",
    alwaysShow: true, 
    component: () => import('@/layout/index.vue'),
    redirect: "/notice/index",
    meta: { title: "公告告警",img: 'authActive'},
    children: [
      {
        path: "/notice/index",
        name: "NoticeIndex",
        component: () => import('@/views/Notice/index.vue'),
        meta: { title: "公告告警" ,img: 'data',icon:"el-icon-lock"},
      }
    ],
  },
  {
    path: "/order",
    name: "Order",
    alwaysShow: true, 
    component: () => import('@/layout/index.vue'),
    redirect: "/order/index",
    meta: { title: "轮询锁单",img: 'authActive'},
    children: [
      {
        path: "/order/index",
        name: "order-index",
        component: () => import('@/views/Order/index.vue'),
        meta: { title: "自动锁单" ,img: 'data',icon:"el-icon-lock"},
      },
      {
        path: "/order/login",
        name: "order-login",
        component: () => import('@/views/Order/relogin.vue'),
        meta: { title: "账号登陆" ,img: 'data',icon:"el-icon-lock"},
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