<script setup lang="ts">
import { defineOptions, ref, computed, watch } from 'vue'
import { generateTOTP, firstVerifyTOTP } from '@/api/totp'
import { getAccountStatus } from '@/api/login'
import { useUserStore } from '@/stores/user'
import { message } from '@/services/message'
import QRCode from 'qrcode'

defineOptions({
    name: 'TotpPage'
})

const userStore = useUserStore()
const isUserReady = computed(() => !!userStore.userInfo.user_name)

// 状态变量
const loading = ref(false)
const totpEnabled = ref(false)
const setupStep = ref(0) // 0: 初始状态, 1: 显示二维码, 2: 输入验证码, 3: 显示恢复码
const qrCodeImageUrl = ref('')

// TOTP 数据
const qrCode = ref('')
const secret = ref('')
const verificationCode = ref('')
const recoveryCodes = ref<string[]>([])

// 生成二维码图像
const generateQRCodeImage = async () => {
    if (!qrCode.value) return

    try {
        // 将文本内容转换为二维码图像数据URL
        const url = await QRCode.toDataURL(qrCode.value, {
            width: 200,
            margin: 2,
            errorCorrectionLevel: 'M'
        })
        qrCodeImageUrl.value = url
    } catch (error) {
        console.error('生成二维码失败:', error)
        message.error('二维码生成失败')
    }
}

// 处理错误的辅助函数
const handleError = (error: any) => {
    console.error('TOTP操作失败:', error)
    const errorMessage = error?.response?.data?.msg || '操作失败，请稍后重试'
    message.error(errorMessage)
    loading.value = false
}

// 获取两步验证状态
const checkTotpStatus = async () => {
    // 如果用户信息不可用，则不执行
    if (!isUserReady.value) return

    loading.value = true
    try {
        console.log('使用用户名:', userStore.userInfo.user_name)
        const { data } = await getAccountStatus({
            username: userStore.userInfo.user_name
        })
        totpEnabled.value = data?.data?.user_info?.enable_totp || false
    } catch (error) {
        handleError(error)
    } finally {
        loading.value = false
    }
}

// 开始设置TOTP
const startSetupTotp = async () => {
    loading.value = true
    setupStep.value = 1

    try {
        const { data } = await generateTOTP()
        if (data?.data) {
            qrCode.value = data.data.qr_code
            secret.value = data.data.secret
            await generateQRCodeImage()
        } else {
            throw new Error('无法获取TOTP数据')
        }
    } catch (error) {
        handleError(error)
        setupStep.value = 0
    } finally {
        loading.value = false
    }
}

// 验证TOTP
const verifyTotp = async () => {
    if (!verificationCode.value || verificationCode.value.length !== 6) {
        message.warning('请输入6位验证码')
        return
    }

    loading.value = true
    try {
        const { data } = await firstVerifyTOTP({ code: verificationCode.value })
        if (data?.data) {
            recoveryCodes.value = data.data.recovery_codes
            totpEnabled.value = true
            setupStep.value = 3
            message.success('两步验证已成功启用')
        } else {
            throw new Error('验证失败')
        }
    } catch (error) {
        handleError(error)
    } finally {
        loading.value = false
    }
}

// 禁用TOTP
const disableTotp = async () => {
    message.info('这是禁用TOTP的功能')
}

// 复制恢复码到剪贴板
const copyRecoveryCodes = () => {
    if (recoveryCodes.value.length) {
        const text = recoveryCodes.value.join('\n')
        navigator.clipboard
            .writeText(text)
            .then(() => message.success('恢复码已复制到剪贴板'))
            .catch(() => message.error('复制失败，请手动复制'))
    }
}

// 重置设置流程
const resetSetup = () => {
    setupStep.value = 0
    verificationCode.value = ''
    qrCode.value = ''
    secret.value = ''
    recoveryCodes.value = []
}

const copyToClipboard = (text: string) => {
    navigator.clipboard
        .writeText(text)
        .then(() => message.success('已复制到剪贴板'))
        .catch(() => message.error('复制失败，请手动复制'))
}

watch(
    () => isUserReady.value,
    (ready) => {
        if (ready) {
            checkTotpStatus()
        }
    },
    { immediate: true } // 组件创建时立即执行一次
)
</script>

