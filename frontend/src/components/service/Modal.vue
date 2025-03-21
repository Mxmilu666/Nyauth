<script setup lang="ts">
import { ref } from 'vue'

defineOptions({
    name: 'ModalComponent'
})

const props = defineProps<{
    title: string
    content: string
    confirmButtonText?: string
    cancelButtonText?: string
    emitResult: (result: boolean) => void
    destroyComponent: () => void
}>()

const dialog = ref<boolean>(true)

const handleConfirm = () => {
    dialog.value = false
    props.emitResult(true)
    props.destroyComponent()
}

const handleCancel = () => {
    dialog.value = false
    props.emitResult(false)
    props.destroyComponent()
}
</script>

<template>
    <v-dialog v-model="dialog" max-width="500px" persistent>
        <v-card>
            <v-card-title>{{ title }}</v-card-title>
            <v-card-text>{{ content }}</v-card-text>
            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="error" variant="text" @click="handleCancel">
                    {{ cancelButtonText || '取消' }}
                </v-btn>
                <v-btn color="primary" variant="text" @click="handleConfirm">
                    {{ confirmButtonText || '确认' }}
                </v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>
