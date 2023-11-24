<template>
  <TheHeader>
    <template #body>
      <div class="ml-24 mt-7 flex flex-row">
        <p class="mr-2 text-lg font-[600] capitalize">{{ state.dataDetail.name }}</p>
        <UBadge
          class="self-center capitalize"
          size="xs"
          :color="statusColors[state.dataDetail.status].color"
          variant="soft"
          :ui="{
            rounded: 'rounded-full',
            font: 'font-small',
            size: {
              xs: 'text-[10px] px-2 pt-1 h-5',
            },
            variant: {
              soft: `${statusColors[state.dataDetail.status].background}  text-{color}-500`,
            },
          }"
          >{{ state.dataDetail.status }}</UBadge
        >
      </div>
      <UTabs
        class="mt-3"
        :items="items"
        :ui="{
          base: 'w-[45px]',
          list:{
            width: 'w-[270px]',
            padding: 'pl-[86px]',
            marker: {
              base: 'w-[45px]',
              wrapper: 'absolute top-[4px] left-[4px] duration-200 ease-out focus:outline-none focus:bg-transparent',
              shadow: 'shadow-none'
          },
            tab: {
              base: 'w-fit',
              background: 'bg-gray-100 dark:bg-gray-800',
              active: 'text-[#1569C4]',
            },
          }
        }"
      >
        <template #default="{ item, selected }">
          <div
            v-if="selected"
            class="align-center relative flex items-center border-b-2 py-2 border-[#1569C4] bg-transparent"
          >
            <span class="align-center"> {{ item.label }}</span>
          </div>
        </template>
        <template #profile>
          <div class="w-screen bg-white h-screen">
            <div class="flex px-3 overflow-y-scroll no-scrollbar h-[82vh] -mt-2 mb-20 place-content-center">
              <div class="max-w-xl flex flex-col py-5">
                <ObjectDetailProfile :data="fetchData()" />
              </div>
            </div>
          </div>
        </template>
        <template #tiket>
          <div class="w-screen bg-white h-screen">
            <div class="flex px-3 overflow-y-scroll no-scrollbar h-[82vh] -mt-2 mb-20 place-content-center">
              <div class="max-w-xl flex flex-col py-5">
                <ObjectDetailTiket />
              </div>
            </div>
          </div>
        </template>
        <template #pengaturan>
          <div class="w-screen bg-white h-screen">
            <div class="flex px-3 overflow-y-scroll no-scrollbar h-[82vh] -mt-2 mb-20 place-content-center">
              <div class="max-w-xl flex flex-col py-5">
                <ObjectDetailPengaturan :data="fetchData()" />
              </div>
            </div>
          </div>
        </template>
      </UTabs>
    </template>
  </TheHeader>
</template>
<script setup lang="ts">
  import { useActivePage } from '@/store/index'
  import { useGetData } from '~/composables/useFetchData'

  const items = [{
    slot: 'profile',
    label: 'Profil'
  }, {
    slot: 'tiket',
    label: 'Tiket'
  }, {
    slot: 'pengaturan',
    label: 'Pengaturan'
  }]

  interface apiDataResponse {
    data: {
      id: string
      name: string
      banner: object
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
      status: string,
      organization_id: string
    }
  }

  interface statusColor {
    color: string,
    background: string
  }

  const statusColors: Record<string, statusColor> = {
    draft: { color: 'orange', background: 'bg-orange-100' },
    published: { color: 'green', background: 'bg-green-50' },
    unpublished: { color: 'red', background: 'bg-red-100' },
  }

  const router = useRoute()

  const state = reactive({
    dataDetail: {} as apiDataResponse['data']
  })
  
  async function fetchData() {
    const responseObject = (await useGetData(
      `/v1/event/object/${router.query.id}`,
    )) as apiDataResponse
    
    return responseObject.data
  }

  state.dataDetail = await fetchData()

  definePageMeta({
    middleware: ['auth'],
  })
  
  onMounted(() => {
    const activePage = useActivePage()
    activePage.page = 'Detail Objek Wisata'
    activePage.navigation = true
  })
</script>
