import router from "@/router";
import axios from "axios";
import { ElMessage } from "element-plus";
const request = axios.create({
  baseURL: "http://localhost:8080", //这里配置的是后端服务提供的接口
  // baseURL: "http://www.hitori.cn:8080",
  timeout: 1000,
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
      localStorage.setItem("token", "");
      // ElMessage({
      // 	message: "登录过期，请重新登录",
      // 	type: 'warning'
      // })
      router.push("/login");
    }
    ElMessage({
      message: err.response.data.msg,
      type: "warning",
    });
    return Promise.reject(err);
  }
);

export default request;
