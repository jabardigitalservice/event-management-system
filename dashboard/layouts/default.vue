<template>
  <div>
    <Head>
      <Title>Event Management System</Title>
    </Head>
    <div :class="`grid h-full min-w-0 overflow-hidden ${state.styleGridSidebar}`">
      <TheSidebar v-if="!activePage.navigation"/>
      <main :class="`h-screen overflow-y-auto ${state.bgColor}`">
        <div :class="`w-full ${state.stylePaddingSidebar}`">
          <p v-if="!activePage.navigation" class="font-[500] text-[22px] mb-6">{{ activePage.page }}</p>
            <slot />
        </div>
      </main>
    </div>
  </div>
</template>
<script setup lang="ts">
  import { useActivePage } from '@/store/index'
  const activePage = useActivePage()
  const state = reactive({
    styleGridSidebar: '',
    stylePaddingSidebar: '',
    bgColor: ''
  })
  
  watch([activePage], ()=>{
    state.styleGridSidebar = !activePage.navigation? 'grid-cols-[260px,1fr]': ''
    state.bgColor = !activePage.navigation? 'bg-white': 'bg-gray-100'
    state.stylePaddingSidebar = !activePage.navigation? 'mt-10 px-24': ''
  })
  
</script>
<style>
  .nuxt-icon svg {
    @apply mb-0;
  }
</style>
