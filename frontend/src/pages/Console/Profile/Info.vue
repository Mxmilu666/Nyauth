<script setup lang="ts">
import { defineOptions, inject } from 'vue'
import BasicInfoCard from './BasicInfoCard.vue'
import MultiaccountsContainer from './MultiaccountsContainer.vue'
import { useUserStore } from '@/stores/user'
import { useMultiAccounts } from '@/hooks/useMultiAccounts'
import Console from '../Console.vue'

defineOptions({
    name: 'InfoPage'
})

const avatar = inject('avatar') as string
const userStore = useUserStore()

// 使用 hook 获取多账户信息
const { accounts, loading, error, fetchMultiAccounts } = useMultiAccounts()

// 处理刷新多账户列表事件
const handleRefreshAccounts = () => {
    fetchMultiAccounts()
}
</script>

<template>
    <v-container class="center">
        <p class="text-h5 pt-4 font-weight-bold">个人信息</p>
        <p class="text-subtitle-1 py-1">您在各种 Nyauth 中的个人信息和偏好设置</p>
        <div class="text-left">
            <div class="ps-1 pb-3 pt-2">
                <p class="text-h5 pt-4">您在 Nyauth 中的个人资料信息</p>
                <p class="text-subtitle-2 py-2 font-weight-thin">
                    管理并修改您的个人信息，您还可以查看您的个人资料的摘要
                </p>
            </div>
            <v-row>
                <v-col cols="12">
                    <BasicInfoCard
                        :avatar="avatar"
                        :username="userStore.userInfo.user_name || 'Baka'"
                        :email="userStore.userInfo.user_email || 'Baka'"
                        :userId="userStore.userInfo.user_uuid || 'Baka'"
                    />
                </v-col>
            </v-row>
            <div class="ps-1 pb-3 pt-2">
                <p class="text-h5 pt-4">您在 Nyauth 中的其他信息和偏好设置</p>
                <p class="text-subtitle-2 py-2 font-weight-thin">用于验证您身份的方式</p>
            </div>
            <v-row>
                <v-col cols="12">
                    <v-card v-if="loading" class="pa-4">
                        <v-progress-circular indeterminate color="primary" />
                        <span class="ml-3">正在加载账户信息...</span>
                    </v-card>
                    <v-card v-else-if="error" class="pa-4">
                        <v-alert type="error" title="加载失败" :text="error.message" />
                    </v-card>
                    <v-card v-else-if="!accounts || accounts.length === 0" class="pa-4">
                        <v-alert
                            type="info"
                            title="暂无多账户"
                            text="您目前没有关联的多账户信息"
                        />
                    </v-card>
                    <MultiaccountsContainer
                        v-else
                        :accounts="accounts"
                        @refresh-accounts="handleRefreshAccounts"
                    />
                </v-col>
            </v-row>
        </div>
    </v-container>
</template>

<style scoped>
.center {
    text-align: center;
    padding-left: 150px;
    padding-right: 150px;
}

:deep(.v-card),
:deep(.v-card > *) {
    text-align: left;
}

@media (max-width: 960px) {
    .center {
        padding-left: 20px;
        padding-right: 20px;
    }
}
</style>
