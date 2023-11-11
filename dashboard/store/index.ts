import { defineStore } from 'pinia'

export const useActivePage = defineStore('active-page', {
  state: () => {
    return { page: '', desc: '' }
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
