<template>
    <div>
    <UModal v-model="props.show">
      <UCard :ui="{ divide: 'divide-y divide-gray-100 dark:divide-gray-800' }">
        <template #header>
            <div class="flex items-center justify-between">
            <h3 class="text-base font-semibold leading-6 text-gray-900 dark:text-white">
              {{title}}
            </h3>
            <UButton 
                color="gray" 
                variant="ghost" 
                icon="i-heroicons-x-mark-20-solid" 
                class="-my-1" 
                @click="handleClose"
            />
          </div>
        </template>
        <iframe
            v-if="props.mimeType === 'application/pdf'"
            :src="fileDocument"
            class="w-full min-h-[600px]"
        />
        <img v-else :src="fileDocument()" class="w-full">
      </UCard>
    </UModal>
    </div>
</template>
<script setup lang="ts">
    import { base64ToBlobUrl } from '~/utils'
    const props = defineProps({
        title: {
            type: String,
            default: '',
        },
        show: {
            type: Boolean,
            default: false
        },
        file: {
            type: String,
            default: ''
        },
        mimeType: {
            type: String,
            default: ''
        },
        withUrlPath: {
            type: Boolean,
            default: false
        }
    })

    const emit = defineEmits([
        'close'
    ])

    const handleClose = () => {
        emit('close')
    }

    const fileDocument = () => {
        if (props.withUrlPath) {
            return props.file ? props.file : ''
        } else {
            return props.file ? base64ToBlobUrl(props.file, props.mimeType) : ''
        }
    }
</script>