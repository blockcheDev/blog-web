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

const defaultData: Article = Object.freeze({
    ID: 0,
    CreatedAt: '',
    UpdatedAt: '',
    CategoryID: '',
    CategoryName: '',
    UserID: '',
    Title: '',
    Content: '',
    Type: 0,
    tags: [],
})

const list: Article[] = reactive([])

export { defaultData, list }
