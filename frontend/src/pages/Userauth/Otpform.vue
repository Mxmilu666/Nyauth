<script setup lang="ts">
import { defineOptions, ref, watch } from 'vue'

defineOptions({
    name: 'Otpform'
})

const props = defineProps({
    email: {
        type: String,
        required: true
    },
    otp: {
        type: String,
        required: true
    },
    loading: {
        type: Boolean,
        default: false
    }
})

const emit = defineEmits(['otpEnter'])

const localOtp = ref('')

// 监听OTP输入变化，实时传递给父组件
watch(localOtp, (newValue) => {
    emit('otpEnter', newValue)
})

</script>

<template>
    <div>
        <v-otp-input 
            v-model="localOtp" 
            type="text"
            :disabled="loading"
        />
        <p class="text-center mt-3">验证码已发送至 {{ email }}</p>
        <p class="text-center">请注意查收并验证</p>
    </div>
</template>