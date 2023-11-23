<template>
  <TransitionRoot appear :show="props.openModal" as="template">
    <Dialog as="div" class="relative z-10" @close="closedModal">
      <TransitionChild
        as="template"
        enter="duration-300 ease-out"
        enter-from="opacity-0"
        enter-to="opacity-100"
        leave="duration-200 ease-in"
        leave-from="opacity-100"
        leave-to="opacity-0"
      >
        <div class="fixed inset-0 bg-black bg-opacity-25" />
      </TransitionChild>

      <div class="fixed inset-0 overflow-y-auto">
        <div
          class="flex min-h-full items-center justify-center p-4 text-center"
        >
          <TransitionChild
            as="template"
            enter="duration-300 ease-out"
            enter-from="opacity-0 scale-95"
            enter-to="opacity-100 scale-100"
            leave="duration-200 ease-in"
            leave-from="opacity-100 scale-100"
            leave-to="opacity-0 scale-95"
          >
            <DialogPanel
              class="w-full max-w-md transform rounded-2xl bg-white p-6 text-left align-middle shadow-xl transition-all"
            >
              <div class="sm:items-start">
                <UFormGroup label="Provinsi" class="mb-4 w-full">
                  <USelectMenu
                    v-model="state.selectedProvince"
                    :options="state.dataProvince"
                    searchable
                    :ui="{ color: 'blue' }"
                    @click="getProvince"
                    @change="changeProvince"
                  >
                    <template #label>
                      <span v-if="!state.selectedProvince?.id" class="truncate">
                        Pilih Provinsi
                      </span>
                    </template>
                    <template #option-empty="{ query }">
                      <q>{{ query }}</q> not found
                    </template>
                  </USelectMenu>
                </UFormGroup>
                <UFormGroup label="Kota/Kabupaten" class="mb-4 w-full">
                  <USelectMenu
                    v-model="state.selectedCity"
                    :disabled="!state.selectedProvince?.id"
                    :options="state.dataCity"
                    searchable
                    :ui="{ color: 'blue' }"
                    @click="getCity"
                    @change="changeCity"
                  >
                    <template #label>
                      <span v-if="!state.selectedCity?.id" class="truncate">
                        Pilih Kota/Kabupaten
                      </span>
                    </template>
                    <template #option-empty="{ query }">
                      <q>{{ query }}</q> not found
                    </template>
                  </USelectMenu>
                </UFormGroup>
                <UFormGroup label="Kecamatan" class="mb-4 w-full">
                  <USelectMenu
                    v-model="state.selectedDistrict"
                    :disabled="!state.selectedCity?.id"
                    :options="state.dataDistrict"
                    searchable
                    :ui="{ color: 'blue' }"
                    @click="getDistrict"
                    @change="state.selectedVillage = []"
                  >
                    <template #label>
                      <span v-if="!state.selectedDistrict?.id" class="truncate">
                        Pilih Kecamatan
                      </span>
                    </template>
                    <template #option-empty="{ query }">
                      <q>{{ query }}</q> not found
                    </template>
                  </USelectMenu>
                </UFormGroup>
                <UFormGroup label="Desa/Kelurahan" class="mb-4 w-full">
                  <USelectMenu
                    v-model="state.selectedVillage"
                    :disabled="!state.selectedDistrict?.id"
                    :options="state.dataVillage"
                    searchable
                    :ui="{ color: 'blue' }"
                    @click="getVillage"
                  >
                    <template #label>
                      <span v-if="!state.selectedVillage?.id" class="truncate">
                        Pilih Desa/Kelurahan
                      </span>
                    </template>
                    <template #option-empty="{ query }">
                      <q>{{ query }}</q> not found
                    </template>
                  </USelectMenu>
                </UFormGroup>
                <UFormGroup label="Alamat Lengkap" class="mb-4 w-full" required>
                  <p class="text-[13px] text-gray-600">
                    Silahkan masukkan detail lokasi tambahan, seperti nama jalan
                    atau no. bangunan
                  </p>
                  <UTextarea
                    v-model="state.address"
                    placeholder="Masukan Alamat Lengkap Anda"
                    class="mt-1"
                    :ui="{ color: 'blue' }"
                  />
                </UFormGroup>
                <UFormGroup
                  label="Link Google Maps"
                  class="mb-4 w-full"
                  required
                >
                  <UInput
                    v-model="state.google_map"
                    placeholder="Masukan Link Google Maps Anda"
                    class="mt-1"
                    :ui="{ color: 'blue' }"
                  >
                    <template #leading>
                      <NuxtIcon
                        name="common/map-icon"
                        class="text-2xl text-gray-300"
                      />
                    </template>
                  </UInput>
                </UFormGroup>
              </div>
              <div class="mt-6 sm:flex sm:flex-row-reverse">
                <UButton
                  color="blue"
                  class="w-28 justify-center"
                  variant="solid"
                  @click="confirmModal"
                  >Simpan
                </UButton>
                <UButton
                  color="blue"
                  class="mr-3 w-20 justify-center"
                  variant="soft"
                  @click="closedModal"
                >
                  Batal
                </UButton>
              </div>
            </DialogPanel>
          </TransitionChild>
        </div>
      </div>
    </Dialog>
  </TransitionRoot>