<template>
    <div class="totp">
        <v-container class="px-md-16 px-4">
            <div class="text-center">
                <p class="text-h5 pt-4 font-weight-bold">两步验证</p>
                <p class="text-subtitle-1 py-1">增强您的账户安全性，启用两步验证</p>
            </div>

            <div class="mt-4">
                <v-card v-if="loading && setupStep === 0" class="my-4">
                    <v-card-text class="text-center">
                        <v-progress-circular indeterminate color="primary" />
                        <p class="mt-2">正在加载...</p>
                    </v-card-text>
                </v-card>

                <v-card v-if="!loading || setupStep > 0" class="my-4">
                    <v-card-text>
                        <div
                            v-if="setupStep === 0"
                            class="d-flex align-center justify-space-between"
                        >
                            <div>
                                <div class="text-subtitle-1 font-weight-medium">
                                    两步验证
                                </div>
                                <div class="d-flex align-center py-2">
                                    <v-icon
                                        :color="totpEnabled ? 'success' : 'error'"
                                        class="me-2"
                                    >
                                        {{
                                            totpEnabled
                                                ? 'mdi-shield-check'
                                                : 'mdi-shield-off-outline'
                                        }}
                                    </v-icon>
                                    <span>{{ totpEnabled ? '已启用' : '未启用' }}</span>
                                </div>
                            </div>
                            <v-btn
                                :color="totpEnabled ? 'error' : 'primary'"
                                @click="totpEnabled ? disableTotp() : startSetupTotp()"
                                :loading="loading"
                            >
                                {{ totpEnabled ? '禁用' : '启用' }}
                            </v-btn>
                        </div>

                        <div v-if="setupStep === 1">
                            <h3 class="text-h6 mb-4">第1步：扫描二维码</h3>

                            <div class="text-center mb-4">
                                <div class="d-flex justify-center mb-2">
                                    <v-img
                                        :src="qrCodeImageUrl"
                                        alt="二维码"
                                        max-width="200"
                                        max-height="200"
                                        v-if="qrCode"
                                        class="border rounded"
                                    ></v-img>
                                    <v-skeleton-loader
                                        v-else
                                        type="image"
                                        height="200"
                                        width="200"
                                    ></v-skeleton-loader>
                                </div>

                                <p class="text-body-2 mt-2">
                                    使用 Google Authenticator、Microsoft Authenticator
                                    等验证器应用扫描上方二维码
                                </p>

                                <v-card class="mt-4 pa-3" elevation="0" variant="tonal">
                                    <div class="d-flex align-center flex-nowrap mb-1">
                                        <span class="text-subtitle-2 text-no-wrap"
                                            >密钥：</span
                                        >
                                        <div
                                            class="d-flex align-center flex-grow-1 overflow-auto"
                                        >
                                            <span
                                                class="text-body-1 font-weight-medium mx-2"
                                                >{{ secret }}</span
                                            >
                                            <v-btn
                                                density="compact"
                                                icon="mdi-content-copy"
                                                size="small"
                                                variant="text"
                                                class="flex-shrink-0"
                                                @click="copyToClipboard(secret)"
                                            ></v-btn>
                                        </div>
                                    </div>
                                    <p class="text-caption">
                                        如果无法扫描二维码，请手动输入上面的密钥
                                    </p>
                                </v-card>
                            </div>

                            <v-btn block color="primary" @click="setupStep = 2">
                                继续
                            </v-btn>
                        </div>

                        <div v-if="setupStep === 2">
                            <h3 class="text-h6 mb-4">第2步：输入验证码</h3>

                            <p class="mb-4">
                                请打开您的验证器应用，获取6位数验证码并输入以下框中：
                            </p>

                            <v-form @submit.prevent="verifyTotp">
                                <v-text-field
                                    v-model="verificationCode"
                                    label="验证码"
                                    :rules="[
                                        (v) => !!v || '请输入验证码',
                                        (v) => v.length === 6 || '验证码应为6位数'
                                    ]"
                                    type="text"
                                    inputmode="numeric"
                                    maxlength="6"
                                    class="mb-4"
                                    :disabled="loading"
                                    autofocus
                                ></v-text-field>

                                <div class="d-flex gap-3">
                                    <v-btn
                                        variant="outlined"
                                        @click="setupStep = 1"
                                        :disabled="loading"
                                    >
                                        返回
                                    </v-btn>
                                    <v-btn
                                        color="primary"
                                        type="submit"
                                        :loading="loading"
                                    >
                                        验证
                                    </v-btn>
                                </div>
                            </v-form>
                        </div>

                        <div v-if="setupStep === 3">
                            <h3 class="text-h6 mb-4">两步验证已启用</h3>

                            <v-alert
                                type="warning"
                                title="保存您的恢复码"
                                text="请保存以下恢复码，每个恢复码只能使用一次。如果您无法访问您的验证器应用，可以使用这些恢复码登录。"
                                variant="tonal"
                                class="mb-4"
                                icon="mdi-alert-circle"
                            ></v-alert>

                            <v-sheet
                                class="pa-4 mb-4 rounded"
                                color="grey-lighten-4"
                                elevation="0"
                                border
                            >
                                <div class="d-flex flex-wrap justify-center gap-3">
                                    <v-chip
                                        v-for="(code, index) in recoveryCodes"
                                        :key="index"
                                        variant="flat"
                                        color="grey-lighten-3"
                                        class="font-weight-medium font-monospace"
                                    >
                                        {{ code }}
                                    </v-chip>
                                </div>
                            </v-sheet>

                            <div class="d-flex justify-space-between">
                                <v-btn
                                    prepend-icon="mdi-content-copy"
                                    @click="copyRecoveryCodes"
                                >
                                    复制恢复码
                                </v-btn>
                                <v-btn color="primary" @click="resetSetup"> 完成 </v-btn>
                            </div>
                        </div>
                    </v-card-text>
                </v-card>

                <div class="pt-2 pb-3" v-if="setupStep === 0">
                    <p class="text-subtitle-1 font-weight-medium">关于两步验证</p>
                    <v-sheet class="pa-1 rounded-lg mt-4">
                        <v-list>
                            <v-list-item prepend-icon="mdi-check-circle">
                                两步验证会在您登录时额外要求输入验证器应用生成的验证码
                            </v-list-item>
                            <v-list-item prepend-icon="mdi-check-circle">
                                即使密码泄露，未经授权的用户也无法访问您的账户
                            </v-list-item>
                            <v-list-item prepend-icon="mdi-check-circle">
                                您需要下载验证器应用（如Google Authenticator、Microsoft
                                Authenticator）
                            </v-list-item>
                            <v-list-item prepend-icon="mdi-check-circle">
                                启用两步验证后，我们会提供恢复码，请妥善保管
                            </v-list-item>
                        </v-list>
                    </v-sheet>
                </div>
            </div>
        </v-container>
    </div>
</template>
