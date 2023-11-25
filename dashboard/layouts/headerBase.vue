<template>
    <div>
      <Head>
        <Title>Event Management System</Title>
      </Head>
      <div class="grid h-full min-w-0 overflow-hidden">
        <main class="h-screen overflow-y-auto bg-gray-100">
          <div class="w-full">
              <slot />
          </div>
        </main>
      </div>
    </div>
  </template>
  <script setup lang="ts">
  </script>
  <style>
    .nuxt-icon svg {
      @apply mb-0;
    }
  </style>
  