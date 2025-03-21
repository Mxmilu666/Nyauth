import { ref } from 'vue'
import { VForm } from 'vuetify/lib/components/index.mjs'
import { getAccountStatus, accountLogin } from '@/api/login'
import { message } from '@/services/message'
import { useRouter } from 'vue-router'

export function useLogin() {
    const router = useRouter()
    const istologin = ref(false)
    const istoregister = ref(false)
    const isLoading = ref(false)

    const email = ref('')
    const password = ref('')
    const otp = ref('')
    const form = ref<InstanceType<typeof VForm>>()

    const emailRules = [
        (value: string) => !!value || '请输入电子邮箱',
        (value: string) => /.+@.+\..+/.test(value) || '请输入有效的电子邮箱'
    ]

    const checkAccount = async () => {
        if (!form.value) return null
        const { valid } = await form.value.validate()
        if (!valid) return null

        isLoading.value = true
        try {
            const { data } = await getAccountStatus({
                username: email.value
            })
            // 确保 data 和 data.data 都存在
            if (data && data.data !== undefined) {
                return data.data.exists
            }
            return null
        } catch (error) {
            console.error('API调用失败:', error)
            return null
        } finally {
            isLoading.value = false
        }
    }

    const login = async (captchaToken: string): Promise<boolean | void> => {
        if (!istoregister.value && !istologin.value) {
            isLoading.value = true
            const accountExists = await checkAccount()
            if (accountExists === true) {
                // 账户存在，显示登录表单
                istologin.value = true
                istoregister.value = false
            } else if (accountExists === false) {
                // 账户不存在，显示注册表单
                istologin.value = false
                istoregister.value = true
            }
            isLoading.value = false
            return
        }

        if (!form.value) return
        const { valid } = await form.value.validate()
        if (!valid) return

        // 登录逻辑
        if (istologin.value) {
            isLoading.value = true
            try {
                const { data } = await accountLogin({
                    username: email.value,
                    password: password.value,
                    turnstile_secretkey: captchaToken
                })

                if (data && data.data) {
                    // 登录成功，保存 token
                    localStorage.setItem('token', data.data.token)
                    localStorage.setItem('tokenExpiry', data.data.exp.toString())
                    message.info('登录成功')
                    router.push('/console')
                    return true
                } else {
                    message.info('登录失败，请重试')
                    return false
                }
            } catch (error: any) {
                console.error('登录失败:', error)
                return false
            } finally {
                isLoading.value = false
            }
        }

        // 注册逻辑
        if (istoregister.value) {
        }
    }

    return {
        istologin,
        istoregister,
        isLoading,
        email,
        password,
        otp,
        form,
        emailRules,
        checkAccount,
        login
    }
}
