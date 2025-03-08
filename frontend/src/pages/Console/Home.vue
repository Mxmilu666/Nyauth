<script setup lang="ts">
import { defineOptions } from 'vue'
import { useConsole } from '@/hooks/useConsole'

defineOptions({
    name: 'HomePage'
})

const { avatar, timeGreeting } = useConsole()
</script>

<template>
    <v-container class="center">
        <v-avatar :image="avatar" size="110"></v-avatar>
        <p class="text-h5 pt-4 font-weight-bold">Mxmilu，{{ timeGreeting }}</p>
        <p class="text-subtitle-1 py-2">
            管理自己的信息、隐私和安全，创建应用程序、开发等更多玩法
        </p>

        <v-row class="pt-6">
            <v-col cols="12" md="6">
                <v-card class="pa-1">
                    <v-card-title>个人信息</v-card-title>
                    <v-card-text>查看并编辑您 Nyauth 中的个人信息</v-card-text>
                    <v-card-actions>
                        <v-btn
                            block
                            color="primary"
                            variant="outlined"
                            append-icon="mdi-chevron-right"
                            >查看</v-btn
                        >
                    </v-card-actions>
                    <v-icon class="card-icon" size="100" icon="mdi-account-circle" />
                </v-card>
            </v-col>
            <v-col cols="12" md="6">
                <v-card class="pa-1">
                    <v-card-title>数据和隐私设置</v-card-title>
                    <v-card-text>管理您在 Nyauth 中的数据和隐私设置</v-card-text>
                    <v-card-actions>
                        <v-btn
                            block
                            color="primary"
                            variant="outlined"
                            append-icon="mdi-chevron-right"
                            >查看</v-btn
                        >
                    </v-card-actions>
                    <v-icon class="card-icon" size="100" icon="mdi-security" />
                </v-card>
            </v-col>
        </v-row>
        <v-row>
            <v-col cols="12">
                <v-card class="pa-1">
                    <v-card-title>查看操作日志</v-card-title>
                    <v-card-text>查看您在 Nyauth 中的操作日志</v-card-text>
                    <v-card-actions>
                        <v-btn
                            block
                            color="primary"
                            variant="outlined"
                            append-icon="mdi-chevron-right"
                            >查看</v-btn
                        >
                    </v-card-actions>
                    <v-icon class="card-icon" size="100" icon="mdi-database" />
                </v-card>
            </v-col>
            <v-col cols="12">
                <v-card class="pa-1">
                    <v-card-title>想要了解其他内容?</v-card-title>
                    <v-card-text>
                        <v-list nav>
                            <v-list-item
                                prepend-icon="mdi-book-open-variant"
                                title="查询 Nyauth 文档"
                                @click="$router.push({ name: 'Docs' })"
                            />
                            <v-list-item
                                prepend-icon="mdi-email-outline"
                                title="联系我们"
                                @click="$router.push({ name: 'Contact' })"
                            />
                        </v-list>
                    </v-card-text>
                </v-card>
            </v-col>
        </v-row>
    </v-container>
</template>

<style scoped>
.center {
    margin-top: 100px;
    text-align: center;
    padding-left: 150px;
    padding-right: 150px;
}

:deep(.v-card),
:deep(.v-card > *) {
    text-align: left;
}

.center .card-icon {
    position: absolute;
    bottom: -36px;
    right: -36px;
    opacity: 0.1;
    transform: scale(2);
}

@media (max-width: 960px) {
    .center {
        padding-left: 20px;
        padding-right: 20px;
    }
}
</style>
