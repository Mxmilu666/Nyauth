<script setup lang="ts">
import { defineOptions, ref, watch, onMounted, onBeforeUnmount } from 'vue'

defineOptions({
    name: 'TotpVerifyForm'
})

const props = defineProps({
    loading: {
        type: Boolean,
        default: false
    }
})

const emit = defineEmits(['totpEnter'])

const totpCode = ref('')
const countdown = ref(0)
const TOTP_PERIOD = 30 // TOTP码通常30秒更新一次
let timer: number | null = null

// 计算并更新倒计时
const updateCountdown = () => {
    // 计算当前时间在30秒周期内的位置
    const now = Math.floor(Date.now() / 1000)
    countdown.value = TOTP_PERIOD - (now % TOTP_PERIOD)
}

// 启动倒计时定时器
const startCountdown = () => {
    updateCountdown() // 立即计算一次当前倒计时值
    timer = window.setInterval(() => {
        updateCountdown()
    }, 1000)
}

// 组件挂载时启动倒计时
onMounted(() => {
    startCountdown()
})

// 组件卸载前清除定时器
onBeforeUnmount(() => {
    if (timer !== null) {
        clearInterval(timer)
        timer = null
    }
})

// 监听TOTP输入变化，当到达6位时传递给父组件
watch(totpCode, (newValue) => {
    if (newValue.length === 6) {
        emit('totpEnter', newValue)
    }
})
</script>

<template>
    <div>
        <v-otp-input v-model="totpCode" :disabled="loading" length="6" />
        <p class="text-center mt-3">请输入两步验证码</p>
        <div class="text-center d-flex align-center justify-center">
            <span>请打开您的验证器应用获取验证码</span>
            <v-progress-circular
                class="ms-2"
                :model-value="(countdown / TOTP_PERIOD) * 100"
                :size="27"
                color="primary"
                width="4"
            >
                <span class="text-caption">{{ countdown }}</span>
            </v-progress-circular>
        </div>
    </div>
</template>
