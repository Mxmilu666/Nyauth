<script setup lang="ts">
import { defineOptions } from 'vue'
import { ref, watch } from 'vue'

import loginform from './Loginform.vue'
import otpform from './Otpform.vue'
import TotpVerifyForm from './TotpVerifyForm.vue'
import passwordSetForm from './PasswordSetForm.vue'
import { useLogin } from '@/hooks/useLogin'
import turnstile from '@/components/turnstile/Turnstile.vue'
import { message } from '@/services/message'
import { verifyCode, sendCode } from '@/api/util'

defineOptions({
    name: 'AuthPage'
})

const showTurnstile = ref(false)
const captchaToken = ref('')
const otpVerifying = ref(false)
const currentOtp = ref('')
const sendingOtp = ref(false)
const turnstilePurpose = ref<'login' | 'register' | 'sendOtp'>('login')
const turnstileVerify = ref(false)

const {
    istologin,
    istoregister,
    isOtpVerified,
    isLoading,
    isTotpEnabled,
    showTotp,
    email,
    password,
    otp,
    totpCode,
    form,
    emailRules,
    rememberMe,
    login,
    handleTotpInput,
    completeOtpVerification
} = useLogin()

// 进入注册流程时发送 OTP
watch(istoregister, (isRegistering) => {
    if (isRegistering && !isOtpVerified.value && !currentOtp.value) {
        turnstilePurpose.value = 'sendOtp'
        showTurnstile.value = true
    }
})

// 验证码回调处理
const handleCaptchaVerify = async (token: string) => {
    captchaToken.value = token
    turnstileVerify.value = true
    showTurnstile.value = false

    switch (turnstilePurpose.value) {
        case 'sendOtp':
            sendUserOtp(token)
            break
        case 'login':
        case 'register':
            const loginResult = await login(token)
            if (loginResult === false) {
                showTotp.value = false
            }
            break
    }
}

// 发送OTP验证码
const sendUserOtp = async (token: string) => {
    if (!email.value) {
        message.warning('请输入有效的邮箱地址')
        return
    }

    sendingOtp.value = true
    try {
        const { data } = await sendCode({
            useremail: email.value,
            turnstile_secretkey: token,
            usefor: 'register'
        })

        if (data) {
            message.success('验证码已发送至您的邮箱，请查收')
        } else {
            message.error('验证码发送失败，请重试')
        }
    } catch (error) {
        console.error('发送OTP出错:', error)
    } finally {
        sendingOtp.value = false
    }
}

const handleCaptchaError = (error: string) => {
    console.error('验证码错误:', error)
}

// 统一处理认证流程
const handleAuthentication = async () => {
    // 检查账户是否存在
    if (!istoregister.value && !istologin.value) {
        turnstilePurpose.value = 'login'
        await login(captchaToken.value)
        return
    }

    // 验证表单并请求登录
    if (istologin.value) {
        if (!form.value) return
        const { valid } = await form.value.validate()
        if (!valid) return

        if (isTotpEnabled.value && !showTotp.value) {
            // 如果启用了 TOTP 且未显示 TOTP 输入框
            showTotp.value = true
            return
        }

        // 如果启用了 TOTP 且已经显示 TOTP 输入框
        if (isTotpEnabled.value && showTotp.value) {
            if (!totpCode.value || totpCode.value.length < 6) {
                message.warning('请输入完整的两步验证码')
                return
            }
            // 验证码已输入，继续验证
            turnstilePurpose.value = 'login'
            showTurnstile.value = true
            return
        }
        return
    }

    // 验证验证码
    if (istoregister.value && !isOtpVerified.value) {
        if (!currentOtp.value || currentOtp.value.length < 6) {
            message.warning('请输入完整的验证码')
            return
        }

        otpVerifying.value = true
        try {
            const { data } = await verifyCode({
                useremail: email.value,
                code: currentOtp.value,
                usefor: 'register'
            })

            if (data?.data) {
                completeOtpVerification(data.data.temp_code)
            } else {
                message.error('验证码验证失败，请检查后重试')
            }
        } catch (error) {
            console.error('OTP验证出错:', error)
            message.error('验证码验证过程中出现错误')
        } finally {
            otpVerifying.value = false
        }
        return
    }

    // 提交注册
    if (istoregister.value && isOtpVerified.value) {
        if (!form.value) return
        const { valid } = await form.value.validate()
        if (!valid) return

        turnstilePurpose.value = 'register'
        showTurnstile.value = true
        return
    }
}

// 处理OTP输入
const handleOtpInput = (otpCode: string) => {
    currentOtp.value = otpCode
    if (otpCode.length === 6) {
        handleAuthentication()
    }
}

