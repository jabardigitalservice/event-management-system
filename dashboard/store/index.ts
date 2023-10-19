import { defineStore } from 'pinia'

export const useActivePage = defineStore('active-page', {
  state: () => {
    return { page: '' }
  },
})