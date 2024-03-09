import router from "@/router";
import axios from "axios";
import { ElMessage } from "element-plus";
const request = axios.create({
  // baseURL: "http://localhost:8080", //这里配置的是后端服务提供的接口
  // baseURL: "http://www.hitori.cn:8080",
  baseURL: import.meta.env.VITE_API_URL as string,
  timeout: 5000,
});

request.interceptors.request.use(
  (req) => {
    req.headers.set("token", localStorage.getItem("token"));
    return req;
  },
  (err) => {
    return Promise.reject(err);
  }
);

request.interceptors.response.use(
  (res) => {
    // 如果状态码是2xx，表示是成功的结果
    if (res.data.msg) {
      ElMessage({
        message: res.data.msg,
        type: "success",
      });
    }
    return res;
  },
  (err) => {
    // 401表示认证失败，可能是token过期或错误
    if (err.response.status === 401) {
      if (err.response.data.msg === "请登录") {
        localStorage.setItem("token", "");
        router.push("/login");
      }
      // ElMessage({
      // 	message: "登录过期，请重新登录",
      // 	type: 'warning'
      // })
    }
    ElMessage({
      message: err.response.data.msg,
      type: "warning",
    });
    return Promise.reject(err);
  }
);

export default request;
