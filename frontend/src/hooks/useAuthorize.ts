import { ref, computed } from 'vue'
import { getClientInfo, getOAuthAuthorize, type OAuthAuthorizeParams } from '@/api/oauth'
import { useRoute } from 'vue-router'
import { modal } from '@/services/modal'
import router from '@/router'

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

export interface AppInfo {
    appName: string
    appCreator: string
    appIcon: string
    appDescription: string
}

export interface Permission {
    title: string
    description: string
}

// 权限映射表
const PERMISSION_MAP: Record<string, { title: string; description: string }> = {
    'user:info': {
        title: '获取个人资料',
        description: '允许应用读取您的个人资料信息，如用户ID、用户名、头像等'
    },
    'user:email': {
        title: '获取邮箱',
        description: '允许应用读取您的邮箱地址'
    },
    'user:*': {
        title: '获取所有用户信息',
        description: '允许应用读取您的所有用户信息'
    }
}

export function useOAuthAuthorize() {
    const route = useRoute()
    const { identities, selectedIdentityId, selectedIdentity } = useIdentities()

    // 应用信息状态
    const appInfo = ref<AppInfo>({
        appName: '',
        appCreator: '',
        appIcon: '',
        appDescription: ''
    })

    // 权限列表
    const permissions = ref<Permission[]>([])

    // 加载状态
    const loading = ref(true)
    // 错误信息
    const error = ref('')
    // OAuth请求参数
    const oauthParams = ref<OAuthAuthorizeParams>({
        client_id: '',
        redirect_uri: '',
        response_type: '',
        scope: '',
        state: ''
    })

    // 授权处理中状态
    const authProcessing = ref(false)
    // 授权成功状态
    const authSuccess = ref(false)
    // 重定向URL
    const redirectUrl = ref('')

    // 初始化 OAuth 流程
    const initOAuthFlow = async () => {
        try {
            // 获取URL参数
            oauthParams.value = {
                client_id: (route.query.client_id as string) || '',
                redirect_uri: (route.query.redirect_uri as string) || '',
                response_type: (route.query.response_type as string) || '',
                scope: (route.query.scope as string) || '',
                state: (route.query.state as string) || ''
            }

            // 验证必要参数是否存在
            if (!oauthParams.value.client_id) {
                const choice = await modal.error<string>({
                    title: '笨蛋！',
                    content: '你还没有填写 client_id 参数呢！',
                    buttons: [
                        {
                            text: '回到主页',
                            color: 'primary',
                            variant: 'text',
                            value: 'back'
                        }
                    ]
                })
                if (choice === 'back') {
                    router.push({ name: 'Home' })
                }
                return
            }

            // 请求应用信息
            const { data: clientResponse } = await getClientInfo({
                client_id: oauthParams.value.client_id
            })

            if (clientResponse && clientResponse.data !== undefined) {
                // 更新应用信息
                appInfo.value = {
                    appName: clientResponse.data.client_name,
                    appCreator: clientResponse.data.created_by,
                    appIcon: clientResponse.data.avatar || 'https://placehold.co/100',
                    appDescription: clientResponse.data.description || '暂无描述'
                }

                // 更新权限列表 - 使用权限映射表
                permissions.value = clientResponse.data.permissions.map((perm) => {
                    // 从映射表中查找对应权限
                    if (PERMISSION_MAP[perm]) {
                        return PERMISSION_MAP[perm]
                    }

                    // 如果映射表中没有，则使用默认解析逻辑
                    const permParts = perm.split(':')
                    return {
                        title: `获取${permParts[1]}`,
                        description: `允许应用${permParts[1] === 'read' ? '读取' : '修改'}您的${permParts[0]}`
                    }
                })
            } else {
                modal.error({
                    title: '出错惹',
                    content: clientResponse.msg || '获取应用信息失败'
                })
            }
        } catch (err) {
            console.error('获取应用信息失败:', err)
            modal.error({
                title: '出错惹',
                content: '获取应用信息失败 qnq'
            })
        } finally {
            loading.value = false
        }
    }

    // 处理授权操作
    const handleAuthorize = async () => {
        try {
            // 设置处理中状态
            authProcessing.value = true

            // 发送授权请求
            const { data: response } = await getOAuthAuthorize(oauthParams.value)
            
            // 添加1秒延迟模拟处理过程
            await new Promise(resolve => setTimeout(resolve, 1000))
            
            if (response && response.data !== undefined) {
                // 保存重定向URL
                redirectUrl.value = response.data.redirect_url
                
                // 切换到授权成功状态
                authProcessing.value = false
                authSuccess.value = true

                // 延迟2秒后跳转
                setTimeout(() => {
                    window.location.href = redirectUrl.value
                }, 2000)
            } else {
                // 处理错误
                authProcessing.value = false
                error.value = response.msg || '授权失败'
                modal.error({
                    title: '授权失败',
                    content: error.value
                })
            }
        } catch (err) {
            console.error('授权请求失败:', err)
            authProcessing.value = false
            error.value = '授权请求失败，请稍后再试'
            modal.error({
                title: '授权失败',
                content: error.value
            })
        }
    }

    // 处理拒绝操作
    const handleReject = () => {
        // 拒绝授权，可以返回到固定页面或带错误信息跳转到redirect_uri
        if (oauthParams.value.redirect_uri) {
            // 带上error参数跳转回应用
            const redirectUrl = new URL(oauthParams.value.redirect_uri)
            redirectUrl.searchParams.append('error', 'access_denied')
            redirectUrl.searchParams.append('state', oauthParams.value.state)
            window.location.href = redirectUrl.toString()
        } else {
            // 如果没有redirect_uri，可以跳转到首页或其他页面
            window.location.href = '/'
        }
    }

    return {
        identities,
        selectedIdentityId,
        selectedIdentity,
        appInfo,
        permissions,
        loading,
        error,
        oauthParams,
        initOAuthFlow,
        handleAuthorize,
        handleReject,
        authSuccess,
        redirectUrl,
        authProcessing // 新增授权处理中状态
    }
}
