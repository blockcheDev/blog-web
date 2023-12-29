import api from "@/api/api";
import { reactive, ref } from "vue";

export interface Article {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  CategoryID: string;
  CategoryName: string;
  UserID: string;
  Title: string;
  Content: string;
  Type: number; // 0-原创 1-转载
  tags: string[];
}
export interface Category {
  ID: number;
  Name: string;
}
export interface Tag {
  ID: number;
  Name: string;
}

const defaultData: Article = Object.freeze({
  ID: 0,
  CreatedAt: "",
  UpdatedAt: "",
  CategoryID: "",
  CategoryName: "",
  UserID: "",
  Title: "",
  Content: "",
  Type: 0,
  tags: [],
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
