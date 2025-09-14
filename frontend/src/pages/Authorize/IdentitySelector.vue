<script setup lang="ts">
import type { Identity } from '@/hooks/useAuthorize'

defineProps({
    identities: {
        type: Array as () => Identity[],
        required: true
    },
    selectedIdentityId: {
        type: Number,
        required: true
    },
    selectedIdentity: {
        type: Object as () => Identity,
        required: true
    }
})

const emit = defineEmits(['update:selectedIdentityId'])

const onIdentityChange = (id: number) => {
    emit('update:selectedIdentityId', id)
}
</script>

<template>
    <div>
        <p class="font-weight-medium mb-2">通过以下身份继续：</p>
        <v-select
            :model-value="selectedIdentityId"
            @update:model-value="onIdentityChange"
            :items="identities"
            item-value="id"
            item-title="userName"
            :return-object="false"
            variant="outlined"
            class="identity-select"
        >
            <template v-slot:selection="{ item }">
                <div class="d-flex align-center identity-selection">
                    <v-avatar size="28" class="me-2 flex-shrink-0">
                        <v-img
                            :src="selectedIdentity.avatar"
                            :alt="selectedIdentity.userName"
                        />
                    </v-avatar>
                    <div class="d-flex flex-wrap align-center">
                        <span class="font-weight-medium me-1 text-truncate" style="max-width: 200px">
                            {{ selectedIdentity.userName }}
                        </span>
                        <v-chip
                            size="small"
                            color="primary"
                            class="me-1 flex-shrink-0 my-1"
                        >
                            <v-icon icon="mdi-label" />
                            <span class="text-truncate" style="max-width: 120px">{{ selectedIdentity.tagText }}</span>
                        </v-chip>
                        <span class="text-caption text-grey text-truncate">
                            {{ selectedIdentity.email }}
                        </span>
                    </div>
                </div>
            </template>
            <template v-slot:item="{ item, props }">
                <v-list-item v-bind="props">
                    <template v-slot:prepend>
                        <v-avatar size="28" class="mr-2">
                            <v-img :src="item.raw.avatar" :alt="item.raw.userName" />
                        </v-avatar>
                    </template>
                    <v-list-item-subtitle
                        >{{ item.raw.email }}
                        <v-chip
                            size="x-small"
                            color="primary"
                            class="ml-1"
                            density="compact"
                        >
                            {{ item.raw.tagText }}
                        </v-chip>
                    </v-list-item-subtitle>
                </v-list-item>
            </template>
        </v-select>
    </div>
</template>
