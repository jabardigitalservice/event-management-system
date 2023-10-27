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
  <UInput
    :id="name"
    v-model="value"
    :name="name"
    :color="errorMessage ? 'red' : 'green'"
    :type="type"
    :placeholder="placeholder"
    class="mt-1"
  />
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
      default: undefined,
    },
    name: {
      type: String,
      required: true,
    },
    label: {
      type: String,
      default: undefined,
    },
    sublabel: {
      type: String,
      default: undefined,
    },
    placeholder: {
      type: String,
      default: '',
    },
  })

  const { value, errorMessage } = useField(() => props.name)
</script>
