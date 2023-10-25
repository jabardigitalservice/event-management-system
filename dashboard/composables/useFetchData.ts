import { Ref } from 'vue'

export const useFetchData = async (path:string, page: Ref<number>) => {
  const config = useRuntimeConfig();
  return new Promise(async (resolve, reject) => {
    try {
      
      await useFetch(config.public.baseURL.concat(path), {
        query: { page: page.value },
        onResponse({ response }) {
          resolve(response._data.data);
        },
      });
    } catch (error) {
      reject(error);
    }
  });
}
