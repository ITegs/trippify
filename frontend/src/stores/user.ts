import { defineStore } from 'pinia'
import { DefaultApi, type User } from 'trippify-client/api'

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
    async getUser(username: string) {
      const api = new DefaultApi()
      return (await api.userGet(username)).data
    },

    async setUser(username: string, password: string) {
      const api = new DefaultApi()
      const user = (await api.userGet(username)).data
      this.$patch({ user })
    }
  }
})
