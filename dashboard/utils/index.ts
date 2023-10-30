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