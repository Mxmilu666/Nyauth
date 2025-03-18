<script setup lang="ts">
import { defineProps } from 'vue'

defineProps({
    appName: {
        type: String,
        default: ''
    },
    appCreator: {
        type: String,
        default: ''
    },
    appDescription: {
        type: String,
        default: ''
    },
    appIcon: {
        type: String,
        default: 'https://placehold.co/100'
    }
})
</script>

<template>
    <div
        class="d-flex flex-column align-center justify-center pa-6"
        :class="{ 'fill-height': $vuetify.display.smAndUp }"
    >
        <v-avatar size="80" class="mb-4">
            <v-img :src="appIcon" alt="应用图标" />
        </v-avatar>
        <h3 class="text-h5 mb-2 text-white text-center">{{ appName }}</h3>
        <p class="text-body-2 mb-4 text-white text-center">
            由 {{ appCreator }} 创建
        </p>
        <p class="text-caption text-white text-center">
            {{ appDescription }}
        </p>
    </div>
</template>
