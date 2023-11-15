<template>
  <OrganizationForm />
</template>
<script setup lang="ts">
  import { useActivePage, useIdData } from '@/store/index'

  definePageMeta({
    middleware: ['auth'],
  })
  onMounted(() => {
    const activePage = useActivePage()
    activePage.page = !!useIdData().id?'Ubah Organisasi' :'Tambah Organisasi'
    activePage.navigation = true
  })
</script>
