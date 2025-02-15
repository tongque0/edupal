import axios from 'axios';

const baseURL = import.meta.env.VITE_APP_API_BASE_URL || "http://127.0.0.1:8080";
// 创建一个 axios 实例
const service = axios.create({
  baseURL: baseURL, // 后端 API 的基础 URL
  timeout: 5000, // 请求超时时间
});

// 请求拦截器
service.interceptors.request.use(
  config => {
    // 在发送请求之前做些什么
    // 例如，在请求头中添加 token
    const token = localStorage.getItem('token');
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
    }
    return config;
  },
  error => {
    // 对请求错误做些什么
    console.error('请求错误:', error);
    return Promise.reject(error);
  }
);

// 响应拦截器
service.interceptors.response.use(
  response => {
    // 对响应数据做些什么
    return response.data;
  },
  error => {
    // 对响应错误做些什么
    console.error('响应错误:', error);
    return Promise.reject(error);
  }
);

export default service;
