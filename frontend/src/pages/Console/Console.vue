<script setup lang="ts">
import { defineOptions, ref, provide } from 'vue'
import defineAvatar from '@/assets/logo/512x.png'

defineOptions({
    name: 'ConsolePage'
})

const drawer = ref(true)

const avatar = defineAvatar

provide('avatar', avatar)
</script>

<template>
    <v-navigation-drawer
        v-model="drawer"
        permanent
        expand-on-hover
        :rail="$vuetify.display.mobile"
    >
        <v-list>
            <v-list-item
                :prepend-avatar="avatar"
                subtitle="milu@milu.ink"
                title="Mxmilu"
            />
        </v-list>

        <v-divider></v-divider>
        <v-divider></v-divider>
        <v-list density="compact" nav>
            <v-list-item
                prepend-icon="mdi-account-circle"
                title="首页"
                value="home"
                :to="{ name: 'ConsoleHome' }"
            ></v-list-item>

            <v-list-item
                prepend-icon="mdi-card-account-details-outline"
                title="个人信息"
                value="info"
            ></v-list-item>

            <v-list-item
                prepend-icon="mdi-lock-outline"
                title="数据和隐私设置"
                value="security"
            ></v-list-item>

            <v-list-group value="Log">
                <template v-slot:activator="{ props }">
                    <v-list-item
                        v-bind="props"
                        prepend-icon="mdi-file-document-outline"
                        title="操作日志"
                    ></v-list-item>
                </template>
                <v-list-item title="登录日志" value="Loginlog"></v-list-item>
                <v-list-item title="授权日志" value="Authlog"></v-list-item>
            </v-list-group>
        </v-list>
    </v-navigation-drawer>
    <router-view />
</template>
