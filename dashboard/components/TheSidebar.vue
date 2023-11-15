<template>
  <aside
    class="sidebar relative grid h-screen grid-rows-[auto,auto,1fr,auto] gap-6 border-r-2 border-solid border-neutral-200 border-l-purple-950 bg-white p-5"
  >
    <section
      ref="sidebar-title"
      data-cy="sidebar__header"
      class="flex gap-3 border-l-purple-950"
    >
      <img
        src="/logos/sapawarga.svg"
        alt="Portal Jabar Logo"
        width="23"
        height="32"
        class="self-center"
      />
      <div>
        <h1
          class="whitespace-nowrap text-[15px] font-[700] leading-5 text-neutral-900"
        >
          Ticketing
        </h1>
        <p class="text-[12px] font-[400] tracking-wide text-neutral-500">
          Management System
        </p>
      </div>
    </section>
    <nav class="-ml-6 w-[calc(100%+48px)] overflow-hidden">
      <ul data-cy="sidebar__navigation" class="px-4">
        <li>
          <div
            v-for="navigation in NAVIGATION_MENU"
            :key="navigation.titleMenu"
          >
            <p
              v-if="!!navigation.titleMenu"
              class="my-3 ml-2 text-[11px] font-[500] uppercase tracking-widest text-[#A3A3A3]"
            >
              {{ navigation.titleMenu }}
            </p>
            <NuxtLink
              v-for="data in navigation.menu"
              :key="data.label"
              :to="data.link"
              class="sidebar__navigation-item mb-1 flex min-h-[50px] items-center rounded-lg p-[10px] text-sm font-[500] text-[#404040] hover:bg-[#f3f3f3]"
            >
              <NuxtIcon
                :name="data.icon"
                filled
                class="body_navigation stroke-[#404040] text-xl"
              />
              <span
                class="body_navigation ml-4 text-sm font-[500] text-[#404040]"
                >{{ data.label }}</span
              >
            </NuxtLink>
          </div>
        </li>
      </ul>
    </nav>
    <section class="absolute bottom-0 mt-auto w-full">
      <ul ref="sidebar-bottom-nav" data-cy="sidebar__bottom-nav">
        <li>
          <div
            class="mb-3 flex min-h-[50px] w-full items-center px-4 font-lato text-sm font-bold text-black"
          >
            <NuxtIcon name="images/profile" filled class="mr-3 text-4xl" />
            <div class="grow">
              <h1
                class="whitespace-nowrap text-[14px] font-[700] leading-5 text-neutral-900"
              >
                {{ session.user?.name }}
              </h1>
              <p class="text-[12px] font-[400] tracking-wide text-neutral-500">
                Superadmin
              </p>
            </div>
            <UDropdown
              :items="items"
              :ui="{ item: { disabled: 'cursor-text select-text' } }"
              :popper="{ placement: 'bottom-start' }"
            >
              <NuxtIcon
                name="navigation/dots-icon"
                filled
                class="justify-self-end text-lg"
              />
              <template #account="{ item }">
                <div class="text-left">
                  <p>Signed in as</p>
                  <p class="truncate font-medium text-gray-900 dark:text-white">
                    {{ item.label }}
                  </p>
                </div>
              </template>

              <template #item="{ item }">
                <span class="truncate">{{ item.label }}</span>

                <UIcon
                  :name="item.icon"
                  class="ms-auto h-4 w-4 flex-shrink-0 text-blue-400 dark:text-blue-600"
                />
              </template>
            </UDropdown>
          </div>
        </li>
      </ul>
    </section>
  </aside>
</template>

<script setup lang="ts">
  import { NAVIGATION_MENU } from '~/common/constant/navigation'
  const { signOut, getSession } = useAuth()
  const session = await getSession()

  const items = [
    [
      {
        label: session.user?.email,
        slot: 'account',
        disabled: true,
      },
    ],
    [
      {
        label: 'Settings',
        icon: 'i-heroicons-cog-8-tooth',
      },
    ],
    [
      {
        label: 'Sign out',
        icon: 'i-heroicons-arrow-left-on-rectangle',
        click: () => {
          signOut({ callbackUrl: '/login' })
        },
      },
    ],
  ]
</script>

<style scoped>
  .sidebar__navigation-item.router-link-exact-active {
    @apply bg-[#E3F2FD];
    .body_navigation {
      @apply text-blue-600;
      @apply stroke-blue-600;
    }
  }
</style>
