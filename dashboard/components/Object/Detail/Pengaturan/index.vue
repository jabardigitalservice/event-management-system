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
    <div class="mt-7 flex w-screen items-center">
      <p class="mr-[2px] text-lg font-[600] capitalize">Status</p>
      <NuxtIcon
        name="common/help"
        filled
        class="mr-2 mt-px justify-self-end text-lg"
      />
      <UBadge
        class="self-center capitalize"
        size="xs"
        :color="statusColors[state.items.status].color"
        variant="soft"
        :ui="{
          rounded: 'rounded-full',
          font: 'font-small',
          size: {
            xs: 'text-[12px] px-2 pt-1 h-6',
          },
          variant: {
            soft: `${
              statusColors[state.items.status].background
            }  text-{color}-500`,
          },
        }"
        >{{ state.items.status }}</UBadge
      >
    </div>
    <div class="mt-4 flex justify-between">
      <p class="w-[70%]">
        Tiket belum dibuat, silahkan buat tiket terlebih dahulu untuk bisa
        diterbitkan pada layanan aplikasi Sapawarga.
      </p>
      <UButton
        class="pointer-events-none h-9"
        variant="solid"
        :ui="{
          rounded: 'rounded-lg',
          variant: { solid: 'bg-[#F5F5F5] text-[#A3A3A3] hover:bg-[#F9F9F9]' },
        }"
      >
        Published</UButton
      >
    </div>
    <div class="mt-4 flex justify-between">
      <p class="w-[70%]">
        Objek wisata sudah diterbitkan, dan sudah bisa diakses pada layanan
        aplikasi Sapawarga.
      </p>
      <UButton
        class="pointer-events-none h-9"
        color="red"
        variant="solid"
        :ui="{ rounded: 'rounded-lg', variant: { solid: 'bg-[#DC2626]' } }"
      >
        Unpublish</UButton
      >
    </div>
    <div class="mt-4 flex justify-between">
      <p class="w-[70%]">
        Objek wisata belum diterbitkan, dan belum bisa diakses pada layanan
        aplikasi Sapawarga.
      </p>
      <UButton
        class="pointer-events-none h-9"
        color="green"
        variant="solid"
        :ui="{ rounded: 'rounded-lg', variant: { solid: 'bg-[#108F5A]' } }"
      >
        Publish</UButton
      >
    </div>
    <hr class="my-8 h-[2px] w-full border-0 bg-gray-200 dark:bg-gray-700" />
    <p class="mr-2 text-lg font-[600] capitalize">Hapus Objek Wisata</p>
    <div class="mt-4 flex justify-between">
      <p class="w-[70%]">
        Apakah Anda yakin akan menghapus objek wisata ini? Objek wisata yang
        telah dihapus tidak dapat dikembalikan.
      </p>
      <UButton
        class="h-9"
        color="red"
        variant="ghost"
        :ui="{ rounded: 'rounded-lg' }"
        @click="
          () => {
            state.isOpenDelete = true
          }
        "
      >
        Hapus</UButton
      >
    </div>
  </div>
  <BaseModal
    :open-modal="state.isOpenDelete"
    title-modal="Confirm to delete"
    desc-modal="Are you sure you want to delete your data ? "
    icon-modal="i-heroicons-exclamation-triangle"
    text-confirm="Delete"
    type-modal="danger"
    @close="state.isOpenDelete = false"
    @confirm="deleteData()"
  />
  <UNotifications />
</template>
<script setup lang="ts">
  interface apiDataResponse {
    id: string
    name: string
    banner: object[]
    logo: string
    status: string
  }
  interface statusColor {
    color: string
    background: string
  }

  const statusColors: Record<string, statusColor> = {
    draft: { color: 'orange', background: 'bg-orange-100' },
    published: { color: 'green', background: 'bg-green-50' },
    unpublished: { color: 'red', background: 'bg-red-100' },
  }

  const router = useRouter()
  const urlAPI: string = '/v1/event/object'
  const toast = useToast()
  const route = useRoute()

  const state = reactive({
    items: {} as apiDataResponse,
    isLoading: true,
    isOpenDelete: false,
  })

  const props = defineProps({
    data: {
      type: Object,
      default: () => ({}),
    },
  })

  function deleteData() {
    let cancelled = false
    toast.add({
      icon: 'i-heroicons-exclamation-triangle',
      title: 'data will be deleted',
      color: 'yellow',
      timeout: 2000,
      callback: async () => {
        if (!cancelled) {
          await useDeleteData(urlAPI, route.query.id)
          toast.add({
            title: 'data successfully deleted',
            icon: 'i-heroicons-check-circle',
            timeout: 1000,
          })
          router.push({ path: '/object' })
        }
      },
      actions: [
        {
          label: 'Undo',
          variant: 'solid',
          color: 'orange',
          click: () => {
            cancelled = true
            toast.add({
              title: 'data cancel deleted',
              color: 'red',
              icon: 'i-heroicons-x-circle',
              timeout: 1000,
            })
          },
        },
      ],
    })
    state.isOpenDelete = false
  }

  onMounted(async () => {
    state.items = await props.data
    state.isLoading = false
  })
</script>
