<template>
  <div class="h-screen overflow-hidden flex flex-col">
    <header
      class="h-[64px] flex w-full items-center bg-white px-3 shadow-md sticky">
      <UButton
        :padded="false"
        color="gray"
        variant="link"
        class="mr-2"
        @click="handleBack"
      >
        <NuxtIcon
          name="navigation/close-icon"
          filled
          class="justify-self-end text-4xl"
        />
      </UButton>
      <img
          src="/logos/sapawarga.svg"
          alt="Portal Jabar Logo"
          width="20"
          height="28"
          class="self-center mr-3"
        />
      <div class="ml-2 grow flex">
        <nav class="flex flex-wrap items-center">
          <div class="inline-block">
            <div
              class="text-[16px] font-inter text-[#171717] mr-4 text-md font-bold">
              {{ activePage.page }}
            </div>
          </div>
        </nav>
      </div>
      <slot name="buttonSave" />
    </header>
    <div class="xl:px-[444px] lg:px-[296px] md:px-[148px] px-[66px] flex flex-col h-full overflow-y-auto py-5 max-h-screen">
      <slot name="body" />
    </div>
  </div>
</template>

<script setup lang="ts">
  import { useActivePage } from '@/store/index'
  const activePage = useActivePage()
  const router = useRouter()
  const handleBack = () => {
    router.go(-1)
  }
 
</script>

<style scoped></style>
