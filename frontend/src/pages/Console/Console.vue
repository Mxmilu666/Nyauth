<script setup lang="ts">
import { defineOptions, ref, provide, onMounted, computed } from 'vue'
import defineAvatar from '@/assets/logo/512x.png'
import { useUser } from '@/hooks/useUser'
import { useUserStore } from '@/stores/user'

defineOptions({
    name: 'ConsolePage'
})

const drawer = ref(true)
const { fetchUserInfo } = useUser()
const userStore = useUserStore()

// 使用默认头像，如果有用户头像则使用用户头像
const avatar = computed(() => {
    return userStore.userInfo.user_avatar || defineAvatar
})

// 提供头像给子组件使用
provide('avatar', avatar)

onMounted(() => {
    fetchUserInfo()
})
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
                :subtitle="userStore.userInfo.user_email || '未登录'"
                :title="userStore.userInfo.user_name || 'Baka'"
            />
        </v-list>

        <v-divider />
        <v-list density="compact" nav>
            <v-list-item
                prepend-icon="mdi-account-circle"
                title="首页"
                value="home"
                :to="{ name: 'HomePage' }"
                exact
            />

            <v-list-item
                prepend-icon="mdi-card-account-details-outline"
                title="个人信息"
                value="info"
                :to="{ name: 'InfoPage' }"
                exact
            />

            <v-list-item
                prepend-icon="mdi-lock-outline"
                title="数据和隐私设置"
                value="security"
                :to="{ name: 'SecurityPage' }"
                exact
            />

            <v-list-group value="Log">
                <template v-slot:activator="{ props }">
                    <v-list-item
                        v-bind="props"
                        prepend-icon="mdi-file-document-outline"
                        title="操作日志"
                    />
                </template>
                <v-list-item title="登录日志" value="Loginlog" />
                <v-list-item title="授权日志" value="Authlog" />
            </v-list-group>
        </v-list>
    </v-navigation-drawer>
    <router-view />
</template>
