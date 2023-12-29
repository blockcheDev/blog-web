import request from "./axios";

export default {
  getArticle: (id: any) => request.get(`/getInfo/article/${id}`),
  getCategory: (id: any) => request.get(`/getInfo/category/${id}`),
  getTag: (id: any) => request.get(`/getInfo/tag/${id}`),
  getTagByArticle: (id: any) => request.get(`/getInfo/article/${id}/tag`), //根据文章ID获取标签
  getTagIDByArticle: (id: any) => request.get(`/getInfo/article/${id}/tagid`), //根据文章ID获取标签ID
  getArticleListByCategory: (id: any) => request.get(`/category/${id}`),
  getArticleListByTag: (id: any) => request.get(`/tag/${id}`),

  pushArticle: (form: any) => request.post("/user/push", form),
  createTag: (data: any) => request.post(`/user/create/tag`, data),
  createCategory: (data: any) => request.post(`/user/create/category`, data),

  deleteUser: (data: any) => request.post(`/user/delete/user`, data),
  deleteArticle: (id: any) => request.post(`/user/delete/article/${id}`),

  modifyArticle: (data: any) => request.post("/user/modify/article", data),
};
