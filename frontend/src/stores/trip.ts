import {defineStore} from 'pinia'
import {DefaultApi, type Spot, type Trip} from 'trippify-client/api'
import type {NewSpot} from "trippify-client";

export const useTripStore = defineStore('trip', {
  state: () => ({
    trip: {} as Trip,
    spots: [] as Spot[],
  }),

  getters: {
    // getSpot: (state) => (id: string) => {
    //   return state.spots.find((a) => a._id === id)
    // },
  },

  actions: {
    async setTrip(tripId: string) {
      const api = new DefaultApi()
      const trip: Trip = (await api.tripGet(tripId)).data
      this.$patch({trip})
    },

    async addSpotToTrip(tripId: string, spot: NewSpot) {
      const api = new DefaultApi()

      try {
        return (await api.addSpotToTripPost(tripId, spot)).data
      } catch (e) {
        return Promise.reject(e)
      }
    },

    async getSpot(spotId: string) {
      const spot = this.spots.find((s) => s._id === spotId)
      if (spot) {
        return spot
      }

      const api = new DefaultApi()
      const fetchedSpot = (await api.spotGet(spotId)).data
      this.$patch(state => {
        state.spots.push(fetchedSpot);
      });
      return fetchedSpot
    }
  }
})
