import { ref } from 'vue'
import { VForm } from 'vuetify/lib/components/index.mjs'
import { getAccountStatus } from '@/api/login'
import { sleep } from '@/utils/common'

export function useLogin() {
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
            if (data !== undefined) {
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

    const login = async () => {
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
