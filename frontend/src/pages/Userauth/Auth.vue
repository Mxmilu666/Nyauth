<script setup lang="ts">
import { defineOptions } from 'vue'
import { ref } from 'vue'

import loginform from './Loginform.vue'
import otpform from './Otpform.vue'
import { useLogin } from '@/hooks/useLogin'
import turnstile from '@/components/turnstile/Turnstile.vue'

defineOptions({
    name: 'AuthPage'
})

const showTurnstile = ref(false)
const captchaToken = ref('')

const {
    istologin,
    istoregister,
    isLoading,
    email,
    password,
    otp,
    form,
    emailRules,
    login
} = useLogin()

// 记得移到 Hook
const handleCaptchaVerify = (token: string) => {
    captchaToken.value = token
    showTurnstile.value = false
    login(captchaToken.value)
}

const handleCaptchaError = (error: string) => {
    console.error('验证码错误:', error)
}

const handleLogin = async () => {
    if (!istoregister.value && !istologin.value) {
        await login(captchaToken.value)
        return
    }

    if (!form.value) return
    const { valid } = await form.value.validate()
    if (!valid) return
    showTurnstile.value = true
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
                <v-card :disabled="isLoading" :loading="isLoading">
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
                            <p v-else class="text-h5">注册到 <strong>Nyauth</strong></p>
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
                                @keyup.enter="login"
                                required
                            />
                            <v-slide-y-transition>
                                <loginform v-if="istologin" v-model="password" />
                                <otpform
                                    v-if="istoregister"
                                    :email="email"
                                    :otp="otp"
                                    @enter="login"
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
                        >
                            继续
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
