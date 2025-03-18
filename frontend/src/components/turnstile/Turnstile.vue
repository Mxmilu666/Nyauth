<script setup lang="ts">
import VueTurnstile from 'vue-turnstile'
import { getCaptcha } from '@/api/captcha'
import { ref, computed, onMounted, watch } from 'vue'

const props = defineProps<{
    show: boolean
}>()

const emit = defineEmits<{
    (e: 'error', error: string): void
    (e: 'callback', token: string): void
    (e: 'update:show', value: boolean): void
}>()

const turnstileToken = ref<string>('')
const siteKey = ref<string>('') // 默认值

// 计算属性监听 props.show 并触发更新
const dialog = computed({
    get: () => props.show,
    set: (value) => emit('update:show', value)
})

const captchaKey = computed(() => siteKey.value)

onMounted(async () => {
    const key = await getCaptchaKey()
    if (key) {
        siteKey.value = key
    }
})

// 从后端获取验证码
const getCaptchaKey = async () => {
    try {
        const { data } = await getCaptcha()
        if (data !== undefined) {
            return data.data.id
        }
        return null
    } catch (error) {
        console.error('验证码获取失败:', error)
        emit('error', '验证码获取失败')
        return null
    }
}

watch(turnstileToken, (token) => {
    if (token) {
        callback(token)
    }
})

function callback(token: string) {
    emit('callback', token)
}
</script>

<template>
    <v-dialog v-model="dialog" max-width="400px">
        <v-card>
            <v-card-title>请完成人机验证</v-card-title>
            <v-card-text
                class="d-flex justify-center align-center"
                style="min-height: 120px"
            >
                <vue-turnstile :site-key="captchaKey" v-model="turnstileToken" />
            </v-card-text>
            <v-card-actions>
                <v-btn color="primary" variant="text" @click="dialog = false">
                    取消
                </v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>
