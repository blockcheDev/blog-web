import { createRouter, createWebHistory } from "vue-router";
import { ElMessage } from "element-plus";
import { drawerMenu, loginDialog } from "@/store/store";

const routes = [
  {
    path: "/oauth/redirect",
    name: "GitHub登录等待",
    component: () => import("../views/githubLogin.vue"),
  },
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
    // meta: { title: '首页 - blockche blog' },
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
  // {
  //   path: "/login",
  //   name: "登录",
  //   component: () => import("../views/login.vue"),
  // },
  // {
  //   path: "/register",
  //   name: "注册",
  //   component: () => import("../views/register.vue"),
  // },
  {
    path: "/article/:id",
    component: () => import("../views/article/index.vue"),
  },
  {
    path: "/category",
    name: "分类",
    component: () => import("../views/category/index.vue"),
  },
  {
    path: "/category/:id",
    component: () => import("../views/category/show.vue"),
  },
  {
    path: "/tag",
    name: "标签",
    component: () => import("../views/tag/index.vue"),
  },
  {
    path: "/tag/:id",
    component: () => import("../views/tag/show.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  // 设置title
  if (to.name) {
    document.title = to.name as string + " - blockche blog";  // 动态设置标题
  } else {
    document.title = "blockche blog"; // title的前导空格会被自动去掉。。。。。。
  }

  const tokenString = localStorage.getItem("token");
  if (to.path.indexOf("/user") === 0 && !tokenString) {
    ElMessage({
      message: "请先登录",
    });
    loginDialog.open();
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
