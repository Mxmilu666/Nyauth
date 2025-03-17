<script setup lang="ts">
import { defineOptions, ref, computed } from 'vue'

defineOptions({
    name: 'AuthPage'
})
// TODO: 把这一堆迁到 Hook 然后记得把组件拆开

// 用户身份列表
const identities = ref([
    {
        id: 1,
        userName: '米露',
        email: 'milu@milu.moe',
        avatar: 'https://placehold.co/100/6200ea/fff?text=M',
        tagText: '主账号'
    },
    {
        id: 2,
        userName: '西米',
        email: 'work@company.com',
        avatar: 'https://placehold.co/100/2962ff/fff?text=X',
        tagText: 'BList'
    },
    {
        id: 3,
        userName: 'Baka',
        email: 'admin@example.com',
        avatar: 'https://placehold.co/100/dd2c00/fff?text=B',
        tagText: 'Blog'
    }
])

// 当前选择的身份ID
const selectedIdentityId = ref(1)

// 获取当前选择的身份对象
const selectedIdentity = computed(() => {
    return (
        identities.value.find((identity) => identity.id === selectedIdentityId.value) ||
        identities.value[0]
    )
})
</script>

<template>
    <v-container
        class="auth-wrapper fill-height d-flex align-center justify-center"
        fluid
    >
        <v-card max-width="900" class="mx-auto" elevation="3">
            <v-row no-gutters>
                <v-col cols="12" sm="4" class="app-info bg-primary">
                    <div
                        class="d-flex flex-column align-center justify-center pa-6"
                        :class="{ 'fill-height': $vuetify.display.smAndUp }"
                    >
                        <v-avatar size="80" class="mb-4">
                            <v-img src="https://placehold.co/100" alt="应用图标" />
                        </v-avatar>
                        <h3 class="text-h5 mb-2 text-white text-center">BList</h3>
                        <p class="text-body-2 mb-4 text-white text-center">
                            由 Baka 创建
                        </p>
                        <p class="text-caption text-white text-center">
                            这是应用描述喵喵喵喵喵喵喵喵喵喵喵喵喵喵喵喵喵喵喵喵喵喵喵喵
                        </p>
                    </div>
                </v-col>

                <v-col cols="12" sm="8" class="pa-4 pa-sm-6">
                    <h2 class="text-h5 mb-2">授权请求</h2>
                    <p class="mb-4 text-body-1">
                        <strong>BList</strong> 请求访问您的账户
                    </p>

                    <v-divider class="mb-4" />

                    <p class="font-weight-medium mb-2">此应用将获得以下权限：</p>

                    <v-list>
                        <v-list-item density="compact">
                            <template v-slot:prepend>
                                <v-icon
                                    color="primary"
                                    size="small"
                                    icon="mdi-check-circle"
                                />
                            </template>
                            <v-list-item-title>获取基础信息</v-list-item-title>
                            <v-list-item-subtitle
                                >读取您的基本信息（uid、用户状态、昵称、头像）</v-list-item-subtitle
                            >
                        </v-list-item>

                        <v-list-item density="compact">
                            <template v-slot:prepend>
                                <v-icon
                                    color="primary"
                                    size="small"
                                    icon="mdi-check-circle"
                                />
                            </template>
                            <v-list-item-title>获取电子邮箱</v-list-item-title>
                            <v-list-item-subtitle
                                >获取您的电子邮箱地址</v-list-item-subtitle
                            >
                        </v-list-item>
                    </v-list>

                    <div class="mt-4">
                        <p class="font-weight-medium mb-2">通过以下身份继续：</p>
                        <v-select
                            v-model="selectedIdentityId"
                            :items="identities"
                            item-value="id"
                            item-title="userName"
                            :return-object="false"
                            variant="outlined"
                            class="identity-select"
                        >
                            <template v-slot:selection="{ item }">
                                <div class="d-flex align-center identity-selection">
                                    <v-avatar size="28" class="me-2 flex-shrink-0">
                                        <v-img
                                            :src="selectedIdentity.avatar"
                                            :alt="selectedIdentity.userName"
                                        />
                                    </v-avatar>
                                    <div class="d-flex flex-wrap align-center">
                                        <span class="font-weight-medium me-1">
                                            {{ selectedIdentity.userName }}
                                        </span>
                                        <v-chip
                                            size="small"
                                            color="primary"
                                            class="me-1 flex-shrink-0 my-1"
                                        >
                                            <v-icon icon="mdi-label" />
                                            {{ selectedIdentity.tagText }}
                                        </v-chip>
                                        <span
                                            class="text-caption text-grey text-truncate"
                                        >
                                            {{ selectedIdentity.email }}
                                        </span>
                                    </div>
                                </div>
                            </template>
                            <template v-slot:item="{ item, props }">
                                <v-list-item v-bind="props">
                                    <template v-slot:prepend>
                                        <v-avatar size="28" class="mr-2">
                                            <v-img
                                                :src="item.raw.avatar"
                                                :alt="item.raw.userName"
                                            />
                                        </v-avatar>
                                    </template>
                                    <v-list-item-subtitle
                                        >{{ item.raw.email }}
                                        <v-chip
                                            size="x-small"
                                            color="primary"
                                            class="ml-1"
                                            density="compact"
                                        >
                                            {{ item.raw.tagText }}
                                        </v-chip>
                                    </v-list-item-subtitle>
                                </v-list-item>
                            </template>
                        </v-select>
                    </div>

                    <v-card-actions>
                        <v-spacer />
                        <v-btn
                            variant="text"
                            color="error"
                            class="mr-2"
                            prepend-icon="mdi-close"
                            >拒绝访问</v-btn
                        >
                        <v-btn color="primary" variant="elevated" prepend-icon="mdi-check"
                            >允许访问</v-btn
                        >
                    </v-card-actions>
                </v-col>
            </v-row>
        </v-card>
    </v-container>
</template>

<style scoped>
/* 移动端 */
@media (max-width: 599px) {
    .app-info {
        border-top-left-radius: inherit;
        border-top-right-radius: inherit;
        padding-top: 1rem;
        padding-bottom: 1rem;
    }

    .v-card {
        width: 95%;
        max-width: 450px;
    }
}

/* v-card 固定宽度 */
:deep(.v-card) {
    width: 95%;
    max-width: 900px;
}
</style>
