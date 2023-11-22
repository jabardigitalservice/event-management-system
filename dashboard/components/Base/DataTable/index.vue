<template>
  <div>
    <div class="mb-5 flex h-auto place-items-end justify-between align-middle">
      <UButton
        class="bg-[#1569C4]"
        icon="i-heroicons-pencil-square"
        size="md"
        color="blue"
        variant="solid"
        :label="`Tambah ${activePage.page}`"
        :trailing="false"
        :to="props.baseRoute + '/form'"
        @click="useIdData().id = ''"
      />
      <UInput
        v-model="search"
        :placeholder="'Search ' + activePage.page"
        icon="i-heroicons-magnifying-glass-20-solid"
        class="w-[386px]"
        :ui="{ icon: { trailing: { pointer: '' } } }"
      >
        <template #trailing>
          <UButton
            v-show="search !== ''"
            color="gray"
            variant="link"
            icon="i-heroicons-x-mark-20-solid"
            :padded="false"
            @click="search = ''"
          />
        </template>
      </UInput>
    </div>
    <div
      class="inline-block min-w-full overflow-hidden rounded-lg border-[2px] border-neutral-100 bg-white shadow-sm"
    >
      <UTable
        :loading="loading"
        :loading-state="{
          icon: 'i-heroicons-arrow-path-20-solid',
          label: 'Loading...',
        }"
        :columns="props.headers"
        :rows="itemTable"
      >
        <template #name-data="{ row }">
          <slot name="customeName" :items="row" />
        </template>
        <template #status-data="{ row }">
          <slot name="customeStatus" :items="row" />
        </template>
        <template #actions-data="{ row }">
          <slot name="customeAction" :items="row" :fetch="fetchData" />
        </template>
      </UTable>
      <div class="border-t-[1px] border-neutral-200 px-4">
        <div class="float-left mt-4 flex w-auto flex-row">
          <div class="mr-2 basis-1/5 self-center">
            <span class="text-sm text-gray-500 dark:text-white"
              >Menampilkan</span
            >
          </div>
          <div class="mr-2 basis-1/4">
            <USelectMenu
              v-model="selectedLimit"
              :options="listLimit"
              name="labels"
            />
          </div>
          <div class="basis-1/2 self-center">
            <span class="text-sm text-gray-500 dark:text-white"
              >Dari <span class="font-semibold">{{ total }}</span> item</span
            >
          </div>
        </div>

        <UPagination
          v-model="page"
          class="my-4 justify-end"
          :page-count="selectedLimit"
          :total="total"
          :value="page"
          :active-button="{ color: 'blue', variant: 'solid' }"
          :inactive-button="{ color: 'blue', variant: 'ghost' }"
          :ui="{
            wrapper: 'flex items-center gap-2',
            rounded: '!rounded-md min-w-[32px] justify-center',
            color: {
              blue: {
                solid: 'bg-white dark:bg-gray-900'
              }
            }
          }"
          @click="onClickPagination(page)"
        >
          <template #prev="{ onClick }">
            <UTooltip text="Previous page">
              <UButton
                variant="outline"
                :ui="{
                  rounded: 'rounded-md',
                  variant: {
                    outline:
                      'ring-1 ring-inset ring-current text-neutral-100 hover:bg-neutral-50 disabled:bg-transparent border-neutral-300 stroke-slate-800',
                  },
                }"
                class="me-2 rtl:[&_span:first-child]:rotate-180"
                @click="onClick"
              >
                <template #trailing>
                  <NuxtIcon
                    name="navigation/previous-icon"
                    filled
                    class="m-[6px] stroke-[#1569C4] text-[10px]"
                  />
                </template>
              </UButton>
            </UTooltip>
          </template>

          <template #next="{ onClick }">
            <UTooltip text="Next page">
              <UButton
                color="blue"
                variant="outline"
                :ui="{
                  rounded: 'rounded-md',
                  variant: {
                    outline:
                      'ring-1 ring-inset ring-current text-neutral-100 hover:bg-neutral-50 disabled:bg-transparent border-neutral-300 stroke-slate-800',
                  },
                }"
                class="ms-2 rtl:[&_span:last-child]:rotate-180"
                @click="onClick"
              >
                <template #trailing>
                  <NuxtIcon
                    name="navigation/next-icon"
                    filled
                    class="m-[6px] stroke-[#1569C4] text-[10px]"
                  />
                </template>
              </UButton>
            </UTooltip>
          </template>
        </UPagination>
      </div>
    </div>
    <UNotifications />
  </div>
</template>

<script setup lang="ts">
  import { useFetchData } from '~/composables/useFetchData'
  import { useActivePage, useIdData } from '@/store/index'

  interface ApiResponse {
    current_page: number
    data: object[]
    total: number
    to: number
    from: number
  }

  interface TableColumn {
    key: string
    sortable?: boolean
    direction?: 'desc' | 'asc'
    class?: string
  }

  const props = defineProps({
    headers: {
      type: Array as () => TableColumn[],
      default: () => [],
    },
    path: {
      type: String,
      default: '',
    },
    baseRoute: {
      type: String,
      default: '',
    },
  })

  const activePage = useActivePage()
  const search = ref('')
  const itemTable = ref()
  const page: Ref<number> = ref(1)
  const loading = ref(true)
  let total: number = 1
  const listLimit = [5, 10, 25, 30]
  const selectedLimit = ref(listLimit[0])

  async function fetchData() {
    loading.value = true
    const result = (await useFetchData(
      props.path,
      page,
      search,
      selectedLimit,
    )) as ApiResponse
    page.value = result.data.meta.page
    itemTable.value = result.data.data
    total = result.data.meta.total_data
    loading.value = false
  }

  function onClickPagination(getPage: number) {
    page.value = getPage
  }

  fetchData()

  watch([page, search], () => {
    fetchData()
  })
  watch([selectedLimit], () => {
    page.value = 1
    fetchData()
  })
</script>

<style>
  thead {
    @apply rounded-lg border-b border-gray-200 px-5 py-3 text-left text-sm font-semibold tracking-wider text-[#737373];
  }
  th > button > span {
    @apply text-[#737373];
  }
  th > span {
    @apply text-[#737373];
  }
  .form-input {
    @apply h-10;
  }
  .hover\:bg-gray-50:hover {
    @apply bg-[#FAFAFA];
  }
</style>
