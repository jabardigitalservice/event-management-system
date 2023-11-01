// file: ~/server/api/auth/[...].ts
import { NuxtAuthHandler } from '#auth'
import KeycloakProvider from 'next-auth/providers/keycloak'

export default NuxtAuthHandler({
  secret: process.env.KEYCLOAK_CLIENT_SECRET,
  pages: {
    signIn: '/login',
  },
  callbacks: {
    jwt: async ({token, account, user}) => {
      const isSignIn = user ? true : false;
      if (isSignIn) {
        token.accessToken = account.access_token
        token.idToken = account.id_token
      }
      return Promise.resolve(token);
    },
    session: async ({token, session}) => {
      session.accessToken = token.accessToken;
      session.idToken = token.idToken;
      return Promise.resolve(session);
    },
  },
  events: {
    signOut: async ({ token }) => {
      await $fetch(`${process.env.KEYCLOAK_ISSUER}/protocol/openid-connect/logout?id_token_hint=${token.idToken}`)
    },
  },
  providers: [
    KeycloakProvider.default({
      clientId: process.env.KEYCLOAK_CLIENT_ID,
      clientSecret: process.env.KEYCLOAK_CLIENT_SECRET,
      issuer: process.env.KEYCLOAK_ISSUER,
    }),
  ],
})

