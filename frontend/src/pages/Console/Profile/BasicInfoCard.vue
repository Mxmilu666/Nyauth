<script setup lang="ts">
import BasicInfoItem from './BasicInfoItem.vue'
import { modal } from '@/services/modal'

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

const changeAvatar = () => {
    modal.show({
        title: '请前往 Cravatar 更换头像',
        content: 'Nyauth 使用 Cravatar 作为头像提供商，您可前往 Cravatar 更换头像'
    })
}
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
                @click="changeAvatar"
            />
            <BasicInfoItem
                label="用户名"
                :content="props.username"
                @click="() => $router.push({ name: 'UsernameInfoPage' })"
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
