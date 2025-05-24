<script setup lang="ts">
import { defineOptions, ref, watch } from 'vue'

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

// 监听TOTP输入变化，当到达6位时传递给父组件
watch(totpCode, (newValue) => {
    if (newValue.length === 6) {
        emit('totpEnter', newValue)
    }
})
</script>

<template>
    <div>
        <v-otp-input 
            v-model="totpCode" 
            :disabled="loading"
            length="6"
        />
        <p class="text-center mt-3">请输入两步验证码</p>
        <p class="text-center">请打开您的验证器应用获取验证码</p>
    </div>
</template>