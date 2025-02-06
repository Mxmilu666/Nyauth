<script setup lang="ts">
import { ref } from 'vue'

defineOptions({
    name: 'MessageComponent'
})

const props = defineProps<{
    message: string
    emitResult: () => void
    destroyComponent: () => void
}>()

const snackbar = ref<boolean>(true)

const handleConfirm = () => {
    snackbar.value = false
    props.emitResult()
    props.destroyComponent()
}
</script>

<template>
    <v-snackbar v-model="snackbar">
        {{ message }}

        <template v-slot:actions>
            <v-btn color="pink" variant="text" @click="handleConfirm()"> OK </v-btn>
        </template>
    </v-snackbar>
</template>
