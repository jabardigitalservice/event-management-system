<template>
  <TheHeader>
    <template #buttonSave>
      <UButton
        size="lg"
        :label="!!state.idData ? `Simpan Perubahan` : `Buat Organisasi`"
        class="justify-self-end bg-[#1569C4]"
        color="blue"
        variant="solid"
        @click="handleSubmit"
      />
    </template>
    <template #body>
      <div class="overflow-y-auto">
        <div class="flex h-full max-h-screen flex-col items-center py-5">
          <div class="max-w-xl w-full">
            <Form
              ref="formContainer"
              class="message-notif-form my-5 grid grid-cols-6 gap-x-6 rounded-lg bg-white px-6 py-4"
              :validation-schema="state.schema"
              @submit="onSubmit"
            >
              <div class="col-span-6 mt-3">
                <p class="style-bold text-xl">
                  {{ useRoute().query?.id ? 'Ubah' : 'Tambah' }} Objek Wisata
                </p>
              </div>
              <div class="col-span-6 mt-5">
                <div>
                  <BaseDragAndDropFile
                    ref="BaseDragAndDropFile"
                    label="Logo"
                    sublabel="Tipe File JPG/JPEG/PNG dengan maksimal ukuran file 2 MB"
                    height-drag-and-drop="h-[224px]"
                    :disabled="state.isFormDisabled"
                    :detail-drag-and-drop="state.detailDragAndDrop"
                    :image-url="state.dataUrlImage"
                    @preview-file="previewFile"
                    @delete-url-file="deleteImageUrl"
                  />
                </div>
              </div>
              <div class="col-span-6 mt-5">
                <div>
                  <BaseTextInput
                    name="name"
                    type="text"
                    label="Nama Organisasi"
                    placeholder="Masukan Nama Organisasi"
                    :disabled="state.isFormDisabled"
                  />
                </div>
                <div class="mt-5 flex justify-between">
                  <div>
                    <label
                      class="text-gray-800"
                    >
                      Alamat
                    </label>
                    <p class="text-[13px] text-gray-600">
                      Isikan alamat lengkap organisasi
                    </p>
                  </div>
                  <UButton
                    size="sm"
                    color="blue"
                    square
                    variant="solid"
                    label="Pilih Alamat"
                    class="mt-2"
                    :disabled="state.isFormDisabled"
                    @click="handleOpenDialogAddress"
                  />
                </div>
              </div>

              <div class="col-span-6 mt-5">
                <BaseTextInput
                  name="pic_name"
                  type="text"
                  label="Nama Pengelola"
                  placeholder="Masukan Nama Pengelola"
                  :disabled="state.isFormDisabled"
                />
              </div>
              <div class="col-span-2 mt-5">
                <BaseTextInput
                  name="email"
                  type="text"
                  label="Email"
                  placeholder="Masukan Email"
                  :disabled="state.isFormDisabled"
                />
              </div>
              <div class="col-span-2 mt-5">
                <BaseTextInput
                  name="pic_position"
                  type="text"
                  label="Posisi Pengelola"
                  placeholder="Masukan Posisi Pengelola"
                  :disabled="state.isFormDisabled"
                />
              </div>
              <div class="col-span-2 mt-5">
                <BaseTextInput
                  name="pic_phone"
                  type="number"
                  label="No. Telp Pengelola"
                  placeholder="Masukan No. Telp Pengelola"
                  :disabled="state.isFormDisabled"
                />
              </div>
              <div class="col-span-6 mt-5">
                <BaseTextareaInput
                  name="description"
                  label="Deskripsi Organisasi"
                  placeholder="Masukan Deskripsi Organisasi"
                  :disabled="state.isFormDisabled"
                />
              </div>
              <button v-show="false" ref="submitForm" type="submit">
                Submit
              </button>
            </Form>
          </div>
        </div>
      </div>
    </template>
  </TheHeader>
  <BaseViewFileModal
    title="Logo"
    :show="state.dataImage.showDialog"
    :file="state.dataImage.fileId"
    :mime-type="state.dataImage.mimeType"
    :with-url-path="state.dataImage.withUrlPath"
    @close="state.dataImage.showDialog = false"
  />
  <BaseModalFormAddress
    :open-modal="state.showDialogAddress"
    :address-data="state.address"
    @close="state.showDialogAddress = false"
    @confirm="handleSubmitAddress"
  />
  <UNotifications />
