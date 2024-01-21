import api from "@/api/api";
import { reactive } from "vue";

export const userInfo = reactive({
  Name: "",
  async getInfo() {
    try {
      const res = await api.getUserName(0);
      this.Name = res.data.Name;
    } catch (err) {
      console.error(err);
    }
  },
});
