<template>
  <div>
    <BaseDataTable :headers="objectHeaders" :path="urlAPI" base-route="/object">
      <template #customeName="{ items }">
        <div class="flex flex-row">
          <div class="basis-[13%]">
            <UAvatar size="xs" :src="items.logo" :alt="items.name" />
          </div>
          <div class="basis-1/2">
            {{ items.name }}
          </div>
        </div>
      </template>
      <template #customeStatus="{ items }">
        <UBadge
          :color="statusColors[items.status].color"
          variant="soft"
          :ui="{
            rounded: 'rounded-full',
            variant: {
              soft: 'bg-{color}-100 dark:bg-{color}-400 dark:bg-opacity-10 text-{color}-500 dark:text-{color}-400',
            },
          }"
          class="capitalize"
          >{{ items.status }}
        </UBadge>
      </template>
      <template #customeAction="{ items, fetch }">
        <UTooltip text="Laporan" :popper="{ offsetDistance: 16 }">
          <UButton
            class="mr-1"
            variant="ghost"
            color="blue"
            :ui="{
              rounded: 'rounded-md',
            }"
          >
            <template #trailing>
              <NuxtIcon
                name="navigation/laporan-icon"
                filled
                class="stroke-[#737373] text-[19px]"
              />
            </template>
          </UButton>
        </UTooltip>
        <UTooltip text="Atur Tiket" :popper="{ offsetDistance: 16 }">
          <UButton
            class="mr-1"
            variant="ghost"
            color="blue"
            :ui="{
              rounded: 'rounded-md',
            }"
          >
            <template #trailing>
              <NuxtIcon
                name="navigation/tiket-icon"
                filled
                class="stroke-[#737373] text-[19px]"
              />
            </template>
          </UButton>
        </UTooltip>
        <UTooltip text="Detail" :popper="{ offsetDistance: 16 }">
          <UButton
            class="mr-1"
            variant="ghost"
            color="blue"
            :ui="{
              rounded: 'rounded-md',
            }"
            @click="
              router.push({ path: 'object/detail', query: { id: items.id } })
            "
          >
            <template #trailing>
              <NuxtIcon
                name="navigation/eye-icon"
                filled
                class="stroke-[#737373] text-[19px]"
              />
            </template>
          </UButton>
        </UTooltip>
        <UDropdown :items="itemActions(items, fetch)">
          <UButton color="blue" variant="ghost">
            <template #trailing>
              <NuxtIcon
                name="navigation/dots-icon"
                filled
                class="stroke-[#737373] p-[2px] text-[15px]"
              />
            </template>
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
    <BaseModal
      :open-modal="state.isOpenObject"
      :title-modal="`Confirm to ${state.titleObject}`"
      :desc-modal="`Are you sure you want to ${state.titleObject}`"
      icon-modal="i-heroicons-question-mark-circle"
      :text-confirm="state.titleObject"
      type-modal="warning"
      @close="state.isOpenObject = false"
      @confirm="updateStatus()"
    />
  </div>
</template>

<script setup lang="ts">
  import { useActivePage, useIdData } from '@/store/index'
  import { objectHeaders } from '~/common/constant/object'

  const router = useRouter()
  const toast = useToast()
  const state = reactive({
    isOpenDelete: false,
    isOpenObject: false,
    statusObject: '',
    titleObject: '',
    idItems: '',
    fetchObject: {},
  })
  const urlAPI: string = '/v1/event/object'

  const statusColors: Record<string, { color: string }> = {
    draft: { color: 'orange' },
    published: { color: 'green' },
    unpublished: { color: 'red' },
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
          iconClass: 'bg-blue-700',
          click: () => {
            useIdData().id = items.id
            router.push({ path: '/object/form' })
          },
        },
        {
          label: 'Detail',
          icon: 'i-heroicons-document-magnifying-glass-20-solid',
          iconClass: 'bg-blue-700',
        },
      ],
      [
        {
          label: 'Delete',
          icon: 'i-heroicons-trash-20-solid',
          iconClass: 'bg-blue-700',
          click: () => openModalDelete(items.id, fetch),
        },
      ],
    ]

    if (items.status === 'draft' || items.status === 'unpublished') {
      actions.push([
        {
          label: 'Publish',
          icon: 'i-heroicons-check-badge-20-solid',
          iconClass: 'bg-blue-700',
          click: () => openModalStatus(items, fetch, 'published'),
        },
      ])
    } else if (items.status === 'published') {
      actions.push([
        {
          label: 'Unpublish',
          icon: 'i-heroicons-x-circle-20-solid',
          iconClass: 'bg-blue-700',
          click: () => openModalStatus(items, fetch, 'unpublished'),
        },
      ])
    }

    return actions
  }

  function openModalDelete(id: string, fetch: object) {
    state.idItems = id
    state.isOpenDelete = true
    state.fetchObject = fetch
  }

  function openModalStatus(row: { id: string }, fetch: object, status: string) {
    state.statusObject = status
    state.titleObject =
      state.statusObject === 'published' ? 'Publish' : 'Unpublish'
    state.idItems = row.id
    state.fetchObject = fetch
    state.isOpenObject = true
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

  async function updateStatus() {
    const status = {
      status: state.statusObject,
    }
    await useUpdatePatchData(urlAPI, state.idItems, status)
    state.fetchObject()
    toast.add({
      icon: 'i-heroicons-exclamation-triangle',
      title: 'data successfully Update',
      color: 'green',
      timeout: 1000,
    })
    state.isOpenObject = false
  }

  onMounted(() => {
    const activePage = useActivePage()
    activePage.page = 'Objek Wisata'
    activePage.navigation = false
  })
</script>
