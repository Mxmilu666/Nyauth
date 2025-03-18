import { ref, computed } from 'vue'

export interface Identity {
    id: number
    userName: string
    email: string
    avatar: string
    tagText: string
}

export function useIdentities() {
    // 用户身份列表
    const identities = ref<Identity[]>([
        {
            id: 1,
            userName: '米露',
            email: 'milu@milu.moe',
            avatar: 'https://placehold.co/100/6200ea/fff?text=M',
            tagText: '主账号'
        },
        {
            id: 2,
            userName: '西米',
            email: 'work@company.com',
            avatar: 'https://placehold.co/100/2962ff/fff?text=X',
            tagText: 'BList'
        },
        {
            id: 3,
            userName: 'Baka',
            email: 'admin@example.com',
            avatar: 'https://placehold.co/100/dd2c00/fff?text=B',
            tagText: 'Blog'
        }
    ])

    // 当前选择的身份ID
    const selectedIdentityId = ref(1)

    // 获取当前选择的身份对象
    const selectedIdentity = computed(() => {
        return (
            identities.value.find(
                (identity) => identity.id === selectedIdentityId.value
            ) || identities.value[0]
        )
    })

    return {
        identities,
        selectedIdentityId,
        selectedIdentity
    }
}
