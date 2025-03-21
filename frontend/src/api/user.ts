import axios, { type Response } from '@/utils/axios'
export const getAccountInfo = () => {
    return axios.get<
        Response<{
            user_info?: {
                email: string
                is_banned: boolean
                register_at: string
                role: string
                user_avatar: string
                user_email: string
                user_id: string
                user_name: string
            }
        }>
    >('/account/info')
}
