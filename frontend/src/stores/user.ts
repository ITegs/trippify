import { defineStore } from 'pinia'
import { DefaultApi, type User } from 'trippify-client/api'

const LOCAL_STORAGE_KEY = 'trippify-user'

const accessToken = localStorage.getItem('JWT') || ''
const apiConfig = {
  isJsonMime(): boolean {
    return false
  },
  accessToken
}
const api = new DefaultApi(apiConfig)

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
      localStorage.setItem('JWT', jwt)
    },

    async logout() {
      this.$reset()
      localStorage.setItem(LOCAL_STORAGE_KEY, '')
    }
  }
})
