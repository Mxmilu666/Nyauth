<template>
    <div>
        <!-- 先拿来占用一下测试个验证码晚点再还回去( -->
        <Turnstile
            v-model:show="showTurnstile"
            @callback="handleCaptchaVerify"
            @error="handleCaptchaError"
        />

        <v-btn @click="showTurnstile = true">进行人机验证</v-btn>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import Turnstile from '@/components/turnstile/Turnstile.vue'

const showTurnstile = ref(false)

const handleCaptchaVerify = (token) => {
    console.log('验证成功，获取到token:', token)

}

const handleCaptchaError = (error) => {
    console.error('验证码错误:', error)
}
</script>
