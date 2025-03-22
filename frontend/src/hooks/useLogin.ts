import { ref, watch } from 'vue'
import { VForm } from 'vuetify/lib/components/index.mjs'
import { getAccountStatus, accountLogin } from '@/api/login'
import { message } from '@/services/message'
import { useRouter, useRoute } from 'vue-router'

// 表单状态和验证相关钩子
export function useLoginForm() {
    const form = ref<InstanceType<typeof VForm>>()
    const email = ref('')
    const password = ref('')
    const otp = ref('')

    const emailRules = [
        (value: string) => !!value || '请输入电子邮箱',
        (value: string) => /.+@.+\..+/.test(value) || '请输入有效的电子邮箱'
    ]

    const validateForm = async (): Promise<boolean> => {
        if (!form.value) return false
        const { valid } = await form.value.validate()
        return valid
    }

    return {
        form,
        email,
        password,
        otp,
        emailRules,
        validateForm
    }
}

// 账户状态检查 Hook
export function useAccountCheck() {
    const isLoading = ref(false)

    const checkAccount = async (email: string) => {
        isLoading.value = true
        try {
            const { data } = await getAccountStatus({
                username: email
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

    return {
        isLoading,
        checkAccount
    }
}

// 认证流程管理 Hook
export function useAuthFlow() {
    const istologin = ref(false)
    const istoregister = ref(false)

    const setAuthMode = (exists: boolean | null) => {
        if (exists === true) {
            // 账户存在，显示登录表单
            istologin.value = true
            istoregister.value = false
        } else if (exists === false) {
            // 账户不存在，显示注册表单
            istologin.value = false
            istoregister.value = true
        }
    }

    return {
        istologin,
        istoregister,
        setAuthMode
    }
}

// 登录操作 Hook
export function useLoginOperation() {
    const isLoading = ref(false)
    const router = useRouter()
    const route = useRoute()

    const performLogin = async (
        email: string,
        password: string,
        captchaToken: string
    ): Promise<boolean> => {
        isLoading.value = true
        try {
            const { data } = await accountLogin({
                username: email,
                password: password,
                turnstile_secretkey: captchaToken
            })

            if (data && data.data) {
                // 登录成功，保存 token
                localStorage.setItem('token', data.data.token)
                localStorage.setItem('tokenExpiry', data.data.exp.toString())
                message.info('登录成功')
                if (route.query.redirect) {
                    router.push(route.query.redirect as string)
                } else {
                    router.push('/console')
                }
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

    return {
        isLoading,
        performLogin
    }
}

export function useLogin() {
    const { form, email, password, otp, emailRules, validateForm } = useLoginForm()
    const { isLoading: isCheckingAccount, checkAccount } = useAccountCheck()
    const { istologin, istoregister, setAuthMode } = useAuthFlow()
    const { isLoading: isLoggingIn, performLogin } = useLoginOperation()

    // 合并loading状态
    const isLoading = ref(false)

    const updateLoadingState = () => {
        isLoading.value = isCheckingAccount.value || isLoggingIn.value
    }

    // 监听两个loading状态的变化
    watch([isCheckingAccount, isLoggingIn], updateLoadingState)

    const login = async (captchaToken: string): Promise<boolean | void> => {
        if (!istoregister.value && !istologin.value) {
            const accountExists = await checkAccount(email.value)
            setAuthMode(accountExists)
            return
        }

        const valid = await validateForm()
        if (!valid) return false

        // 登录逻辑
        if (istologin.value) {
            return await performLogin(email.value, password.value, captchaToken)
        }

        // 注册逻辑
        if (istoregister.value) {
            // 注册逻辑待实现
            return false
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
        login
    }
}
