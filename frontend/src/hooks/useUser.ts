import { ref } from 'vue'
import { getAccountInfo } from '@/api/user'
import { updateUsername } from '@/api/user'
import { useUserStore } from '@/stores/user'

export function useUser() {
    const isLoading = ref(false)
    const isError = ref(false)
    const errorMessage = ref('')
    const userStore = useUserStore()

    const fetchUserInfo = async () => {
        isLoading.value = true
        isError.value = false

        try {
            const { data } = await getAccountInfo()
            if (data && data.data !== undefined) {
                if (data.data.user_info) {
                    userStore.updateUserInfo(data.data.user_info)
                } else {
                    isError.value = true
                    errorMessage.value = '用户信息未定义'
                }
            } else {
                isError.value = true
                errorMessage.value = '获取用户信息失败'
            }
        } catch (error) {
            isError.value = true
            errorMessage.value =
                error instanceof Error ? error.message : '获取用户信息失败'
        } finally {
            isLoading.value = false
        }
    }

    const updateUserName = async (newUsername: string) => {
        isLoading.value = true
        isError.value = false
        errorMessage.value = ''

        try {
            const data = await updateUsername({ username: newUsername.trim() })
            if (data && data.data !== undefined) {
                userStore.userInfo.user_name = newUsername
                return { success: true, message: '用户名已成功更新' }
            } else {
                isError.value = true
                errorMessage.value = '用户名更新失败'
                return { success: false, message: '用户名更新失败' }
            }
        } catch (error: any) {
            isError.value = true
            const message = error.response?.data?.message || '用户名更新失败'
            errorMessage.value = message
            return { success: false, message }
        } finally {
            isLoading.value = false
        }
    }
    return {
        userInfo: userStore.userInfo,
        isLoading,
        isError,
        errorMessage,
        fetchUserInfo,
        updateUserName
    }
}
