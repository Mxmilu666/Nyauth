import { ref } from 'vue'
import { getAccountInfo } from '@/api/user'
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

    return {
        userInfo: userStore.userInfo,
        isLoading,
        isError,
        errorMessage,
        fetchUserInfo
    }
}
