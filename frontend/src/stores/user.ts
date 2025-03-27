import { defineStore } from 'pinia'
import { DefaultApi, type User } from 'trippify-client/api'

const LOCAL_STORAGE_KEY = 'trippify-user'

let apiConfig = createApiConfig()
let api = new DefaultApi(apiConfig)

function refreshConfig() {
  apiConfig = createApiConfig()
  api = new DefaultApi(apiConfig)
}

function createApiConfig(): any {
  return {
    isJsonMime: () => false,
    accessToken: getAccessToken()
  }
}

function getAccessToken(): string {
  return localStorage.getItem('accessToken') || ''
}

export const useUserStore = defineStore('user', {
  state: () => ({
    user: {} as User
  }),

  getters: {
    isTravelerOf(tripId) {
      return true
    }
  },

  actions: {
    async getUserFromLocalStorage() {
      refreshConfig()
      const userRaw = localStorage.getItem(LOCAL_STORAGE_KEY)
      if (!userRaw) {
        return
      }
      const user = JSON.parse(userRaw)
      this.$patch({ user })
    },

    async login(username: string, password: string) {
      const jwt = (await api.loginPost({ username, password })).data
      const user = (await api.userGet(username)).data

      this.$patch({ user })
      localStorage.setItem(LOCAL_STORAGE_KEY, JSON.stringify(user))
      localStorage.setItem('accessToken', jwt)
    },

    async logout() {
      this.$reset()
      localStorage.setItem(LOCAL_STORAGE_KEY, '')
      localStorage.setItem('accessToken', '')
    }
  }
})
