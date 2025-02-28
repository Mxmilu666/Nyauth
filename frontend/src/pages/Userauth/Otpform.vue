<script setup>
import { ref } from 'vue'
const otp = ref('')

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
    }
})
</script>

<template>
    <div>
        <v-sheet color="surface">
            <v-otp-input v-model="otp" type="" />
        </v-sheet>
        <p class="text-center">验证码已发送至 {{ email }}</p>
        <p class="text-center">请注意查收并验证</p>
    </div>
</template>
