<template>
  <div>
    <Head>
      <Title>Event Management System</Title>
    </Head>
    <div class="grid h-screen min-w-0 grid-cols-[260px,1fr] overflow-hidden">
      <TheSidebar />
      <div class="w-full">
        <TheHeader />
        <main class="h-[calc(100vh-72px)] overflow-y-auto bg-gray-100">
          <div class="w-full pl-16 pr-16 mt-5">
              <div class="pb-[25px]">
                <nav class="flex flex-wrap items-center h-[45px]">
                  <div class="inline-block">
                    <NuxtLink
                      class="font-roboto text-[28px] font-bold text-blue-gray-700 breadcrumb__item"
                    >
                      {{ activePage.page }}
                    </NuxtLink>
                  </div>
                </nav>
                <p class="text-[16px] font-roboto mt-1 text-blue-gray-800 ">{{ activePage.desc }}</p>
              </div>
              <slot />
          </div>
        </main>
      </div>
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
