import { message } from 'antd'
import axios, { AxiosRequestConfig, AxiosError, AxiosResponse } from 'axios'
import qs from 'qs'
import { LoginTokenStore } from '../stores/local'
import { getUUId } from '../utils/utils'
import { BaseResponse } from '../types/baseType'
import { getHttpDebug } from '@/config/url'

const httpService = axios.create({
  timeout: 2 * 60 * 1000, // 默认请求超时时间
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/x-www-form-urlencoded'
  }
})

// request 请求拦截器
httpService.interceptors.request.use(
  (config: AxiosRequestConfig) => {
    if (config.params) {
      for (const item in config.params) {
        if (item === null) {
          delete config.params[item]
        }
      }
    }
    if (config && config.headers) {
      const loginToken = LoginTokenStore.getToken()
      if (loginToken) {
        config.headers['X-KT-Login-Token'] = loginToken
      }
      config.headers['X-KT-Request-Id'] = getUUId()
      config.headers['X-KT-Timestamp'] = new Date().getTime().toString()
      config.headers['X-KT-Debug'] = getHttpDebug()
    }
    return config
  },
  (error: AxiosError) => {
    return Promise.reject(error)
  }
)

// response 返回结果拦截器
httpService.interceptors.response.use(
  async (response: AxiosResponse<BaseResponse<any>>) => {
    const res = response.data
    switch (res.code) {
      case 0: // 正常数据
        return response
      case 10104:
        LoginTokenStore.removeToken()
        await message.error('账号登录超时！')
        window.location.href = `${window.location.origin}`
        return Promise.reject(new Error(res.message))
    }
    console.log('请求错误:', res)
    if (res.message) {
      await message.error(res.message)
    }
    return Promise.reject(new Error(res.message))
  },
  async (error) => {
    const status = error.response?.status
    switch (status) {
      case 401:
        await message.error('账号登录态失效！', 2)
        LoginTokenStore.removeToken()
        window.location.href = `${window.location.origin}`
        return Promise.reject(error)
      case 403:
        await message.error('暂无数据访问权限！', 2)
        return Promise.reject(error)
      case 500:
        await message.error('服务器未知错误！', 2)
        return Promise.reject(error)
    }
    return Promise.reject(error)
  }
)

const httpRequest = {
  get<T>(url: string, params = {}): Promise<T> {
    const options: AxiosRequestConfig = {
      url: url,
      method: 'GET',
      params: params,
      responseType: 'json'
    }
    return new Promise((resolve, reject) => {
      httpService(options)
        .then((res: AxiosResponse<BaseResponse<T>>) => {
          return resolve(res.data.data)
        })
        .catch((err) => {
          return reject(err)
        })
    })
  },
  post<T>(url: string, params = {}, data = {}): Promise<T> {
    const options: AxiosRequestConfig = {
      url: url,
      method: 'POST',
      params: params,
      data: qs.stringify(data, { arrayFormat: 'repeat' }),
      responseType: 'json'
    }
    return new Promise((resolve, reject) => {
      httpService(options)
        .then((res: AxiosResponse<BaseResponse<T>>) => {
          return resolve(res.data.data)
        })
        .catch((err) => {
          return reject(err)
        })
    })
  },
  postFile<T>(url: string, params = {}): Promise<T> {
    const options: AxiosRequestConfig = {
      url: url,
      method: 'POST',
      data: qs.stringify(params),
      responseType: 'json'
    }
    return new Promise((resolve, reject) => {
      httpService(options)
        .then((res: AxiosResponse<BaseResponse<T>>) => {
          return resolve(res.data.data)
        })
        .catch((err) => {
          return reject(err)
        })
    })
  }
}

export default httpRequest
