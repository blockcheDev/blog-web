import request from "@/api/axios";
import { createRouter, createWebHistory } from "vue-router";
import { ElMessage } from "element-plus";
import { drawerMenu } from "@/store/store";

const routes = [
  {
    path: "/",
    // 重定向
    // redirect: '/welcome',
    redirect: "/home",
  },
  {
    path: "/home",
    name: "首页",
    component: () => import("../views/home.vue"),
  },
  {
    path: "/user/info",
    name: "个人资料",
    component: () => import("../views/user/user.vue"),
  },
  {
    path: "/user/setting",
    name: "账户设置",
    component: () => import("../views/user/userSetting.vue"),
  },
  {
    path: "/user/push",
    name: "发布文章",
    component: () => import("../views/user/article/pushArticle.vue"),
  },
  {
    path: "/user/article",
    name: "文章管理",
    component: () => import("../views/user/article/articleManage.vue"),
  },
  {
    path: "/user/edit/article/:id",
    name: "修改文章",
    component: () => import("../views/user/article/editArticle.vue"),
  },
  {
    path: "/login",
    name: "登录",
    component: () => import("../views/login.vue"),
  },
  {
    path: "/register",
    name: "注册",
    component: () => import("../views/register.vue"),
  },
  {
    path: "/article/:id",
    component: () => import("../views/article/index.vue"),
  },
  {
    path: "/category",
    component: () => import("../views/category/index.vue"),
  },
  {
    path: "/category/:id",
    component: () => import("../views/category/show.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  const tokenString = localStorage.getItem("token");
  if (to.path.indexOf("/user") === 0 && !tokenString) {
    ElMessage({
      message: "请先登录",
    });
    next("/login");
  } else {
    next();
  }
});
router.afterEach((to, from) => {
  drawerMenu.close();
});

// async function isAccept() {
//   try {
//     const res = await api.post("/auth")
//     console.log(res)
//     ElMessage({
//       message: res.data.msg,
//       type: res.data.msg === "成功" ? 'success' : 'warning',
//     })
//     return (res.data.msg === "成功")
//   } catch (error) {
//     console.error('请求失败', error)
//     throw error
//   }
// }

export default router;
