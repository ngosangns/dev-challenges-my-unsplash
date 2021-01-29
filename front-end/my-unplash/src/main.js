import Vue from 'vue'
import App from './App.vue'

import Vuesax from 'vuesax'
import 'vuesax/dist/vuesax.css' //Vuesax styles

import 'boxicons'

import Vuex from 'vuex'

// Vuex
Vue.use(Vuex)
const store = new Vuex.Store({
  state: {
    images: [],
  },
  getters: {
    getImages: (state) => {
      return state.images
    },
  },
  mutations: {
    appendImages: (state, arr) => {
      state.images.push(arr)
    },
    setImages: (state, arr) => {
      state.images = arr
    },
    unshiftImage: (state, el) => {
      state.images.unshift(el)
    },
    removeImage: (state, el) => {
      state.images.splice(state.images.indexOf(el), 1)
    },
  },
})

// Vusax
Vue.use(Vuesax, {
  // options here
})


Vue.config.productionTip = false
new Vue({
  render: h => h(App),
  store: store,
}).$mount('#app')

