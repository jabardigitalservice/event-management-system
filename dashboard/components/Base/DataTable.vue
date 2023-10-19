<template>
  <div>
    <div class="inline-block min-w-full shadow-md rounded-lg overflow-hidden"
       >
       <UTable 
         :columns="props.headers"
         :rows="itemTable"
       />
     </div>
     <span class="text-sm text-gray-700 dark:text-gray-400 float-left mt-5">
        Showing <span class="font-semibold text-gray-900 dark:text-white">{{ from }}</span> to <span class="font-semibold text-gray-900 dark:text-white">{{ to }}</span> of <span class="font-semibold text-gray-900 dark:text-white">{{ total }}</span> Entries
    </span>
     <UPagination 
     v-model="currentPage"
     class="mt-5 justify-end" 
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
    });

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
    @apply px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-sm font-semibold text-gray-700 tracking-wider rounded-lg
  }
</style>

