<script setup lang="ts">
defineOptions({
    name: 'AppComponent'
})
</script>

<template>
    <v-app>
        <v-main>
            <v-container>
                <router-view />
            </v-container>
        </v-main>
    </v-app>
</template>
