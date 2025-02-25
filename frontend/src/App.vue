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
            <RouterView v-slot="{ Component, route }">
                <v-fade-transition>
                    <div :key="route.name">
                        <component :is="Component"></component>
                    </div>
                </v-fade-transition>
            </RouterView>
        </v-main>
    </v-app>
</template>
