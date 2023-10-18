<template>
  <div>
    <div class="flex justify-center">
      <div
        class="h-[72px] w-[72px] rounded-lg bg-[#F6F6F9] px-[18px] py-[10px]"
      >
        <!-- <Logo class="h-[52px] w-[37px]" /> -->
        <img
          src="/logos/sapawarga.svg"
          alt="Portal Jabar Logo"
          width="37"
          height="52"
          class="mb-1"
        />
        <NuxtIcon
          name="logo/sapawarga"
          class="h-[52px] w-[37px]"
          aria-hidden="true"
        />
      </div>
    </div>
    <h1
      class="mt-[24px] text-center font-roboto text-[24px] font-bold text-blue-gray-900"
    >
      Selamat Datang!
    </h1>
    <p class="text-center font-lato text-gray-700">
      Silakan login ke akun admin Anda.
    </p>
  </div>
</template>

<script>
  // import Logo from '~/assets/logo/sapawarga.svg?inline'
  export default {
    name: 'HeaderLogin',
    // components: { Logo },
  }
</script>
