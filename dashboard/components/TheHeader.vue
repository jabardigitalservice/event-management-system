<template>
  <header
    data-cy="header__container"
    class="flex h-[72px] w-full items-center justify-between bg-white px-6">
    <div class="ml-7">
      <nav class="flex flex-wrap items-center">
        <div class="inline-block">
          <div
            class="font-roboto text-blue-gray-800 mr-4 text-md font-bold">
            {{ activePage.page }}
          </div>
        </div>
      </nav>
      <p class="font-roboto text-blue-gray-600 mr-4 text-xs font-light ">{{ activePage.desc }}</p>
    </div>
    <HeaderMenu />
  </header>
</template>

<script setup lang="ts">
  import { useActivePage } from '@/store/index'
  const activePage = useActivePage()
 
</script>

<style scoped></style>