</template>

<script setup lang="ts">
  import {
    TransitionRoot,
    TransitionChild,
    Dialog,
    DialogPanel,
  } from '@headlessui/vue'

  const state = reactive({
    selectedProvince: [],
    dataProvince: [],
    selectedCity: [],
    dataCity: [],
    selectedDistrict: [],
    dataDistrict: [],
    selectedVillage: [],
    dataVillage: [],
    address: '',
    google_map: '',
  })

  const getProvince = async () => {
    const result = await useFetchAddress('/v1/area/province')
    state.dataProvince = result.data.map((data) => ({
      id: data.id,
      label: data.name,
    }))
  }

  const changeProvince = async () => {
    state.selectedCity = []
    state.selectedDistrict = []
    state.selectedVillage = []
  }

  const getCity = async () => {
    const result = await useFetchAddress('/v1/area/city', {
      provinceId: state.selectedProvince.id,
    })
    state.dataCity = result.data.map((data) => ({
      id: data.id,
      label: data.name,
    }))
  }

  const changeCity = async () => {
    state.selectedDistrict = []
    state.selectedVillage = []
  }

  const getDistrict = async () => {
    const result = await useFetchAddress('/v1/area/district', {
      cityId: state.selectedCity.id,
    })
    state.dataDistrict = result.data.map((data) => ({
      id: data.id,
      label: data.name,
    }))
  }

  const getVillage = async () => {
    const result = await useFetchAddress('/v1/area/village', {
      districtId: state.selectedDistrict.id,
    })
    state.dataVillage = result.data.map((data) => ({
      id: data.id,
      label: data.name,
    }))
  }

  const props = defineProps({
    openModal: {
      type: Boolean,
      default: false,
    },
    addressData: {
      type: Object,
      default: () => ({}),
    },
  })

  const emit = defineEmits(['close', 'confirm'])

  onUpdated(() => {
    state.address = props.addressData.address
    state.selectedProvince = {
      label: props.addressData.province,
      id: props.addressData.province_id,
    }
    state.selectedCity = {
      label: props.addressData.city,
      id: props.addressData.city_id,
    }
    state.selectedDistrict = {
      label: props.addressData.district,
      id: props.addressData.district_id,
    }
    state.selectedVillage = {
      label: props.addressData.village,
      id: props.addressData.village_id,
    }

    state.google_map = props.addressData.google_map
  })

  function closedModal() {
    emit('close')
  }
  function confirmModal() {
    emit('confirm', {
      province: state.selectedProvince.label,
      province_id: state.selectedProvince.id,
      city: state.selectedCity.label,
      city_id: state.selectedCity.id,
      district: state.selectedDistrict.label,
      district_id: state.selectedDistrict.id,
      village: state.selectedVillage.label,
      village_id: state.selectedVillage.id,
      address: state.address,
      google_map: state.google_map,
    })
  }
</script>
