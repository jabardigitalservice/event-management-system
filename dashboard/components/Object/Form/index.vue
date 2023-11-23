<template>
  <TheHeader>
    <template #buttonSave>
      <UButton
        size="lg"
        :label="!!state.idData ? `Simpan Perubahan` : `Buat Objek Wisata`"
        class="justify-self-end bg-[#1569C4]"
        color="blue"
        variant="solid"
        @click="handleSubmit"
      />
    </template>
    <template #body>
      <div class="overflow-y-auto">
        <div class="flex h-full max-h-screen flex-col items-center py-5">
          <div class="max-w-xl">
            <Form
              ref="formContainer"
              class="message-notif-form my-5 grid grid-cols-2 gap-x-6 rounded-lg bg-white px-6 py-4"
              :validation-schema="state.schema"
              @submit="onSubmit"
            >
              <div class="col-span-2 mt-3">
                <p class="style-bold text-xl">
                  {{ useRoute().query?.id ? 'Ubah' : 'Tambah' }} Objek Wisata
                </p>
              </div>
              <div class="col-span-2 mt-4">
                <div>
                  <BaseDragAndDropFileMultiple
                    ref="BaseDragAndDropFileMultiple"
                    label="Banner"
                    sublabel="Jenis file (.jpg, .png) resolusi 360 x 170 px. Maksimal 1 MB."
                    height-drag-and-drop="h-[165px]"
                    :disabled="state.isFormDisabled"
                    :detail-drag-and-drop="state.detailDragAndDropMultiple"
                    :image-url-multiple="state.dataUrlImageMultiple"
                    @preview-file="previewFile"
                    @delete-url-file-multiple="deleteImageMultipleUrl"
                  />
                </div>
              </div>
              <div class="col-span-2 mt-4">
                <div>
                  <BaseDragAndDropFile
                    ref="BaseDragAndDropFile"
                    label="Logo"
                    sublabel="Jenis file (.png, .svg) resolusi 46x46, maksimal 1 MB"
                    height-drag-and-drop="h-[200px]"
                    :disabled="state.isFormDisabled"
                    :detail-drag-and-drop="state.detailDragAndDrop"
                    :image-url="state.dataUrlImage"
                    @preview-file="previewFile"
                    @delete-url-file="deleteImageUrl"
                  />
                </div>
              </div>
              <div class="col-span-2">
                <div class="mt-6">
                  <BaseTextInput
                    name="name"
                    type="text"
                    label="Nama Objek Wisata"
                    placeholder="Masukan Nama Objek Wisata"
                    :disabled="state.isFormDisabled"
                  />
                </div>
              </div>
              <div class="col-span-2">
                <div class="mt-6">
                  <BaseDropdown
                    name="organization_id"
                    type="text"
                    label="Organisasi"
                    placeholder="Pilih Organisasi"
                    :data-selected="state.selectedOrganisasi"
                    :data-dropdown="state.dataOrganisasi"
                    :disabled="state.isFormDisabled"
                  />
                </div>
              </div>
              <div class="col-span-2">
                <div class="mt-6">
                  <BaseTextareaInput
                    name="description"
                    label="Deskripsi Objek"
                    placeholder="Masukan Deskripsi Objek"
                    :disabled="state.isFormDisabled"
                  />
                </div>
              </div>
              <div class="col-span-2">
                <div class="mt-6 flex justify-between">
                  <div>
                    <label
                      class="text-gray-800"
                    >
                      Alamat
                    </label>
                    <p class="text-[13px] text-gray-600">
                      Isikan alamat lengkap tempat objek wisata
                    </p>
                  </div>
                  <UButton
                    size="sm"
                    color="blue"
                    square
                    variant="solid"
                    label="Pilih Alamat"
                    :disabled="state.isFormDisabled"
                    @click="handleOpenDialogAddress"
                  />
                </div>
              </div>
              <div class="col-span-2 my-8">
                <hr />
              </div>
              <div class="col-span-2">
                <BaseTextInput
                  name="organizer"
                  type="text"
                  label="Nama Pengelola"
                  placeholder="Masukan Nama Pengelola"
                  :disabled="state.isFormDisabled"
                />
              </div>
              <div class="col-span-2 mt-6">
                <BaseTextInput
                  name="organizer_email"
                  type="text"
                  label="Email Pengelola"
                  placeholder="Masukan Email Pengelola"
                  :disabled="state.isFormDisabled"
                />
              </div>
              <div class="col-span-2 mt-6">
                <BaseTextInput
                  name="organizer_phone"
                  type="number"
                  label="No. Telp Objek Wisata"
                  placeholder="Masukan No. Telp Objek Wisata"
                  :disabled="state.isFormDisabled"
                />
              </div>
              <div class="col-span-2 my-8">
                <hr />
              </div>
              <div class="col-span-2">
                <BaseTextInputGroup
                  name="instagram"
                  label="Instagram"
                  placeholder="https://instagram.com/"
                  icon="common/instagram"
                  :disabled="state.isFormDisabled"
                />
              </div>
              <div class="col-span-2 mt-6">
                <BaseTextInputGroup
                  name="tiktok"
                  label="Tiktok"
                  placeholder="https://tiktok.com/"
                  icon="common/tiktok"
                  :disabled="state.isFormDisabled"
                />
              </div>
              <div class="col-span-2 mt-6">
                <BaseTextInputGroup
                  name="facebook"
                  label="Facebook"
                  placeholder="https://facebook.com/"
                  icon="common/facebook"
                  :disabled="state.isFormDisabled"
                />
              </div>

              <div class="col-span-2 mt-6">
                <BaseTextInputGroup
                  name="x"
                  label="X"
                  placeholder="https://twitter.com/"
                  icon="common/twitter"
                  :disabled="state.isFormDisabled"
                />
              </div>
              <div class="col-span-2 mt-6">
                <BaseTextInputGroup
                  name="youtube"
                  label="Youtube"
                  placeholder="https://youtube.com/"
                  icon="common/youtube"
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
  import { object, string } from 'yup'
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
      name: string().required('Nama Objek Wisata wajib diisi'),
      organizer: string().required('Nama Pengelola wajib diisi'),
      organizer_email: string()
        .email('Format Email Tidak Sesuai')
        .required('Email Pengelola wajib diisi'),
      organizer_phone: string()
        .max(13)
        .required('No. Telp Objek Wisata wajib diisi'),
      description: string().required('Deskripsi Objek Wisata wajib diisi'),
      facebook: string(),
      instagram: string(),
      x: string(),
      organization_id: string().required('Organisasi Wajib Dipilih'),
    },
    detailDragAndDrop: {
      informationSizeCompatible:
        'Ukuran file dokumen SK tidak boleh melebihi 1 MB.',
      informationFormatCompatible:
        'Hanya file yang berformat PNG/SVG yang dapat diupload.',
      formatTypeFile: ['image/png', 'image/svg+xml'],
      maxSizeFile: 1097152,
      acceptFile: '.png,.svg',
      maxResolution: '46',
    },
    detailDragAndDropMultiple: {
      informationSizeCompatible:
        'Ukuran file dokumen SK tidak boleh melebihi 1 MB.',
      informationFormatCompatible:
        'Hanya file yang berformat JPG/JPEG/PNG yang dapat diupload.',
      formatTypeFile: ['image/jpeg', 'image/png', 'image/jpg'],
      maxSizeFile: 1097152,
      acceptFile: '.jpg,.jpeg,.png',
      maxResolution: [360, 170],
    },
    dataImage: {
      showDialog: false,
      fileId: '',
      mimeType: '',
      withUrlPath: false,
    },
    dataUrlImage: '',
    dataUrlImageMultiple: [],
    idData: useRoute().query?.id,
    showDialogAddress: false,
    address: {},
    dataOrganisasi: [],
    isFormDisabled: true,
    selectedOrganisasi: {},
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
    banner: object
    logo: string
    organizer: string
    organizer_email: string
    organizer_phone: string
    social_media: object[]
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
    status: string
    organization_id: string
    organization_name: string
  }

  onMounted(async () => {
    const id = state.idData
    if (id) {
      const responseObject = (await useGetData(
        `/v1/event/object/${id}`,
      )) as apiResponse
      if (responseObject.code === '2000800') {
        const responseObjectData =
          responseObject.data as object as apiDataResponse

        formContainer.value?.setValues({
          name: responseObjectData?.name,
          organizer: responseObjectData?.organizer,
          organizer_email: responseObjectData?.organizer_email,
          organizer_phone: responseObjectData?.organizer_phone,
          instagram: responseObjectData?.social_media[0]?.link,
          tiktok: responseObjectData?.social_media[1]?.link,
          facebook: responseObjectData?.social_media[2]?.link,
          x: responseObjectData?.social_media[3]?.link,
          youtube: responseObjectData?.social_media[4]?.link,
          description: responseObjectData?.description,
          organization_id: responseObjectData?.organization_id,
        })
        state.selectedOrganisasi = {
          id: responseObjectData?.organization_id,
          label: responseObjectData?.organization_name,
        }

        state.address = {
          address: responseObjectData?.address,
          province: responseObjectData?.province,
          province_id: responseObjectData?.province_id,
          city: responseObjectData?.city,
          city_id: responseObjectData?.city_id,
          district: responseObjectData?.district,
          district_id: responseObjectData?.district_id,
          village: responseObjectData?.village,
          village_id: responseObjectData?.village_id,
          google_map: responseObjectData?.google_map,
        }

        state.dataUrlImage = responseObjectData.logo
        state.dataUrlImageMultiple = responseObjectData.banner
      }
    }

    getOrganisasi(1, 10)

    state.isFormDisabled = !state.isFormDisabled
  })

  const getOrganisasi = async (page: number, pageSize: number) => {
    const responseOrganisasi = await useGetData(
      `/v1/event/organization/?page=${page}&pageSize=${pageSize}`,
    )
    if (responseOrganisasi.code === '2000800') {
      const responseOrganisasiData = responseOrganisasi.data.data
      for (let i = 0; i < responseOrganisasiData.length; i++) {
        state.dataOrganisasi.push({
          id: responseOrganisasiData[i].id,
          label: responseOrganisasiData[i].name,
        })
      }
    }
  }

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
    data.banner = await postPhotoServiceMultiple()
    data.social_media = [
      {
        name: 'instagram',
        link: data?.instagram,
      },
      {
        name: 'tiktok',
        link: data?.tiktok,
      },
      {
        name: 'facebook',
        link: data?.facebook,
      },
      {
        name: 'x',
        link: data?.x,
      },
      {
        name: 'youtube',
        link: data?.youtube,
      },
    ]
    data.status = 'draft'

    const isAlamatValidate: boolean = validateAlamat(state.address)
    if (isAlamatValidate) {
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
          usePutData(`/v1/event/object/${state.idData}`, data).then((res) => {
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
      default:
        try {
          usePostData('/v1/event/object', data).then((res) => {
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

  const deleteImageMultipleUrl = (index: number) => {
    state.dataUrlImageMultiple.splice(index, 1)

    state.dataUrlImageMultiple = { ...state.dataUrlImageMultiple }
  }

  const handleBack = () => {
    useRouter().push({ path: '/object' })
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
        title: 'Image Melanggar Rules Atau Kosong',
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

  const postPhotoServiceMultiple = async () => {
    const imageMultiple: object[] = JSON.parse(
      JSON.stringify(useDataImage().dataImageMultiple),
    )
    const imageMultipleLength: number = Object.keys(imageMultiple).length

    const [dataToUpload, dataAreExist] = multipleImageValidate(
      imageMultiple,
      imageMultipleLength,
    )

    if (dataToUpload.length >= 1) {
      const result = await usePostServicePhoto(dataToUpload[0])
      dataToUpload[0] = result.path
    }
    return multipleImageConcatToSameIndex(dataToUpload, dataAreExist)
  }

  const multipleImageValidate = (
    imageMultiple: object[],
    imageMultipleLength: number,
  ) => {
    // this function for validating where the image that get from response or not
    const dataToUpload = []
    const dataAreExist = []
    if (imageMultipleLength >= 1) {
      for (let i = 0; i < imageMultipleLength; i++) {
        if (!imageMultiple[i]?.url) {
          if (!imageMultiple[i].fileCorrect) {
            return toast.add({
              title: 'Banner Image Melanggar Rules',
              color: 'red',
              icon: 'i-heroicons-x-circle',
              timeout: 2500,
            })
          }
          dataToUpload.push(imageMultiple[i])
        } else {
          dataAreExist.push({
            index: i,
            value: imageMultiple[i].url.replace(
              'https://file.digitalservice.id/superapp-utilities-public/',
              '',
            ),
          })
        }
      }
    }

    return [dataToUpload, dataAreExist]
  }

  const multipleImageConcatToSameIndex = (
    dataToUpload: object[],
    dataAreExist: object[],
  ) => {
    // this function for concating to spesfic index
    if (dataAreExist.length >= 1) {
      for (let i = 0; i < dataAreExist.length; i++) {
        dataToUpload.splice(dataAreExist[i].index, 0, dataAreExist[i].value)
      }
    }
    return dataToUpload
  }
</script>
