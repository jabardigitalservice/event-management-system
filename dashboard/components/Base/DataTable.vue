<template>
    <div>
      <EasyDataTable
        :headers="props.headers"
        :items="props.items"
        table-class-name="customize-table"
        header-text-direction="center"
        body-text-direction="center"
        show-index
      />
    </div>
  </template>
  
<script setup lang="ts">
    const props = defineProps({
      headers: {
            type: String,
            default: ''
        },
        items: {
            type: String,
            default: ''
        }
    });
</script>
<style>
</style>

