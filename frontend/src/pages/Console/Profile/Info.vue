<script setup lang="ts">
import { defineOptions, inject } from 'vue'

defineOptions({
    name: 'ConsoleInfo'
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
            <v-card class="pa-1">
                <v-card-title>基本信息</v-card-title>
                <v-card-subtitle>查看并编辑您 Nyauth 中的基本信息</v-card-subtitle>

                <v-list nav>
                    <v-list-item value="avatar">
                        <v-list-item-content>
                            <v-row>
                                <v-col cols="12" md="6">
                                    <p class="text-subtitle-2">用户头像</p>
                                </v-col>
                                <v-col cols="12" md="6">
                                    <p>更换用户头像可帮助您个性化您的账号</p>
                                </v-col>
                            </v-row>
                        </v-list-item-content>
                        <template v-slot:append>
                            <v-avatar :image="avatar"></v-avatar>
                        </template>
                    </v-list-item>
                    <v-list-item value="name">
                        <v-list-item-content class="text-center">
                            <p>Mxmilu</p>
                        </v-list-item-content>

                        <template v-slot:prepend>
                            <p class="text-subtitle-2">用户名</p>
                        </template>

                        <template v-slot:append>
                            <v-avatar :image="avatar"></v-avatar>
                        </template>
                    </v-list-item>
                </v-list>
            </v-card>
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
