<template>
  <label
    v-show="label"
    :for="name"
    class="message-notif-form__label-required text-gray-800"
    >{{ label }}</label
  >
  <p v-show="sublabel" :for="name" class="text-[13px] text-gray-600">
    {{ sublabel }}
  </p>
  <UFormGroup>
    <UInput
      :id="name"
      v-model="value"
      :name="name"
      :color="errorMessage ? 'red' : 'white'"
      :type="type"
      :placeholder="placeholder"
      class="mt-1"
      :disabled="disabled"
    >
    <template #leading>
      <NuxtIcon
        :name="icon"
        filled
        class="text-xl text-gray-300"
      />
    </template>
  </UInput>
  </UFormGroup>
  <p v-show="errorMessage" class="help-message text-sm text-red-500">
    {{ errorMessage }}
  </p>
</template>

<script setup lang="ts">
  import { useField } from 'vee-validate'

  const props = defineProps({
    type: {
      type: String,
      default: 'text',
    },
    value: {
      type: String,
      default: '',
    },
    name: {
      type: String,
      required: true,
    },
    label: {
      type: String,
      default: '',
    },
    sublabel: {
      type: String,
      default: '',
    },
    placeholder: {
      type: String,
      default: '',
    },
    icon: {
      type: String,
      default: '',
    },
    disabled: {
      type: Boolean, 
      default: false
    }
  })

  const { value, errorMessage } = useField(() => props.name)
</script>
