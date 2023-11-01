<template>
  <div>
    <BaseDataTable :headers="headerTable" :path="urlAPI" base-route="/object">
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
      :title-modal="'Confirm to ' + statusObject"
      :desc-modal="'Are you sure you want to ' + statusObject"
      icon-modal="i-heroicons-question-mark-circle"
      :text-confirm="statusObject"
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
  const idItems = ref('')
  const fetchObject = ref()
  const headerTable = objectHeaders
  const urlAPI: string = '/object'

  const statusColors: Record<string, StatusColor> = {
    draft: { color: 'red', variant: 'soft' },
    published: { color: 'green', variant: 'soft' },
    unpublished: { color: 'yellow', variant: 'soft' },
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

    if (items.status === 'draft') {
      actions.push([
        {
          label: 'Publish',
          icon: 'i-heroicons-check-badge-20-solid',
          iconClass: 'bg-green-500',
          click: () => openModalStatus(items.id, 'Publish'),
        },
      ])
    } else if (items.status === 'published') {
      actions.push([
        {
          label: 'Unpublish',
          icon: 'i-heroicons-x-circle-20-solid',
          iconClass: 'bg-green-500',
          click: () => openModalStatus(items.id, 'Unpublish'),
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

  function openModalStatus(row: string, status: string) {
    statusObject.value = status
    idItems.value = row
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

  function updateStatus() {
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
