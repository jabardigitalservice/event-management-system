<template>
  <Menu v-slot="{ open }" as="div" class="w-fit">
    <div class="mr-2 grid grid-cols-[34px,_minmax(0,_1fr)] gap-4">
      <!-- @todo: change this user avatar to be dynamic -->
      <img
        data-cy="header__user-avatar"
        :src="userAvatar"
        alt="user avatar"
        width="34"
        height="34"
        class="h-[34px] w-[34px] max-w-full rounded-full border-2 border-green-700 bg-gray-500 object-cover object-center"
      />
      <MenuButton
        data-cy="header__user-dropdown-button"
        class="inline-flex w-full items-center"
      >
        <!-- @todo: change this user name to be dynamic -->
        <span
          class="font-lato text-blue-gray-800 mr-2 line-clamp-1 max-w-[100px] text-left text-sm font-semibold"
        >
          Superadmin
        </span>
        <NuxtIcon
          name="common/chevron-right"
          :class="{
            'eas pointer-events-none text-xl text-green-700 transition-transform duration-150 ease-in': true,
            'rotate-90': open,
          }"
          aria-hidden="true"
        />
      </MenuButton>
    </div>

    <transition
      enter-active-class="transition duration-100 ease-out"
      enter-from-class="transform scale-95 opacity-0"
      enter-to-class="transform scale-100 opacity-100"
      leave-active-class="transition duration-75 ease-in"
      leave-from-class="transform scale-100 opacity-100"
      leave-to-class="transform scale-95 opacity-0"
    >
      <MenuItems
        data-cy="header__user-dropdown-menu"
        class="absolute right-6 mt-2 w-[170px] origin-top-right rounded-lg bg-white shadow-lg"
      >
        <ul class="p-2">
          <MenuItem as="li">
            <NuxtLink
              data-cy="header__setting-menu-button"
              :to="`${config.public.portalJabarCMSBaseURL}/pengaturan`"
              class="font-lato group flex w-full items-center gap-2 rounded-md p-2 text-sm font-medium text-gray-800 hover:bg-green-50 hover:text-green-700"
            >
              <NuxtIcon
                name="navigation/account-settings-icon"
                class="text-lg text-gray-800 group-hover:text-green-700"
              />
              Pengaturan Akun
            </NuxtLink>
          </MenuItem>
          <!-- @todo: add logout popup modal -->
          <MenuItem as="li">
            <button
              data-cy="header__user-logout-button"
              class="group w-full rounded-md p-2 text-gray-800 hover:bg-green-50"
              @click="signOut({ callbackUrl: '/login' })"
            >
              <p
                class="font-lato flex items-center gap-2 text-sm font-medium group-hover:text-green-700"
              >
                <NuxtIcon
                  name="common/logout"
                  class="text-lg text-gray-800 group-hover:text-green-700"
                  aria-hidden="true"
                />
                Keluar
              </p>
            </button>
          </MenuItem>
        </ul>
      </MenuItems>
    </transition>
  </Menu>
</template>

<script setup lang="ts">
  import { Menu, MenuButton, MenuItems, MenuItem } from '@headlessui/vue'
  const { signOut } = useAuth()

  const userAvatar = 'https://placehold.co/34'
  const config = useRuntimeConfig()
</script>
