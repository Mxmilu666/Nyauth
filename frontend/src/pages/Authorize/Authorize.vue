<script setup lang="ts">
import { defineOptions, ref } from 'vue'
import { useIdentities } from '@/hooks/useAuthorize'
import AppInfo from './AppInfo.vue'
import PermissionList from './PermissionList.vue'
import IdentitySelector from './IdentitySelector.vue'

defineOptions({
    name: 'AuthPage'
})

const { identities, selectedIdentityId, selectedIdentity } = useIdentities()

const appInfo = {
    appName: 'BList',
    appCreator: 'Baka',
    appIcon: 'https://placehold.co/100',
    appDescription: '这是应用描述喵喵喵喵喵喵喵喵喵喵喵喵喵喵喵喵喵喵喵喵喵喵喵喵'
}

// 权限列表
const permissions = [
    {
        title: '获取基础信息',
        description: '读取您的基本信息（uid、用户状态、昵称、头像）'
    },
    {
        title: '获取电子邮箱',
        description: '获取您的电子邮箱地址'
    }
]

// 处理授权操作
const handleAuthorize = () => {
    console.log(`已授权，使用ID: ${selectedIdentityId.value}`)
}

// 处理拒绝操作
const handleReject = () => {
    console.log('已拒绝')
}
</script>

<template>
    <v-container
        class="auth-wrapper fill-height d-flex align-center justify-center"
        fluid
    >
        <v-card max-width="900" class="mx-auto" elevation="3">
            <v-row no-gutters>
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
