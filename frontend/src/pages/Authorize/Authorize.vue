<script setup lang="ts">
import { defineOptions, onMounted, ref } from 'vue'
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
    redirectUrl,
    authProcessing
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
            <v-progress-linear v-if="loading" color="primary" height="4" indeterminate />

            <v-fade-transition mode="out-in">
                <!-- 授权处理中状态 -->
                <div v-if="authProcessing" class="pa-6 text-center auth-processing">
                    <v-progress-circular
                        indeterminate
                        color="primary"
                        size="50"
                        class="mb-3"
                    />
                    <p class="text-h6 mb-2">正在处理您的授权请求...</p>
                </div>

                <!-- 授权成功状态 -->
                <div v-else-if="authSuccess" class="pa-6 text-center auth-success">
                    <v-scale-transition>
                        <v-icon size="64" color="success" class="mb-3"
                            >mdi-check-circle</v-icon
                        >
                    </v-scale-transition>
                    <p class="text-h6 mb-2">授权已完成，正在跳转回应用...</p>
                </div>

                <!-- 授权请求状态 -->
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
            </v-fade-transition>
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

.auth-processing,
.auth-success {
    min-height: 200px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}
</style>
