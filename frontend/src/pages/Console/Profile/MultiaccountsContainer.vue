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
const step = ref(1)

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
        (v: string) => /.+@.+\..+$/.test(v) || '请输入有效的邮箱地址'
    ],
    code: [(v: string) => !!v || '验证码不能为空']
}

const resetForm = () => {
    formData.display_name = ''
    formData.description = ''
    formData.email = ''
    formData.code = ''
    dialog.value = false
    step.value = 1 // 重置步骤
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
    step.value = 1 // 确保从第一步开始
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
        fetchCaptcha()
    } finally {
        loading.value = false
    }
}

// 步骤控制函数
const nextStep = () => {
    if (step.value === 1) {
        if (!formData.display_name) {
            message.warning('请输入显示名称')
            return
        }
        step.value = 2
    } else if (step.value === 2) {
        if (!formData.email) {
            message.warning('请输入邮箱地址')
            return
        }
        step.value = 3
    }
}

const prevStep = () => {
    if (step.value > 1) {
        step.value--
    }
}

// 检查当前步骤表单是否有效
const isStepValid = () => {
    if (step.value === 1) {
        return !!formData.display_name
    } else if (step.value === 2) {
        return !!formData.email
    } else if (step.value === 3) {
        return !!formData.code
    }
    return true
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

    <v-dialog v-model="dialog" max-width="550">
        <v-card>
            <v-card-title class="text-h5">创建新的身份</v-card-title>
            <v-card-text>
                <v-stepper v-model="step" class="elevation-0">
                    <v-stepper-header class="elevation-0">
                        <v-stepper-item :value="1" title="基本信息" />
                        <v-divider />
                        <v-stepper-item :value="2" title="邮箱验证" />
                        <v-divider />
                        <v-stepper-item :value="3" title="确认创建" />
                    </v-stepper-header>

                    <v-stepper-window>
                        <v-stepper-window-item :value="1">
                            <v-form>
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
                            </v-form>
                        </v-stepper-window-item>

                        <v-stepper-window-item :value="2">
                            <v-form>
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
                        </v-stepper-window-item>

                        <v-stepper-window-item :value="3">
                            <v-card variant="text" class="mb-4">
                                <div
                                    class="d-flex align-center pa-4 bg-primary-lighten-5"
                                >
                                    <v-avatar color="primary" class="me-4">
                                        <v-icon color="white">mdi-check-circle</v-icon>
                                    </v-avatar>
                                    <div>
                                        <div class="text-h6">即将创建新身份</div>
                                        <div class="text-subtitle-2">
                                            请确认以下信息无误
                                        </div>
                                    </div>
                                </div>

                                <v-divider />

                                <div class="pa-4">
                                    <v-row>
                                        <v-col cols="12">
                                            <div class="d-flex align-center mb-3">
                                                <v-icon color="primary" class="me-2"
                                                    >mdi-account</v-icon
                                                >
                                                <div>
                                                    <div
                                                        class="text-caption text-medium-emphasis"
                                                    >
                                                        显示名称
                                                    </div>
                                                    <div
                                                        class="text-body-1 font-weight-medium"
                                                    >
                                                        {{ formData.display_name }}
                                                    </div>
                                                </div>
                                            </div>

                                            <div
                                                v-if="formData.description"
                                                class="d-flex align-center mb-3"
                                            >
                                                <v-icon color="primary" class="me-2"
                                                    >mdi-comment-text-outline</v-icon
                                                >
                                                <div>
                                                    <div
                                                        class="text-caption text-medium-emphasis"
                                                    >
                                                        备注
                                                    </div>
                                                    <div class="text-body-1">
                                                        {{ formData.description }}
                                                    </div>
                                                </div>
                                            </div>

                                            <div class="d-flex align-center">
                                                <v-icon color="primary" class="me-2"
                                                    >mdi-email</v-icon
                                                >
                                                <div>
                                                    <div
                                                        class="text-caption text-medium-emphasis"
                                                    >
                                                        邮箱
                                                    </div>
                                                    <div class="text-body-1">
                                                        {{ formData.email }}
                                                    </div>
                                                </div>
                                            </div>
                                        </v-col>
                                    </v-row>
                                </div>
                            </v-card>

                            <v-alert
                                color="info"
                                variant="tonal"
                                icon="mdi-information"
                                border="start"
                                density="compact"
                            >
                                创建身份后，您可以随时在多身份管理中切换不同的账户身份。
                            </v-alert>
                        </v-stepper-window-item>
                    </v-stepper-window>
                </v-stepper>
            </v-card-text>
            <v-card-actions>
                <v-btn color="grey-darken-1" variant="text" @click="resetForm">
                    取消
                </v-btn>
                <v-spacer />
                <v-btn
                    v-if="step > 1"
                    color="grey-darken-1"
                    variant="text"
                    @click="prevStep"
                    :disabled="loading"
                >
                    上一步
                </v-btn>
                <v-btn
                    v-if="step < 3"
                    color="primary"
                    variant="text"
                    @click="nextStep"
                    :disabled="!isStepValid() || loading"
                >
                    下一步
                </v-btn>
                <v-btn
                    v-if="step === 3"
                    color="primary"
                    variant="elevated"
                    @click="createIdentity"
                    :loading="loading"
                    :disabled="!formData.code"
                >
                    创建
                </v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>
