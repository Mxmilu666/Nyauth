<script setup lang="ts">
defineOptions({
    name: 'AppComponent'
})
</script>

<template>
    <v-app>
        <v-main>
            <router-view />
            <v-btn prepend-icon="$vuetify" append-icon="$vuetify"> Button </v-btn>
        </v-main>
    </v-app>
</template>
