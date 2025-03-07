<script setup lang="ts">
import BasicInfoItem from './BasicInfoItem.vue'

defineOptions({
    name: 'BasicInfoCard'
})

interface Props {
    avatar: string
    username: string
    email: string
    userId: string
}
const props = defineProps<Props>()
</script>


<template>
    <v-card class="pa-1">
        <v-card-title>基本信息</v-card-title>
        <v-card-subtitle>查看并编辑您 Nyauth 中的基本信息</v-card-subtitle>
        <v-list nav>
            <BasicInfoItem
                label="用户头像"
                content="更换用户头像可帮助您个性化您的账号"
                :avatar="avatar"
            />
            <BasicInfoItem
                label="用户名"
                :content="props.username"
                @click="() => $router.push({ name: 'UserProfile' })"
            />
            <BasicInfoItem
                label="电子邮件"
                :content="props.email"
                @click="() => $router.push({ name: 'UserEmail' })"
            />
            <BasicInfoItem
                label="用户 ID"
                :content="props.userId"
                @click="() => $router.push({ name: 'UserID' })"
            />
        </v-list>
    </v-card>
</template>