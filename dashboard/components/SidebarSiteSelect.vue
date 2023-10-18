<template>
  <Listbox v-model="selected" as="div">
    <div class="relative">
      <ListboxButton
        class="relative flex w-full cursor-default items-center rounded-md bg-[#35B472] px-4 py-3 text-left font-lato text-sm text-white shadow-sm"
      >
        <NuxtIcon
          name="common/globe"
          class="mr-3 text-xl"
          filled
          aria-hidden="true"
        />
        <div class="flex flex-col">
          <span class="text-xxs leading-none">Situs saya</span>
          <span class="block w-full truncate text-sm font-bold">
            {{ selected.name }}
          </span>
        </div>
        <NuxtIcon
          name="common/chevron-down"
          class="pointer-events-none absolute right-0 pr-3 text-xl"
          filled
          aria-hidden="true"
        />
      </ListboxButton>

      <transition
        leave-active-class="transition ease-in duration-100"
        leave-from-class="opacity-100"
        leave-to-class="opacity-0"
      >
        <ListboxOptions
          class="absolute z-10 mt-1 max-h-56 w-full overflow-auto rounded-md bg-[#35B472] py-1 font-lato text-sm font-normal text-white shadow-lg"
        >
          <ListboxOption
            v-for="site in availableSites"
            :key="site.id"
            v-slot="{ selected }"
            as="template"
            :value="site"
          >
            <li class="relative cursor-pointer select-none py-2 pl-3 pr-9">
              <div class="flex items-center">
                <span class="ml-3 block truncate">
                  {{ site.name }}
                </span>
                <NuxtIcon
                  v-show="selected"
                  name="common/check-circle"
                  class="absolute inset-y-0 right-0 flex items-center pr-4 text-xl"
                  filled
                  aria-hidden="true"
                />
              </div>
            </li>
          </ListboxOption>
          <div class="mt-2 flex justify-center border-t border-[#5BC08B] py-3">
            <!-- @todo: change this link href -->
            <NuxtLink to="/" class="inline-flex">
              <NuxtIcon
                name="navigation/account-settings-icon"
                class="mr-2 text-lg"
                filled
                aria-hidden="true"
              />
              <span class="font-bold">Kelola Event CMS</span>
            </NuxtLink>
          </div>
        </ListboxOptions>
      </transition>
    </div>
  </Listbox>
</template>

<script setup>
  import { ref } from 'vue'
  import {
    Listbox,
    ListboxButton,
    ListboxOption,
    ListboxOptions,
  } from '@headlessui/vue'

  // @todo: change avaiable site dynamically
  const availableSites = [
    {
      id: 1,
      name: 'Portal Jabar 1',
    },
    {
      id: 2,
      name: 'Portal Jabar 2',
    },
  ]

  const selected = ref(availableSites[0])
</script>
