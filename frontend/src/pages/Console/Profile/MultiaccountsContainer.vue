<script setup lang="ts">
import { defineProps, defineEmits, ref, reactive } from 'vue'
import MultiaccountsCard from './MultiaccountsCard.vue'
import { getCaptcha } from '@/api/captcha'
import { createMultiAccounts } from '@/api/multi'
import { message } from '@/services/message'
import { sendCode } from '@/api/util'
import Turnstile from '@/components/turnstile/Turnstile.vue'

const emits = defineEmits(['refresh-accounts'])

defineProps<{
    accounts: Array<{
        avatar: string
        userName: string
        lastActiveTime: string
        tagText: string
    }>
}>()

const dialog = ref(false)
const loading = ref(false)
const captchaId = ref('')
const showTurnstile = ref(false)
const sendingCode = ref(false)

const formData = reactive({
    display_name: '',
    description: '',
    email: '',
    code: ''
})

const rules = {
    display_name: [(v: string) => !!v || '显示名称不能为空'],
    email: [
        (v: string) => !!v || '邮箱不能为空',
        (v: string) => /.+@.+\..+/.test(v) || '请输入有效的邮箱地址'
    ],
    code: [(v: string) => !!v || '验证码不能为空']
}

const resetForm = () => {
    formData.display_name = ''
    formData.description = ''
    formData.email = ''
    formData.code = ''
    dialog.value = false
}

const fetchCaptcha = async () => {
    try {
        const { data } = await getCaptcha()
        if (data && data.data) {
            captchaId.value = data.data.id
        }
    } catch (error) {
        message.error('获取验证码失败')
    }
}

const openDialog = () => {
    dialog.value = true
    fetchCaptcha()
}

// 处理验证码验证成功的回调
const handleCaptchaVerify = async (token: string) => {
    showTurnstile.value = false
    await sendVerificationCode(token)
}

// 处理验证码验证失败的回调
const handleCaptchaError = (error: string) => {
    console.error('验证码错误:', error)
    message.error('人机验证失败，请重试')
}

// 请求发送验证码
const requestVerificationCode = () => {
    if (!formData.email) {
        message.warning('请先填写邮箱地址')
        return
    }

    if (!/^.+@.+\..+$/.test(formData.email)) {
        message.warning('请输入有效的邮箱地址')
        return
    }

    showTurnstile.value = true
}

// 发送验证码
const sendVerificationCode = async (token: string) => {
    sendingCode.value = true
    try {
        const { data } = await sendCode({
            useremail: formData.email,
            turnstile_secretkey: token,
            usefor: 'multi_identity'
        })

        if (data) {
            message.success('验证码已发送至您的邮箱，请查收')
        } else {
            message.error('验证码发送失败，请重试')
        }
    } catch (error) {
        console.error('发送验证码出错:', error)
        message.error('验证码发送失败，请稍后重试')
    } finally {
        sendingCode.value = false
    }
}

const createIdentity = async () => {
    if (!formData.display_name || !formData.email || !formData.code) {
        message.error('请填写必填字段')
        return
    }

    loading.value = true
    try {
        const { data } = await createMultiAccounts({
            display_name: formData.display_name,
            email: formData.email,
            description: formData.description || '',
            code: formData.code
        })

        if (data && data.data !== undefined) {
            message.success('创建新身份成功')
            resetForm()
            // 刷新账户列表
            emits('refresh-accounts')
        }
    } catch (error) {
        message.error('创建新身份失败')
        fetchCaptcha()
    } finally {
        loading.value = false
    }
}
</script>

<template>
    <v-card class="pa-1">
        <v-card-title>多身份管理</v-card-title>
        <v-card-subtitle>管理您在 Nyauth 中的多个身份</v-card-subtitle>
        <v-card-text>
            <v-row>
                <v-col
                    v-for="(account, index) in accounts"
                    :key="index"
                    cols="12"
                    sm="6"
                    md="4"
                    lg="4"
                >
                    <MultiaccountsCard
                        :avatar="account.avatar"
                        :userName="account.userName"
                        :last-active-time="account.lastActiveTime"
                        :tagText="account.tagText"
                    />
                </v-col>
                <v-col cols="12" sm="6" md="4" lg="4">
                    <v-card height="92" link @click="openDialog">
                        <v-card-text
                            class="d-flex align-center justify-center"
                            style="height: 100%"
                        >
                            <v-icon size="50" class="me-2" icon="mdi-plus" />
                            <strong>创建新的身份</strong>
                        </v-card-text>
                    </v-card>
                </v-col>
            </v-row>
        </v-card-text>
    </v-card>

    <!-- Turnstile 验证组件 -->
    <turnstile
        v-model:show="showTurnstile"
        @callback="handleCaptchaVerify"
        @error="handleCaptchaError"
    />

    <!-- 创建新身份的模态框 -->
    <v-dialog v-model="dialog" max-width="500">
        <v-card>
            <v-card-title class="text-h5">创建新的身份</v-card-title>
            <v-card-text>
                <v-form @submit.prevent="createIdentity">
                    <v-text-field
                        v-model="formData.display_name"
                        label="显示名称"
                        :rules="rules.display_name"
                        required
                        variant="outlined"
                        class="mt-3"
                        prepend-inner-icon="mdi-account"
                    />

                    <v-text-field
                        v-model="formData.description"
                        label="备注（可选）"
                        variant="outlined"
                        class="mt-3"
                        hint="用于区分不同身份的标签，如'工作账号'、'个人账号'等"
                        prepend-inner-icon="mdi-comment-text-outline"
                    />

                    <v-text-field
                        v-model="formData.email"
                        label="邮箱"
                        :rules="rules.email"
                        required
                        variant="outlined"
                        class="mt-3"
                        type="email"
                        prepend-inner-icon="mdi-email"
                    >
                        <template v-slot:append-inner>
                            <v-btn
                                color="primary"
                                variant="text"
                                @click.stop="requestVerificationCode"
                                :disabled="
                                    loading ||
                                    sendingCode ||
                                    !formData.email ||
                                    !/^.+@.+\..+$/.test(formData.email)
                                "
                                :loading="sendingCode"
                            >
                                获取验证码
                            </v-btn>
                        </template>
                    </v-text-field>

                    <v-text-field
                        v-model="formData.code"
                        label="验证码"
                        :rules="rules.code"
                        required
                        variant="outlined"
                        class="mt-3"
                        prepend-inner-icon="mdi-shield-check"
                    />
                </v-form>
            </v-card-text>
            <v-card-actions>
                <v-spacer />
                <v-btn color="grey-darken-1" variant="text" @click="resetForm">
                    取消
                </v-btn>
                <v-btn
                    color="primary"
                    variant="elevated"
                    @click="createIdentity"
                    :loading="loading"
                >
                    创建
                </v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>
