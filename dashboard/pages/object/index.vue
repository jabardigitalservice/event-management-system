<template>
  <BaseDataTable
    :headers="headerTable"
    :actions="commonItems"
    :path="urlAPI"
    base-route="/object"
  >
    <template #customeStatus="{items}">
      <UBadge
        :color="statusColors[items.status].color"
        :variant="statusColors[items.status].variant"
        class="uppercase"
        >{{ items.status }}
      </UBadge>
    </template>
  </BaseDataTable>
</template>

<script setup lang="ts">
  import { useActivePage } from '@/store/index'
  import { objectHeaders } from '~/common/constant/object'

  interface StatusColor {
    color: string
    variant: string
  }

  const statusColors: Record<string, StatusColor> = {
    draft: { color: 'red', variant: 'soft' },
    published: { color: 'green', variant: 'soft' },
    unpublished: { color: 'yellow', variant: 'soft' },
  }

  const headerTable = objectHeaders
  const urlAPI: string = '/object'
  const commonItems = [
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
    {
      label: 'Delete',
      icon: 'i-heroicons-trash-20-solid',
      iconClass: 'bg-green-500',
    },
    {
      label: 'Publish',
      icon: 'i-heroicons-check-badge-20-solid',
      iconClass: 'bg-green-500',
    },
    {
      label: 'Unpublish',
      icon: 'i-heroicons-x-circle-20-solid',
      iconClass: 'bg-green-500',
    },
  ]

  onMounted(() => {
    const activePage = useActivePage()
    activePage.page = 'Objek Wisata'
    activePage.desc = 'Berisi semua daftar terkait objek tempat wisata.'
  })
</script>
