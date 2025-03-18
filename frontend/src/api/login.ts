import axios, { type Response } from '@/utils/axios'
export const getAccountStatus = (data: {
    username: string
    captcha?: {
        randstr: string
        ticket: string
    }
}) => {
    return axios.post<
        Response<{
            exists: boolean
            user_info?: {
                email: string
            }
        }>
    >('/account/getaccountstatus', data)
}
