import axios, { AxiosResponse } from 'axios';
import Cookies from 'js-cookie';

// 在这里定义你的基础URL
const BASE_URL = 'http://127.0.0.1:8080';
const CSRF_URL = BASE_URL + 'csrf';
const LOGIN_URL = BASE_URL + 'login';

// 创建 axios 实例
const axiosInstance = axios.create({
  baseURL: BASE_URL,
  timeout: 30000 // 请求超时时间
});

let csrfToken = Cookies.get('csrfToken');
let authToken = Cookies.get('authToken');

// 请求拦截器
axiosInstance.interceptors.request.use(
  async (config) => {
    // 获取 csrf_token
    // if (!csrfToken) {
    //   try {
    //     const response = await axios.get(CSRF_URL);
    //     csrfToken = "response.data.csrfToken";
    //     Cookies.set('csrfToken', csrfToken);
    //   } catch (error) {
    //     console.log('Failed to update CSRF token');
    //     return Promise.reject(error);
    //   }
    // }

    // config.headers['X-CSRF-Token'] = csrfToken;

    // // 获取 auth token
    // if (!authToken) {
    //   try {
    //     const response = await axios.get(LOGIN_URL);
    //     authToken = "response.data.authToken";
    //     Cookies.set('authToken', authToken);
    //   } catch (error) {
    //     console.log('Failed to update Auth token');
    //     return Promise.reject(error);
    //   }
    // }

    // config.headers['Authorization'] = `Bearer ${authToken}`;

    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器
axiosInstance.interceptors.response.use(
  (response: AxiosResponse) => {
    // 在这里你可以对响应进行处理
    return response.data;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export const getRequest = (url: string, config?: object) => {
  return axiosInstance.get(url, config);
};

export const postRequest = (url: string, data?: any, config?: object) => {
  return axiosInstance.post(url, data, config);
};

export const putRequest = (url: string, data?: any, config?: object) => {
  return axiosInstance.put(url, data, config);
};

export const deleteRequest = (url: string, config?: object) => {
  return axiosInstance.delete(url, config);
};
