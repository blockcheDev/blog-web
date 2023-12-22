import request from "./axios";

export default {
    getArticle: (id: any) => request.get(`/getInfo/article/${id}`),
    getCategory: (id: any) => request.get(`/getInfo/category/${id}`),
    getTag: (id: any) => request.get(`/getInfo/tag/${id}`),
    getTagByArticle: (id: any) => request.get(`/getInfo/article/${id}/tag`), //根据文章ID获取标签
    pushArticle: (form: any) => request.post("/user/push", form),
    createTag: (data: any) => request.post(`/user/create/tag`, data),
    createCategory: (data: any) => request.post(`/user/create/category`, data)
}