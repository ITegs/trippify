import { defineStore } from 'pinia'
import { DefaultApi, type Spot, type Trip } from 'trippify-client/api'
import type { NewSpot } from 'trippify-client'

const apiConfig = createApiConfig()
const api = new DefaultApi(apiConfig)

function createApiConfig(): any {
  return {
    isJsonMime: () => false,
    accessToken: getAccessToken()
  }
}

function getAccessToken(): string {
  return localStorage.getItem('accessToken') || ''
}

export const useTripStore = defineStore('trip', {
  state: () => ({
    trip: {} as Trip,
    spots: [] as Spot[]
  }),

  getters: {
    getTrip: (state) => state.trip,
    getSpot: (state) => (id: string) => {
      return state.spots.find((spot) => spot._id === id)
    }
  },

  actions: {
    async init() {
      await this.setToFirstTrip()
    },

    async setToFirstTrip() {
      try {
        const trip: Trip = (await api.firstTripGet()).data
        this.$patch({ trip })
      } catch (e) {
        console.error('Error fetching trip:', e)
        return Promise.reject(e)
      }
    },

    async setTrip(tripId: string) {
      try {
        const trip: Trip = (await api.tripGet(tripId)).data
        this.$patch({ trip })
      } catch (e) {
        console.error('Error fetching trip:', e)
        return Promise.reject(e)
      }
    },

    async addSpotToTrip(tripId: string, spot: NewSpot) {
      try {
        const newSpot = (await api.addSpotToTripPost(tripId, spot)).data
        this.spots.push(newSpot)
        return newSpot
      } catch (e) {
        console.error('Error adding spot to trip:', e)
        return Promise.reject(e)
      }
    },

    async fetchSpot(spotId: string) {
      const existingSpot = this.spots.find((s) => s._id === spotId)
      if (existingSpot) {
        return existingSpot
      }

      try {
        const spot: Spot = (await api.spotGet(spotId)).data
        this.$patch((state) => {
          state.spots.push(spot)
        })
        return spot
      } catch (e) {
        console.error('Error fetching spot:', e)
        return Promise.reject(e)
      }
    }
  }
})
