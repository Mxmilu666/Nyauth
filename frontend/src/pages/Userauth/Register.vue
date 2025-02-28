<script setup lang="ts">
import { ref } from 'vue'
import { defineOptions } from 'vue'

defineOptions({
    name: 'RegisterCard'
})
</script>

<template>
    <v-form class="px-4">
        <v-text-field
            label="电子邮箱"
            color="primary"
            prepend-inner-icon="mdi-email-outline"
            variant="outlined"
            required
        />
        <v-row>
            <v-col cols="6">
                <v-text-field
                    label="密码"
                    color="primary"
                    type="password"
                    prepend-inner-icon="mdi-key-outline"
                    variant="outlined"
                    required
                />
            </v-col>
            <v-col cols="6">
                <v-text-field
                    label="确认密码"
                    color="primary"
                    type="password"
                    variant="outlined"
                    required
                />
            </v-col>
        </v-row>
    </v-form>
</template>
