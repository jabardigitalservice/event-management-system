<template>
  <div class="w-full text-center">
    <img src="/images/tiket.svg" class="mt-7 h-[150px] w-[400px]" />
    <p class="mt-4 text-[14px] font-[500]">Belum Ada Tiket</p>
    <p class="mt-1 text-[12px] font-[400]">
      Silahkan atur terlebih dahulu tiket pada tombol dibawah.
    </p>
    <UButton
      class="mt-4 h-9"
      color="blue"
      variant="outline"
      :ui="{ 
        rounded: 'rounded-lg', 
        variant: { 
          outline: 'text-[#1569C4]'
          } 
        }"
    >
      Buat Tiket</UButton
    >
  </div>
</template>
