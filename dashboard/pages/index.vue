<template>
  <div class="h-[2000px] p-8">
    <h1 class="text-3xl">Jds Builder</h1>
  </div>
</template>

<script setup lang="ts">
  // const { status, data, lastRefreshedAt, getCsrfToken } = useAuth()

  definePageMeta({
    title: 'Beranda',
    middleware: ['auth'],
  })
  // console.log(status, data, lastRefreshedAt, await getCsrfToken())
</script>