// 处理TOTP输入
const handleTotpVerify = (code: string) => {
    handleTotpInput(code)
    if (code.length === 6) {
        handleAuthentication()
    }
}
</script>

<template>
    <v-container
        class="auth-wrapper fill-height d-flex align-center justify-center"
        fluid
    >
        <turnstile
            v-model:show="showTurnstile"
            @callback="handleCaptchaVerify"
            @error="handleCaptchaError"
        />

        <v-row align="center" justify="center">
            <v-col cols="12" sm="8" md="4">
                <v-card
                    :disabled="isLoading || otpVerifying"
                    :loading="isLoading || otpVerifying"
                >
                    <template v-slot:loader="{ isActive }">
                        <v-progress-linear
                            :active="isActive"
                            color="primary"
                            indeterminate
                        />
                    </template>
                    <v-card-title class="text-center">
                        <div class="py-5">
                            <v-lazy>
                                <img
                                    src="@/assets/sticker/yuzu_serious.png"
                                    class="logo"
                                />
                            </v-lazy>
                            <p v-if="!istoregister" class="text-h5">
                                登录到 <strong>Nyauth</strong>
                            </p>
                            <p v-else-if="istoregister && !isOtpVerified" class="text-h5">
                                注册到 <strong>Nyauth</strong>
                            </p>
                            <p v-else class="text-h5">完成注册 <strong>Nyauth</strong></p>
                        </div>
                    </v-card-title>
                    <v-card-text>
                        <v-form ref="form" class="px-4" @submit.prevent>
                            <v-text-field
                                v-if="!(istoregister && turnstileVerify) && !showTotp"
                                v-model="email"
                                label="电子邮箱"
                                color="primary"
                                prepend-inner-icon="mdi-email-outline"
                                variant="outlined"
                                :rules="emailRules"
                                @keyup.enter="handleAuthentication"
                                required
                            />
                            <v-slide-y-transition :leave-absolute="true">
                                <loginform
                                    v-if="istologin && !isTotpEnabled"
                                    v-model="password"
                                    @enter="handleAuthentication"
                                />
                                <loginform
                                    v-if="istologin && isTotpEnabled && !showTotp"
                                    v-model="password"
                                    @enter="handleAuthentication"
                                />
                                <TotpVerifyForm
                                    v-if="istologin && isTotpEnabled && showTotp"
                                    :loading="isLoading"
                                    @totpEnter="handleTotpVerify"
                                />
                                <otpform
                                    v-if="istoregister && !isOtpVerified && turnstileVerify"
                                    :email="email"
                                    :otp="otp"
                                    :loading="otpVerifying || sendingOtp"
                                    @otpEnter="handleOtpInput"
                                />
                                <passwordSetForm
                                    v-if="istoregister && isOtpVerified"
                                    v-model:password="password"
                                    @enter="handleAuthentication"
                                />
                            </v-slide-y-transition>
                        </v-form>
                    </v-card-text>
                    <v-card-actions
                        class="px-8 pb-6 d-flex flex-column align-items-center"
                    >
                        <v-btn
                            block
                            append-icon="mdi-chevron-right"
                            color="primary"
                            variant="flat"
                            @click="handleAuthentication"
                            :loading="otpVerifying"
                            :disabled="otpVerifying"
                        >
                            {{
                                istoregister && !isOtpVerified
                                    ? '验证'
                                    : istoregister && isOtpVerified
                                      ? '注册'
                                      : '继续'
                            }}
                        </v-btn>
                        <div class="d-flex justify-space-between w-100 mt-1">
                            <v-btn
                                color="primary"
                                variant="text"
                                @click="$router.push({ name: 'ResetPassword' })"
                                >忘记密码</v-btn
                            >
                            <div v-if="istologin" class="d-flex align-center">
                                <v-checkbox-btn
                                    v-model="rememberMe"
                                    density="comfortable"
                                    color="primary"
                                />
                                <div class="text-subtitle-2 text-primary">保持登录</div>
                            </div>
                        </div>
                        <p
                            v-if="!istologin && !istoregister"
                            class="text-center text-subtitle-2"
                        >
                            未注册用户输入邮箱将自动注册
                        </p>
                    </v-card-actions>
                </v-card>
                <div class="text-center mt-8">
                    <v-btn
                        v-if="!istoregister"
                        color="primary"
                        prepend-icon="mdi-fingerprint"
                        variant="text"
                    >
                        使用外部验证器登录
                    </v-btn>
                </div>
            </v-col>
        </v-row>
    </v-container>
</template>
