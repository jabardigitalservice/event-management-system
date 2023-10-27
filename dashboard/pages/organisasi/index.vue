<template>
    <BaseDataTable :headers="headerTable" :actions="commonItems" :path="urlAPI">
      <template #customePhoto="">
        <UButton  
          color="green" 
          icon="i-heroicons-photo" 
          variant="outline"  
          @click="(openModalPhoto = true)"> Lihat Gambar 
        </UButton>
      </template>
    </BaseDataTable>
    <UModal v-model="openModalPhoto">
      <div class="p-4">
        <div class="relative overflow-hidden rounded border border-dashed border-gray-400 dark:border-gray-500 opacity-75 h-48">
          
        </div>
      </div>
    </UModal>
</template>

<script setup lang="ts">
  import { useActivePage } from '@/store/index'
  import { organisasiHeaders } from "~/common/constant/organisasi"

  const headerTable = organisasiHeaders
  const urlAPI: string = "/organizer"
  const openModalPhoto = ref(false)

  const commonItems = [
    {
      label: 'Edit',
      icon: 'i-heroicons-pencil-square-20-solid',
      iconClass: 'bg-green-500'
    },
    {
      label: 'Detail',
      icon: 'i-heroicons-document-magnifying-glass-20-solid',
      iconClass: 'bg-green-500'
    }
  ];
  
  onMounted(() => {
    const activePage = useActivePage()
    activePage.page = 'Organisasi'
    activePage.desc = 'Berisi semua daftar terkait organisasi yang mengelola event.'
  })

</script>