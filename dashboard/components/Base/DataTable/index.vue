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
        <template #status-data="{ row }">
          <slot name="customeStatus" :data-row="row" />
        </template>
        <template #photo-data="{ row }">
          <slot name="customePhoto" :data-row="row" />
        </template>
        <template #actions-data="{ row }">
          <UDropdown :items="itemActions(row)">
            <UButton
              color="green"
              variant="outline"
              icon="i-heroicons-chevron-down"
              >Actions
            </UButton>
          </UDropdown>
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
        <USelectMenu v-model="selectedLimit" :options="listLimitOptions" />
      </div>
      <div class="basis-1/2 self-center">
        <span class="text-sm font-semibold text-gray-500 dark:text-white"
          >Entries from {{ total }}</span
        >
      </div>
    </div>
    <UPagination
      v-model="currentPage"
      class="mb-10 mt-5 justify-end"
      :page-count="5"
      :total="total"
      :value="currentPage"
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
      @click="onClickPagination(currentPage)"
    />
    <BaseModal
      :open-modal="isOpenDelete"
      title-modal="Confirm to delete"
      :desc-modal="descModalDelete"
      icon-modal="i-heroicons-exclamation-triangle"
      text-confirm="Delete"
      type-modal="danger"
      @close="isOpenDelete = false"
      @confirm="deleteData()"
    />
    <BaseModal
      :open-modal="isOpenConfirm"
      title-modal="Confirm to publish"
      desc-modal="Are you sure you want to publish object"
      icon-modal="i-heroicons-question-mark-circle"
      text-confirm="Publish"
      type-modal="warning"
      @close="isOpenConfirm = false"
      @confirm="updateStatus()"
    />
    <UNotifications />
  </div>
</template>

<script setup lang="ts">
  import { useFetchData, useDeleteData } from '~/composables/useFetchData'
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
    actions: {
      type: Object,
      default: () => ({}),
    },
    path: {
      type: String,
      default: '',
    },
    baseRoute: {
      type: String,
      default: '',
    }
  })

  const toast = useToast()
  const activePage = useActivePage()
  const search = ref('')
  const itemTable = ref()
  const idRow = ref(1)
  const statusPublish = ref('')
  const page: Ref<number> = ref(1)
  const isOpenDelete = ref(false)
  const isOpenConfirm = ref(false)
  const loading = ref(true)
  let currentPage: number = 1
  let total: number = 1
  const listLimit = [5, 10, 25, 30]
  const selectedLimit = ref(listLimit[0])
  const listLimitOptions = listLimit.map((value) => ({
    label: value.toString(),
    value: value.toString(),
  }))
  const descModalDelete =
    'Are you sure you want to delete your data? All of your data will be permanently removed.'

  async function fetchData() {
    loading.value = true
    const result = (await useFetchData(props.path, page, search)) as ApiResponse
    currentPage = result.current_page
    itemTable.value = result.data
    total = result.total
    loading.value = false
  }

  const itemActions = (row: { id: number; status: string }) => {
    let dataItems = JSON.parse(JSON.stringify(props.actions))

    const getIndexAndHandleClick = (label: string, onClick: () => void) => {
      const index = dataItems.findIndex(
        (data: { label: string }) => data.label === label,
      )

      if (index !== -1) {
        dataItems[index].click = onClick
      }
    }

    getIndexAndHandleClick('Delete', () => openModalDelete(row.id))
    getIndexAndHandleClick('Unpublish', () =>
      openModalStatus(row.id, 'publish'),
    )
    getIndexAndHandleClick('Publish', () =>
      openModalStatus(row.id, 'unpublish'),
    )

    if (row.status === 'draft') {
      dataItems = dataItems.filter(
        (data: { label: string }) => data.label !== 'Unpublish',
      )
    } else if (row.status === 'published') {
      dataItems = dataItems.filter(
        (data: { label: string }) => data.label !== 'Publish',
      )
    }
    return [dataItems]
  }

  function deleteData() {
    let cancelled = false
    toast.add({
      icon: 'i-heroicons-exclamation-triangle',
      title: 'data will be deleted',
      color: 'yellow',
      timeout: 2000,
      callback: async () => {
        if (!cancelled) {
          await useDeleteData(props.path, idRow)
          fetchData()
          toast.add({
            title: 'data successfully deleted',
            icon: 'i-heroicons-check-circle',
            timeout: 1000,
          })
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
    isOpenDelete.value = false
  }

  function updateStatus() {
    toast.add({
      icon: 'i-heroicons-exclamation-triangle',
      title: 'data successfully Update',
      color: 'green',
      timeout: 1000,
    })
    isOpenConfirm.value = false
  }

  function onClickPagination(getPage: number) {
    page.value = getPage
  }

  function openModalDelete(row: number) {
    idRow.value = row
    isOpenDelete.value = true
  }

  function openModalStatus(row: number, status: string) {
    statusPublish.value = status
    idRow.value = row
    isOpenConfirm.value = true
  }

  fetchData()

  watch([page, search], () => {
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
