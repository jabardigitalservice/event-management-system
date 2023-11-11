<template>
    <BaseDataTable :headers="organisasiHeaders" :path="urlAPI">
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
      :open-modal="state.isOpenDelete"
      title-modal="Confirm to delete"
      desc-modal="Are you sure you want to delete your data ? "
      icon-modal="i-heroicons-exclamation-triangle"
      text-confirm="Delete"
      type-modal="danger"
      @close="state.isOpenDelete = false"
      @confirm="deleteData()"
    />
</template>

<script setup lang="ts">
  import { useActivePage } from '@/store/index'
  import { organisasiHeaders } from "~/common/constant/organisasi"

  const toast = useToast()
  const state = reactive({
    isOpenDelete: false,
    idItems: '',
    fetchObject: {},
  })
  const urlAPI: string = "/organizer"

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

    return actions
  }

  function openModalDelete(id: string, fetch: object) {
    state.idItems = id
    state.isOpenDelete = true
    state.fetchObject = fetch
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
          await useDeleteData(urlAPI, state.idItems)
          state.fetchObject()
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
    state.isOpenDelete = false
  }
  
  onMounted(() => {
    const activePage = useActivePage()
    activePage.page = 'Organisasi'
    activePage.desc = 'Berisi semua daftar terkait organisasi yang mengelola event.'
  })

</script>