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

export const createMultiAccounts = (data: {
    display_name: string
    email: string
    description?: string
    code: string
}) => {
    return axios.post<
        Response<{
            identity_id: string
        }>
    >('/account/multi/create', data)
}
