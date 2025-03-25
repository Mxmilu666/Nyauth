<script setup lang="ts">
import { defineOptions, ref, watch } from 'vue'
import { useUserStore } from '@/stores/user'
import { message } from '@/services/message'

defineOptions({
    name: 'UserNamePage'
})

const userStore = useUserStore()

const isEditing = ref(false)
const loading = ref(false)

const userName = ref('Baka！ 你还没设置用户名呢')  // 默认值
const newUserName = ref('')

watch(
  () => userStore.userInfo.user_name,
  (newValue) => {
    if (newValue) {
      userName.value = newValue
      if (!isEditing.value) {
        newUserName.value = newValue
      }
    }
  },
  { immediate: true }
)

const startEditing = () => {
    isEditing.value = true
    newUserName.value = userName.value
}

const cancelEditing = () => {
    isEditing.value = false
    newUserName.value = userName.value
}

const saveUserName = async () => {
    if (!newUserName.value.trim()) {
        message.info('用户名不能为空')
        return
    }

    if (newUserName.value === userName.value) {
        isEditing.value = false
        return
    }

    loading.value = true
    try {
        // 这里添加实际的API

        // 模拟API调用
        await new Promise((resolve) => setTimeout(resolve, 800))

        userName.value = newUserName.value
        userStore.userInfo.user_name = newUserName.value
        message.info('用户名已成功更新')
        isEditing.value = false
    } catch (error) {
        message.info('用户名更新失败')
    } finally {
        loading.value = false
    }
}
</script>

<template>
    <v-container class="center">
        <p class="text-h5 pt-4 font-weight-bold">用户名</p>
        <p class="text-subtitle-1 py-1">查看和修改您的用户名</p>

        <div class="text-left">
            <div class="ps-1 pb-3 pt-2">
                <p class="text-h5 pt-4">您当前的用户名</p>
                <p class="text-subtitle-2 py-2 font-weight-thin">
                    这是您在 Nyauth 中显示的用户名，其他用户可以通过该名称识别您
                </p>
            </div>

            <v-card class="mb-6">
                <v-card-text>
                    <div
                        v-if="!isEditing"
                        class="d-flex align-center justify-space-between"
                    >
                        <div>
                            <div class="text-subtitle-1 font-weight-medium">用户名</div>
                            <div class="text-body-1 py-2">{{ userName }}</div>
                        </div>
                        <v-btn color="primary" variant="text" @click="startEditing">
                            修改
                        </v-btn>
                    </div>
                    <div v-else>
                        <div class="text-subtitle-1 font-weight-medium">修改用户名</div>
                        <v-form @submit.prevent="saveUserName">
                            <v-text-field
                                v-model="newUserName"
                                label="新用户名"
                                variant="outlined"
                                :rules="[(v) => !!v || '用户名不能为空']"
                                class="mt-2"
                                hide-details="auto"
                            ></v-text-field>

                            <div class="d-flex justify-end mt-4 gap-2">
                                <v-btn
                                    variant="text"
                                    @click="cancelEditing"
                                    :disabled="loading"
                                >
                                    取消
                                </v-btn>
                                <v-btn color="primary" type="submit" :loading="loading">
                                    保存
                                </v-btn>
                            </div>
                        </v-form>
                    </div>
                </v-card-text>
            </v-card>

            <div class="ps-1 pb-3 pt-2">
                <p class="text-subtitle-1 font-weight-medium">关于用户名</p>
                <v-card variant="flat" color="background" class="pa-0">
                    <v-card-text>
                        <ul class="ps-4">
                            <li class="mb-2">用户名是您在 Nyauth 上的公开标识</li>
                            <li class="mb-2">修改用户名不会影响您的登录凭据</li>
                            <li class="mb-2">
                                用户名应避免使用敏感信息，如真实姓名、电话号码等
                            </li>
                            <li class="mb-2">
                                用户名一旦被其他用户看到，就可能被记住，请谨慎选择
                            </li>
                        </ul>
                    </v-card-text>
                </v-card>
            </div>
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
