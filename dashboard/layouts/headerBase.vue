<template>
    <div>
      <div class="grid h-full min-w-0 overflow-hidden">
        <main class="h-screen overflow-y-auto bg-gray-100">
          <div class="w-full">
              <slot />
          </div>
        </main>
      </div>
    </div>
  </template>
  <style>
    .nuxt-icon svg {
      @apply mb-0;
    }
  </style>
  