<template>
    <BaseDataTable :headers="headerTable" :path="urlAPI"/>
</template>

<script setup lang="ts">
  import { onMounted} from 'vue'
  import { useActivePage } from '@/store/index'
  import { organisasiHeaders } from "~/common/constant/organisasi"

  const headerTable = organisasiHeaders
  const urlAPI: string = "/users"
  
  onMounted(() => {
    const activePage = useActivePage()
    activePage.page = 'Organisasi'
  })

</script>