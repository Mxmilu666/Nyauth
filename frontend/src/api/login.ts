import axios, { type Response } from '@/utils/axios'
export const getAccountStatus = (data: { username: string }) => {
    return axios.post<
        Response<{
            exists: boolean
            user_info?: {
                email: string
                enable_totp: boolean
            }
        }>
    >('/account/getaccountstatus', data)
}

export const accountLogin = (data: {
    username: string
    password: string
    turnstile_secretkey: string
    totp_code?: string
}) => {
    return axios.post<
        Response<{
            token: string
            exp: number
        }>
    >('/account/auth/login', data)
}

export const accountRegister = (data: {
    username: string
    useremail: string
    password: string
    code: string
    turnstile_secretkey: string
}) => {
    return axios.post<
        Response<{
            token: string
            exp: number
        }>
    >('/account/auth/register', data)
}
