<script setup lang="ts">
defineOptions({
    name: 'BasicInfoItem'
})
interface Props {
    label: string
    content: string
    avatar?: string
    onClick?: () => void
}
const props = defineProps<Props>()
const { onClick } = props
</script>

<template>
    <v-list-item @click="onClick && onClick()">
        <v-row no-gutters>
            <v-col cols="12" md="3">
                <p class="text-subtitle-2 font-weight-light">{{ label }}</p>
            </v-col>
            <v-col cols="12" md="9">
                <p class="font-weight-medium">{{ content }}</p>
            </v-col>
        </v-row>
        <template v-slot:append>
            <template v-if="avatar">
                <v-avatar :image="avatar" />
            </template>
            <template v-else>
                <v-icon icon="mdi-chevron-right" />
            </template>
        </template>
    </v-list-item>
</template>
