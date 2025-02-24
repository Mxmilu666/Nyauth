<script setup lang="ts">
import { useTheme } from 'vuetify';

const theme = useTheme();
function toggleTheme() {
    theme.global.name.value = theme.global.current.value.dark ? 'lightTheme' : 'darkTheme';
}
</script>

<template>
    <v-app-bar app color="appbar">
        <v-toolbar-title>Nyauth</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-btn icon @click="toggleTheme">
            <v-icon>{{ theme.global.current.value.dark ? 'mdi-weather-night' : 'mdi-weather-sunny' }}</v-icon>
        </v-btn>
        <v-btn icon>
            <v-icon>mdi-home</v-icon>
        </v-btn>
        <v-btn icon>
            <v-icon>mdi-account</v-icon>
        </v-btn>
    </v-app-bar>
</template>