<template>
  <OrganizationForm />
</template>
<script setup lang="ts">
  import { useActivePage } from '@/store/index'

  definePageMeta({
    layout:'header-base',
    middleware: ['auth'],
  })
  onMounted(() => {
    const activePage = useActivePage()
    activePage.page = useRoute().query?.id ? 'Ubah Organisasi' :'Tambah Organisasi'
    activePage.navigation = true
  })
</script>
