<template>
    <BaseDataTable :headers="headerTable" :items="itemTable"/>
</template>

<script setup lang="ts">
  import type { Header, Item } from "vue3-easy-data-table"
  import { onMounted, ref } from 'vue'
  import { useActivePage } from '@/store/index'
  import { organisasiHeaders } from "~/common/constant/organisasi";

  const config = useRuntimeConfig()
  const headerTable: Header[] = organisasiHeaders
  const itemTable = ref<Item[]>([]);
  
  async function fetchData() {
    await useFetch(`${config.public.baseURL}/users?page=1`, {
      onResponse({ response }) {
        itemTable.value = response._data.data.data as Item[];
      }
    });
  }

  fetchData();
  
  onMounted(() => {
    const activePage = useActivePage()
    activePage.page = 'Organisasi'
  })

</script>