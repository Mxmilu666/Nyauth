<script setup lang="ts">
import { defineOptions } from 'vue'
import { ref } from 'vue'

import loginform from './Loginform.vue'
import otpform from './Otpform.vue'
import passwordSetForm from './PasswordSetForm.vue'
import { useLogin } from '@/hooks/useLogin'
import turnstile from '@/components/turnstile/Turnstile.vue'
import { message } from '@/services/message'

defineOptions({
    name: 'AuthPage'
})

const showTurnstile = ref(false)
const captchaToken = ref('')
const otpVerifying = ref(false)
const currentOtp = ref('')

const {
    istologin,
    istoregister,
    isOtpVerified,
    isLoading,
    email,
    password,
    otp,
    form,
    emailRules,
    login,
    completeOtpVerification,
    
} = useLogin()

// 处理验证码验证
const handleCaptchaVerify = (token: string) => {
    captchaToken.value = token
    showTurnstile.value = false
    login(captchaToken.value)
}

const handleCaptchaError = (error: string) => {
    console.error('验证码错误:', error)
}

// 统一处理所有登录/注册流程
const handleLogin = async () => {
    // 邮箱输入阶段：检查账户状态并进入相应流程
    if (!istoregister.value && !istologin.value) {
        await login(captchaToken.value)
        return
    }

    // 登录阶段：验证表单并请求登录
    if (istologin.value) {
        if (!form.value) return
        const { valid } = await form.value.validate()
        if (!valid) return
        showTurnstile.value = true
        return
    }

    // 注册阶段（OTP验证）：验证OTP
    if (istoregister.value && !isOtpVerified.value) {
        if (!currentOtp.value || currentOtp.value.length < 6) {
            message.info('请输入完整的验证码')
            return
        }

        otpVerifying.value = true
        try {
            // const verified = await verifyOtp(email.value, currentOtp.value)
            const verified = true
            if (verified) {
                // OTP 验证成功，进入密码设置环节
                completeOtpVerification()
                message.info('验证码验证成功')
            } else {
                // OTP 验证失败
                message.info('验证码验证失败，请检查后重试')
            }
        } catch (error) {
            console.error('OTP验证出错:', error)
            message.info('验证码验证过程中出现错误')
        } finally {
            otpVerifying.value = false
        }
        return
    }

    // 注册阶段（密码设置）：验证表单并提交注册
    if (istoregister.value && isOtpVerified.value) {
        if (!form.value) return
        const { valid } = await form.value.validate()
        if (!valid) return
        showTurnstile.value = true
        return
    }
}

// 从OTP组件接收验证码
const handleOtpInput = (otpCode: string) => {
    currentOtp.value = otpCode
    // 如果已经输入了完整的6位验证码，可以自动触发验证
    if (otpCode.length === 6) {
        handleLogin()
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
                <v-card :disabled="isLoading || otpVerifying" :loading="isLoading || otpVerifying">
                    <template v-slot:loader="{ isActive }">
                        <v-progress-linear
                            :active="isActive"
                            color="primary"
                            indeterminate
                        ></v-progress-linear>
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
                            <p v-else-if="istoregister && !isOtpVerified" class="text-h5">注册到 <strong>Nyauth</strong></p>
                            <p v-else class="text-h5">完成注册 <strong>Nyauth</strong></p>
                        </div>
                    </v-card-title>
                    <v-card-text>
                        <v-form ref="form" class="px-4" @submit.prevent>
                            <v-text-field
                                v-if="!istoregister"
                                v-model="email"
                                label="电子邮箱"
                                color="primary"
                                prepend-inner-icon="mdi-email-outline"
                                variant="outlined"
                                :rules="emailRules"
                                @keyup.enter="handleLogin"
                                required
                            />
                            <v-slide-y-transition :leave-absolute="true">
                                <loginform v-if="istologin" v-model="password" />
                                <otpform
                                    v-if="istoregister && !isOtpVerified"
                                    :email="email"
                                    :otp="otp"
                                    :loading="otpVerifying"
                                    @otpEnter="handleOtpInput"
                                />
                                <passwordSetForm
                                    v-if="istoregister && isOtpVerified"
                                    v-model:password="password"
                                    @enter="handleLogin"
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
                            @click="handleLogin"
                            :loading="otpVerifying"
                            :disabled="otpVerifying"
                        >
                            {{ 
                                istoregister && !isOtpVerified ? '验证' : 
                                istoregister && isOtpVerified ? '注册' : '继续' 
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
                                <v-checkbox-btn density="comfortable" color="primary" />
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