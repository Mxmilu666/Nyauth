<script setup lang="ts">
import { message } from './services/message'

defineOptions({
    name: 'AppComponent'
})

message.info('Hello from App Component')
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
