import axios, { type Response } from '@/utils/axios'

export const generateTOTP = () => {
    return axios.get<
        Response<{
            account: string
            exp_time: number
            issuer: string
            qr_code: string
            secret: string
        }>
    >('/account/totp/generate')
}

export const firstVerifyTOTP = (data: { code: string }) => {
    return axios.post<
        Response<{
            recovery_codes: string[]
        }>
    >('/account/totp/verify', data)
}
