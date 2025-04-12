import axios, { type Response } from '@/utils/axios'

export const getMultiAccountsInfo = () => {
    return axios.get<
        Response<{
            identities: Array<{
                avatar: string
                created_at: boolean
                description: string
                display_name: string
                email: string
                identity_id: string
                is_primary: boolean
                uuid: string
            }> | null
        }>
    >('/account/multi/info')
}
