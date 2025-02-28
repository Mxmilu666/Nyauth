<script setup lang="ts">
import { ref } from 'vue'
import { defineOptions } from 'vue'

defineOptions({
    name: 'ResetPage'
})
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
                                <img src="@/assets/sticker/yuzu_sad.png" class="logo" />
                            </v-lazy>
                            <p class="text-h5">找回密码</p>
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
                            下一步
                        </v-btn>
                        <div class="d-flex justify-space-between w-100 mt-1">
                            <v-btn
                                color="primary"
                                variant="text"
                                @click="$router.push('/login')"
                                >返回登录</v-btn
                            >
                        </div>
                    </v-card-actions>
                </v-card>
            </v-col>
        </v-row>
    </v-container>
</template>
