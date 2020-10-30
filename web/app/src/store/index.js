import Vue from 'vue'
import Vuex from 'vuex'
import template from "./modules/template"
import auth from "./modules/auth"
import setting from "./modules/setting"
import notification from "./modules/notification"

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    auth,
    template,
    setting,
    notification
  },
  strict: true,
})
