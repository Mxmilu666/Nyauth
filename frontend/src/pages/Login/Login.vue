<script setup lang="ts">
import { defineOptions } from 'vue'

defineOptions({
    name: 'LoginPage'
})
</script>

<template>
    <v-container
        class="fill-height d-flex align-center justify-center grey-background"
        fluid
    >
        <v-row align="center" justify="center">
            <v-col cols="12" sm="8" md="4">
                <v-card>
                    <v-card-title class="text-center">
                        <div class="py-5">
                            <v-avatar color="primary">
                                <v-icon icon="mdi-lock-outline" size="small" />
                            </v-avatar>
                            <p class="pt-2 text-h5">登录到 <strong>Nyauth</strong></p>
                        </div>
                    </v-card-title>
                    <v-card-text>
                        <v-form class="px-4">
                            <v-text-field
                                label="电子邮箱"
                                color="primary"
                                prepend-inner-icon="mdi-email-outline"
                                variant="outlined"
                                required
                            />
                            <v-text-field
                                label="密码"
                                color="primary"
                                type="password"
                                prepend-inner-icon="mdi-key-outline"
                                variant="outlined"
                                required
                            />
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
                        >
                            登录
                        </v-btn>
                        <div class="d-flex justify-space-between w-100 mt-1">
                            <v-btn color="primary" variant="text">忘记密码</v-btn>
                            <v-btn color="primary" variant="text">注册账号</v-btn>
                        </div>
                    </v-card-actions>
                </v-card>
                <div class="text-center mt-8">
                    <v-btn color="primary" prepend-icon="mdi-fingerprint" variant="text">
                        使用外部验证器登录
                    </v-btn>
                </div>
            </v-col>
        </v-row>
    </v-container>
</template>
