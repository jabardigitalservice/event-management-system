<template>
  <OrganizationForm />
</template>
<script setup lang="ts">
  import { useActivePage } from '@/store/index'

  definePageMeta({
    middleware: ['auth'],
  })
  onMounted(() => {
    const activePage = useActivePage()
    activePage.page = 'Tambah Organisasi'
    activePage.desc = 'Form Pengisian Data Organisasi'
  })
</script>
