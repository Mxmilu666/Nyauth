<script setup lang="ts">
interface Permission {
    title: string
    description: string
}

defineProps({
    permissions: {
        type: Array as () => Permission[],
        default: () => []
    }
})
</script>

<template>
    <div>
        <p class="font-weight-medium mb-2">此应用将获得以下权限：</p>
        <v-list>
            <v-list-item v-for="(permission, index) in permissions" :key="index" density="compact">
                <template v-slot:prepend>
                    <v-icon color="primary" size="small" icon="mdi-check-circle" />
                </template>
                <v-list-item-title>{{ permission.title }}</v-list-item-title>
                <v-list-item-subtitle>{{ permission.description }}</v-list-item-subtitle>
            </v-list-item>
        </v-list>
    </div>
</template>
