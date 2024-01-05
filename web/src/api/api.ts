import request from "./axios";

export default {
  getUser: (id: any) => request.get(`/getInfo/user/${id}`),
  getArticle: (id: any) => request.get(`/getInfo/article/${id}`),
  getCategory: (id: any) => request.get(`/getInfo/category/${id}`),
  getTag: (id: any) => request.get(`/getInfo/tag/${id}`),
  getTagByArticle: (id: any) => request.get(`/getInfo/article/${id}/tag`), //根据文章ID获取标签
  getTagIDByArticle: (id: any) => request.get(`/getInfo/article/${id}/tagid`), //根据文章ID获取标签ID
  getArticleListByCategory: (id: any) => request.get(`/category/${id}`),
  getArticleListByTag: (id: any) => request.get(`/tag/${id}`),
  getCommentListByArticle: (id: any) => request.get(`/article/${id}/comment`),

  pushArticle: (form: any) => request.post("/user/push", form),
  createTag: (data: any) => request.post(`/user/create/tag`, data),
  createCategory: (data: any) => request.post(`/user/create/category`, data),
  pushComment: (data: any) => request.post("/user/create/comment", data),

  deleteUser: (data: any) => request.post(`/user/delete/user`, data),
  deleteArticle: (id: any) => request.post(`/user/delete/article/${id}`),
  deleteComment: (id: any) => request.delete(`/user/delete/comment/${id}`),

  modifyArticle: (data: any) => request.post("/user/modify/article", data),
};
