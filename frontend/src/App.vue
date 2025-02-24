<script setup lang="ts">
import { message } from './services/message'
import AppBar from '@/components/AppBar.vue'

defineOptions({
    name: 'AppComponent'
})
</script>

<template>
    <v-app>
        <AppBar />
        <v-main>
            <v-fade-transition>
                <router-view />
            </v-fade-transition>
        </v-main>
    </v-app>
</template>