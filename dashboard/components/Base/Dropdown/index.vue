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
  <USelectMenu
    v-model="state.selected"
    :disabled="disabled"
    :options="dataDropdown"
    searchable
    @change="onChange"
  >
    <template #label>
      <span v-if="state.selected?.id" class="truncate">{{ placeholder }}</span>
    </template>
    <template #option-empty="{ query }">
      <q>{{ query }}</q> not found
    </template>
  </USelectMenu>
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
      default: false,
    },
    dataDropdown: {
      type: Array,
      default: () => [],
    },
    dataSelected: {
      type: String,
      default: '',
    },
  })

  const state = reactive({
    selected: [],
  })

  const { value, errorMessage } = useField(() => props.name)

  onMounted(async () => {
    if (value) {
      const filterData = await props.dataDropdown.filter(
        (el) => el.id === value.value,
      )
      state.selected = filterData[0]
    }
  })

  const onChange = (e) => {
    value.value = e.id
    state.selected = e
  }
</script>
