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
                    v-model="selectedProvince"
                    :options="dataProvince"
                    searchable
                    @click="getProvince"
                    @change="changeProvince"
                  >
                    <template #label>
                      <span
                        v-if="selectedProvince.length === 0"
                        class="truncate"
                        >Pilih Provinsi</span
                      >
                    </template>
                    <template #option-empty="{ query }">
                      <q>{{ query }}</q> not found
                    </template>
                  </USelectMenu>
                </UFormGroup>
                <UFormGroup label="Kota/Kabupaten" class="mb-4 w-full">
                  <USelectMenu
                    v-model="selectedCity"
                    :disabled="!!selectedProvince.values"
                    :options="dataCity"
                    searchable
                    @click="getCity"
                    @change="changeCity"
                  >
                    <template #label>
                      <span 
                        v-if="selectedCity.length === 0" 
                        class="truncate"
                        >Pilih Kota/Kabupaten</span
                      >
                    </template>
                    <template #option-empty="{ query }">
                      <q>{{ query }}</q> not found
                    </template>
                  </USelectMenu>
                </UFormGroup>
                <UFormGroup label="Kecamatan" class="mb-4 w-full">
                  <USelectMenu
                    v-model="selectedDistrict"
                    :disabled="!!selectedCity.values"
                    :options="dataDistrict"
                    searchable
                    @click="getDistrict"
                    @change="selectedVillage.value = []"
                  >
                    <template #label>
                      <span
                        v-if="selectedDistrict.length === 0"
                        class="truncate"
                        >Pilih Kecamatan</span
                      >
                    </template>
                    <template #option-empty="{ query }">
                      <q>{{ query }}</q> not found
                    </template>
                  </USelectMenu>
                </UFormGroup>
                <UFormGroup label="Desa/Kelurahan" class="w-full">
                  <USelectMenu
                    v-model="selectedVillage"
                    :disabled="!!selectedDistrict.values"
                    :options="dataVillage"
                    searchable
                    @click="getVillage"
                  >
                    <template #label>
                      <span 
                        v-if="selectedVillage.length === 0"
                        class="truncate"
                        >Pilih Desa/Kelurahan</span
                      >
                    </template>
                    <template #option-empty="{ query }">
                      <q>{{ query }}</q> not found
                    </template>
                  </USelectMenu>
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

  const selectedProvince = ref([])
  const dataProvince = ref([])

  const selectedCity = ref([])
  const dataCity = ref([])

  const selectedDistrict = ref([])
  const dataDistrict = ref([])

  const selectedVillage = ref([])
  const dataVillage = ref([])

  const getProvince = async () => {
    const result = await useFetchAddress('/v1/area/province')
    dataProvince.value = result.data.map((data) => ({
      id: data.id,
      label: data.name,
    }))
  }

  const changeProvince = async () => {
    selectedCity.value = []
    selectedDistrict.value = []
    selectedVillage.value = []
  }

  const getCity = async () => {
    const result = await useFetchAddress('/v1/area/city', {
      provinceId: selectedProvince.value.id,
    })
    dataCity.value = result.data.map((data) => ({
      id: data.id,
      label: data.name,
    }))
  }

  const changeCity = async () => {
    selectedDistrict.value = []
    selectedVillage.value = []
  }

  const getDistrict = async () => {
    const result = await useFetchAddress('/v1/area/district', {
      cityId: selectedCity.value.id,
    })
    dataDistrict.value = result.data.map((data) => ({
      id: data.id,
      label: data.name,
    }))
  }

  const getVillage = async () => {
    const result = await useFetchAddress('/v1/area/village', {
      districtId: selectedDistrict.value.id,
    })
    dataVillage.value = result.data.map((data) => ({
      id: data.id,
      label: data.name,
    }))
  }

  const props = defineProps({
    openModal: {
      type: Boolean,
      default: false,
    },
  })

  const emit = defineEmits(['close', 'confirm'])

  function closedModal() {
    emit('close')
  }
  function confirmModal() {
    emit('confirm')
  }
</script>
