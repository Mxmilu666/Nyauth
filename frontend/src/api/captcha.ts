import axios, { type Response } from '@/utils/axios'

export const getCaptcha = () => {
    return axios.get<
        Response<{
            id: string
            type: string
        }>
    >('/captcha')
}
