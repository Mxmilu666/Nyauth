<script setup lang="ts">
import AppBar from '@/components/AppBar.vue'

defineOptions({
    name: 'AppComponent'
})
</script>

<template>
    <v-app>
        <app-bar />
        <v-main>
            <v-fade-transition>
                <router-view />
            </v-fade-transition>
        </v-main>
    </v-app>
</template>
