import axios, { type Response } from '@/utils/axios'

export const verifyCode = (data: { useremail: string; code: string; usefor: string }) => {
    return axios.post<
        Response<{
            exp: number
            temp_code: string
        }>
    >(`/account/verifycode?usefor=${data.usefor}`, {
        useremail: data.useremail,
        code: data.code
    })
}

export const sendCode = (data: {
    useremail: string
    turnstile_secretkey: string
    usefor: string
}) => {
    return axios.post<Response<null>>(`/account/sendcode?usefor=${data.usefor}`, {
        useremail: data.useremail,
        turnstile_secretkey: data.turnstile_secretkey
    })
}
