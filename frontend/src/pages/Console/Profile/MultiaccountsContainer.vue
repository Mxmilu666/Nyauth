<script setup lang="ts">
import { defineProps } from 'vue'
import MultiaccountsCard from './MultiaccountsCard.vue'

defineProps<{
    accounts: Array<{
        avatar: string
        userName: string
        lastActiveTime: string
        tagText: string
    }>
}>()
</script>

<template>
    <v-card class="pa-1">
        <v-card-title>多身份管理</v-card-title>
        <v-card-subtitle>管理您在 Nyauth 中的多个身份</v-card-subtitle>
        <v-card-text>
            <v-row>
                <v-col
                    v-for="(account, index) in accounts"
                    :key="index"
                    cols="12"
                    sm="6"
                    md="4"
                    lg="4"
                >
                    <MultiaccountsCard
                        :avatar="account.avatar"
                        :userName="account.userName"
                        :last-active-time="account.lastActiveTime"
                        :tagText="account.tagText"
                    />
                </v-col>
                <v-col cols="12" sm="6" md="4" lg="4">
                    <v-card>
                        <v-card-text>
                            <div class="d-flex align-center">
                                <v-icon size="60" class="me-4" icon="mdi-plus" />
                                <div class="d-flex flex-column justify-center">
                                    创建新的身份
                                </div>
                            </div>
                        </v-card-text>
                    </v-card>
                </v-col>
            </v-row>
        </v-card-text>
    </v-card>
</template>
