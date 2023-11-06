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
          'Authorization': `Bearer ${session.accessToken}`,
        },
        query: { page: page?.value, q: search, pageSize: selectedLimit },

        onResponse({ response }) {
          // TODO: remove dummy data after API read or integration with API finished
          const responseOrganizer = {
            data: {
              data: [
                {
                  id: 1,
                  name: 'Sapawarga',
                  address: 'Jl Sultan Hasanuddin',
                  photo: 'daengweb.png',
                  email: 'aljabar@gmail.com',
                  phone_number: '085343966997',
                  status: 1,
                  created_at: '2023-10-02T13:40:13.000000Z',
                  updated_at: '2023-10-18T07:11:28.000000Z',
                },
                {
                  id: 2,
                  name: 'Aljabar',
                  address: 'Jl Sultan Hasanuddin2',
                  photo: 'daengweb.png2',
                  email: 'aljabar1@gmail.com',
                  phone_number: '085343966997',
                  status: 1,
                  created_at: '2023-10-02T13:40:13.000000Z',
                  updated_at: '2023-10-18T07:11:28.000000Z',
                },
                {
                  id: 3,
                  name: 'Saparua',
                  address: 'Jl Sultan Hasanuddin2',
                  photo: 'daengweb.png2',
                  email: 'aljabar2@gmail.com',
                  phone_number: '085343966997',
                  status: 1,
                  created_at: '2023-10-02T13:40:13.000000Z',
                  updated_at: '2023-10-18T07:11:28.000000Z',
                },
                {
                  id: 4,
                  name: 'Sapawarga',
                  address: 'Jl Sultan Hasanuddin4',
                  photo: 'daengweb.png4',
                  email: 'aljabar4@gmail.com',
                  phone_number: '085343966997',
                  status: 1,
                  created_at: '2023-10-02T13:40:13.000000Z',
                  updated_at: '2023-10-18T07:11:28.000000Z',
                },
                {
                  id: 5,
                  name: 'Sapawarga',
                  address: 'Jl Sultan Hasanuddin5',
                  photo: 'daengweb.png5',
                  email: 'aljabar5@gmail.com',
                  phone_number: '085343966997',
                  status: 1,
                  created_at: '2023-10-02T13:40:13.000000Z',
                  updated_at: '2023-10-18T07:11:28.000000Z',
                },
              ],
              meta: {
                page: 1,
                page_size: 10,
                total_data: 1,
                total_pages: 1,
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

export const useDeleteData = async (path: string, id?: Ref<string>) => {
  const config = useRuntimeConfig()
  const session = await getSession()
  return new Promise(async (resolve, reject) => {
    try {
      await useFetch(config.public.baseURL.concat(`${path}/${id?.value}`), {
        method: 'delete',
        headers: {
          'Api-Key': config.public.apiKey,
          'Authorization': `Bearer ${session.accessToken}`,
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

export const usePostData = async (path: string, body: object) => {
  const config = useRuntimeConfig()
  const session = await getSession()
  return new Promise(async (resolve, reject) => {
    try {
      await useFetch(config.public.baseURL.concat(`${path}`), {
        method: 'post',
        headers: {
          'Api-Key': config.public.apiKey,
          'Authorization': `Bearer ${session.accessToken}`
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

export const useUpdatePatchData = async (path: string, id?: Ref<string>, body?: object) => {
  const config = useRuntimeConfig()
  const session = await getSession()
  return new Promise(async (resolve, reject) => {
    try {
      await useFetch(config.public.baseURL.concat(`${path}/${id?.value}`), {
        method: 'patch',
        headers: {
          'Api-Key': config.public.apiKey,
          'Authorization': `Bearer ${session.accessToken}`,
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
