<script setup lang="ts">
import { ref, type VNodeRef } from 'vue'
import { defineOptions } from 'vue'
import { VForm } from 'vuetify/lib/components/index.mjs'

import loginform from './Loginform.vue'
import otpform from './Otpform.vue'

defineOptions({
    name: 'AuthPage'
})

const istologin = ref(false)
const istoregister = ref(false)

const email = ref('')
const password = ref('')
const otp = ref('')
const form = ref<InstanceType<typeof VForm>>()

const emailRules = [
    (value: string) => !!value || '电子邮箱是必须的',
    (value: string) => /.+@.+\..+/.test(value) || '电子邮箱格式非法'
]

const login = async () => {
    if (!form.value) return
    const { valid } = await form.value.validate()
    if (!valid) return
    istoregister.value = true
}
</script>

<template>
    <v-container
        class="auth-wrapper fill-height d-flex align-center justify-center"
        fluid
    >
        <v-row align="center" justify="center">
            <v-col cols="12" sm="8" md="4">
                <v-card>
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
                                <loginform v-if="istologin" :password="password" />
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
                            @click="login"
                        >
                            继续
                        </v-btn>
                        <div class="d-flex justify-space-between w-100 mt-1">
                            <v-btn
                                color="primary"
                                variant="text"
                                @click="$router.push('/reset-password')"
                                >忘记密码</v-btn
                            >
                        </div>
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
