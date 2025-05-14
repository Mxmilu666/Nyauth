<script setup lang="ts">
import SecurityInfoItem from './SecurityInfoItem.vue'
import { modal } from '@/services/modal'

defineOptions({
    name: 'SecurityInfoCard'
})

interface Props {
    totpTitle: string
}
const props = defineProps<Props>()
</script>

<template>
    <v-card class="pa-1">
        <v-card-title>登录选项</v-card-title>
        <v-card-subtitle>查看并编辑您 Nyauth 中的登录选项</v-card-subtitle>
        <v-list nav>
            <SecurityInfoItem
                label="两步验证"
                :content="props.totpTitle"
                @click="() => $router.push({ name: 'TotpSecurityPage' })"
            />
        </v-list>
    </v-card>
</template>
