<template>
  <label
    v-show="label"
    :for="name"
    class="text-gray-800"
    >{{ label }}</label
  >
  <p v-show="sublabel" :for="name" class="text-[13px] text-gray-600">
    {{ sublabel }}
  </p>
  <UTextarea
    :id="name"
    v-model="value"
    :name="name"
    :color="errorMessage ? 'red' : 'white'"
    :type="type"
    :placeholder="placeholder"
    class="mt-1"
    :disabled="disabled"
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
    disabled: {
      type: Boolean, 
      default: false
    }
  })

  const { value, errorMessage } = useField(() => props.name)

  onUpdated( () => {
    if (!!props?.value) {
      value.value = props.value
    }
  })
</script>
