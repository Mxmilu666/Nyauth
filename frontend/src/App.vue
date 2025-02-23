<script setup lang="ts">
import { message } from './services/message'
import AppBar from './components/AppBar.vue'

defineOptions({
    name: 'AppComponent'
})

message.info('Hello from App Component')
</script>

<template>
    <v-app>
        <AppBar />
        <v-main>
            <router-view />
        </v-main>
    </v-app>
</template>