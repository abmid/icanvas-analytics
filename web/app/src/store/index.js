import Vue from 'vue'
import Vuex from 'vuex'
import template from "./modules/template"
import auth from "./modules/auth"
import setting from "./modules/setting"

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    auth,
    template,
    setting
  },
  strict: true,
})
