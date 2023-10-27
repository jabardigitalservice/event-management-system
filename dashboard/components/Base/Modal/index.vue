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
      leave-to="opacity-0">
        <div class="fixed inset-0 bg-black bg-opacity-25" />
      </TransitionChild>

      <div class="fixed inset-0 overflow-y-auto">
        <div class="flex min-h-full items-center justify-center p-4 text-center">
          <TransitionChild
          as="template"
          enter="duration-300 ease-out"
          enter-from="opacity-0 scale-95"
          enter-to="opacity-100 scale-100"
          leave="duration-200 ease-in"
          leave-from="opacity-100 scale-100"
          leave-to="opacity-0 scale-95">
            <DialogPanel class="w-full max-w-md transform overflow-hidden rounded-2xl bg-white p-6 text-left align-middle shadow-xl transition-all">
              <div class="sm:flex sm:items-start">
                <div :class=circleAlert>
                  <UIcon :name= props.iconModal :class= colorIcon />
                </div>
                <div class="mt-3 text-center sm:ml-4 sm:mt-0 sm:text-left">
                    <h3 id="modal-title" class="text-base font-semibold leading-6 text-gray-900">{{ props.titleModal }}</h3>
                    <div class="mt-2">
                    <p class="text-sm text-gray-500">{{ props.descModal }}</p>
                    </div>
                </div>
              </div>
              <div class="mt-6 sm:flex sm:flex-row-reverse">
                <UButton 
                    color="red" 
                    class="w-28 justify-center" 
                    variant="solid" 
                    @click="confirmModal">{{ props.textConfirm }}
                </UButton>
                <UButton 
                    color="green" 
                    class="mr-3 w-20 justify-center" 
                    variant="soft"
                    @click="closedModal"> Cancel
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

  let colorIcon: string
  let circleAlert: string

  const props = defineProps({
    openModal: {
      type: Boolean,
      default: false
    },
    titleModal: {
      type: String,
      default: ''
    },
    descModal: {
      type: String,
      default: ''
    },
    iconModal: {
      type: String,
      default: ''
    },
    textConfirm: {
      type: String,
      default: ''
    },
    typeModal: {
      type: String,
      default: ''
    }
  })

  switch (props.typeModal) {
    case 'danger':
        colorIcon = 'bg-red-600'
        circleAlert = 'mx-auto flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-full bg-red-100 sm:mx-0 sm:h-10 sm:w-10'
      break;
    case 'warning':
        colorIcon = 'bg-orange-600'
        circleAlert = 'mx-auto flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-full bg-orange-100 sm:mx-0 sm:h-10 sm:w-10'
      break;
    default:
        colorIcon = 'bg-green-600'
        circleAlert = 'mx-auto flex h-12 w-12 flex-shrink-0 items-center justify-center rounded-full bg-green-100 sm:mx-0 sm:h-10 sm:w-10'
      break;
  }

  const emit = defineEmits(['close', 'confirm']);
  
  function closedModal(){
    emit('close')
  }
  function confirmModal(){
    emit('confirm')
  }

</script>