</template>
<script setup lang="ts">
  import { Form } from 'vee-validate'
  import { useDataImage } from '@/store/index'
  import { string } from 'yup'
  import {
    usePostServicePhoto,
    usePostData,
    useGetData,
    usePutData,
  } from '~/composables/useFetchData'

  const formContainer = ref<HTMLInputElement | null>(null)
  const submitForm = ref<HTMLInputElement>()
  const toast = useToast()

  const state = reactive({
    schema: {
      name: string().required('Nama Organisasi wajib diisi'),
      email: string()
        .email('Format Email Tidak Sesuai')
        .required('Email Organisasi wajib diisi'),
      pic_name: string().required('Nama Pengelola wajib diisi'),
      pic_position: string().required('Posisi Pengelola wajib diisi'),
      pic_phone: string().max(13).required('No. Telp Objek Wisata wajib diisi'),
      description: string().required('Deskripsi Objek Wisata wajib diisi'),
    },
    detailDragAndDrop: {
      informationSizeCompatible:
        'Ukuran file dokumen SK tidak boleh melebihi 2 MB.',
      informationFormatCompatible:
        'Hanya file yang berformat JPG/JPEG/PNG yang dapat diupload.',
      formatTypeFile: ['image/jpeg', 'image/png', 'image/jpg'],
      maxSizeFile: 2097152,
      acceptFile: '.jpg,.jpeg,.png',
      maxResolution: '46',
    },
    dataImage: {
      showDialog: false,
      fileId: '',
      mimeType: '',
      withUrlPath: false,
    },
    dataUrlImage: '',
    idData: useRoute().query?.id,
    showDialogAddress: false,
    address: {},
    isFormDisabled: true,
  })

  interface apiResponse {
    code: string
    data: apiDataResponse[]
    message: string
    status: boolean
  }

  interface apiDataResponse {
    id: string
    name: string
    logo: string
    email: string
    pic_name: string
    pic_position: string
    pic_phone: string
    description: string
    address: string
    province: string
    province_id: string
    city: string
    city_id: string
    district: string
    district_id: string
    village: string
    village_id: string
    google_map: string
  }

  onMounted(async () => {
    const id = state.idData
    if (id) {
      const response = (await useGetData(
        `/v1/event/organization/${id}`,
      )) as apiResponse
      if (response.code === '2000800') {
        const responseData = response.data as object as apiDataResponse

        formContainer.value?.setValues({
          name: responseData?.name,
          email: responseData?.email,
          pic_name: responseData?.pic_name,
          pic_position: responseData?.pic_position,
          pic_phone: responseData?.pic_phone,
          description: responseData?.description,
        })

        state.address = {
          address: responseData?.address,
          province: responseData?.province,
          province_id: responseData?.province_id,
          city: responseData?.city,
          city_id: responseData?.city_id,
          district: responseData?.district,
          district_id: responseData?.district_id,
          village: responseData?.village,
          village_id: responseData?.village_id,
          google_map: responseData?.google_map,
        }

        state.dataUrlImage = responseData.logo
      }
    }
    state.isFormDisabled = !state.isFormDisabled
  })

  const previewFile = () => {
    state.dataImage.showDialog = true
    state.dataImage.fileId =
      state.dataUrlImage !== ''
        ? state.dataUrlImage
        : useDataImage().dataImage?.data
    state.dataImage.withUrlPath = state.dataUrlImage !== '' ? true : false
    state.dataImage.mimeType = useDataImage().dataImage?.mimeType
  }

  const handleOpenDialogAddress = () => {
    state.showDialogAddress = true
  }

  const handleSubmitAddress = (address: object) => {
    state.address = address
    state.showDialogAddress = false
  }

  const handleSubmit = async () => {
    submitForm.value?.click()
  }

  const onSubmit = async (values: object) => {
    const method = state.idData ? 'edit' : 'add'

    let data = values as apiDataResponse
    data.logo = state.dataUrlImage
      ? state.dataUrlImage.replace(
          'https://file.digitalservice.id/superapp-utilities-public/',
          '',
        )
      : await postPhotoService()

    const isAddressValid: boolean = validateAlamat(state.address)
    if (isAddressValid) {
      data.address = state.address.address
      data.province = state.address.province
      data.province_id = state.address.province_id
      data.city = state.address.city
      data.city_id = state.address.city_id
      data.district = state.address.district
      data.district_id = state.address.district_id
      data.village = state.address.village
      data.village_id = state.address.village_id
      data.google_map = state.address.google_map
    } else {
      return toast.add({
        title: 'Data Alamat Wajib Diisi',
        color: 'red',
        icon: 'i-heroicons-x-circle',
        timeout: 2500,
      })
    }

    switch (method) {
      case 'edit':
        try {
          usePutData(`/v1/event/organization/${state.idData}`, data).then(
            (res) => {
              if (res.code === '2010800') {
                toast.add({
                  icon: 'i-heroicons-exclamation-triangle',
                  title: 'Data Successfully Added',
                  color: 'green',
                  timeout: 2000,
                })
                handleBack()
              }
            },
          )
        } catch (error) {
          toast.add({
            title: 'Data Failed to Add',
            color: 'red',
            icon: 'i-heroicons-x-circle',
            timeout: 2000,
          })
        }
        break
      default:
        try {
          usePostData('/v1/event/organization', data).then((res) => {
            if (res.code === '2010800') {
              toast.add({
                icon: 'i-heroicons-exclamation-triangle',
                title: 'Data Successfully Added',
                color: 'green',
                timeout: 2000,
              })
              handleBack()
            }
          })
        } catch (error) {
          toast.add({
            title: 'Data Failed to Add',
            color: 'red',
            icon: 'i-heroicons-x-circle',
            timeout: 2000,
          })
        }
        break
    }
  }

  const deleteImageUrl = () => {
    state.dataUrlImage = ''
  }

  const handleBack = () => {
    useRouter().push({ path: '/organisasi' })
  }

  const validateAlamat = (address: object) => {
    if (Object.keys(address).length > 0) {
      for (const key in address) {
        if (address[key] === undefined) {
          return false
        }
      }
      return true
    }
    return false
  }

  const postPhotoService = async () => {
    const image = JSON.parse(JSON.stringify(useDataImage().dataImage))
    const imageLength = Object.keys(image).length

    if (!image.fileCorrect) {
      return toast.add({
        title: 'Image Melanggar Rules',
        color: 'red',
        icon: 'i-heroicons-x-circle',
        timeout: 2500,
      })
    }
    if (imageLength > 0 && image.fileCorrect) {
      const result = await usePostServicePhoto(image)
      return result.path
    }

    return ''
  }
</script>
