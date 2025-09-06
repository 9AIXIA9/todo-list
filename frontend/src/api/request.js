import axios from "axios";

//创建axios实例

const request = axios.create({
  baseURL: "http://localhost:8088/api/v1",
  timeout: 5000,
  headers: {
    "Content-Type": "application/json",
  },
});

//请求拦截器
request.interceptors.request.use(
  (config) => {
    console.log("send request:", config.method?.toUpperCase(), config.url);
    return config;
  },
  (error) => {
    console.error("request error:", error);
    return Promise.reject(error);
  }
);

//响应拦截器
request.interceptors.response.use(
  (response) => {
    console.log("receive response:", response.data);
    return response.data;
  },
  (error) => {
    console.error("响应错误:", error);
    return Promise.reject(error);
  }
);

export default request;
