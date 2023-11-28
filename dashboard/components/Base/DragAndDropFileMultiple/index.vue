<template>
  <div class="flex justify-between gap-6">
    <div class="w-[60%] flex-1 font-lato text-gray-800">
      <div class="mt-2 flex h-full w-full items-center justify-center">
      
      <div
      v-if="fileInputIsChange"
      :class="[
        dataFilesMultiple[0]?.fileCorrect
        ? 'border-green-300 bg-green-50'
        : 'border-red-300 bg-red-50',
        heightDragAndDrop,
      ]"
        class="flex h-full w-full flex-col items-center justify-center rounded-lg border-2 border-dashed px-2"
        >
        <img
          class="max-w-[250px] max-h-[150px]"
          :src="dataFilesMultiple[0]?.url !== '' ? dataFilesMultiple[0].url : fileDocument(dataFilesMultiple[0])"
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
        class="flex h-full w-full flex-col items-center justify-center rounded-lg border-2 border-dashed border-gray-300 bg-gray-50 px-4"
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
    </div>
    <div class="ml-2 mt-3 w-[40%]">
      <label v-show="label" class="font-bold text-gray-800">{{ label }}</label>
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
        <div v-if="fileInputIsChange" class="flex flex-row">
          <UButton
            class="ml-2"
            color="red"
            variant="soft"
            size="lg"
            @click="imageDelete(0, dataFilesMultiple[0]?.url)"
          >
            <NuxtIcon name="common/trash" filled class="text-lg text-red-600" />
          </UButton>
        </div>
      </div>
      <input
        id="dragAndDropFileMultiple"
        ref="dragAndDropFileMultiple"
        type="file"
        class="hidden"
        :disabled="dataFilesMultiple.length === 0 ? false : disabled"
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
    imageUrlMultiple: {
      type: Array,
      default: () => ([]),
    },
    disabled: {
      type: Boolean,
      default: false,
    },
  })

  const emit = defineEmits(['previewFile', 'deleteUrlFileMultiple'])
  const files = ref()
  const dragAndDropFileMultiple = ref()
  const dataFilesMultiple: Ref<object[]> = ref([])
  const fileInputIsChange = ref(false)
  const proggresBarIsSuccess = ref(false)
  const percentageProggres = ref(0)
  const intervalPercentage = ref(null)
  const formatSizeFile = ref(['Bytes', 'KB', 'MB', 'GB', 'TB'])
  const responseImage = ref('')
  const fileIsCorrect = ref(false)
  const disabledButton = ref(true)
  const dataImage = useDataImage()
  const counterImageLoad = ref(0)

  interface dataFile {
    name: string
    isConfidental: boolean
    mimeType: string
    roles: object
    data: string
    fileSize: string
    fileCorrect: boolean
    url: string
    height: number
    width: number
    fileSizeNumber: number
  }

  onUpdated(async () => {
    if (counterImageLoad.value < props.imageUrlMultiple.length) {
      for (let i = 0; i < props.imageUrlMultiple.length; i++) {
        const dataFile: dataFile = {
          name: '',
          isConfidental: false,
          mimeType: '',
          roles: ['admin'],
          data: '',
          fileSize: '',
          fileCorrect: false,
          url: '',
        }
        fileInputIsChange.value = true
        fileIsCorrect.value = true
        dataFile.fileCorrect = true
        dataFile.url = props.imageUrlMultiple[i]

        dataFilesMultiple.value.push(dataFile)
        counterImageLoad.value++
      }

      dataImage.dataImageMultiple = dataFilesMultiple.value
    }
  })

  const inputFileHandle = () => {
    dragAndDropFileMultiple.value.click()
  }

  const onChangeUpload = async (e: Event): Promise<void> => {
    let element = e.target as HTMLInputElement
    let fileTarget = element?.files ? element.files : []

    if (fileTarget?.length >= 1) {
      for (let i = 0; i < fileTarget.length; i++) {
        const dataFile: dataFile = {
          name: '',
          isConfidental: false,
          mimeType: '',
          roles: ['admin'],
          data: '',
          fileSize: '',
          fileSizeNumber: 0,
          fileCorrect: false,
          url: '',
          width: 0,
          height: 0,
        }

        files.value = fileTarget[i]
        dataFile.name = files.value.name
        dataFile.mimeType = files.value.type
        dataFile.fileSizeNumber = files.value.size
        dataFile.fileSize = convertSize(files.value.size)
        fileInputIsChange.value = true
        convertFileToBase64(files.value, files.value.name)
        let img = new Image()
        img.src = await window.URL.createObjectURL(fileTarget[i])
        img.onload = () => {
          dataFile.width = img.width
          dataFile.height = img.height

          dataFile.fileCorrect = checkFileValidation(dataFile)
        }

        dataFilesMultiple.value.push(dataFile)
      }

      dataImage.dataImageMultiple = dataFilesMultiple.value
    }
  }

  const dragover = (e: Event): void => {
    let element = e.target as HTMLInputElement

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

  const convertSize = (sizeFile: string) => {
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
    useDataImage().dataImageMultiple = []
  }

  const previewFile = () => {
    emit('previewFile')
  }

  const convertFileToBase64 = (FileObject: Blob, name: string) => {
    const reader = new FileReader()

    reader.onload = () => {
      dataFilesMultiple.value.filter((x) => x.name === name)[0].data =
        reader.result.split(',')[1]

      useDataImage().dataImageMultiple = { ...dataFilesMultiple.value }
    }
    reader.readAsDataURL(FileObject)
  }

  const checkFileValidation = (file: Array) => {
    if (file) {
      if (
        fileSizeIsCompatible(file) &&
        formatFileIsCompatible(file) &&
        fileResolutionIsCompatible(file)
      ) {
        return true
      } else {
        return false
      }
    } else {
      return false
    }
  }

  const fileSizeIsCompatible = (file: Array) => {
    return file.fileSizeNumber <= props.detailDragAndDrop.maxSizeFile
  }

  const formatFileIsCompatible = (file: Array) => {
    return props.detailDragAndDrop.formatTypeFile.includes(file.mimeType)
  }

  const fileResolutionIsCompatible = (file: Array) => {
    const validate =
      file.width > props.detailDragAndDrop.maxResolution[0] &&
      file.height > props.detailDragAndDrop.maxResolution[1]

    return !validate
  }

  const fileDocument = (file: Array) => {
    return file.data ? base64ToBlobUrl(file.data, file.mimeType) : ''
  }

  const imageClicked = (data: object, index: number) => {
    // TODO : It for define the main of banner
  }

  const imageDelete = (index: number, dataUrl: string) => {
    if (dataUrl !== '') {
      emit('deleteUrlFileMultiple', index)
    }
    
    fileInputIsChange.value = false
    dataFilesMultiple.value.splice(index, 1)
    dataImage.dataImageMultiple = { ...dataFilesMultiple.value }
  }
</script>
