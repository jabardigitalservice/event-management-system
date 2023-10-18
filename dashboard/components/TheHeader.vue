<template>
  <header
    data-cy="header__container"
    class="flex h-[72px] w-full items-center justify-between bg-white px-6"
  >
    <h1
      data-cy="header__title"
      class="font-roboto text-blue-gray-800 mr-4 text-xl font-bold"
    >
      {{ pageTitle }}
    </h1>
    <HeaderMenu />
  </header>
</template>

<script setup lang="ts">
  const route = useRoute()
  const pageTitle = route.meta?.title ?? ''
</script>

<style scoped></style>
