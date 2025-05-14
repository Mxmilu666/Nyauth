<script setup lang="ts">
import SecurityInfoCard from './SecurityInfoCard.vue'

defineOptions({
    name: 'securityPage'
})
</script>

<template>
    <v-container class="center">
        <p class="text-h5 pt-4 font-weight-bold">数据和隐私设置</p>
        <p class="text-subtitle-1 py-1">您在各种 Nyauth 中的数据和隐私设置</p>
        <div class="text-left">
            <div class="ps-1 pb-3 pt-2">
                <p class="text-h5 pt-4">您的 Nyauth 账号登录选项</p>
                <p class="text-subtitle-2 py-2 font-weight-thin">
                    请务必及时更新这些信息，确保您始终都能访问自己的 Nyauth 账号
                </p>
            </div>
            <v-row>
                <v-col cols="12">
                    <SecurityInfoCard :totpTitle="'启用时间：2022-22-22'" />
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
