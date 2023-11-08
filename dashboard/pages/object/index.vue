<template>
  <div>
    <BaseDataTable :headers="headerTable" :path="urlAPI" base-route="/object">
      <template #customeName="{ items }">
        <div class="flex flex-row">
          <div class="basis-[14%]">
            <img
              :src= items.logo
              :alt=items.name
              width="20"
              height="20"
              class="self-start"
            />
          </div>
          <div class="basis-1/2">
            {{ items.name }}
          </div>
        </div>
      </template>
      <template #customeStatus="{ items }">
        <UBadge
          :color="statusColors[items.status].color"
          :variant="statusColors[items.status].variant"
          class="uppercase"
          >{{ items.status }}
        </UBadge>
      </template>
      <template #customeAction="{ items, fetch }">
        <UDropdown :items="itemActions(items, fetch)">
          <UButton
            color="green"
            variant="outline"
            icon="i-heroicons-chevron-down"
            >Actions
          </UButton>
        </UDropdown>
      </template>
    </BaseDataTable>
    <BaseModal
      :open-modal="isOpenDelete"
      title-modal="Confirm to delete"
      desc-modal="Are you sure you want to delete your data ? "
      icon-modal="i-heroicons-exclamation-triangle"
      text-confirm="Delete"
      type-modal="danger"
      @close="isOpenDelete = false"
      @confirm="deleteData()"
    />
    <BaseModal
      :open-modal="isOpenObject"
      :title-modal="`Confirm to ${titleObject}`"
      :desc-modal="`Are you sure you want to ${titleObject}`" 
      icon-modal="i-heroicons-question-mark-circle"
      :text-confirm="titleObject"
      type-modal="warning"
      @close="isOpenObject = false"
      @confirm="updateStatus()"
    />
  </div>
</template>

<script setup lang="ts">
  import { useActivePage } from '@/store/index'
  import { objectHeaders } from '~/common/constant/object'

  interface StatusColor {
    color: string
    variant: string
  }

  const toast = useToast()
  const isOpenDelete = ref(false)
  const isOpenObject = ref(false)
  const statusObject = ref('')
  const titleObject = ref('')
  const idItems = ref('')
  const fetchObject = ref()
  const headerTable = objectHeaders
  const urlAPI: string = '/v1/event/object'

  const statusColors: Record<string, StatusColor> = {
    draft: { color: 'red', variant: 'subtle' },
    published: { color: 'green', variant: 'subtle' },
    unpublished: { color: 'orange', variant: 'subtle' },
  }
  const itemActions = (
    items: { id: string; status: string },
    fetch: object,
  ) => {
    let actions = [
      [
        {
          label: 'Edit',
          icon: 'i-heroicons-pencil-square-20-solid',
          iconClass: 'bg-green-500',
        },
        {
          label: 'Detail',
          icon: 'i-heroicons-document-magnifying-glass-20-solid',
          iconClass: 'bg-green-500',
        },
      ],
      [
        {
          label: 'Delete',
          icon: 'i-heroicons-trash-20-solid',
          iconClass: 'bg-green-500',
          click: () => openModalDelete(items.id, fetch),
        },
      ],
    ]

    if (items.status === 'draft' || items.status === 'unpublished') {
      actions.push([
        {
          label: 'Publish',
          icon: 'i-heroicons-check-badge-20-solid',
          iconClass: 'bg-green-500',
          click: () => openModalStatus(items, fetch, 'published'),
        },
      ])
    } else if (items.status === 'published') {
      actions.push([
        {
          label: 'Unpublish',
          icon: 'i-heroicons-x-circle-20-solid',
          iconClass: 'bg-green-500',
          click: () => openModalStatus(items, fetch, 'unpublished'),
        },
      ])
    }

    return actions
  }

  function openModalDelete(id: string, fetch: object) {
    idItems.value = id
    isOpenDelete.value = true
    fetchObject.value = fetch
  }

  function openModalStatus(row: {id: string}, fetch: object, status: string) {
    statusObject.value = status
    titleObject.value = (statusObject.value === 'published') ? 'Publish': 'Unpublish'
    idItems.value = row.id
    fetchObject.value = fetch
    isOpenObject.value = true
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
          await useDeleteData(urlAPI, idItems)
          fetchObject.value()
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

  async function updateStatus() {
    const status = {
      status: statusObject.value
    }
    await useUpdatePatchData(urlAPI, idItems, status)
    fetchObject.value()
    toast.add({
      icon: 'i-heroicons-exclamation-triangle',
      title: 'data successfully Update',
      color: 'green',
      timeout: 1000,
    })
    isOpenObject.value = false
  }

  onMounted(() => {
    const activePage = useActivePage()
    activePage.page = 'Objek Wisata'
    activePage.desc = 'Berisi semua daftar terkait objek tempat wisata.'
  })
</script>
