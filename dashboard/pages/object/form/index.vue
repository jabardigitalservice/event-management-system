<template>
  <ObjectForm />
</template>
<script setup lang="ts">
  import { useActivePage } from '@/store/index'

  definePageMeta({
    middleware: ['auth'],
  })
  onMounted(() => {
    const activePage = useActivePage()
    activePage.page = 'Tambah Objek Wisata'
    activePage.desc = 'Form Pengisian Data Objek Wisata'
  })
</script>
