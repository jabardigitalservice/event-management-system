<template>
  <div
    v-if="state.isLoading"
    role="status"
    class="absolute left-1/2 top-[35%] -translate-x-1/2 -translate-y-1/2"
  >
    <div class="flex">
      <div class="relative">
        <div
          class="absolute h-10 w-10 rounded-full border-4 border-solid border-gray-200"
        ></div>
        <div
          class="absolute h-10 w-10 animate-spin rounded-full border-4 border-solid border-blue-500 border-t-transparent shadow-md"
        ></div>
      </div>
    </div>
  </div>
  <div v-if="!state.isLoading" class="pb-14 text-[14px] font-[400]">
    <img :src="state.items.banner" class="mt-5 h-[264px] w-full rounded-md" />
    <div class="mt-[30px] flex w-full items-center bg-white">
      <UAvatar
        icon="i-heroicons-photo"
        class="h-12 w-12 border border-gray-300 object-none"
        :src="state.items.logo"
        size="md"
        alt="Logo"
      />
      <div class="ml-3 grow">
        <h1
          class="whitespace-nowrap text-[18px] font-[600] capitalize leading-5 text-neutral-900"
        >
          {{ state.items.name }}
        </h1>
        <UBadge
          color="gray"
          variant="soft"
          size="xs"
          class="bg-gray-100 text-[10px]"
          >Museum Ilmu Sejarah</UBadge
        >
      </div>
      <UButton
        class="h-9"
        color="blue"
        variant="outline"
        :ui="{ rounded: 'rounded-lg', 
        variant: { 
          outline: 'text-[#1569C4]'
          }
        }"
        @click="
          router.push({ path: '/object/form', query: { id: state.items.id } })
        "
      >
        <NuxtIcon
          name="common/pencil"
          filled
          class="justify-self-end text-xl"
        />
        Ubah</UButton
      >
    </div>
    <p class="mt-6">
      {{ state.items.description }}
    </p>
    <div class="mt-[15px] flex w-full">
      <NuxtIcon
        name="common/map-icon"
        filled
        class="mr-2 justify-self-end text-xl"
      />
      <p class="capitalize">
        {{
          `${state.items.address}, kel. ${state.items.village}, kec. ${state.items.district}, Kota ${state.items.city}, ${state.items.province}`
        }}
      </p>
    </div>

    <div class="mt-5 text-[16px] font-[600]">Pengelola</div>
    <p class="mt-2 capitalize">{{ state.items.organizer }}</p>
    <div class="mt-2 flex w-full">
      <NuxtIcon
        name="common/mail"
        filled
        class="mr-2 justify-self-end text-xl"
      />
      <p class="capitalize">{{ state.items.organizer_email }}</p>
    </div>
    <div class="mt-2 flex w-full">
      <NuxtIcon
        name="common/phone"
        filled
        class="mr-2 justify-self-end text-xl"
      />
      <p class="capitalize">{{ state.items.organizer_phone }}</p>
    </div>
    <div class="mt-5 text-[16px] font-[600]">Sosial Media</div>
    <div
      v-for="rows in state.items.social_media"
      :key="rows.name"
      class="mt-2 flex w-full"
    >
      <NuxtIcon
        :name="iconSocialMedia(rows.name)"
        filled
        class="mr-2 justify-self-end text-xl"
      />
      <p class="capitalize">{{ rows.link }}</p>
    </div>
  </div>
</template>
<script setup lang="ts">
  interface apiDataResponse {
    id: string
    name: string
    banner: object[]
    logo: string
    organizer: string
    organizer_email: string
    organizer_phone: string
    social_media: object[]
    description: string
    address: string
    province: string
    province_id: string
    city: string
    city_id: string
    district: string
    district_id: string
    village: string
    village_id: string
    google_map: string
    status: string
    organization_id: string
  }
  const router = useRouter()
  const state = reactive({
    items: {} as apiDataResponse,
    isLoading: true,
  })

  const props = defineProps({
    data: {
      type: Object,
      default: () => ({}),
    },
  })

  function iconSocialMedia(label: string) {
    let icon: string = ''
    switch (label) {
      case 'youtube':
        icon = 'common/youtube'
        break
      case 'instagram':
        icon = 'common/instagram'
        break
      case 'x':
        icon = 'common/twitter'
        break
      case 'facebook':
        icon = 'common/facebook'
        break
      case 'tiktok':
        icon = 'common/tiktok'
        break
    }
    return icon
  }

  onMounted(async () => {
    state.items = await props.data
    state.isLoading = false
  })
</script>
