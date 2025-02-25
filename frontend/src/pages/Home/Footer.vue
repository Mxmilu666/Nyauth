<script setup lang="ts">
defineOptions({
    name: 'FooterComponent'
})
</script>

<template>
    <v-footer class="custom-footer" color="appbar">
        <v-row justify="center" no-gutters>
            <v-col cols="12">
                版权所有 © {{ new Date().getFullYear() }}
                <strong>Nyauth</strong> 保留所有权利
            </v-col>
            <v-col class="mt-2" cols="12"> By Mxmilu & Alongw with ❤ </v-col>
            <v-col class="mt-2" cols="12">
                <v-btn
                    color="while"
                    prepend-icon="mdi-github"
                    href="https://github.com/Mxmilu666/Nyauth"
                    variant="outlined"
                >
                    GitHub
                </v-btn>
            </v-col>
        </v-row>
    </v-footer>
</template>

<style scoped>
.custom-footer {
    padding: 70px;
}
</style>
