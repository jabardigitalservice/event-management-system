// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  typescript: {
    strict: true,
  },
  ssr: false,
  devtools: { enabled: true },
  modules: [
    '@nuxt/ui',
    '@nuxtjs/tailwindcss',
    'nuxt-icons',
    '@pinia/nuxt',
    '@sidebase/nuxt-auth',
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
  auth: {
    globalAppMiddleware: true,
  },
  colorMode: {
    preference: 'light'
  },
  app: {
    head: {
      title: "Ticketing Management System",
      link: [
        { rel: "icon", type: "image/svg", href: "/logos/sapawarga.svg" }
      ]
    }
  },
  runtimeConfig: {
    // Public keys that are exposed to the client
    public: {
      baseURL: process.env.BASE_URL_API,
      apiKey: process.env.API_KEY,
      baseURLService: process.env.BASE_URL_API_SERVICE,
      apiServiceKey: process.env.API_SERVICE_KEY,
      portalJabarCMSBaseURL: process.env.NUXT_PUBLIC_PORTAL_JABAR_CMS_BASE_URL,
    }
  },
})
