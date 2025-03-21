import axios, { type Response } from '@/utils/axios'
export const getAccountStatus = (data: { username: string }) => {
    return axios.post<
        Response<{
            exists: boolean
            user_info?: {
                email: string
            }
        }>
    >('/account/getaccountstatus', data)
}

export const accountLogin = (data: {
    username: string
    password: string
    turnstile_secretkey: string
}) => {
    return axios.post<
        Response<{
            token: string
            exp: number
        }>
    >('/account/auth/login', data)
}
