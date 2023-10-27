<template>
  <div class="flex justify-between gap-6">
    <div class="w-50 flex-1 font-lato text-gray-800">
      <slot name="header" />
      <label
        v-show="label"
        class="message-notif-form__label-required text-gray-800"
        >{{ label }}</label
      >
      <p v-show="sublabel" class="mb-1 text-[13px] text-gray-600">
        {{ sublabel }}
      </p>
      <div class="mt-2 flex w-full items-center justify-center">
        <label
          :class="heightDragAndDrop"
          for="drag-and-drop-file-multiple"
          class="flex w-full cursor-pointer flex-col items-center justify-center rounded-lg border-2 border-dashed border-gray-300 bg-gray-50 px-4 hover:bg-gray-200"
          @dragover="dragover"
          @dragleave="dragleave"
          @drop="drop"
        >
          <div class="flex flex-col items-center justify-center">
            <NuxtIcon
              name="common/upload-file"
              class="mb-3 text-4xl text-gray-300"
            />
            <p class="mb-2 font-lato text-[14px] text-gray-700">
              <span>Drag File kesini</span>
              <span class="ml-1 text-gray-500">atau</span>
              <span class="ml-1 text-green-600 underline">Pilih File</span>
            </p>
          </div>
          <input
            id="drag-and-drop-file-multiple"
            ref="file"
            type="file"
            class="hidden"
            :disabled="dataFilesMultiple.length === 0 ? false : true"
            :accept="detailDragAndDrop.acceptFile"
            @change="onChangeUpload"
          />
        </label>
      </div>
    </div>
    <div class="w-50 mt-5 max-h-[300px] flex-1 overflow-auto pl-3 pr-3 pt-5">
      <RadioGroup v-show="dataFilesMultiple" v-model="selected">
        <RadioGroupLabel class="sr-only">Server size</RadioGroupLabel>
        <div class="grid gap-4 space-y-2 md:grid-cols-2 lg:grid-cols-3">
          <RadioGroupOption
            v-for="(data, index) in dataFilesMultiple"
            v-slot="{ active, checked }"
            :key="index"
            :value="index"
            as="template"
            @click="imageClicked(data, index)"
          >
            <div
              :class="[
                active
                  ? 'ring-2 ring-white ring-opacity-60 ring-offset-2 ring-offset-sky-300'
                  : '',
                checked ? 'bg-sky-900 bg-opacity-75 text-white ' : 'bg-white ',
              ]"
              class="relative flex cursor-pointer rounded-lg p-2 shadow-md focus:outline-none"
            >
              <div class="absolute -translate-x-3 translate-y-6">
                <UBadge
                  :label="index + 1"
                  size="sm"
                  color="green"
                  square
                  variant="solid"
                  class="absolute"
                />
              </div>
              <div class="absolute -translate-x-3 -translate-y-3">
                <UButton
                  icon="i-heroicons-trash"
                  size="sm"
                  color="red"
                  square
                  variant="solid"
                  class="absolute"
                  @click="imageDelete(index)"
                />
              </div>

              <div class="flex items-center">
                <div class="flex items-center">
                  <img :src="fileDocument(data)" class="" />
                </div>
              </div>
            </div>
          </RadioGroupOption>
        </div>
      </RadioGroup>
    </div>
  </div>
</template>

<script setup lang="ts">
  import { useDataImage } from '@/store/index'
  import { base64ToBlobUrl } from '~/utils'
  import {
    RadioGroup,
    RadioGroupLabel,
    RadioGroupOption,
  } from '@headlessui/vue'

  const props = defineProps({
    detailDragAndDrop: {
      type: Object,
      default: () => {},
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
  })

  const emit = defineEmits(['previewFile'])
  const files = ref()
  const dataFilesMultiple: Ref<object[]> = ref([])
  const fileInputIsChange = ref(false)
  const proggresBarIsSuccess = ref(false)
  const percentageProggres = ref(0)
  const intervalPercentage = ref(null)
  const formatSizeFile = ['Bytes', 'KB', 'MB', 'GB', 'TB']
  const responseImage = ref('')
  const fileIsCorrect = ref(false)
  const disabledButton = ref(true)
  const dataImage = useDataImage()

  const selected = ref(dataFilesMultiple.value[0])

  const onChangeUpload = async (e: Event): Promise<void> => {
    let element = e.target as HTMLInputElement
    let fileTarget = element?.files ? element.files : []

    if (fileTarget?.length >= 1) {
      for (let i = 0; i < fileTarget.length; i++) {
        const dataFile = {
          name: '',
          isConfidental: false,
          mimeType: '',
          roles: ['admin'],
          data: '',
          fileSize: '',
          fileCorrect: false,
        }

        files.value = fileTarget[i]
        dataFile.name = files.value.name
        dataFile.mimeType = files.value.type
        dataFile.fileSize = convertSize(files.value.size)
        fileInputIsChange.value = true
        convertFileToBase64(files.value, files.value.name)
        runProgressBar()
        checkFileValidation()
        dataFile.fileCorrect = fileIsCorrect.value

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
      return sizeFile + ' ' + formatSizeFile[indexFileSize]
    }
    return (
      (sizeFile / Math.pow(1024, indexFileSize)).toFixed(1) +
      ' ' +
      formatSizeFile[indexFileSize]
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
      dataFilesMultiple.value.filter((x) => x.name == name)[0].data =
        reader.result.split(',')[1]

      useDataImage().dataImageMultiple = { ...dataFilesMultiple.value }
    }
    reader.readAsDataURL(FileObject)
  }

  const checkFileValidation = () => {
    if (files.value) {
      if (FileSizeIsCompatible() && FormatFileIsCompatible()) {
        fileIsCorrect.value = true
        disabledButton.value = false
      } else {
        fileIsCorrect.value = false
      }
    } else {
      fileIsCorrect.value = false
    }
  }

  const FileSizeIsCompatible = () => {
    return files.value.size <= props.detailDragAndDrop.maxSizeFile
  }

  const FormatFileIsCompatible = () => {
    return props.detailDragAndDrop.formatTypeFile.includes(files.value.type)
  }

  const fileDocument = (file: Array) => {
    return file.data ? base64ToBlobUrl(file.data, file.mimeType) : ''
  }

  const imageClicked = (data: object, index: number) => {
    // TODO : It for define the main of banner
  }

  const imageDelete = (index: number) => {
    dataFilesMultiple.value.splice(index, 1)

    useDataImage().dataImageMultiple = { ...dataFilesMultiple.value }
  }
</script>
