const { getSession } = useAuth()

export const useFetchData = async (
  path: string,
  page?: Ref<number>,
  search?: Ref<string>,
  selectedLimit?: Ref<number>,
) => {
  const config = useRuntimeConfig()
  const session = await getSession()
  return new Promise(async (resolve, reject) => {
    try {
      await useFetch(config.public.baseURL.concat(`${path}`), {
        method: 'get',
        headers: {
          'Api-Key': config.public.apiKey,
          Authorization: `Bearer ${session.accessToken}`,
        },
        query: { page: page?.value, q: search, pageSize: selectedLimit },

        onResponse({ response }) {
          // TODO: remove dummy data after API read or integration with API finished
          const responseOrganizer = {
            data: {
              data: [
                {
                  id: 'a2f7fc23-054f-4037-86c8-cd9a60cf51c6',
                  name: 'Sapawarga',
                  email: 'sapawarga@gmail.com',
                  address: 'bandung, City',
                  description: 'This is a sample organization.',
                  logo: 'google.com/https://example.com/logo.png',
                  created_at: '2023-11-10T14:38:03.945726Z',
                  updated_at: '2023-11-10T14:38:03.945726Z',
                  province: 'SampleProvince',
                  city: 'SampleCity',
                  district: 'SampleDistrict',
                  village: 'SampleVillage',
                  google_map: 'SampleGoogleMap',
                  pic_name: 'SamplePicName',
                  pic_position: 'SamplePicPosition',
                  pic_phone: '+1-123-456-789a0',
                },
              ],
              meta: {
                page: 1,
                page_size: 1,
                total_data: 6,
                total_pages: 6,
              },
            },
          }
          if (path === '/organizer') {
            resolve(responseOrganizer)
          } else {
            resolve(response._data)
          }
        },
      })
    } catch (error) {
      reject(error)
    }
  })
}

export const useFetchAddress = async (
  path: string,
  params: { provinceId?: string; cityId?: string; districtId?: string } = {},
) => {
  const config = useRuntimeConfig()
  const { provinceId, cityId, districtId } = params
  return new Promise(async (resolve, reject) => {
    try {
      await useFetch(config.public.baseURL.concat(`${path}`), {
        method: 'get',
        headers: {
          'Api-Key': config.public.apiKey,
          'Content-Type': 'application/json',
        },
        query: { provinceId, cityId, districtId },

        onResponse({ response }) {
          resolve(response._data)
        },
      })
    } catch (error) {
      reject(error)
    }
  })
}

export const useDeleteData = async (path: string, id?: string) => {
  const config = useRuntimeConfig()
  const session = await getSession()
  return new Promise(async (resolve, reject) => {
    try {
      await useFetch(config.public.baseURL.concat(`${path}/${id}`), {
        method: 'delete',
        headers: {
          'Api-Key': config.public.apiKey,
          Authorization: `Bearer ${session.accessToken}`,
        },
        onResponse({ response }) {
          resolve(response._data.data)
        },
      })
    } catch (error) {
      reject(error)
    }
  })
}

export const usePostServicePhoto = async (body: object) => {
  const config = useRuntimeConfig()
  return new Promise(async (resolve, reject) => {
    try {
      await useFetch(config.public.baseURLService, {
        method: 'post',
        headers: {
          'Api-Key': config.public.apiServiceKey,
          'Content-Type': 'application/json',
        },
        body: body,
        onResponse({ response }) {
          resolve(response._data.data)
        },
      })
    } catch (error) {
      reject(error)
    }
  })
}

export const useGetData = async (path: string) => {
  const config = useRuntimeConfig()
  const session = await getSession()
  return new Promise(async (resolve, reject) => {
    try {
      await useFetch(config.public.baseURL.concat(`${path}`), {
        method: 'get',
        headers: {
          'Api-Key': config.public.apiKey,
          Authorization: `Bearer ${session.accessToken}`,
        },
        onResponse({ response }) {
          resolve(response._data)
        },
      })
    } catch (error) {
      reject(error)
    }
  })
}

export const usePostData = async (path: string, body: object) => {
  const config = useRuntimeConfig()
  const session = await getSession()
  return new Promise(async (resolve, reject) => {
    try {
      await useFetch(config.public.baseURL.concat(`${path}`), {
        method: 'post',
        headers: {
          'Api-Key': config.public.apiKey,
          Authorization: `Bearer ${session.accessToken}`,
        },
        body: body,
        onResponse({ response }) {
          resolve(response._data)
        },
      })
    } catch (error) {
      reject(error)
    }
  })
}

export const usePutData = async (path: string, body: object) => {
  const config = useRuntimeConfig()
  const session = await getSession()
  return new Promise(async (resolve, reject) => {
    try {
      await useFetch(config.public.baseURL.concat(`${path}`), {
        method: 'put',
        headers: {
          'Api-Key': config.public.apiKey,
          Authorization: `Bearer ${session.accessToken}`,
        },
        body: body,
        onResponse({ response }) {
          resolve(response._data)
        },
      })
    } catch (error) {
      reject(error)
    }
  })
}

export const useUpdatePatchData = async (
  path: string,
  id?: string,
  body?: object,
) => {
  const config = useRuntimeConfig()
  const session = await getSession()
  return new Promise(async (resolve, reject) => {
    try {
      await useFetch(config.public.baseURL.concat(`${path}/${id}`), {
        method: 'patch',
        headers: {
          'Api-Key': config.public.apiKey,
          Authorization: `Bearer ${session.accessToken}`,
        },
        body: body,
        onResponse({ response }) {
          resolve(response._data.data)
        },
      })
    } catch (error) {
      reject(error)
    }
  })
}
