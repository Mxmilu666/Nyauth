import { ref, onMounted } from 'vue'
import { getMultiAccountsInfo } from '@/api/multi'

export interface MultiAccount {
    avatar: string
    userName: string
    lastActiveTime: string
    tagText: string
    userId?: string
}

export function useMultiAccounts() {
    const accounts = ref<MultiAccount[]>([])
    const loading = ref(false)
    const error = ref<Error | null>(null)

    const fetchMultiAccounts = async () => {
        loading.value = true
        error.value = null
        try {
            const { data } = await getMultiAccountsInfo()
            if (data && data.data !== undefined && data.data.identities) {
                accounts.value = data.data.identities.map((identity, index) => {
                    let tagText = '账号'
                    if (identity.is_primary) {
                        tagText = '主账号'
                    } else if (identity.description) {
                        tagText = identity.description
                    }
                    
                    const lastActiveTime = '我不知道'

                    return {
                        avatar: identity.avatar,
                        userName: identity.display_name,
                        lastActiveTime,
                        tagText,
                        userId: identity.uuid
                    }
                })
            } else {
                error.value = new Error('获取多账户信息失败')
            }
        } catch (e) {
            error.value = e instanceof Error ? e : new Error('未知错误')
        } finally {
            loading.value = false
        }
    }

    onMounted(fetchMultiAccounts)

    return {
        accounts,
        loading,
        error,
        fetchMultiAccounts
    }
}
