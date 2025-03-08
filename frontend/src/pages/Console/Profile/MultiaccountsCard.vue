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
    <v-card class="d-flex flex-column align-center justify-center">
        <v-avatar :image="avatar" />
        <v-card-text class="username text-center">
            {{ userName }}
        </v-card-text>
    </v-card>
</template>

<style scoped></style>
