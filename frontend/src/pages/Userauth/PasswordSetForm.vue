<script setup lang="ts">
import { defineOptions, ref, computed } from 'vue'

defineOptions({
    name: 'PasswordSetForm'
})

const props = defineProps({
    password: {
        type: String,
        required: true
    }
})

const emit = defineEmits(['update:password', 'otpEnter', 'keyup'])

const passwordModel = computed({
    get: () => props.password,
    set: (value) => emit('update:password', value)
})

const confirmPassword = ref('')
const showPassword = ref(false)
const showConfirmPassword = ref(false)
const passwordError = ref('')

// 密码强度规则
const passwordRules = [
    (v: string) => !!v || '请输入密码',
    (v: string) => v.length >= 8 || '密码长度至少为8位',
    (v: string) => /[A-Z]/.test(v) || '密码需要包含至少一个大写字母',
    (v: string) => /[a-z]/.test(v) || '密码需要包含至少一个小写字母',
    (v: string) => /[0-9]/.test(v) || '密码需要包含至少一个数字',
    (v: string) => /[!@#$%^&*(),.?":{}|<>]/.test(v) || '密码需要包含至少一个特殊字符'
]

// 确认密码规则
const confirmRules = [
    (v: string) => !!v || '请确认密码',
    (v: string) => v === passwordModel.value || '两次输入的密码不一致'
]

const handleKeyup = (event: KeyboardEvent) => {
    emit('keyup', event) // 转发原始keyup事件
    if (event.key === 'Enter') {
        if (passwordModel.value && confirmPassword.value === passwordModel.value) {
            emit('otpEnter')
        }
    }
}
</script>

<template>
    <div>
        <v-text-field
            v-model="passwordModel"
            :append-icon="showPassword ? 'mdi-eye-off' : 'mdi-eye'"
            :type="showPassword ? 'text' : 'password'"
            label="设置密码"
            color="primary"
            prepend-inner-icon="mdi-lock-outline"
            variant="outlined"
            :rules="passwordRules"
            @click:append="showPassword = !showPassword"
            autocomplete="new-password"
            required
        />

        <v-text-field
            v-model="confirmPassword"
            :append-icon="showConfirmPassword ? 'mdi-eye-off' : 'mdi-eye'"
            :type="showConfirmPassword ? 'text' : 'password'"
            label="确认密码"
            color="primary"
            prepend-inner-icon="mdi-lock-check-outline"
            variant="outlined"
            :rules="confirmRules"
            @click:append="showConfirmPassword = !showConfirmPassword"
            @keyup="handleKeyup"
            autocomplete="new-password"
            required
        />
    </div>
</template>
