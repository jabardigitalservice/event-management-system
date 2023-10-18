// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  ssr: false,
  devtools: { enabled: true },
  modules: [
    '@nuxtjs/tailwindcss',
    'nuxt-icons',
    '@pinia/nuxt',
    [
      '@nuxtjs/google-fonts',
      {
        families: {
          Lato: true,
          Roboto: true,
          Lora: true,
          Inter: true,
        },
      },
    ],
  ],
  runtimeConfig: {
    // Public keys that are exposed to the client
    public: {
      portalJabarCMSBaseURL: process.env.NUXT_PUBLIC_PORTAL_JABAR_CMS_BASE_URL,
    },
  },
})
