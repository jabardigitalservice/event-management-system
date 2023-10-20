<template>
  <div>
    <UInput 
      v-model="search" 
      placeholder="Cari nama organisasi"
      icon="i-heroicons-magnifying-glass-20-solid" 
      class="w-1/5 mb-5" 
      :ui="{ icon: { trailing: { pointer: '' } } }">
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
    <div class="inline-block min-w-full rounded-lg overflow-hidden bg-white">
       <UTable 
         :columns="props.headers"
         :rows="itemTable" >
          <template #photo-data="{ row }">
            <UButton  color="green" icon="i-heroicons-photo" variant="outline"> Lihat Gambar </UButton>
          </template>
          <template #actions-data="{ row }">
            <UDropdown :items="items()">
              <UButton color="green" variant="outline" icon="i-heroicons-eye-solid" >Details </UButton> 
            </UDropdown>
          </template>
        </UTable>
     </div>
     <span class="text-sm text-gray-700 dark:text-gray-400 float-left mt-5">
        Showing 
        <span class="font-semibold text-gray-900 dark:text-white">{{ from }}</span> 
        to 
        <span class="font-semibold text-gray-900 dark:text-white">{{ to }}</span> 
        of 
        <span class="font-semibold text-gray-900 dark:text-white">{{ total }}</span> 
        Entries
    </span>
     <UPagination 
     v-model="currentPage"
     class="mt-5 mb-10 justify-end" 
     :page-count="10" 
     :total=total
     :value="currentPage"
     :prev-button="{ icon: 'i-heroicons-arrow-small-left-20-solid', label: 'Prev', color: 'gray' }" 
     :next-button="{ icon: 'i-heroicons-arrow-small-right-20-solid', trailing: true, label: 'Next', color: 'gray' }"
     @click="onClick(currentPage)"
     />
  </div>
</template>
  
<script setup lang="ts">

    import { ref, watch } from 'vue'
    import { useFetchData } from "~/composables/useFetchData"

    const props = defineProps({
        headers: {
            type: String,
            default: ''
        },
        path: {
            type: String,
            default: ''
        }
    })

    const items = () => [
      [{
        label: 'Edit',
        icon: 'i-heroicons-pencil-square-20-solid',
      }, {
        label: 'Detail',
        icon: 'i-heroicons-document-duplicate-20-solid'
      }], [{
        label: 'Delete',
        icon: 'i-heroicons-trash-20-solid'
      }]
    ]

    const search = ref('')
    const itemTable = ref()
    const page:Ref<number> = ref(1)
    let currentPage:number = 1
    let total:number = 1
    let to:number = 1
    let from:number = 1

    async function fetchData() {
      const result= await useFetchData(props.path, page)
      currentPage = result.current_page
      itemTable.value = result.data
      total = result.total
      to = result.to
      from = result.from
    }

    function onClick(getPage: number){
      page.value = getPage
    }

    fetchData()
    
    watch(page, () => {
      fetchData()
    })

</script>
<style>
  thead {
    @apply px-5 py-3 border-b-2 border-gray-200 text-left text-sm bg-green-700 font-semibold text-white tracking-wider rounded-lg
  }
  th>button>span {
    @apply text-white
  }
  th>span{
    @apply text-white
  }
  .form-input{
    @apply h-10
  }
  .hover\:bg-gray-50:hover {
    @apply bg-green-800
  }
</style>

