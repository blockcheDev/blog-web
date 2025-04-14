import request from "./axios";

export default {
  getUserName: (id: any) => request.get(`/api/user/${id}`),
  getArticle: (id: any) => request.get(`/api/article/${id}`),
  getCategory: (id: any) => request.get(`/api/category/${id}`),
  getTag: (id: any) => request.get(`/api/tag/${id}`),
  // getTagByArticle: (id: any) => request.get(`/api/article/${id}/tag`), //根据文章ID获取标签
  // getTagIDByArticle: (id: any) => request.get(`/api/article/${id}/tagid`), //根据文章ID获取标签ID
  getArticleListByCategory: (id: any) =>
    request.get(`/api/category/${id}/list`), // 根据分类id获取文章列表
  getArticleListByTag: (id: any) => request.get(`/api/tag/${id}/list`), // 根据标签id获取文章列表
  getCommentListByArticle: (id: any) =>
    request.get(`/api/article/${id}/comment`), // 根据文章id获取评论列表

  login: (form: any) => request.post("/api/login", form), // 用户登录
  register: (form: any) => request.post("/api/register", form), // 用户注册
  getUser: () => request.get("/api/user"), // 获取用户信息
  modifyUser: (form: any) => request.put("/api/user", form), // 更新用户信息
  modifyUserPassword: (form: any) => request.put("/api/user/password", form), // 修改用户密码
  pushArticle: (form: any) => request.post("/api/article", form),
  createTag: (data: any) => request.post(`/api/tag`, data),
  createCategory: (data: any) => request.post(`/api/category`, data),
  pushComment: (data: any) => request.post("/api/comment", data),

  // deleteUser: (data: any) => request.delete(`/api/user`, {data: data}),
  deleteUser: (data: any) =>
    request({ url: `/api/user`, method: "delete", data: data }), // 删除用户

  deleteArticle: (id: any) => request.delete(`/api/article/${id}`),
  deleteComment: (id: any) => request.delete(`/api/comment/${id}`),

  modifyArticle: (data: any) => request.put("/api/article", data),

  // 使用github登录
  loginWithGithub: (code: string) =>
    request.get(`/api/login/github?code=${code}`),

  getRecentVisitorsCount: () => request.get("/api/recent_visitors_count"), // 获取最近访客数量
  getRecentVisitors: () => request.get("/api/recent_visitors") // 获取最近访客ip和region
};
