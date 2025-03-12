<script setup lang="ts">
defineOptions({
    name: 'MultiaccountsCard'
})

const props = defineProps<{
    avatar: string
    userName: string
    lastActiveTime?: string
    tagText?: string
}>()
</script>

<template>
    <v-card>
        <v-card-text>
            <div class="d-flex align-center">
                <v-avatar :image="avatar" size="60" class="me-4" />
                <div class="d-flex flex-column justify-center flex-grow-1">
                    <div class="text-subtitle-1 font-weight-medium d-flex align-center">
                        <span class="username-text text-truncate">{{ userName }}</span>
                        <v-chip
                            v-if="tagText"
                            size="small"
                            color="primary"
                            text-color="white"
                            class="ms-2 flex-shrink-0"
                        >
                            <v-icon icon="mdi-label" start />
                            {{ tagText }}
                        </v-chip>
                    </div>
                    <div class="text-caption text-grey" v-if="lastActiveTime">
                        最近授权: {{ lastActiveTime }}
                    </div>
                </div>
            </div>
        </v-card-text>
    </v-card>
</template>

<style scoped>
.v-chip {
    margin-top: 1px;
}

.username-text {
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    max-width: 100px;
}
</style>
