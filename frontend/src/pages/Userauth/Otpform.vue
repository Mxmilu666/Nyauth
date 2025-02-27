<script setup>
import { ref } from 'vue'
const otp = ref('')
</script>

<template>
    <v-sheet color="surface">
        <v-otp-input v-model="otp" type="" />
    </v-sheet>
</template>
