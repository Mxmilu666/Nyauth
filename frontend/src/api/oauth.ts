import axios, { type Response } from '@/utils/axios'

export interface OAuthAuthorizeParams {
    client_id: string
    redirect_uri: string
    response_type: string
    scope: string
    state: string
}

export interface OAuthAuthorizeResponse {
    redirect_url: string
}

export const getOAuthAuthorize = (params: OAuthAuthorizeParams) => {
    return axios.get<Response<OAuthAuthorizeResponse>>('/oauth/authorize', {
        params
    })
}

export const getClientInfo = (data: { client_id: string }) => {
    return axios.post<
        Response<{
            avatar: string
            client_id: string
            client_name: string
            description : string
            create_at: string
            created_by: string
            status: number
            permissions: string[]
        }>
    >('/oauth/getclientinfo', data)
}
