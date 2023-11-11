<template>
  <div>
    <div class="mb-5 flex h-auto place-items-end justify-between align-middle">
      <UInput
        v-model="search"
        :placeholder="'Search ' + activePage.page"
        icon="i-heroicons-magnifying-glass-20-solid"
        class="w-1/5"
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
      <UButton
        class="bg-green-700"
        icon="i-heroicons-pencil-square"
        size="md"
        color="primary"
        variant="solid"
        label="Create"
        :trailing="false"
        :to="props.baseRoute + '/form'"
      />
    </div>
    <div class="inline-block min-w-full overflow-hidden rounded-lg bg-white">
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
    </div>
    <div class="float-left mt-4 flex w-56 flex-row">
      <div class="basis-1/5 self-center">
        <span class="text-sm font-semibold text-gray-500 dark:text-white"
          >Show</span
        >
      </div>
      <div class="mr-2 basis-1/4">
        <USelectMenu v-model="selectedLimit" :options="listLimit" />
      </div>
      <div class="basis-1/2 self-center">
        <span class="text-sm font-semibold text-gray-500 dark:text-white"
          >Entries from {{ total }}</span
        >
      </div>
    </div>
    <UPagination
      v-model="page"
      class="mb-10 mt-5 justify-end"
      :page-count="selectedLimit"
      :total="total"
      :value="page"
      :prev-button="{
        icon: 'i-heroicons-arrow-small-left-20-solid',
        label: 'Prev',
        color: 'gray',
      }"
      :next-button="{
        icon: 'i-heroicons-arrow-small-right-20-solid',
        trailing: true,
        label: 'Next',
        color: 'gray',
      }"
      @click="onClickPagination(page)"
    />
    <UNotifications />
  </div>
</template>

<script setup lang="ts">
  import { useFetchData } from '~/composables/useFetchData'
  import { useActivePage } from '@/store/index'

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
    @apply rounded-lg border-b-2 border-gray-200 bg-green-700 px-5 py-3 text-left text-sm font-semibold tracking-wider text-white;
  }
  th > button > span {
    @apply text-white;
  }
  th > span {
    @apply text-white;
  }
  .form-input {
    @apply h-10;
  }
  .hover\:bg-gray-50:hover {
    @apply bg-green-800;
  }
</style>
