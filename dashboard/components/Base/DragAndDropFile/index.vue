<template>
  <div class="flex font-lato text-gray-800">
    <div class="mt-2 flex h-[120px] w-[120px] items-center justify-center">
      <div
        v-if="fileInputIsChange"
        :class="[
          dataFiles.fileCorrect
            ? 'border-green-300 bg-green-50'
            : 'border-red-300 bg-red-50',
          heightDragAndDrop,
        ]"
        class="flex h-full w-full flex-col items-center justify-center rounded-full border-2 border-dashed px-2"
      >
        <img
          class="max-w-11 max-h-11"
          :src="dataFiles?.url !== '' ? dataFiles.url : fileDocument(dataFiles)"
        />

        <div class="flex w-full flex-row items-center">
          <div class="w-[100%]">
            <template v-if="proggresBarIsSuccess">
              <div class="mb-1 h-1.5 w-full rounded-full bg-gray-200">
                <div
                  class="h-1.5 rounded-full bg-green-500"
                  :style="{ width: percentageProggres + '%' }"
                />
              </div>
              <div class="mb-1">
                <span class="font-lato text-[11px] font-normal text-gray-600">
                  Diupload ... {{ percentageProggres }} %
                </span>
              </div>
            </template>
          </div>
        </div>
      </div>
      <label
        v-else
        :class="heightDragAndDrop"
        class="flex h-full w-full flex-col items-center justify-center rounded-full border-2 border-dashed border-gray-300 bg-gray-50 px-4"
        @dragover="dragover"
        @dragleave="dragleave"
        @drop="drop"
      >
        <div class="flex flex-col items-center justify-center">
          <NuxtIcon
            filled
            name="common/default-photo"
            class="text-3xl text-gray-300"
          />
        </div>
      </label>
    </div>
    <div class="ml-4 mt-3">
      <label v-show="label" class="text-gray-800 font-bold">{{ label }}</label>
      <p v-show="sublabel" class="mb-2 text-[13px] text-gray-600">
        {{ sublabel }}
      </p>
      <div class="flex flex-row">
        <UButton
          for="dragAndDropFile"
          color="blue"
          variant="outline"
          size="lg"
          :disabled="fileInputIsChange"
          @click="inputFileHandle"
        >
          Pilih File
        </UButton>
        <div  v-if="fileInputIsChange" class="flex flex-row">
          <UButton
            class="ml-2"
            color="red"
            variant="soft"
            size="lg"
            @click="resetDataFile"
          >
            <NuxtIcon name="common/trash" filled class="text-red-600 text-lg" />
          </UButton>
        </div>
      </div>
      <input
        id="dragAndDropFile"
        ref="dragAndDropFile"
        type="file"
        :disabled="disabled"
        class="hidden"
        :accept="detailDragAndDrop.acceptFile"
        @change="onChangeUpload"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
  import { useDataImage } from '@/store/index'
  import { base64ToBlobUrl } from '~/utils'

  const props = defineProps({
    detailDragAndDrop: {
      type: Object,
      default: () => ({}),
    },
    heightDragAndDrop: {
      type: String,
      default: 'h-40',
    },
    label: {
      type: String,
      default: '',
    },
    sublabel: {
      type: String,
      default: '',
    },
    imageUrl: {
      type: String,
      default: '',
    },
    disabled: {
      type: Boolean,
      default: false,
    },
  })

  const emit = defineEmits(['previewFile', 'deleteUrlFile'])

  const files = ref()
  const dragAndDropFile = ref()
  const dataFiles = reactive({
    name: '',
    isConfidental: false,
    mimeType: '',
    roles: ['admin'],
    data: '',
    fileSize: '',
    fileCorrect: false,
    height: 0,
    width: 0,
    size: 0,
    url: '',
  })
  const fileInputIsChange = ref(false)
  const proggresBarIsSuccess = ref(false)
  const percentageProggres = ref(0)
  const intervalPercentage = ref(null)
  const formatSizeFile = ref(['Bytes', 'KB', 'MB', 'GB', 'TB'])
  const responseImage = ref('')
  const fileIsCorrect = ref(false)
  const disabledButton = ref(true)

  onUpdated(async () => {
    if (props.imageUrl) {
      const res = await $fetch(props.imageUrl)

      dataFiles.name = ''
      dataFiles.mimeType = res.type
      dataFiles.size = res.size
      dataFiles.url = props.imageUrl
      dataFiles.fileSize = convertSize(res.size)
      fileInputIsChange.value = true
      fileIsCorrect.value = true

      dataFiles.fileCorrect = true
    }
  })

  const inputFileHandle = () => {
    dragAndDropFile.value.click()
  }

  const onChangeUpload = async (e) => {
    if (e.target.files[0]) {
      files.value = e.target.files[0]
      dataFiles.url = ''
      dataFiles.name = files.value.name
      dataFiles.mimeType = files.value.type
      dataFiles.fileSize = convertSize(files.value.size)
      fileInputIsChange.value = true
      convertFileToBase64(files.value)
      let img = new Image()
      img.src = await window.URL.createObjectURL(e.target.files[0])
      img.onload = () => {
        dataFiles.width = img.width
        dataFiles.height = img.height

        dataFiles.fileCorrect = checkFileValidation(dataFiles)
      }
      runProgressBar()
      disabledButton.value = true
      useDataImage().dataImage = JSON.parse(JSON.stringify(dataFiles))
    }
  }

  const dragover = (e: Event): void => {
    let element = e.target as HTMLInputElement

    // add style drag and drop
    if (!element.classList.contains('bg-gray-200')) {
      element.classList.remove('bg-gray-50')
      element.classList.add('bg-gray-200')
    }
    e.preventDefault()
  }

  const dragleave = (e: Event): void => {
    let element = e.target as HTMLInputElement
    // clear style drag and drop
    element.classList.add('bg-gray-50')
    element.classList.remove('bg-gray-200')
  }

  const drop = (e: Event): void => {
    e.preventDefault()
    let element = e.target as HTMLInputElement

    element.files = e?.dataTransfer.files
    onChangeUpload(e)
    // clear style drag and drop
    element.classList.add('bg-gray-50')
    element.classList.remove('bg-gray-200')
  }

  const convertSize = (sizeFile: number) => {
    if (sizeFile === 0) {
      return 'n/a'
    }
    const indexFileSize = parseInt(
      Math.floor(Math.log(sizeFile) / Math.log(1024)),
    )
    if (indexFileSize === 0) {
      return sizeFile + ' ' + formatSizeFile.value[indexFileSize]
    }
    return (
      (sizeFile / Math.pow(1024, indexFileSize)).toFixed(1) +
      ' ' +
      formatSizeFile.value[indexFileSize]
    )
  }

  const runProgressBar = () => {
    if (percentageProggres.value === 0) {
      proggresBarIsSuccess.value = true
      percentageProggres.value = 1
      intervalPercentage.value = setInterval(setProggresBar, 10)
    }
  }

  const setProggresBar = () => {
    if (percentageProggres.value >= 100) {
      clearInterval(intervalPercentage.value)
      percentageProggres.value = 0
      proggresBarIsSuccess.value = false
    } else {
      percentageProggres.value++
    }
  }

  const resetDataFile = () => {
    percentageProggres.value = 0
    proggresBarIsSuccess.value = false
    fileInputIsChange.value = false
    files.value = ''
    responseImage.value = ''
    fileIsCorrect.value = false
    disabledButton.value = true
    dataFiles.url = ''
    dataFiles.name = ''
    dataFiles.mimeType = ''
    dataFiles.fileSize = ''

    emit('deleteUrlFile')

    useDataImage().dataImage = {}
  }

  const previewFile = () => {
    emit('previewFile')
  }

  const convertFileToBase64 = (FileObject: Blob) => {
    const reader = new FileReader()
    reader.onload = () => {
      dataFiles.data = reader.result.split(',')[1]

      useDataImage().dataImage = { ...dataFiles }
    }
    reader.readAsDataURL(FileObject)
  }

  const checkFileValidation = (file: Array) => {
    if (file) {
      if (fileSizeIsCompatible(file) && formatFileIsCompatible(file) && fileResolutionIsCompatible(file)) {
        return true
      } else {
        return false
      }
    } else {
      return false
    }
  }

  const fileSizeIsCompatible = (file: Array) => {
    return file.size <= props.detailDragAndDrop.maxSizeFile
  }

  const fileResolutionIsCompatible = (file: Array) => {
    const validate =
    file.width > props.detailDragAndDrop.maxResolution &&
    file.height > props.detailDragAndDrop.maxResolution
    
      return !validate
  }

  const formatFileIsCompatible = (file: Array) => {
    return props.detailDragAndDrop.formatTypeFile.includes(
      file.mimeType,
    )
  }

  const fileDocument = (file: Array) => {
    return file?.data ? base64ToBlobUrl(file.data, file.mimeType) : ''
  }
</script>
