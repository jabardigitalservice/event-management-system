<template>
  <div>
    <Head>
      <Title>Event Management System</Title>
    </Head>
    <div class="grid h-screen min-w-0 grid-cols-[260px,1fr] overflow-hidden">
      <TheSidebar />
      <div class="w-full">
        <TheHeader />
        <main class="h-[calc(100vh-72px)] overflow-y-auto bg-gray-100">
          <div class="h-full w-full p-4 mt-5">
            <div class="mx-10 bg-white p-7 rounded-lg">
              <slot />
            </div>
          </div>
        </main>
      </div>
    </div>
  </div>
</template>

<style>
  .nuxt-icon svg {
    @apply mb-0;
  }
</style>
