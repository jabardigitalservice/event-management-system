<template>
  <ObjectForm />
</template>
<script setup lang="ts">
  import { useActivePage } from '@/store/index'

  definePageMeta({
    layout:'header-base',
    middleware: ['auth'],
  })
  onMounted(() => {
    const activePage = useActivePage()
    activePage.page = useRoute().query?.id ? 'Ubah Objek Wisata' :'Tambah Objek Wisata'
    activePage.navigation = true
  })
</script>
