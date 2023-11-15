<template>
  <div class="h-[2000px] p-8">
  </div>
</template>

<script setup lang="ts">
  import { useActivePage } from '@/store/index'
  definePageMeta({
    title: 'Beranda',
    middleware: ['auth'],
  })

  onMounted(() => {
    const activePage = useActivePage()
    activePage.page = 'Dashboard'
    activePage.navigation = false
    
  })
</script>
