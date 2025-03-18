import axios from 'axios'
import requestEvent from '@/event/request'

export type Response<T = any> = { status: number; msg: string; data: T; type?: string }

// 默认配置
axios.defaults.baseURL = import.meta.env.VITE_HTTP_BASE_URL || '/api/v0'

// 请求拦截器
axios.interceptors.request.use(
    async (config) => {
        // 从 localStorage 获取 token
        const token = localStorage.getItem('token')

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
        const responseData = response.data as Response

        // 处理服务器错误
        if (responseData?.status >= 500 && responseData?.status <= 599) {
            requestEvent.emit('UnknownError')
        }

        // 处理未授权错误
        if (responseData?.status == 401) {
            requestEvent.emit('Unauthorized')
            // 把 token 删掉
            localStorage.removeItem('token')
        }

        // 处理特定消息错误
        if (responseData?.status === 418) {
            requestEvent.emit('Message', responseData?.type, responseData?.msg)
        }

        return response
    },
    async (error) => {
        console.error('请求错误:', error)

        // 网络异常处理
        if (error.code === 'ERR_NETWORK') {
            // 网络异常
            requestEvent.emit('NetworkError')
        } else if (error.response) {
            // 其他响应码
            const { status } = error.response

            if (status === 401) {
                requestEvent.emit('Unauthorized')
                // 清除本地保存的 token
                localStorage.removeItem('token')
            } else {
                requestEvent.emit('UnknownError')
            }

            console.error(`请求错误: ${status}`, error.response.data)
        } else {
            // 其他错误
            requestEvent.emit('UnknownError')
        }

        return Promise.reject(error)
    }
)

export default axios
