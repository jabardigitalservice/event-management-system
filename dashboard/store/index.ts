import { defineStore } from 'pinia'

export const useActivePage = defineStore('active-page', {
  state: () => {
    return { page: '', navigation: false }
  },
})

export const useDataImage = defineStore('data-image', {
  state: () => {
    return {
      dataImage: {},
      dataImageMultiple: [],
    }
  },
})

export const useIdData = defineStore('id-data', {
  state: () => {
    return {
      id: '',
    }
  },
})
