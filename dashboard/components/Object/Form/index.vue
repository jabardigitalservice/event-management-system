<template>
  <div>
    <div class="flex justify-between">
      <div>
        <UButton
          size="lg"
          label="Kembali"
          color="green"
          variant="outline"
          icon="i-heroicons-arrow-left-20-solid"
          @click="handleBack"
        />
      </div>
      <div>
        <UButton
          size="lg"
          label="Simpan"
          color="green"
          variant="solid"
          @click="handleSubmit"
        />
      </div>
    </div>
    <Form
      ref="formContainer"
      class="message-notif-form my-5 grid grid-cols-6 gap-x-6 rounded-lg bg-white px-6 py-4"
      :validation-schema="schema"
      @submit="onSubmit"
    >
      <div class="col-span-3">
        <div>
          <BaseTextInput
            name="name"
            type="text"
            label="Nama Objek Wisata"
            placeholder="Masukan Nama Objek Wisata"
          />
        </div>
        <div class="mt-5">
          <BaseTextareaInput
            name="address"
            type="text"
            label="Alamat Objek Wisata"
            placeholder="Masukan Alamat Objek Wisata"
          />
        </div>
        <div class="mt-5">
          <BaseTextInput
            name="organizer"
            type="text"
            label="Organisasi"
            placeholder="Masukan Organisasi Wisata"
          />
        </div>
      </div>
      <div class="col-span-3">
        <div>
          <BaseDragAndDropFile
            ref="BaseDragAndDropFile"
            label="Logo"
            sublabel="Tipe File JPG/JPEG/PNG dengan maksimal ukuran file 2 MB"
            height-drag-and-drop="h-[224px]"
            :detail-drag-and-drop="detailDragAndDrop"
            @preview-file="previewFile"
          />
        </div>
      </div>
      <div class="col-span-2 mt-5">
        <BaseTextInput
          name="facebook"
          label="Facebook"
          placeholder="Masukan Link Facebook"
        />
      </div>
      <div class="col-span-2 mt-5">
        <BaseTextInput
          name="instagram"
          label="Instagram"
          placeholder="Masukan Link Instagram"
        />
      </div>
      <div class="col-span-2 mt-5">
        <BaseTextInput name="x" label="X" placeholder="Masukan Link X" />
      </div>
      <div class="col-span-6 mt-5">
        <BaseTextareaInput
          name="description"
          label="Deskripsi Objek"
          placeholder="Masukan Deskripsi Objek"
        />
        <BaseDragAndDropFileMultiple
          ref="BaseDragAndDropFileMultiple"
          class="mt-5"
          label="Banner"
          sublabel="Tipe File JPG/JPEG/PNG dengan maksimal ukuran file 2 MB"
          height-drag-and-drop="h-[165px]"
          :detail-drag-and-drop="detailDragAndDrop"
          @preview-file="previewFile"
        />
      </div>
      <button v-show="false" ref="submitForm" type="submit">Submit</button>
    </Form>
  </div>
  <BaseViewFileModal
    title="Logo"
    :show="dataImage.showDialog"
    :file="dataImage.fileId"
    :mime-type="dataImage.mimeType"
    @close="dataImage.showDialog = false"
  />
</template>
<script setup lang="ts">
  import { Form } from 'vee-validate'
  import { useDataImage } from '@/store/index'
  import { object, string } from 'yup'
  import { usePostServicePhoto } from '~/composables/useFetchData'

  const router = useRouter()
  const formContainer = ref<HTMLInputElement>()
  const submitForm = ref<HTMLInputElement>()

  const schema = object({
    name: string().required('Nama Objek Wisata wajib diisi'),
    address: string().required('Alamat Objek Wisata wajib diisi'),
    organizer: string().required('Organisasi Objek Wisata wajib diisi'),
    description: string().required('Deskripsi Objek Wisata wajib diisi'),
    facebook: string(),
    instagram: string(),
    x: string(),
  })

  const detailDragAndDrop = ref({
    informationSizeCompatible:
      'Ukuran file dokumen SK tidak boleh melebihi 2 MB.',
    informationFormatCompatible:
      'Hanya file yang berformat JPG/JPEG/PNG yang dapat diupload.',
    formatTypeFile: ['image/jpeg', 'image/png', 'image/jpg'],
    maxSizeFile: 2097152,
    acceptFile: '.jpg,.jpeg,.png',
  })

  const dataImage = ref({
    showDialog: false,
    fileId: '',
    mimeType: '',
  })

  const previewFile = () => {
    dataImage.value.showDialog = true
    dataImage.value.fileId = 'loading'
    dataImage.value.fileId = useDataImage().dataImage?.data
    dataImage.value.mimeType = useDataImage().dataImage?.mimeType
  }

  const handleSubmit = async () => {
    submitForm.value?.click()
  }

  const onSubmit = async (values: unknown) => {
    let data = values
    data.logo = await postPhotoService()
    data.banner = await postPhotoServiceMultiple()
    data.social_media = [
      {
        name: 'x',
        link: data.x,
      },
      {
        name: 'facebook',
        link: data.facebook,
      },
      {
        name: 'instagram',
        link: data.instagram,
      },
    ]
    data.status = 'draft'

    // console.log("waiting for API", data)
  }

  const handleBack = () => {
    router.push({ path: '/object' })
  }

  const postPhotoService = async () => {
    const image = JSON.parse(
      JSON.stringify(useDataImage().dataImage))
    const imageLength = Object.keys(image).length
    
    
    if (imageLength) {
      const result = await usePostServicePhoto(image)
      return result.path
    }

    return ''
  }

  const postPhotoServiceMultiple = async () => {
    const imageMultiple = JSON.parse(
      JSON.stringify(useDataImage().dataImageMultiple),
    )
    const imageMultipleLength = Object.keys(imageMultiple).length
    const data = []
    if (imageMultipleLength >= 1) {
      const result = await usePostServicePhoto(imageMultiple[0])
      data.push(result.path)
      return data
    }
    return data
  }
</script>
