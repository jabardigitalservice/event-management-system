<template>
  <aside
    data-cy="sidebar__container"
    class="sidebar grid h-screen grid-rows-[auto,auto,1fr,auto] gap-8 bg-green-600 p-6"
  >
    <section ref="sidebar-title" data-cy="sidebar__header" class="flex gap-2">
      <img
        src="/logos/logo.svg"
        alt="Portal Jabar Logo"
        width="34"
        height="34"
        class="self-start"
      />
      <div>
        <h1
          data-cy="sidebar__header-title"
          class="font-lato whitespace-nowrap text-sm font-bold leading-5 text-white"
        >
          EVENT MANAGEMENT<br />
          SYSTEM
        </h1>
      </div>
    </section>
    <SidebarSiteSelect />
    <nav class="-ml-6 w-[calc(100%+48px)] overflow-hidden">
      <ul data-cy="sidebar__navigation" class="px-6">
        <li>
          <NuxtLink
            v-for="navigation in NAVIGATION_MENU"
            :key="navigation.label"
            :to="navigation.link"
            class="sidebar__navigation-item font-lato mb-2 flex min-h-[50px] items-center rounded-lg p-[15px] text-sm font-bold text-white hover:bg-green-700"
          >
            <NuxtIcon :name="navigation.icon" filled class="text-xl" />
            <span class="ml-3">{{ navigation.label }}</span>
          </NuxtLink>
        </li>
      </ul>
    </nav>
    <section class="mt-auto">
      <ul ref="sidebar-bottom-nav" data-cy="sidebar__bottom-nav">
        <li>
          <NuxtLink
            :to="`${config.public.portalJabarCMSBaseURL}/pengaturan`"
            class="font-lato mb-2 flex min-h-[50px] w-full items-center rounded-lg p-[15px] text-sm font-bold text-white hover:bg-green-700"
          >
            <NuxtIcon
              name="navigation/account-settings-icon"
              filled
              class="text-xl"
            />
            <span class="ml-3">Pengaturan Akun</span>
          </NuxtLink>
        </li>
      </ul>
    </section>
  </aside>
</template>

<script setup lang="ts">
  import { NAVIGATION_MENU } from '~/common/constant/navigation'

  const config = useRuntimeConfig()
</script>

<style scoped>
  .sidebar__navigation-item.router-link-exact-active {
    @apply bg-green-700;
  }
</style>
