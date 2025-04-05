<script setup lang="ts">
import { defineOptions, onMounted } from 'vue'
import { useOAuthAuthorize } from '@/hooks/useAuthorize'
import AppInfo from './AppInfo.vue'
import PermissionList from './PermissionList.vue'
import IdentitySelector from './IdentitySelector.vue'

defineOptions({
    name: 'AuthPage'
})

const {
    appInfo,
    permissions,
    loading,
    error,
    identities,
    selectedIdentityId,
    selectedIdentity,
    initOAuthFlow,
    handleAuthorize,
    handleReject,
    authSuccess,
    redirectUrl
} = useOAuthAuthorize()

// 从URL获取参数并请求数据
onMounted(async () => {
    await initOAuthFlow()
})
</script>

<template>
    <v-container
        class="auth-wrapper fill-height d-flex align-center justify-center"
        fluid
    >
        <v-card max-width="900" class="mx-auto" elevation="3">
            <!-- 显示加载或错误状态 -->
            <div v-if="loading" class="pa-6 text-center">
                <v-progress-circular indeterminate color="primary" />
                <p class="mt-4">正在加载应用信息...</p>
            </div>

            <div v-else-if="error" class="pa-6 text-center">
                <v-alert type="error" title="出错了">
                    {{ error }}
                </v-alert>
                <v-btn class="mt-4" color="primary" to="/"> 返回首页 </v-btn>
            </div>

            <!-- 授权成功状态 -->
            <div v-else-if="authSuccess" class="pa-6 text-center">
                <v-icon size="64" color="success" class="mb-3">mdi-check-circle</v-icon>
                <p class="text-h6 mb-2">授权已完成，正在跳转回应用...</p>
            </div>

            <v-row v-else no-gutters>
                <v-col cols="12" sm="4" class="app-info bg-primary">
                    <AppInfo v-bind="appInfo" />
                </v-col>

                <v-col cols="12" sm="8" class="pa-4 pa-sm-6">
                    <h2 class="text-h5 mb-2">授权请求</h2>
                    <p class="mb-4 text-body-1">
                        <strong>{{ appInfo.appName }}</strong> 请求访问您的账户
                    </p>

                    <v-divider class="mb-4" />

                    <PermissionList :permissions="permissions" />

                    <div class="mt-4">
                        <IdentitySelector
                            :identities="identities"
                            v-model:selectedIdentityId="selectedIdentityId"
                            :selectedIdentity="selectedIdentity"
                        />
                    </div>

                    <v-card-actions>
                        <v-spacer />
                        <v-btn
                            variant="text"
                            color="error"
                            class="mr-2"
                            prepend-icon="mdi-close"
                            @click="handleReject"
                        >
                            拒绝访问
                        </v-btn>
                        <v-btn
                            color="primary"
                            variant="elevated"
                            prepend-icon="mdi-check"
                            @click="handleAuthorize"
                        >
                            允许访问
                        </v-btn>
                    </v-card-actions>
                </v-col>
            </v-row>
        </v-card>
    </v-container>
</template>

<style scoped>
/* 移动端 */
@media (max-width: 599px) {
    .app-info {
        border-top-left-radius: inherit;
        border-top-right-radius: inherit;
        padding-top: 1rem;
        padding-bottom: 1rem;
    }

    .v-card {
        width: 95%;
        max-width: 450px;
    }
}

/* v-card 固定宽度 */
:deep(.v-card) {
    width: 95%;
    max-width: 900px;
}
</style>
