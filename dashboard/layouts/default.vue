<template>
  <div>
    <Head>
      <Title>Event Management System</Title>
    </Head>
    <div class="grid h-full min-w-0 overflow-hidden grid-cols-[260px,1fr]">
      <TheSidebar />
      <main class="h-screen overflow-y-auto bg-gray-100">
        <div class="w-full mt-10 px-24">
          <p class="font-[500] text-[22px] mb-6">{{ activePage.page }}</p>
            <slot />
        </div>
      </main>
    </div>
  </div>
</template>
<script setup lang="ts">
  import { useActivePage } from '@/store/index'
  const activePage = useActivePage()
  
</script>
<style>
  .nuxt-icon svg {
    @apply mb-0;
  }
</style>
