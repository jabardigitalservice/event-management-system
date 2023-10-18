<template>
    <div class="min-h-screen bg-[#E5E5E5]">
      <div class="flex h-screen items-center justify-center">
        <div
          class="h-fit w-[470px] rounded-lg border border-[#EAEAEF] bg-white px-[24px] py-[40px] drop-shadow-md"
        >
          <div class="container mx-auto">
            <HeaderLogin />
            <div class="w-100 mt-[44px] flex items-center justify-center">
              <button
                class="center px-100 h-10 w-[300px] rounded-lg bg-green-600 text-white hover:bg-green-700"
                @click="loginKeyclock"
              >
                <div
                  class="ml-4 mr-4 flex items-center justify-between font-lato text-[14px] font-bold"
                >
                  Login Akun
                  <NuxtIcon
                    name="common/chevron-right"
                    :class="{
                      'eas text-white-700 pointer-events-none text-xl ease-in': true,
                    }"
                    aria-hidden="true"
                  />
                </div>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script lang="ts">
    import HeaderLogin from '~/components/Login/Header.vue'
    const { signIn } = useAuth()
  
    definePageMeta({
      layout: false,
      auth: {
        navigateAuthenticatedTo: '/',
      },
    })
    export default {
      components: {
        HeaderLogin,
      },
      methods: {
        loginKeyclock() {
          signIn('keycloak')
        },
      },
    }
  </script>
  