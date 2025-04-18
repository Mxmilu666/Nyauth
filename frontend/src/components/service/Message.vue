<script setup lang="ts">
import { ref, computed } from 'vue'

defineOptions({
    name: 'MessageComponent'
})

type MessageType = 'info' | 'success' | 'error' | 'warning'

const props = defineProps<{
    message: string
    type?: MessageType
    emitResult: () => void
    destroyComponent: () => void
}>()

const snackbar = ref<boolean>(true)

const handleConfirm = () => {
    snackbar.value = false
    props.emitResult()
    props.destroyComponent()
}

const messageColor = computed(() => {
    switch (props.type) {
        case 'success':
            return 'success'
        case 'error':
            return 'error'
        case 'warning':
            return 'warning'
        default:
            return 'primary'
    }
})

const messageIcon = computed(() => {
    switch (props.type) {
        case 'success':
            return 'mdi-check-circle'
        case 'error':
            return 'mdi-alert-circle'
        case 'warning':
            return 'mdi-alert'
        default:
            return 'mdi-information'
    }
})
</script>

<template>
    <v-snackbar v-model="snackbar" location="top" :color="messageColor">
        <div class="d-flex align-center">
            <v-icon :icon="messageIcon" class="mr-2" />
            {{ message }}
        </div>

        <template v-slot:actions>
            <v-btn :color="messageColor" variant="text" @click="handleConfirm()">
                OK
            </v-btn>
        </template>
    </v-snackbar>
</template>
