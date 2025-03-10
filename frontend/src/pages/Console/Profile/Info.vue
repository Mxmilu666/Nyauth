<script setup lang="ts">
import { defineOptions, inject } from 'vue'
import BasicInfoCard from './BasicInfoCard.vue'
import MultiaccountsCard from './MultiaccountsCard.vue'

defineOptions({
    name: 'InfoPage'
})

const avatar = inject('avatar') as string
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
                        username="米露"
                        email="milu@milu.moe"
                        userId="1145141919810"
                    />
                </v-col>
            </v-row>
            <div class="ps-1 pb-3 pt-2">
                <p class="text-h5 pt-4">Nyauth 中的其他信息和偏好设置</p>
                <p class="text-subtitle-2 py-2 font-weight-thin">用于验证您身份的方式</p>
            </div>
            <v-row>
                <v-col cols="12">
                    <v-card class="pa-1">
                        <v-card-title>多身份管理</v-card-title>
                        <v-card-subtitle>管理您在 Nyauth 中的多个身份</v-card-subtitle>
                        <v-card-text>
                            <v-row>
                                <v-col cols="12" sm="6" md="4" lg="2">
                                    <MultiaccountsCard :avatar="avatar" userName="米露" />
                                </v-col>
                                <v-col cols="12" sm="6" md="4" lg="2">
                                    <MultiaccountsCard :avatar="avatar" userName="米露" />
                                </v-col>
                            </v-row>
                        </v-card-text>
                    </v-card>
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
