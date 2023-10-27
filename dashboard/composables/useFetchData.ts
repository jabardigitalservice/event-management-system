
import { Ref } from 'vue'

export const useFetchData = async (path:string, page?: Ref<number>, search?: Ref<string>) => {
  const config = useRuntimeConfig();
  return new Promise(async (resolve, reject) => {
    try {
      await useFetch(path, {
        query: { page: page?.value, q: search },
        onResponse({ response }) {
           // TODO: remove dummy data after API read or integration with API finished
          const responseOrganizer = {
            "current_page": 1,
            "data": [
                {
                    "id": 1,
                    "name": "Sapawarga",
                    "address": "Jl Sultan Hasanuddin",
                    "photo": "daengweb.png",
                    "email": "aljabar@gmail.com",
                    "phone_number": "085343966997",
                    "status": 1,
                    "created_at": "2023-10-02T13:40:13.000000Z",
                    "updated_at": "2023-10-18T07:11:28.000000Z"
                },
                {
                    "id": 2,
                    "name": "Aljabar",
                    "address": "Jl Sultan Hasanuddin2",
                    "photo": "daengweb.png2",
                    "email": "aljabar1@gmail.com",
                    "phone_number": "085343966997",
                    "status": 1,
                    "created_at": "2023-10-02T13:40:13.000000Z",
                    "updated_at": "2023-10-18T07:11:28.000000Z"
                },
                {
                    "id": 3,
                    "name": "Saparua",
                    "address": "Jl Sultan Hasanuddin2",
                    "photo": "daengweb.png2",
                    "email": "aljabar2@gmail.com",
                    "phone_number": "085343966997",
                    "status": 1,
                    "created_at": "2023-10-02T13:40:13.000000Z",
                    "updated_at": "2023-10-18T07:11:28.000000Z"
                },
                {
                    "id": 4,
                    "name": "Sapawarga",
                    "address": "Jl Sultan Hasanuddin4",
                    "photo": "daengweb.png4",
                    "email": "aljabar4@gmail.com",
                    "phone_number": "085343966997",
                    "status": 1,
                    "created_at": "2023-10-02T13:40:13.000000Z",
                    "updated_at": "2023-10-18T07:11:28.000000Z"
                },
                {
                    "id": 5,
                    "name": "Sapawarga",
                    "address": "Jl Sultan Hasanuddin5",
                    "photo": "daengweb.png5",
                    "email": "aljabar5@gmail.com",
                    "phone_number": "085343966997",
                    "status": 1,
                    "created_at": "2023-10-02T13:40:13.000000Z",
                    "updated_at": "2023-10-18T07:11:28.000000Z"
                }
            ],
            "from": 1,
            "last_page": 1,
            "per_page": 10,
            "to": 5,
            "total": 5
          }
          const responseObject =  {
            "current_page": 1,
            "data": [
                {
                    "id": 1,
                    "name": "Al-Jabar",
                    "description": "object Masjid Aljabar",
                    "created_at": null,
                    "updated_at": null,
                    "location_name": null,
                    "address": "Gedebage",
                    "date": null,
                    "organizer": "Sapawarga",
                    "status": "draft"
                },
                {
                    "id": 2,
                    "name": "Museum Sribaduga",
                    "description": "ini museum",
                    "created_at": null,
                    "updated_at": null,
                    "location_name": null,
                    "address": "bandung",
                    "date": null,
                    "organizer": "Sapawarga",
                    "status": "published"
                }
            ],
            "from": 1,
            "last_page": 1,
            "per_page": 10,
            "to": 2,
            "total": 2
          }
          if(path === '/organizer'){
            resolve(responseOrganizer)
          }else{
            resolve(responseObject)
          }
        },
      });
    } catch (error) {
      reject(error);
    }
  });
}

export const useDeleteData = async (path:string, id?: Ref<number>) => {
  const config = useRuntimeConfig();
  return new Promise(async (resolve, reject) => {
    try {
      
      await useFetch(config.public.baseURL.concat(`${path}/${id?.value}`), {
        method: 'delete',
        onResponse({ response }) {
          resolve(response._data.data);
        },
      });
    } catch (error) {
      reject(error);
    }
  });
}