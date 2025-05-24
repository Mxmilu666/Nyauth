import { ref, watch } from 'vue'
import { VForm } from 'vuetify/lib/components/index.mjs'
import { getAccountStatus, accountLogin, accountRegister } from '@/api/login'
import { message } from '@/services/message'
import { useRouter, useRoute } from 'vue-router'
import { Cookie } from '@/utils/cookie'

// 表单状态和验证管理
export function useLoginForm() {
    const form = ref<InstanceType<typeof VForm>>()
    const email = ref('')
    const password = ref('')
    const otp = ref('')
    const totpCode = ref('') // 添加TOTP验证码

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
        totpCode, // 返回TOTP验证码
        emailRules,
        validateForm
    }
}

// 账户状态检查
export function useAccountCheck() {
    const isLoading = ref(false)

    const checkAccount = async (email: string) => {
        isLoading.value = true
        try {
            const { data } = await getAccountStatus({ username: email })
            return {
                exists: data?.data?.exists ?? null,
                isTotpEnabled: data?.data?.user_info?.enable_totp ?? false
            }
        } catch (error) {
            console.error('账户检查失败:', error)
            return {
                exists: null,
                isTotpEnabled: false
            }
        } finally {
            isLoading.value = false
        }
    }

    return {
        isLoading,
        checkAccount
    }
}

// 认证流程状态管理
export function useAuthFlow() {
    const istologin = ref(false)
    const istoregister = ref(false)
    const isOtpVerified = ref(false)
    const isTotpEnabled = ref(false)
    const showTotp = ref(false)
    const tempCode = ref('')

    const setAuthMode = (accountInfo: {
        exists: boolean | null
        isTotpEnabled: boolean
    }) => {
        if (accountInfo.exists === true) {
            // 账户存在，进入登录流程
            istologin.value = true
            istoregister.value = false
            isTotpEnabled.value = accountInfo.isTotpEnabled
        } else if (accountInfo.exists === false) {
            // 账户不存在，进入注册流程
            istologin.value = false
            istoregister.value = true
        }
        isOtpVerified.value = false
    }

    const completeOtpVerification = (temp_code: string) => {
        tempCode.value = temp_code
        isOtpVerified.value = true
    }

    // 用于完成TOTP验证并显示登录表单
    const completeTotpForm = () => {
        showTotp.value = true
    }

    return {
        istologin,
        istoregister,
        isOtpVerified,
        isTotpEnabled,
        tempCode,
        showTotp,
        setAuthMode,
        completeOtpVerification,
        completeTotpForm
    }
}

// 登录部分处理
export function useLoginOperation() {
    const isLoading = ref(false)
    const router = useRouter()
    const route = useRoute()
    const performLogin = async (
        email: string,
        password: string,
        captchaToken: string,
        totpCode: string = '',
        rememberMe: boolean = true
    ) => {
        isLoading.value = true
        try {
            const { data } = await accountLogin({
                username: email,
                password: password,
                turnstile_secretkey: captchaToken,
                totp_code: totpCode // 添加TOTP验证码参数
            })

            if (data?.data) {
                // 保存登录状态到cookie
                const expirationDays = rememberMe ? 30 : 1 // 保持登录30天，否则1天
                Cookie.set('token', data.data.token, expirationDays)
                Cookie.set('tokenExpiry', data.data.exp.toString(), expirationDays)

                // 如果选择了保持登录，存储额外信息
                if (rememberMe) {
                    Cookie.set('rememberMe', 'true', expirationDays)
                } else {
                    Cookie.remove('rememberMe')
                }

                const redirectPath = (route.query.redirect as string) || '/console'
                router.push(redirectPath)
                message.success('登录成功')
                return true
            } else {
                message.error('登录失败，请重试')
                return false
            }
        } catch (error) {
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

/**
 * 注册操作处理
 */
export function useRegisterOperation() {
    const isLoading = ref(false)
    const router = useRouter()
    const route = useRoute()

    const performRegister = async (
        username: string,
        email: string,
        password: string,
        temp_code: string,
        captchaToken: string
    ) => {
        isLoading.value = true
        try {
            const { data } = await accountRegister({
                username,
                useremail: email,
                password,
                code: temp_code,
                turnstile_secretkey: captchaToken
            })

            if (data?.data) {
                // 保存登录状态到cookie
                Cookie.set('token', data.data.token, 30) // 注册后默认保持登录30天
                Cookie.set('tokenExpiry', data.data.exp.toString(), 30)
                Cookie.set('rememberMe', 'true', 30)

                message.success('注册成功')

                // 重定向处理
                const redirectPath = (route.query.redirect as string) || '/console'
                router.push(redirectPath)
                return true
            }
            return false
        } catch (error: any) {
            console.error('注册失败:', error)
            const errorMessage = error.response?.data?.message || '注册失败，请稍后重试'
            message.error(errorMessage)
            return false
        } finally {
            isLoading.value = false
        }
    }

    return {
        isLoading,
        performRegister
    }
}

// 整合登录注册流程的主Hook
export function useLogin() {
    const { form, email, password, otp, totpCode, emailRules, validateForm } =
        useLoginForm()
    const { isLoading: isCheckingAccount, checkAccount } = useAccountCheck()
    const {
        istologin,
        istoregister,
        isOtpVerified,
        isTotpEnabled,
        tempCode,
        showTotp,
        setAuthMode,
        completeOtpVerification,
        completeTotpForm
    } = useAuthFlow()
    const { isLoading: isLoggingIn, performLogin } = useLoginOperation()
    const { isLoading: isRegistering, performRegister } = useRegisterOperation()

    // 合并loading状态
    const isLoading = ref(false)

    // 添加记住登录状态
    const rememberMe = ref(true)

    // 更新全局加载状态
    watch([isCheckingAccount, isLoggingIn, isRegistering], () => {
        isLoading.value =
            isCheckingAccount.value || isLoggingIn.value || isRegistering.value
    })

    // 处理登录/注册流程
    const login = async (captchaToken: string) => {
        // 判断账户是否存在
        if (!istoregister.value && !istologin.value) {
            const accountExists = await checkAccount(email.value)
            setAuthMode(accountExists)
            return
        }

        // 表单验证
        const valid = await validateForm()
        if (!valid) return false

        // 启用 TOTP 且需要验证
        if (istologin.value && isTotpEnabled.value && !showTotp.value) {
            // 设置状态为需要输入TOTP
            completeTotpForm()
            return false
        }

        // 登录处理(如果启用了TOTP，则需要提供TOTP验证码)
        if (istologin.value) {
            return await performLogin(
                email.value,
                password.value,
                captchaToken,
                isTotpEnabled.value ? totpCode.value : '',
                rememberMe.value
            )
        }

        // 注册处理 验证码已验证后
        if (istoregister.value && isOtpVerified.value) {
            const username = email.value.split('@')[0]
            return await performRegister(
                username,
                email.value,
                password.value,
                tempCode.value,
                captchaToken
            )
        }
    }

    // 处理TOTP输入
    const handleTotpInput = (code: string) => {
        totpCode.value = code
    }

    return {
        istologin,
        istoregister,
        isOtpVerified,
        isTotpEnabled,
        showTotp,
        isLoading,
        email,
        password,
        otp,
        totpCode,
        form,
        emailRules,
        rememberMe,
        login,
        handleTotpInput,
        completeOtpVerification
    }
}
