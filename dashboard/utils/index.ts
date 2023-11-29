const toast = useToast()
const router = useRouter()

export function base64ToBlobUrl (base64: string, type: string) {
    const binStr = atob(base64)
    const len = binStr.length
    const arr = new Uint8Array(len)
    for (let i = 0; i < len; i++) {
      arr[i] = binStr.charCodeAt(i)
    }
    const blob = new Blob([arr], { type })
    const url = URL.createObjectURL(blob)
    return url
  }

  
export function deleteData({ urlAPI, id, path, fetch }: { urlAPI: string, id?: string, path?: string, fetch?: object }) {
  let cancelled = false
  toast.add({
    icon: 'i-heroicons-exclamation-triangle',
    title: 'data will be deleted',
    color: 'yellow',
    timeout: 2000,
    callback: async () => {
      if (!cancelled) {
        await useDeleteData(urlAPI, id)
        toast.add({
          title: 'data successfully deleted',
          icon: 'i-heroicons-check-circle',
          timeout: 1000,
        })
        if(!!path){
          router.push({ path })
        }
        if (fetch && typeof fetch === 'function') {
          fetch()
        }
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
}

export function alertDanger(label: string){
  toast.add({
    title: label,
    color: 'red',
    icon: 'i-heroicons-x-circle',
    timeout: 2500,
  })
}

export function alertSuccess(label: string){
  toast.add({
    title: label,
    color: 'green',
    icon: 'i-heroicons-exclamation-triangle',
    timeout: 2500,
  })
}
