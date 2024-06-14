import { defineStore } from 'pinia'

export const useUserStore = defineStore({
  id: 'user',
  state: () => ({
    user: localStorage.getItem("user") == null? null: JSON.parse(localStorage.getItem("user"))
  }),
  getters: {
    get: (state) => state.user
  },
  actions: {
    update(user){
      this.user = user
      localStorage.setItem("user", JSON.stringify(user))
    }
  }
})