<template>
  <div class="font-lato text-gray-800">
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
      <div
        v-if="fileInputIsChange"
        :class="[
          fileIsCorrect
            ? 'border-green-300 bg-green-50'
            : 'border-red-300 bg-red-50',
          heightDragAndDrop,
        ]"
        class="flex w-full flex-col justify-center rounded-lg border-2 border-dashed px-4"
      >
        <div class="mb-3 flex items-center justify-center">
          <NuxtIcon
            name="common/document"
            class="h-9 w-9 text-4xl text-gray-600"
          />
        </div>

        <div class="flex w-full flex-row items-center justify-between">
          <div class="w-[85%]">
            <p
              class="mb-2 truncate font-lato text-[14px] font-normal text-black"
            >
              {{ dataFiles?.name }}
            </p>
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
            <template v-else>
              <p class="mb-2 font-lato text-[11px] font-normal text-gray-600">
                Ukuran {{ dataFiles?.fileSize }}
              </p>
            </template>
            <div v-if="imageUrl === ''">
              <p
                v-if="!fileSizeIsCompatible()"
                class="font-lato text-[11px] font-bold text-red-600"
              >
                {{ detailDragAndDrop.informationSizeCompatible }}
              </p>
              <p
                v-if="!formatFileIsCompatible()"
                class="font-lato text-[11px] font-bold text-red-600"
              >
                {{ detailDragAndDrop.informationFormatCompatible }}.
              </p>
              <p
                v-if="fileResolutionIsCompatible()"
                class="font-lato text-[11px] font-bold text-red-600"
              >
                Resolusi gambar max {{ detailDragAndDrop.maxResolution }}px
              </p>
            </div>
          </div>
          <div class="flex flex-row">
            <button class="w-4" @click="resetDataFile">
              <NuxtIcon name="common/trash" class="h-4 w-4 text-red-600" />
            </button>
            <button class="ml-5 w-4" @click.prevent="previewFile">
              <NuxtIcon
                v-if="!proggresBarIsSuccess"
                name="common/eyes"
                class="h-4 w-4 text-green-600"
              />
            </button>
          </div>
        </div>
      </div>
      <label
        v-else
        :class="heightDragAndDrop"
        for="drag-and-drop-file"
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
          id="drag-and-drop-file"
          ref="file"
          type="file"
          class="hidden"
          :accept="detailDragAndDrop.acceptFile"
          @change="onChangeUpload"
        />
      </label>
    </div>
  </div>
</template>

<script setup lang="ts">
  import { useDataImage } from '@/store/index'

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
    imageUrl: {
      type: String,
      default: '',
    },
  })

  const emit = defineEmits(['previewFile', 'deleteUrlFile'])

  const files = ref()
  const dataFiles = ref({
    name: '',
    isConfidental: false,
    mimeType: '',
    roles: ['admin'],
    data: '',
    fileSize: '',
    fileCorrect: false,
    height: 0,
    width: 0,
  })
  const fileInputIsChange = ref(false)
  const proggresBarIsSuccess = ref(false)
  const percentageProggres = ref(0)
  const intervalPercentage = ref(null)
  const formatSizeFile = ['Bytes', 'KB', 'MB', 'GB', 'TB']
  const responseImage = ref('')
  const fileIsCorrect = ref(false)
  const disabledButton = ref(true)

  onUpdated(async () => {
    if (props.imageUrl != '') {
      const res = await $fetch(props.imageUrl)

      dataFiles.value.name = ''
      dataFiles.value.mimeType = res.type
      dataFiles.value.fileSize = convertSize(res.size)
      fileInputIsChange.value = true
      fileIsCorrect.value = true

      dataFiles.value.fileCorrect = true
    }
  })

  const onChangeUpload = (e) => {
    if (e.target.files[0]) {
      files.value = e.target.files[0]
      let img = new Image()
      img.src = window.URL.createObjectURL(e.target.files[0])
      img.onload = () => {
        dataFiles.value.width = img.width
        dataFiles.value.height = img.height
      }
      dataFiles.value.name = files.value.name
      dataFiles.value.mimeType = files.value.type
      dataFiles.value.fileSize = convertSize(files.value.size)
      fileInputIsChange.value = true
      convertFileToBase64(files.value)
      runProgressBar()
      checkFileValidation()
      dataFiles.value.fileCorrect = fileIsCorrect.value
      useDataImage().dataImage = JSON.parse(JSON.stringify(dataFiles.value))
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

    emit('deleteUrlFile')

    useDataImage().dataImage = {}
  }

  const previewFile = () => {
    emit('previewFile')
  }

  const convertFileToBase64 = (FileObject: Blob) => {
    const reader = new FileReader()
    reader.onload = () => {
      dataFiles.value.data = reader.result.split(',')[1]

      useDataImage().dataImage = { ...dataFiles.value }
    }
    reader.readAsDataURL(FileObject)
  }

  const checkFileValidation = () => {
    if (files.value) {
      if (fileSizeIsCompatible() && formatFileIsCompatible()) {
        fileIsCorrect.value = true
        disabledButton.value = false
      } else {
        fileIsCorrect.value = false
      }
    } else {
      fileIsCorrect.value = false
    }
  }

  const fileSizeIsCompatible = () => {
    return files.value.size <= props.detailDragAndDrop.maxSizeFile
  }

  const fileResolutionIsCompatible = () => {
    const validate =
      dataFiles.value.width > props.detailDragAndDrop.maxResolution &&
      dataFiles.value.height > props.detailDragAndDrop.maxResolution
    fileIsCorrect.value = !validate
    return validate
  }

  const formatFileIsCompatible = () => {
    return props.detailDragAndDrop.formatTypeFile.includes(files.value.type)
  }
</script>
