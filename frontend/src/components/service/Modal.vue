<script setup lang="ts">
import { ref } from 'vue'
import type { ModalButton } from './index'

defineOptions({
    name: 'ModalComponent'
})

const props = defineProps<{
    title: string
    content: string
    prependIcon: string
    buttons: ModalButton[]
    emitResult: (result: any) => void
    destroyComponent: () => void
}>()

const dialog = ref<boolean>(true)

const handleButtonClick = (value: any) => {
    dialog.value = false
    props.emitResult(value)
    props.destroyComponent()
}
</script>

<template>
    <v-dialog v-model="dialog" max-width="500" persistent>
        <v-card max-width="500" :prepend-icon="prependIcon" :text="content" :title="title">
            <template v-slot:actions>
                <v-btn
                    v-for="(button, index) in buttons"
                    :key="index"
                    :color="button.color || 'default'"
                    :variant="
                        (button.variant as
                            | 'flat'
                            | 'text'
                            | 'elevated'
                            | 'tonal'
                            | 'outlined'
                            | 'plain') || 'text'
                    "
                    @click="handleButtonClick(button.value)"
                >
                    {{ button.text }}
                </v-btn>
            </template>
        </v-card>
    </v-dialog>
</template>