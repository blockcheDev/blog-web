import api from "@/api/api";
import { reactive, ref } from "vue";

export interface Article {
  ID: number;
  CreatedAt: string;
  ModifiedAt: string;
  Category: Category;
  User: User;
  Title: string;
  Content: string;
  Type: number; // 0-原创 1-转载
  Tags: Tag[];
  PageViews: number;
  Likes: number; // 点赞数
}
export interface Category {
  ID: number;
  Name: string;
}
export interface User {
  ID: number;
  Name: string;
}
export interface Tag {
  ID: number;
  Name: string;
}
export interface Comment {
  ID: number;
  CreatedAt: string;
  User: User;
  ArticleID: number;
  ReplyTopCommentID: number;
  ReplyCommentID: number;
  Content: string;
  Favor: number;

  Floor: number;
}

const defaultData: Article = Object.freeze({
  ID: 0,
  CreatedAt: "",
  ModifiedAt: "",
  Category: {
    ID: 0,
    Name: "",
  },
  User: {
    ID: 0,
    Name: "",
  },
  Title: "",
  Content: "",
  Type: 0,
  Tags: [],
  PageViews: 0,
  Likes: 0,
});

const list: Article[] = reactive([]);

const allCategory: Category[] = reactive([]);
const getAllCategory = async () => {
  try {
    const res = await api.getCategory("all");
    Object.assign(allCategory, res.data);
  } catch (err) {
    console.error(err);
  }
};
const allTag: Tag[] = reactive([]);
const getAllTag = async () => {
  try {
    const res = await api.getTag("all");
    Object.assign(allTag, res.data);
  } catch (err) {
    console.error(err);
  }
};

export { defaultData, list, allCategory, getAllCategory, allTag, getAllTag };
