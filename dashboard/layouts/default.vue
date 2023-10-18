<template>
  <div class="grid h-screen min-w-0 grid-cols-[260px,1fr] overflow-hidden">
    <TheSidebar />
    <div class="w-full">
      <TheHeader />
      <main class="h-[calc(100vh-72px)] overflow-y-auto bg-gray-100">
        <slot />
      </main>
    </div>
  </div>
</template>

<style>
  .nuxt-icon svg {
    @apply mb-0;
  }
</style>
