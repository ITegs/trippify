import {defineStore} from 'pinia'
import {DefaultApi, type Spot, type Trip} from 'trippify-client/api'

export const useTripStore = defineStore('trip', {
  state: () => ({
    trip: {} as Trip,
    spots: [] as Spot[],
  }),

  getters: {},

  actions: {
    async setTrip(tripId: string) {
      const api = new DefaultApi()
      const trip: Trip = (await api.tripGet(tripId)).data
      this.$patch({trip})
    },

    async addSpotToTrip(tripId: string, spot: Spot) {
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
