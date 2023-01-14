import axios from "axios";

const request = axios.create({
    // baseURL 需要设置为反向代理前缀，不能设置绝对路径URL
    baseURL: import.meta.env.VITE_BASE_URL,
    timeout: 5000,
    withCredentials: false,
    headers: {'X-Custom-Header': 'zuiyu'}
});


export default request;