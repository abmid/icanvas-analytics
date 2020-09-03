import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import {middleware} from './helpers/middleware';
import vSelect from 'vue-select'

Vue.component('v-select', vSelect)
Vue.config.productionTip = false

// AdminLTE
require("./bootstrap")

// middleware
middleware(store, router);

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
