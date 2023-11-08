// file: ~/server/api/auth/[...].ts
import { NuxtAuthHandler } from '#auth'
import KeycloakProvider from 'next-auth/providers/keycloak'


async function refreshAccessToken(refreshToken: {
  accessToken: string;
  accessTokenExpires: string;
  refreshToken: string;
}) {
  try {
    const refreshedTokens = await $fetch<{
      data: {
        access_token: string;
        expires_in: number;
        refresh_token: string;
      };
    } | null>(`${process.env.KEYCLOAK_ISSUER}/protocol/openid-connect/token`, {
      method: "POST",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      body: new URLSearchParams({
        'grant_type': 'refresh_token',
        'client_id': process.env.KEYCLOAK_CLIENT_ID,
        'client_secret': process.env.KEYCLOAK_CLIENT_SECRET,
        'refresh_token': refreshToken.refreshToken
      }),
    });
    return {
      ...refreshToken,
      accessToken: refreshedTokens.access_token,
      accessTokenExpires: Date.now() + refreshedTokens.expires_in,
      refreshToken: refreshedTokens.refresh_token,
    };
  } catch (error) {
    return {
      ...refreshToken,
      error: "RefreshAccessTokenError",
    };
  }
}

export default NuxtAuthHandler({
  secret: process.env.KEYCLOAK_CLIENT_SECRET,
  pages: {
    signIn: '/login',
  },
  callbacks: {
    jwt: async ({token, account, user}) => {
      const isSignIn = user ? true : false;
      
      if (isSignIn) {
        token.accessToken = account?.access_token
        token.accessTokenExpires = account?.expires_at
        token.refreshToken = account?.refresh_token
        token.idToken = account?.id_token
      }

      if (token.accessTokenExpires && Date.now() > token.accessTokenExpires) {
        return refreshAccessToken(token)
      }

      return Promise.resolve(token);
    },
    session: async ({token, session}) => {
      session.accessToken = token.accessToken
      session.refreshToken = token.refreshToken
      session.idToken = token.idToken
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

