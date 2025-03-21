<script setup lang="ts">
import { defineOptions } from 'vue'

defineOptions({
    name: 'LoginForm'
})

const props = defineProps({
    modelValue: {
        type: String,
        required: true
    }
})

const emit = defineEmits(['update:modelValue', 'enter'])
</script>

<template>
    <v-text-field
        label="密码"
        :value="modelValue"
        @update:model-value="value => emit('update:modelValue', value)"
        color="primary"
        type="password"
        prepend-inner-icon="mdi-key-outline"
        variant="outlined"
        @keyup.enter="emit('enter')"
        required
    />
</template>
