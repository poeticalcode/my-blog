import axios from 'axios'
import { ElMessage } from 'element-plus'

// 创建一个 axios 实例
const request = axios.create({
  baseURL: import.meta.env.VITE_BASE_URL + "/api", // 所有的请求地址前缀部分
  timeout: 8000, // 请求超时时间毫秒
})


// 添加请求拦截器
request.interceptors.request.use((req) => {
  // 在发送请求之前做些什么
  return req
}, (error) => {
  // 对请求错误做些什么
  console.error('请求错误', error)
  return Promise.reject(error)
})


// 添加响应拦截器
request.interceptors.response.use((res) => {
  return res.data
}, (error) => {
  return Promise.reject(error)
})

export default request