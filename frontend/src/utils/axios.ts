import axios from 'axios'
import requestEvent from '@/event/request'
import { Cookie } from '@/utils/cookie'

export type Response<T = any> = { msg?: string; data?: T; type?: string }

// 默认配置
axios.defaults.baseURL = import.meta.env.VITE_HTTP_BASE_URL || '/api/v0'

// 请求拦截器
axios.interceptors.request.use(
    async (config) => {
        // 从 cookie 获取 token
        const token = Cookie.get('token')

        // 如果 token 存在，加到 header
        if (token && config.headers) {
            config.headers.Authorization = 'Bearer ' + token
        }

        return config
    },
    (error) => {
        console.error('请求拦截器错误:', error)
        return Promise.reject(error)
    }
)

// 响应拦截器
axios.interceptors.response.use(
    (response) => {
        // 200状态码不显示消息，直接返回
        return response
    },
    async (error) => {
        console.error('请求错误:', error)

        // 网络异常处理
        if (error.code === 'ERR_NETWORK') {
            // 网络异常
            requestEvent.emit('NetworkError')
            return Promise.reject(error)
        }

        // 有响应但状态码不是2xx
        if (error.response) {
            const { status } = error.response
            const responseData = error.response.data as Response
            const errorMsg = responseData?.msg || `错误 ${status}`

            // 服务器错误
            if (status >= 500 && status <= 599) {
                requestEvent.emit('Message', 'error', errorMsg)
                requestEvent.emit('UnknownError')
            }
            // 未授权错误
            else if (status === 401) {
                requestEvent.emit('Unauthorized')
                // 清除cookie中的token
                Cookie.remove('token')
                Cookie.remove('tokenExpiry')
                Cookie.remove('rememberMe')
            }
            // 其他客户端错误
            else {
                // 显示错误消息
                requestEvent.emit('Message', 'error', errorMsg)
            }

            console.error(`请求错误: ${status}`, error.response.data)
        } else {
            // 其他错误
            requestEvent.emit('UnknownError')
            requestEvent.emit('Message', 'error', '未知错误，请稍后重试')
        }

        return Promise.reject(error)
    }
)

export default axios
