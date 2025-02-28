<script setup lang="ts">
import { defineOptions } from 'vue'

defineOptions({
    name: 'LoginForm'
})

const props = defineProps({
    password: {
        type: String,
        required: true
    }
})

const emit = defineEmits(["enter"])
</script>

<template>
    <v-text-field
        label="密码"
        :v-model="props.password"
        color="primary"
        type="password"
        prepend-inner-icon="mdi-key-outline"
        variant="outlined"
        @keyup.enter="emit('enter')"
        required
    />
</template>
