// file: ~/server/api/auth/[...].ts
import { NuxtAuthHandler } from '#auth'
import KeycloakProvider from 'next-auth/providers/keycloak'

export default NuxtAuthHandler({
  secret: process.env.KEYCLOAK_CLIENT_SECRET,
  pages: {
    signIn: '/login',
  },
  providers: [
    KeycloakProvider.default({
      clientId: process.env.KEYCLOAK_CLIENT_ID,
      clientSecret: process.env.KEYCLOAK_CLIENT_SECRET,
      issuer: process.env.KEYCLOAK_ISSUER,
    }),
  ],
})
