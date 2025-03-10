<script setup lang="ts">
defineOptions({
    name: 'MultiaccountsCard'
})

const props = defineProps<{
    avatar: string
    userName: string
}>()
</script>

<template>
    <v-card class="d-flex flex-column align-center justify-center pa-1">
        <v-avatar :image="avatar" size="70" />
        <v-card-text class="text-center text-h6">
            {{ userName }}
        </v-card-text>
    </v-card>
</template>

<style scoped></style>
