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
            <router-view v-slot="{ Component, route }">
                <v-fade-transition>
                    <div :key="route.name">
                        <component :is="Component" />
                    </div>
                </v-fade-transition>
            </router-view>
        </v-main>
    </v-app>
</template>
