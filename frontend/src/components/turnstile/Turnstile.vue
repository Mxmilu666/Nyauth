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
const siteKey = ref<string>('')

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
        if (data && data.data !== undefined) {
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

function closeDialog() {
    dialog.value = false
}
</script>

<template>
    <v-dialog v-model="dialog" max-width="400px" persistent>
        <v-card>
            <v-card-title class="d-flex justify-space-between align-center">
                请完成人机验证
                <v-btn icon size="small" variant="text" @click="closeDialog">
                    <v-icon icon="mdi-close" />
                </v-btn>
            </v-card-title>
            <v-card-text
                class="d-flex justify-center align-center"
                style="min-height: 100px"
            >
                <vue-turnstile
                    :site-key="captchaKey"
                    v-model="turnstileToken"
                    class="rounded"
                />
            </v-card-text>
        </v-card>
    </v-dialog>
</template>
