<template>
  <ObjectForm />
</template>
<script setup lang="ts">
  import { useActivePage, useIdData } from '@/store/index'

  definePageMeta({
    middleware: ['auth'],
  })
  onMounted(() => {
    const activePage = useActivePage()
    activePage.page = !!useIdData().id?'Ubah Objek Wisata' :'Tambah Objek Wisata'
    activePage.navigation = true
  })
</script>
