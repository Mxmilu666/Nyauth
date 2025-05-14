<script setup lang="ts">
defineOptions({
    name: 'TotpPage'
})
</script>

<template>
    <div class="totp">
        <v-container class="center">
            <p class="text-h5 pt-4 font-weight-bold">两步验证</p>
            <p class="text-subtitle-1 py-1">查看和修改您的两步验证码</p>

            <div class="content"></div>
        </v-container>
    </div>
</template>
